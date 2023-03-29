package client

import (
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/gorilla/websocket"
	"github.com/multiversx/mx-chain-core-go/core"
	"github.com/multiversx/mx-chain-core-go/core/check"
	"github.com/multiversx/mx-chain-core-go/core/closing"
	"github.com/multiversx/mx-chain-core-go/websocketOutportDriver/data"
	"github.com/multiversx/mx-chain-core-go/websocketOutportDriver/sender"
	logger "github.com/multiversx/mx-chain-logger-go"
)

const closedConnection = "use of closed network connection"

var log = logger.GetOrCreate("wsClient")

type client struct {
	url                      string
	blockingAckOnError       bool
	retryDuration            time.Duration
	wsConn                   WSConnClient
	payloadParser            PayloadParser
	payloadProcessor         PayloadProcessor
	uint64ByteSliceConverter sender.Uint64ByteSliceConverter
	safeCloser               core.SafeCloser
}

// ArgsWsClient holds the arguments required to create a new websocket client handler
type ArgsWsClient struct {
	Url                      string
	RetryDurationInSec       uint32
	BlockingAckOnError       bool
	PayloadProcessor         PayloadProcessor
	PayloadParser            PayloadParser
	Uint64ByteSliceConverter sender.Uint64ByteSliceConverter
	WSConnClient             WSConnClient
}

// NewWsClientHandler will create a ws client to receive data from an observer/light client
func NewWsClientHandler(args ArgsWsClient) (*client, error) {
	err := checkArgs(args)
	if err != nil {
		return nil, err
	}

	urlReceiveData := url.URL{Scheme: "ws", Host: args.Url, Path: data.WSRoute}
	retryDuration := time.Duration(args.RetryDurationInSec) * time.Second

	return &client{
		url:                      urlReceiveData.String(),
		blockingAckOnError:       args.BlockingAckOnError,
		retryDuration:            retryDuration,
		wsConn:                   args.WSConnClient,
		payloadParser:            args.PayloadParser,
		payloadProcessor:         args.PayloadProcessor,
		uint64ByteSliceConverter: args.Uint64ByteSliceConverter,
		safeCloser:               closing.NewSafeChanCloser(),
	}, nil
}

func checkArgs(args ArgsWsClient) error {
	if check.IfNil(args.PayloadProcessor) {
		return errNilPayloadProcessor
	}
	if check.IfNil(args.PayloadParser) {
		return errNilPayloadParser
	}
	if check.IfNil(args.WSConnClient) {
		return errNilWsConnReceiver
	}
	if check.IfNil(args.Uint64ByteSliceConverter) {
		return errNilUint64ByteSliceConverter
	}
	if len(args.Url) == 0 {
		return errEmptyUrl
	}
	if args.RetryDurationInSec == 0 {
		return errZeroValueRetryDuration
	}

	return nil
}

// Start will start the client listening ws process
func (c *client) Start() {
	log.Info("connecting to", "url", c.url)

	timer := time.NewTimer(c.retryDuration)
	defer timer.Stop()

	closed := false
	for !closed {
		err := c.wsConn.OpenConnection(c.url)
		if err == nil {
			closed = c.listenOnWebSocket()
		} else {
			log.Warn(fmt.Sprintf("c.openConnection(), retrying in %v...", c.retryDuration), "error", err)
		}

		timer.Reset(c.retryDuration)

		select {
		case <-timer.C:
		case <-c.safeCloser.ChanClose():
			return
		}
	}
}

func (c *client) listenOnWebSocket() (closed bool) {
	for {
		_, message, err := c.wsConn.ReadMessage()
		if err == nil {
			c.verifyPayloadAndSendAckIfNeeded(message)
			continue
		}

		_, isConnectionClosed := err.(*websocket.CloseError)
		if !isConnectionClosed {
			if strings.Contains(err.Error(), closedConnection) {
				log.Info("connection closed by server")
				return true
			}
			log.Warn("c.listenOnWebSocket()-> connection problem, retrying", "error", err.Error())
			return
		}

		log.Warn(fmt.Sprintf("websocket terminated by the server side, retrying in %v...", c.retryDuration), "error", err.Error())
		return
	}
}

func (c *client) verifyPayloadAndSendAckIfNeeded(payload []byte) {
	if len(payload) == 0 {
		log.Error("empty payload")
		return
	}

	payloadData, err := c.payloadParser.ExtractPayloadData(payload)
	if err != nil {
		log.Error("error while extracting payload data: ", "error", err)
		return
	}

	log.Info("processing payload",
		"counter", payloadData.Counter,
		"operation type", payloadData.OperationType,
		"payload", payloadData.Payload,
		"message length", len(payloadData.Payload),
	)

	err = c.payloadProcessor.ProcessPayload(payloadData)
	c.sendAckIfNeeded(payloadData, err)
}

func (c *client) sendAckIfNeeded(payloadData *data.PayloadData, err error) {
	log.LogIfError(err)

	if !payloadData.WithAcknowledge {
		return
	}

	if err != nil && c.blockingAckOnError {
		return
	}

	c.waitForAckSignal(payloadData.Counter)
}

func (c *client) waitForAckSignal(counter uint64) {
	timer := time.NewTimer(c.retryDuration)
	defer timer.Stop()

	counterBytes := c.uint64ByteSliceConverter.ToByteSlice(counter)
	for {
		timer.Reset(c.retryDuration)

		err := c.wsConn.WriteMessage(websocket.BinaryMessage, counterBytes)
		if err == nil {
			return
		}

		log.Error("could not write acknowledge message", "error", err.Error(), "retrying in", c.retryDuration)

		select {
		case <-timer.C:
		case <-c.safeCloser.ChanClose():
			return
		}
	}
}

// Close will close the underlying ws connection
func (c *client) Close() {
	defer c.safeCloser.Close()

	log.Info("closing all components...")
	err := c.wsConn.Close()
	if err != nil {
		log.Error("cannot close ws connection", "error", err)
	}

	err = c.payloadProcessor.Close()
	if err != nil {
		log.Error("cannot close the operations handler", "error", err)
	}
}
