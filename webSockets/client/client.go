package client

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/gorilla/websocket"
	"github.com/multiversx/mx-chain-core-go/core"
	"github.com/multiversx/mx-chain-core-go/core/check"
	"github.com/multiversx/mx-chain-core-go/core/closing"
	"github.com/multiversx/mx-chain-core-go/webSockets"
	"github.com/multiversx/mx-chain-core-go/webSockets/connection"
	"github.com/multiversx/mx-chain-core-go/webSockets/data"
	"github.com/multiversx/mx-chain-core-go/webSockets/receiver"
	"github.com/multiversx/mx-chain-core-go/webSockets/sender"
)

// ArgsWebSocketsClient  holds the arguments needed for creating a client
type ArgsWebSocketsClient struct {
	RetryDurationInSeconds   int
	WithAcknowledge          bool
	BlockingAckOnError       bool
	URL                      string
	Uint64ByteSliceConverter connection.Uint64ByteSliceConverter
	Log                      core.Logger
}

type client struct {
	url           string
	retryDuration time.Duration
	safeCloser    core.SafeCloser
	log           core.Logger
	wsConn        connection.WSConClient
	sender        Sender
	receiver      Receiver
}

// NewWebSocketsClient will create a new instance of websockets client
func NewWebSocketsClient(args ArgsWebSocketsClient) (*client, error) {
	if err := checkArgs(args); err != nil {
		return nil, err
	}

	webSocketsSender, err := sender.NewSender(sender.ArgsSender{
		WithAcknowledge:          args.WithAcknowledge,
		RetryDurationInSeconds:   args.RetryDurationInSeconds,
		Uint64ByteSliceConverter: args.Uint64ByteSliceConverter,
		Log:                      args.Log,
	})
	if err != nil {
		return nil, err
	}

	webSocketsReceiver, err := receiver.NewReceiver(receiver.ArgsReceiver{
		Uint64ByteSliceConverter: args.Uint64ByteSliceConverter,
		Log:                      args.Log,
		RetryDurationInSec:       args.RetryDurationInSeconds,
		BlockingAckOnError:       args.BlockingAckOnError,
	})
	if err != nil {
		return nil, err
	}

	return &client{
		url:           args.URL,
		sender:        webSocketsSender,
		wsConn:        connection.NewWSConnClient(),
		retryDuration: time.Duration(args.RetryDurationInSeconds) * time.Second,
		safeCloser:    closing.NewSafeChanCloser(),
		receiver:      webSocketsReceiver,
		log:           args.Log,
	}, nil
}

// Send will send the provided payload from args
func (c *client) Send(args data.WsSendArgs) error {
	closed := c.openConnection()
	if closed {
		return nil
	}

	_ = c.sender.AddConnection(c.wsConn)

	timer := time.NewTimer(c.retryDuration)
	defer timer.Stop()
	for {
		err := c.sender.Send(args.Payload)
		if err == nil {
			return nil
		}

		_, isConnectionClosed := err.(*websocket.CloseError)
		shouldOpenNewConnection := isConnectionClosed ||
			strings.Contains(err.Error(), data.ClosedConnectionMessage) ||
			strings.Contains(err.Error(), data.CloseSent)
		if shouldOpenNewConnection {
			// open a new connection
			c.log.Warn("clientSender: the previous connection was closed -> trying to open a new connection")
			c.wsConn = connection.NewWSConnClient()
			closed = c.openConnection()
			if closed {
				return nil
			}
			// the old connection will be rewritten in the internal map
			_ = c.sender.AddConnection(c.wsConn)
			continue
		}

		c.log.Warn(fmt.Sprintf("client.writeMessage: connection problem retrying in %v", c.retryDuration), "error", err)

		timer.Reset(c.retryDuration)

		select {
		case <-timer.C:
		case <-c.safeCloser.ChanClose():
			return nil
		}
	}
}

// RegisterPayloadHandler register the payload handler
func (c *client) RegisterPayloadHandler(handler webSockets.PayloadHandler) {
	c.receiver.SetPayloadHandler(handler)
}

// Listen will listen from messages
func (c *client) Listen() {
	timer := time.NewTimer(c.retryDuration)
	defer timer.Stop()

	closed := false
	for !closed {
		c.openConnection()

		closed = c.receiver.Listen(c.wsConn)

		timer.Reset(c.retryDuration)

		select {
		case <-timer.C:
		case <-c.safeCloser.ChanClose():
			return
		}
	}
}

func (c *client) openConnection() (closed bool) {
	timer := time.NewTimer(c.retryDuration)
	defer timer.Stop()

	for {
		err := c.wsConn.OpenConnection(c.url)
		if err == nil || errors.Is(err, data.ErrConnectionAlreadyOpened) {
			return
		} else {
			c.log.Warn(fmt.Sprintf("c.openConnection(), retrying in %v...", c.retryDuration), "error", err)
		}

		timer.Reset(c.retryDuration)

		select {
		case <-timer.C:
		case <-c.safeCloser.ChanClose():
			return true
		}
	}
}

// Close will close the component
func (c *client) Close() error {
	defer c.safeCloser.Close()

	c.log.Info("closing all components...")
	if c.sender != nil {
		err := c.sender.Close()
		if err != nil {
			c.log.Warn("client.Close() sender", "error", err)
		}
	}
	if c.receiver != nil {
		err := c.receiver.Close()
		if err != nil {
			c.log.Warn("client.Close() receiver", "error", err)
		}
	}

	return nil
}

// IsInterfaceNil returns true if there is no value under the interface
func (c *client) IsInterfaceNil() bool {
	return c == nil
}

func checkArgs(args ArgsWebSocketsClient) error {
	if check.IfNil(args.Log) {
		return core.ErrNilLogger
	}
	if check.IfNil(args.Uint64ByteSliceConverter) {
		return data.ErrNilUint64ByteSliceConverter
	}
	if args.URL == "" {
		return data.ErrEmptyUrl
	}
	if args.RetryDurationInSeconds == 0 {
		return data.ErrZeroValueRetryDuration
	}
	return nil
}
