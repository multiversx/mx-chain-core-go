package client

import (
	"fmt"
	"strings"
	"time"

	"github.com/gorilla/websocket"
	"github.com/multiversx/mx-chain-core-go/core/check"
	"github.com/multiversx/mx-chain-core-go/websocketOutportDriver"
	"github.com/multiversx/mx-chain-core-go/websocketOutportDriver/data"
	"github.com/multiversx/mx-chain-core-go/websocketOutportDriver/sender"
	logger "github.com/multiversx/mx-chain-logger-go"
)

const closedConnection = "use of closed network connection"

var (
	log           = logger.GetOrCreate("process/wsclient")
	retryDuration = time.Second * 5
)

type client struct {
	operationHandler         OperationHandler
	wsConn                   data.WSConn
	url                      string
	uint64ByteSliceConverter sender.Uint64ByteSliceConverter
}

type ArgsWsClient struct {
	Url                      string
	OperationHandler         OperationHandler
	Uint64ByteSliceConverter sender.Uint64ByteSliceConverter
}

// NewWsClient will create a ws client to receive data from an observer/light client
func NewWsClient(args *ArgsWsClient) (*client, error) {
	if args.OperationHandler == nil {
		return nil, errNilOperationHandler
	}

	return &client{
		operationHandler:         args.OperationHandler,
		url:                      args.Url,
		uint64ByteSliceConverter: args.Uint64ByteSliceConverter,
	}, nil
}

// Start will start the client listening ws process
func (c *client) Start() {
	log.Info("connecting to", "url", c.url)

	for {
		err := c.openConnection()
		if err != nil {
			log.Warn(fmt.Sprintf("c.openConnection(), retrying in %v...", retryDuration), "error", err.Error())
			time.Sleep(retryDuration)
			continue
		}

		closed := c.listeningOnWebSocket()
		if closed {
			return
		}
	}
}

func (c *client) openConnection() error {
	var err error
	c.wsConn, _, err = websocket.DefaultDialer.Dial(c.url, nil)
	if err != nil {
		return err
	}

	return nil
}

func (c *client) listeningOnWebSocket() (closed bool) {
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
			log.Warn("c.listeningOnWebSocket()-> connection problem, retrying", "error", err.Error())
		} else {
			log.Warn(fmt.Sprintf("websocket terminated by the server side, retrying in %v...", retryDuration), "error", err.Error())
		}
		return
	}

}

func (c *client) verifyPayloadAndSendAckIfNeeded(payload []byte) {
	if len(payload) == 0 {
		log.Error("empty payload")
		return
	}

	payloadParser, _ := websocketOutportDriver.NewWebSocketPayloadParser(c.uint64ByteSliceConverter)
	payloadData, err := payloadParser.ExtractPayloadData(payload)
	if err != nil {
		log.Error("error while extracting payload data: " + err.Error())
		return
	}

	log.Info("processing payload",
		"counter", payloadData.Counter,
		"operation type", payloadData.OperationType,
		"message length", len(payloadData.Payload),
	)

	function, ok := c.operationHandler.GetOperationHandler(payloadData.OperationType)
	if !ok {
		log.Warn("invalid operation", "operation type", payloadData.OperationType.String())
	}

	err = function(payloadData.Payload)
	if err != nil {
		log.Error("something went wrong", "error", err.Error())
	}

	if payloadData.WithAcknowledge {
		counterBytes := c.uint64ByteSliceConverter.ToByteSlice(payloadData.Counter)
		err = c.wsConn.WriteMessage(websocket.BinaryMessage, counterBytes)
		if err != nil {
			log.Error("write acknowledge message", "error", err.Error())
		}
	}
}

func (c *client) closeWsConnection() {
	log.Debug("closing ws connection...")
	if check.IfNilReflect(c.wsConn) {
		return
	}

	//Cleanly close the connection by sending a close message and then
	//waiting (with timeout) for the server to close the connection.
	err := c.wsConn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		log.Error("cannot send close message", "error", err)
	}
	err = c.wsConn.Close()
	if err != nil {
		log.Error("cannot close ws connection", "error", err)
	}
}

// Close will close the underlying ws connection
func (c *client) Close() {
	log.Info("closing all components...")
	c.closeWsConnection()

	err := c.operationHandler.Close()
	if err != nil {
		log.Error("cannot close the operations handler", "error", err)
	}
}
