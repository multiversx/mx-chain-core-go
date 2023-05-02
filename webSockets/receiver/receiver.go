package receiver

import (
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/multiversx/mx-chain-core-go/core"
	"github.com/multiversx/mx-chain-core-go/core/check"
	"github.com/multiversx/mx-chain-core-go/core/closing"
	"github.com/multiversx/mx-chain-core-go/webSockets"
	"github.com/multiversx/mx-chain-core-go/webSockets/data"
)

// ArgsReceiver holds the arguments that are needed for a receiver
type ArgsReceiver struct {
	Uint64ByteSliceConverter webSockets.Uint64ByteSliceConverter
	Log                      core.Logger
	RetryDurationInSec       int
	BlockingAckOnError       bool
}

type receiver struct {
	payloadParser      webSockets.PayloadParser
	payloadHandler     webSockets.PayloadHandler
	log                core.Logger
	safeCloser         core.SafeCloser
	retryDuration      time.Duration
	blockingAckOnError bool
	mutex              sync.RWMutex
}

// NewReceiver will create a new instance of receiver
func NewReceiver(args ArgsReceiver) (*receiver, error) {
	err := checkArgs(args)
	if err != nil {
		return nil, err
	}

	payloadParser, err := webSockets.NewWebSocketPayloadParser(args.Uint64ByteSliceConverter)
	if err != nil {
		return nil, err
	}

	return &receiver{
		log:                args.Log,
		retryDuration:      time.Duration(args.RetryDurationInSec) * time.Second,
		blockingAckOnError: args.BlockingAckOnError,
		safeCloser:         closing.NewSafeChanCloser(),
		payloadHandler:     webSockets.NewNilPayloadHandler(),
		payloadParser:      payloadParser,
	}, nil
}

func checkArgs(args ArgsReceiver) error {
	if check.IfNil(args.Log) {
		return core.ErrNilLogger
	}
	if check.IfNil(args.Uint64ByteSliceConverter) {
		return data.ErrNilUint64ByteSliceConverter
	}
	if args.RetryDurationInSec == 0 {
		return data.ErrZeroValueRetryDuration
	}
	return nil
}

// SetPayloadHandler will set the payload handler
func (r *receiver) SetPayloadHandler(handler webSockets.PayloadHandler) error {
	if check.IfNilReflect(handler) {
		return data.ErrNilPayloadProcessor
	}

	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.payloadHandler = handler
	return nil
}

// Listen will listen for messages from the provided connection
func (r *receiver) Listen(connection webSockets.WSConClient) (closed bool) {
	timer := time.NewTimer(r.retryDuration)
	defer timer.Stop()

	isClosed := false
	connection.SetCloseHandler(func(code int, text string) error {
		isClosed = true
		return nil
	})

	for !isClosed {
		_, message, err := connection.ReadMessage()
		if err == nil {
			r.verifyPayloadAndSendAckIfNeeded(connection, message)
			continue
		}

		// TODO will handle the error in the PR with the integration tests
		timer.Reset(r.retryDuration)
		r.log.Warn("r.Listen()-> connection problem", "error", err.Error())

		select {
		case <-r.safeCloser.ChanClose():
			return
		case <-timer.C:
		}
	}

	return isClosed
}

func (r *receiver) verifyPayloadAndSendAckIfNeeded(connection webSockets.WSConClient, payload []byte) {
	if len(payload) == 0 {
		r.log.Debug("r.verifyPayloadAndSendAckIfNeeded(): empty payload")
		return
	}

	payloadData, err := r.payloadParser.ExtractPayloadData(payload)
	if err != nil {
		r.log.Warn("r.verifyPayloadAndSendAckIfNeeded: cannot extract payload data", "error", err.Error())
		return
	}

	err = r.payloadHandler.ProcessPayload(payloadData.Payload)
	if err != nil && r.blockingAckOnError {
		r.log.Debug("r.payloadHandler.ProcessPayload: cannot handler payload", "error", err)
		return
	}

	r.sendAckIfNeeded(connection, payloadData)
}

func (r *receiver) sendAckIfNeeded(connection webSockets.WSConClient, payloadData *data.PayloadData) {
	if !payloadData.WithAcknowledge {
		return
	}

	timer := time.NewTimer(r.retryDuration)
	defer timer.Stop()

	counterBytes := r.payloadParser.EncodeUint64(payloadData.Counter)
	for {
		timer.Reset(r.retryDuration)

		err := connection.WriteMessage(websocket.BinaryMessage, counterBytes)
		if err == nil {
			return
		}

		if !strings.Contains(err.Error(), data.ErrConnectionNotOpened.Error()) {
			r.log.Error("could not write acknowledge message", "error", err.Error(), "retrying in", r.retryDuration)
		}

		r.log.Debug("r.sendAckIfNeeded(): cannot write ack", "error", err)

		select {
		case <-timer.C:
		case <-r.safeCloser.ChanClose():
			return
		}
	}
}

// Close will close the underlying ws connection
func (r *receiver) Close() error {
	defer r.safeCloser.Close()

	err := r.payloadHandler.Close()
	if err != nil {
		r.log.Error("cannot close the operations handler", "error", err)
	}

	return nil
}
