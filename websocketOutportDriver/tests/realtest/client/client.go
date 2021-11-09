package client

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/ElrondNetwork/elrond-go-core/core/check"
	"github.com/ElrondNetwork/elrond-go-core/core/mock"
	"github.com/ElrondNetwork/elrond-go-core/data/indexer"
	"github.com/ElrondNetwork/elrond-go-core/marshal"
	"github.com/gorilla/websocket"
)

// WSConn defines what a sender shall do
type WSConn interface {
	io.Closer
	ReadMessage() (messageType int, p []byte, err error)
	WriteMessage(messageType int, data []byte) error
}

var (
	log              = &mock.LoggerMock{}
	errNilMarshaller = errors.New("nil marshaller")
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

	urlReceiveData := url.URL{Scheme: "ws", Host: fmt.Sprintf("127.0.0.1:%d", port), Path: "/operations"}
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

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:
			_ = t
		case <-interrupt:
			log.Info(tc.name + " -> interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := wsConnection.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
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

	withAck := false
	if payload[0] == byte(1) {
		withAck = true
	}

	payload = payload[1:]
	counter := payload[:8]
	payload = payload[8:]
	opType := payload[:4]
	payload = payload[4:]
	msgLength := payload[:4]
	payload = payload[4:]

	log.Info(tc.name+" -> processing payload",
		"counter", counter,
		"operation type", opType,
		"message length", msgLength,
		"data", payload,
	)

	if bytes.Compare(opType, []byte{0, 0, 0, 0}) == 0 {
		log.Debug(tc.name + " -> save block operation")
		var argsBlock indexer.ArgsSaveBlockData
		err := tc.marshaller.Unmarshal(&argsBlock, payload)
		if err != nil {
			log.Error(tc.name+" -> cannot unmarshal block", "error", err)
		} else {
			log.Info(tc.name+" -> successfully unmarshalled block", "hash", argsBlock.HeaderHash)
		}

	}

	if withAck {
		err := ackHandler.WriteMessage(websocket.BinaryMessage, counter)
		if err != nil {
			log.Error(tc.name + " -> " + err.Error())
		}
	}
}

func (tc *tempClient) Stop() {
	tc.chanStop <- true
}
