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
		_, message, err := connection.ReadMessage()
		if err == nil {
			wt.verifyPayloadAndSendAckIfNeeded(connection, message)
			continue
		}

		// TODO will handle the error in the PR with the integration tests
		timer.Reset(wt.retryDuration)
		wt.log.Warn("wt.Listen()-> connection problem", "error", err.Error())

		select {
		case <-wt.safeCloser.ChanClose():
			return
		case <-timer.C:
		}
	}
}

func (wt *wsTransceiver) verifyPayloadAndSendAckIfNeeded(connection webSocket.WSConClient, payload []byte) {
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

func (wt *wsTransceiver) sendAckIfNeeded(connection webSocket.WSConClient, payloadData *data.PayloadData) {
	if !payloadData.WithAcknowledge {
		return
	}

	timer := time.NewTimer(wt.retryDuration)
	defer timer.Stop()

	counterBytes := wt.payloadParser.EncodeUint64(payloadData.Counter)
	for {
		timer.Reset(wt.retryDuration)

		err := connection.WriteMessage(websocket.BinaryMessage, counterBytes)
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

	return wt.sendPayload(newPayload, assignedCounter, connection)
}

func (wt *wsTransceiver) sendPayload(payload []byte, assignedCounter uint64, connection webSocket.WSConClient) error {
	errSend := connection.WriteMessage(websocket.BinaryMessage, payload)
	if errSend != nil {
		return errSend
	}

	if !wt.withAcknowledge {
		return nil
	}

	return wt.waitForAck(assignedCounter, connection)
}

func (wt *wsTransceiver) waitForAck(assignedCounter uint64, connection webSocket.WSConClient) error {
	for {
		select {
		case <-wt.safeCloser.ChanClose():
			return nil
		default:
		}

		mType, message, err := connection.ReadMessage()
		if err != nil {
			wt.log.Debug("s.waitForAck(): cannot read message", "id", connection.GetID(), "error", err)
			continue
		}

		if mType != websocket.BinaryMessage {
			wt.log.Debug("received message is not binary message", "id", connection.GetID(), "message type", mType)
			continue
		}

		wt.log.Trace("received ack", "remote addr", connection.GetID(), "message", message)

		receivedCounter, err := wt.payloadParser.DecodeUint64(message)
		if err != nil {
			wt.log.Warn("cannot decode counter: bytes to uint64",
				"id", connection.GetID(),
				"counter bytes", message,
				"error", err,
			)
			continue
		}

		if receivedCounter != assignedCounter {
			wt.log.Debug("s.waitForAck invalid counter", "expected", assignedCounter, "received", receivedCounter, "id", connection.GetID())
			continue
		}

		return nil

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
