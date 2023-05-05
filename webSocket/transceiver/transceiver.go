package transceiver

import (
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gorilla/websocket"
	"github.com/multiversx/mx-chain-core-go/core"
	"github.com/multiversx/mx-chain-core-go/core/check"
	"github.com/multiversx/mx-chain-core-go/core/closing"
	"github.com/multiversx/mx-chain-core-go/webSocket"
	"github.com/multiversx/mx-chain-core-go/webSocket/data"
)

// ArgsTransceiver holds the arguments that are needed for a transceiver
type ArgsTransceiver struct {
	PayloadConverter   webSocket.PayloadConverter
	Log                core.Logger
	RetryDurationInSec int
	BlockingAckOnError bool
	WithAcknowledge    bool
}

type wsTransceiver struct {
	payloadParser      webSocket.PayloadConverter
	payloadHandler     webSocket.PayloadHandler
	log                core.Logger
	safeCloser         core.SafeCloser
	retryDuration      time.Duration
	mutex              sync.RWMutex
	ackChan            chan struct{}
	counter            uint64
	blockingAckOnError bool
	withAcknowledge    bool
}

// NewTransceiver will create a new instance of transceiver
func NewTransceiver(args ArgsTransceiver) (*wsTransceiver, error) {
	err := checkArgs(args)
	if err != nil {
		return nil, err
	}

	return &wsTransceiver{
		log:                args.Log,
		retryDuration:      time.Duration(args.RetryDurationInSec) * time.Second,
		blockingAckOnError: args.BlockingAckOnError,
		safeCloser:         closing.NewSafeChanCloser(),
		payloadHandler:     webSocket.NewNilPayloadHandler(),
		payloadParser:      args.PayloadConverter,
		withAcknowledge:    args.WithAcknowledge,
		ackChan:            make(chan struct{}),
	}, nil
}

func checkArgs(args ArgsTransceiver) error {
	if check.IfNil(args.Log) {
		return core.ErrNilLogger
	}
	if check.IfNil(args.PayloadConverter) {
		return data.ErrNilPayloadConverter
	}
	if args.RetryDurationInSec == 0 {
		return data.ErrZeroValueRetryDuration
	}
	return nil
}

// SetPayloadHandler will set the payload handler
func (wt *wsTransceiver) SetPayloadHandler(handler webSocket.PayloadHandler) error {
	if check.IfNil(handler) {
		return data.ErrNilPayloadProcessor
	}

	wt.mutex.Lock()
	defer wt.mutex.Unlock()

	wt.payloadHandler = handler
	return nil
}

// Listen will listen for messages from the provided connection
func (wt *wsTransceiver) Listen(connection webSocket.WSConClient) (closed bool) {
	timer := time.NewTimer(wt.retryDuration)
	defer timer.Stop()

	for {
		msgType, message, err := connection.ReadMessage()
		if err == nil {
			wt.verifyPayloadAndSendAckIfNeeded(connection, message, msgType)
			continue
		}

		_, isConnectionClosed := err.(*websocket.CloseError)
		if !strings.Contains(err.Error(), data.ClosedConnectionMessage) && !isConnectionClosed {
			wt.log.Warn("wt.Listen()-> connection problem", "error", err.Error())
		}
		if isConnectionClosed {
			wt.log.Info("received connection close")
			return true
		}

		timer.Reset(wt.retryDuration)

		select {
		case <-wt.safeCloser.ChanClose():
			return
		case <-timer.C:
		}
	}
}

func (wt *wsTransceiver) verifyPayloadAndSendAckIfNeeded(connection webSocket.WSConClient, payload []byte, msgType int) {
	if msgType == websocket.TextMessage {
		wt.handleAckMessage(payload)
		return
	}

	if len(payload) == 0 {
		wt.log.Debug("wt.verifyPayloadAndSendAckIfNeeded(): empty payload")
		return
	}

	payloadData, err := wt.payloadParser.ExtractPayloadData(payload)
	if err != nil {
		wt.log.Warn("wt.verifyPayloadAndSendAckIfNeeded: cannot extract payload data", "error", err.Error())
		return
	}

	err = wt.payloadHandler.ProcessPayload(payloadData.Payload)
	if err != nil && wt.blockingAckOnError {
		wt.log.Debug("wt.payloadHandler.ProcessPayload: cannot handler payload", "error", err)
		return
	}

	wt.sendAckIfNeeded(connection, payloadData)
}

func (wt *wsTransceiver) handleAckMessage(payload []byte) {
	counter, err := wt.payloadParser.DecodeUint64(payload)
	if err != nil {
		wt.log.Debug("wsTransceiver.handleAckMessage cannot decode the ack message", "error", err)
		return
	}

	wt.log.Trace("wt.handleAckMessage: received ack", "counter", counter)
	expectedCounter := atomic.LoadUint64(&wt.counter)
	if expectedCounter == counter {
		wt.ackChan <- struct{}{}
		return
	}

	wt.log.Debug("wsTransceiver.handleAckMessage invalid counter received", "expected", expectedCounter, "received", counter)

	return
}

func (wt *wsTransceiver) sendAckIfNeeded(connection webSocket.WSConClient, payloadData *data.PayloadData) {
	if !payloadData.WithAcknowledge {
		return
	}

	timer := time.NewTimer(wt.retryDuration)
	defer timer.Stop()

	counterBytes := wt.payloadParser.EncodeUint64(payloadData.Counter)
	for {
		timer.Reset(wt.retryDuration)

		err := connection.WriteMessage(websocket.TextMessage, counterBytes)
		if err == nil {
			return
		}

		if !strings.Contains(err.Error(), data.ErrConnectionNotOpen.Error()) {
			wt.log.Error("could not write acknowledge message", "error", err.Error(), "retrying in", wt.retryDuration)
		}

		wt.log.Debug("wt.sendAckIfNeeded(): cannot write ack", "error", err)

		select {
		case <-timer.C:
		case <-wt.safeCloser.ChanClose():
			return
		}
	}
}

// Send will prepare and send the provided WsSendArgs
func (wt *wsTransceiver) Send(args data.WsSendArgs, connection webSocket.WSConClient) error {
	assignedCounter := atomic.AddUint64(&wt.counter, 1)
	newPayload := wt.payloadParser.ConstructPayloadData(args, assignedCounter, wt.withAcknowledge)

	return wt.sendPayload(newPayload, connection)
}

func (wt *wsTransceiver) sendPayload(payload []byte, connection webSocket.WSConClient) error {
	errSend := connection.WriteMessage(websocket.BinaryMessage, payload)
	if errSend != nil {
		return errSend
	}

	if !wt.withAcknowledge {
		return nil
	}

	wt.waitForAck()
	return nil
}

func (wt *wsTransceiver) waitForAck() {
	// wait for ack
	select {
	case <-wt.ackChan:
		return
	case <-wt.safeCloser.ChanClose():
		return
	}
}

// Close will close the underlying ws connection
func (wt *wsTransceiver) Close() error {
	defer wt.safeCloser.Close()

	err := wt.payloadHandler.Close()
	if err != nil {
		wt.log.Debug("cannot close the operations handler", "error", err)
	}

	return err
}
