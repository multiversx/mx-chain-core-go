package client

import (
	"errors"
	"fmt"
	"io"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
	"github.com/multiversx/mx-chain-core-go/core/check"
	"github.com/multiversx/mx-chain-core-go/data/outport"
	"github.com/multiversx/mx-chain-core-go/data/typeConverters/uint64ByteSlice"
	"github.com/multiversx/mx-chain-core-go/marshal"
	"github.com/multiversx/mx-chain-core-go/marshal/factory"
	"github.com/multiversx/mx-chain-core-go/webSocket"
	"github.com/multiversx/mx-chain-core-go/webSocket/data"
	logger "github.com/multiversx/mx-chain-logger-go"
)

// WSConn defines what a sender shall do
type WSConn interface {
	io.Closer
	ReadMessage() (messageType int, p []byte, err error)
	WriteMessage(messageType int, data []byte) error
}

var (
	log                      = logger.GetOrCreate("test-client")
	errNilMarshaller         = errors.New("nil marshaller")
	uint64ByteSliceConverter = uint64ByteSlice.NewBigEndianConverter()
)

type tempClient struct {
	name       string
	marshaller marshal.Marshalizer
	chanStop   chan bool
}

// NewTempClient will return a new instance of tempClient
func NewTempClient(name string, marshaller marshal.Marshalizer) (*tempClient, error) {
	if check.IfNil(marshaller) {
		return nil, errNilMarshaller
	}

	return &tempClient{
		name:       name,
		marshaller: marshaller,
		chanStop:   make(chan bool),
	}, nil
}

// Run will start the client on the provided port
func (tc *tempClient) Run(port int) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	urlReceiveData := url.URL{Scheme: "ws", Host: fmt.Sprintf("127.0.0.1:%d", port), Path: data.WSRoute}
	log.Info(tc.name+" -> connecting to", "url", urlReceiveData.String())
	wsConnection, _, err := websocket.DefaultDialer.Dial(urlReceiveData.String(), nil)
	if err != nil {
		log.Error(tc.name+" -> dial", "error", err)
	}
	defer func() {
		err = wsConnection.Close()
		log.LogIfError(err)
	}()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := wsConnection.ReadMessage()
			if err != nil {
				log.Error(tc.name+" -> error read message", "error", err)
				return
			}

			tc.verifyPayloadAndSendAckIfNeeded(message, wsConnection)
		}
	}()

	timer := time.NewTimer(time.Second)
	defer timer.Stop()

	for {
		select {
		case <-done:
			return
		case <-timer.C:
		case <-interrupt:
			log.Info(tc.name + " -> interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err = wsConnection.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Error(tc.name+" -> write close", "error", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}

func (tc *tempClient) verifyPayloadAndSendAckIfNeeded(payload []byte, ackHandler WSConn) {
	if len(payload) == 0 {
		log.Error(tc.name + " -> empty payload")
		return
	}

	marshaller, _ := factory.NewMarshalizer("json")
	payloadParser, _ := webSocket.NewWebSocketPayloadConverter(marshaller)
	wsMessage, err := payloadParser.ExtractWsMessage(payload)
	if err != nil {
		log.Error(tc.name + " -> error while extracting payload data: " + err.Error())
		return
	}

	log.Info(tc.name+" -> processing payload",
		"counter", wsMessage.Counter,
		"topic", wsMessage.Topic,
		"message length", len(wsMessage.Payload),
		"data", wsMessage.Payload,
		"with ack", wsMessage.WithAcknowledge,
	)

	if wsMessage.Topic == outport.TopicSaveBlock {
		log.Debug(tc.name + " -> save block topic")
		var argsBlock outport.OutportBlock
		err = tc.marshaller.Unmarshal(&argsBlock, wsMessage.Payload)
		if err != nil {
			log.Error(tc.name+" -> cannot unmarshal block", "error", err)
		} else {
			log.Info(tc.name+" -> successfully unmarshalled block", "hash", argsBlock.BlockData.HeaderHash)
		}
	}

	if wsMessage.WithAcknowledge {
		counterBytes := uint64ByteSliceConverter.ToByteSlice(wsMessage.Counter)
		err = ackHandler.WriteMessage(websocket.BinaryMessage, counterBytes)
		if err != nil {
			log.Error(tc.name + " -> " + err.Error())
		}
	}
}

// Stop -
func (tc *tempClient) Stop() {
	tc.chanStop <- true
}
