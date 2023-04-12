package payload

import (
	"strings"
	"time"

	"github.com/gorilla/websocket"
	"github.com/multiversx/mx-chain-core-go/core"
	"github.com/multiversx/mx-chain-core-go/core/check"
	"github.com/multiversx/mx-chain-core-go/core/closing"
	"github.com/multiversx/mx-chain-core-go/webSockets/common"
	"github.com/multiversx/mx-chain-core-go/webSockets/data"
)

type ArgsMessagesProcessor struct {
	Log                      core.Logger
	PayloadParser            common.PayloadParser
	PayloadProcessor         common.PayloadProcessor
	WsClient                 common.WSConClient
	Uint64ByteSliceConverter common.Uint64ByteSliceConverter
	RetryDurationInSec       uint32
	BlockingAckOnError       bool
}

type messagesListener struct {
	safeCloser               core.SafeCloser
	log                      core.Logger
	payloadParser            common.PayloadParser
	payloadProcessor         common.PayloadProcessor
	wsClient                 common.WSConClient
	uint64ByteSliceConverter common.Uint64ByteSliceConverter
	retryDuration            time.Duration
	blockingAckOnError       bool
}

// NewMessagesListener will create a new instance of messagesListener
func NewMessagesListener(args ArgsMessagesProcessor) (*messagesListener, error) {
	err := checkArgs(args)
	if err != nil {
		return nil, err
	}

	return &messagesListener{
		log:                      args.Log,
		payloadParser:            args.PayloadParser,
		payloadProcessor:         args.PayloadProcessor,
		wsClient:                 args.WsClient,
		uint64ByteSliceConverter: args.Uint64ByteSliceConverter,
		retryDuration:            time.Duration(args.RetryDurationInSec) * time.Second,
		blockingAckOnError:       args.BlockingAckOnError,
		safeCloser:               closing.NewSafeChanCloser(),
	}, nil
}

// Listen will listen for messages
func (ml *messagesListener) Listen() (closed bool) {
	for {
		_, message, err := ml.wsClient.ReadMessage()
		if err == nil {
			ml.verifyPayloadAndSendAckIfNeeded(message)
			continue
		}

		_, isConnectionClosed := err.(*websocket.CloseError)
		if !isConnectionClosed {
			if strings.Contains(err.Error(), data.ClosedConnectionMessage) {
				ml.log.Info("connection closed by server")
				return true
			}
			if strings.Contains(err.Error(), data.ErrConnectionNotOpened.Error()) {
				return
			}

			ml.log.Warn("c.listenOnWebSocket()-> connection problem, retrying", "error", err.Error())
			return
		}

		ml.log.Warn("websocket terminated", "error", err.Error())
		return
	}
}

func (ml *messagesListener) verifyPayloadAndSendAckIfNeeded(payload []byte) {
	if len(payload) == 0 {
		ml.log.Error("empty payload")
		return
	}

	payloadData, err := ml.payloadParser.ExtractPayloadData(payload)
	if err != nil {
		ml.log.Error("error while extracting payload data: ", "error", err)
		return
	}

	ml.log.Info("processing payload",
		"counter", payloadData.Counter,
		"operation type", payloadData.OperationType,
		"message length", len(payloadData.Payload),
	)

	ml.log.Trace("processing payload data", "payload", payloadData.Payload)

	err = ml.payloadProcessor.ProcessPayload(payloadData)
	ml.sendAckIfNeeded(payloadData, err)
}

func (ml *messagesListener) sendAckIfNeeded(payloadData *data.PayloadData, err error) {
	ml.log.LogIfError(err)

	if !payloadData.WithAcknowledge {
		return
	}

	if err != nil && ml.blockingAckOnError {
		return
	}

	ml.waitForAckSignal(payloadData.Counter)
}

func (ml *messagesListener) waitForAckSignal(counter uint64) {
	timer := time.NewTimer(ml.retryDuration)
	defer timer.Stop()

	counterBytes := ml.uint64ByteSliceConverter.ToByteSlice(counter)
	for {
		timer.Reset(ml.retryDuration)

		err := ml.wsClient.WriteMessage(websocket.BinaryMessage, counterBytes)
		if err == nil {
			return
		}

		if !strings.Contains(err.Error(), data.ErrConnectionNotOpened.Error()) {
			ml.log.Error("could not write acknowledge message", "error", err.Error(), "retrying in", ml.retryDuration)
		}

		select {
		case <-timer.C:
		case <-ml.safeCloser.ChanClose():
			return
		}
	}
}

// Close will close the underlying ws connection
func (ml *messagesListener) Close() {
	defer ml.safeCloser.Close()

	ml.log.Info("closing all components...")
	err := ml.wsClient.Close()
	if err != nil {
		ml.log.Error("cannot close ws connection", "error", err)
	}

	err = ml.payloadProcessor.Close()
	if err != nil {
		ml.log.Error("cannot close the operations handler", "error", err)
	}
}

func checkArgs(args ArgsMessagesProcessor) error {
	if check.IfNil(args.PayloadProcessor) {
		return data.ErrNilPayloadProcessor
	}
	if check.IfNil(args.PayloadParser) {
		return data.ErrNilPayloadParser
	}
	if check.IfNil(args.Uint64ByteSliceConverter) {
		return data.ErrNilUint64ByteSliceConverter
	}
	if args.RetryDurationInSec == 0 {
		return data.ErrZeroValueRetryDuration
	}
	if check.IfNil(args.Log) {
		return data.ErrNilLogger
	}
	if check.IfNilReflect(args.WsClient) {
		return data.ErrNilWebSocketClient
	}

	return nil
}
