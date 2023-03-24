package client

import (
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/gorilla/websocket"
	"github.com/multiversx/mx-chain-core-go/websocketOutportDriver"
	"github.com/multiversx/mx-chain-core-go/websocketOutportDriver/data"
	"github.com/multiversx/mx-chain-core-go/websocketOutportDriver/sender"
	logger "github.com/multiversx/mx-chain-logger-go"
)

const closedConnection = "use of closed network connection"

var log = logger.GetOrCreate("process/wsclient")

type client struct {
	url                      string
	blockingOnError          bool
	retryDuration            time.Duration
	wsConn                   WSConnClient
	payloadParser            PayloadParser
	operationHandler         OperationHandler
	uint64ByteSliceConverter sender.Uint64ByteSliceConverter
}

type ArgsWsClient struct {
	Url                      string
	RetryDurationInSec       uint32
	BlockingAckOnError       bool
	OperationHandler         OperationHandler
	PayloadParser            PayloadParser
	Uint64ByteSliceConverter sender.Uint64ByteSliceConverter
	WSConnClient             WSConnClient
}

// NewWsClientHandler will create a ws client to receive data from an observer/light client
func NewWsClientHandler(args *ArgsWsClient) (*client, error) {
	if args.OperationHandler == nil {
		return nil, errNilOperationHandler
	}
	if args.PayloadParser == nil {
		return nil, errNilPayloadParser
	}
	if args.WSConnClient == nil {
		return nil, errNilWsConnReceiver
	}
	if args.Uint64ByteSliceConverter == nil {
		return nil, errNilUint64ByteSliceConverter
	}
	if len(args.Url) == 0 {
		return nil, errEmptyUrlProvided
	}

	urlReceiveData := url.URL{Scheme: "ws", Host: args.Url, Path: data.WSRoute}
	retryDuration := time.Duration(args.RetryDurationInSec) * time.Second

	return &client{
		url:                      urlReceiveData.String(),
		blockingOnError:          args.BlockingAckOnError,
		retryDuration:            retryDuration,
		wsConn:                   args.WSConnClient,
		payloadParser:            args.PayloadParser,
		operationHandler:         args.OperationHandler,
		uint64ByteSliceConverter: args.Uint64ByteSliceConverter,
	}, nil
}

// Start will start the client listening ws process
func (c *client) Start() {
	log.Info("connecting to", "url", c.url)

	for {
		err := c.wsConn.OpenConnection(c.url)
		if err != nil {
			log.Warn(fmt.Sprintf("c.openConnection(), retrying in %v...", c.retryDuration), "error", err.Error())
			time.Sleep(c.retryDuration)
			continue
		}

		closed := c.listenOnWebSocket()
		if closed {
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
				return true
			}
			log.Warn("c.listenOnWebSocket()-> connection problem, retrying", "error", err.Error())
		} else {
			log.Warn(fmt.Sprintf("websocket terminated by the server side, retrying in %v...", c.retryDuration), "error", err.Error())
		}
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
		log.Error("error while extracting payload data: " + err.Error())
		return
	}

	log.Info("processing payload",
		"counter", payloadData.Counter,
		"operation type", payloadData.OperationType,
		"payload", payloadData.Payload,
		"message length", len(payloadData.Payload),
	)

	function, ok := c.operationHandler.GetOperationHandler(payloadData.OperationType)
	if !ok {
		log.Warn("invalid operation", "operation type", payloadData.OperationType.String())
		c.sendAckIfNeeded(payloadData, true)
		return
	}

	err = function(payloadData.Payload)
	if err != nil {
		log.Error("could not process payload", "error", err.Error())
		c.sendAckIfNeeded(payloadData, true)
		return
	}

	c.sendAckIfNeeded(payloadData, false)
}

func (c *client) sendAckIfNeeded(payloadData *websocketOutportDriver.PayloadData, hadError bool) {
	if !payloadData.WithAcknowledge {
		return
	}

	if hadError && c.blockingOnError {
		return
	}

	c.waitForAckSignal(payloadData.Counter)
}

func (c *client) waitForAckSignal(counter uint64) {
	for {
		counterBytes := c.uint64ByteSliceConverter.ToByteSlice(counter)
		err := c.wsConn.WriteMessage(websocket.BinaryMessage, counterBytes)
		if err != nil {
			log.Error("could not write acknowledge message",
				"error", err.Error(), "retrying in", c.retryDuration)

			time.Sleep(time.Second * c.retryDuration)
			continue
		}

		return
	}
}

// Close will close the underlying ws connection
func (c *client) Close() {
	log.Info("closing all components...")
	err := c.wsConn.Close()
	if err != nil {
		log.Error("cannot close ws connection", "error", err)
	}

	err = c.operationHandler.Close()
	if err != nil {
		log.Error("cannot close the operations handler", "error", err)
	}
}
