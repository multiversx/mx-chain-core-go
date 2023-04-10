package client

import (
	"fmt"
	"time"

	"github.com/gorilla/websocket"
	"github.com/multiversx/mx-chain-core-go/core"
	"github.com/multiversx/mx-chain-core-go/websocketOutportDriver/common"
)

// WebSocketClientSenderArgs holds the arguments needed for creating a new instance of client
type WebSocketClientSenderArgs struct {
	Uint64ByteSliceConverter Uint64ByteSliceConverter
	Log                      core.Logger
	RetryDurationInSec       int
	WithAcknowledge          bool
	URL                      string
}

type client struct {
	log                      core.Logger
	uint64ByteSliceConverter Uint64ByteSliceConverter
	wsConn                   WSConnClient
	retryDuration            time.Duration
	safeCloser               core.SafeCloser
	url                      string
	withAcknowledge          bool
}

func NewClient(args WebSocketClientSenderArgs) (*client, error) {
	return &client{
		uint64ByteSliceConverter: args.Uint64ByteSliceConverter,
		retryDuration:            time.Duration(args.RetryDurationInSec) * time.Second,
		url:                      args.URL,
		wsConn:                   common.NewWSConnClient(),
		withAcknowledge:          args.WithAcknowledge,
		log:                      args.Log,
	}, nil
}

func (c *client) Send(counter uint64, payload []byte) error {
	err := c.writeMessage(payload)
	if err != nil {
		c.log.Warn("cannot write message", "error", err)
	}

	if !c.withAcknowledge {
		return nil
	}

	return c.waitForAckSignal(counter)
}

func (c *client) waitForAckSignal(counter uint64) error {
	for {
		mType, message, err := c.wsConn.ReadMessage()
		if err != nil {
			c.log.Error("cannot read message", "error", err)
			break
		}

		if mType != websocket.BinaryMessage {
			c.log.Warn("received message is not binary message", "message type", mType)
			continue
		}

		receivedCounter, err := c.uint64ByteSliceConverter.ToUint64(message)
		if err != nil {
			c.log.Warn("cannot decode counter: bytes to uint64",
				"counter bytes", message,
				"error", err,
			)
			continue
		}

		if receivedCounter == counter {
			return nil
		}
	}
	return nil
}

func (c *client) writeMessage(payload []byte) error {
	timer := time.NewTimer(c.retryDuration)
	defer timer.Stop()

	for {
		err := c.wsConn.WriteMessage(websocket.BinaryMessage, payload)
		if err == nil {
			return nil
		}

		c.log.Warn("client.WriteMessage: connection problem", "error", err)

		err = c.wsConn.OpenConnection(c.url)
		if err != nil {
			timer.Reset(c.retryDuration)
		}

		select {
		case <-timer.C:
		case <-c.safeCloser.ChanClose():
			return nil
		}
	}
}

func (c *client) openConnection() error {
	timer := time.NewTimer(c.retryDuration)
	defer timer.Stop()

	for {
		err := c.wsConn.OpenConnection(c.url)
		if err == nil {
			return nil
		} else {
			c.log.Warn(fmt.Sprintf("c.openConnection(), retrying in %v...", c.retryDuration), "error", err)
		}

		timer.Reset(c.retryDuration)

		select {
		case <-timer.C:
		case <-c.safeCloser.ChanClose():
			return nil
		}
	}
}

func (c *client) Close() error {
	defer c.safeCloser.Close()

	c.log.Info("closing all components...")

	return c.wsConn.Close()
}
