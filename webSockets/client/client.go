package client

import (
	"errors"
	"fmt"
	"net/url"
	"time"

	"github.com/multiversx/mx-chain-core-go/core"
	"github.com/multiversx/mx-chain-core-go/core/check"
	"github.com/multiversx/mx-chain-core-go/core/closing"
	"github.com/multiversx/mx-chain-core-go/webSockets"
	"github.com/multiversx/mx-chain-core-go/webSockets/connection"
	"github.com/multiversx/mx-chain-core-go/webSockets/data"
	"github.com/multiversx/mx-chain-core-go/webSockets/receiver"
	"github.com/multiversx/mx-chain-core-go/webSockets/sender"
)

// ArgsWebSocketsClient holds the arguments needed for creating a client
type ArgsWebSocketsClient struct {
	RetryDurationInSeconds int
	WithAcknowledge        bool
	BlockingAckOnError     bool
	URL                    string
	PayloadConverter       webSockets.PayloadConverter
	Log                    core.Logger
}

type client struct {
	url           string
	retryDuration time.Duration
	safeCloser    core.SafeCloser
	log           core.Logger
	wsConn        webSockets.WSConClient
	sender        Sender
	receiver      Receiver
}

// NewWebSocketsClient will create a new instance of websockets client
func NewWebSocketsClient(args ArgsWebSocketsClient) (*client, error) {
	err := checkArgs(args)
	if err != nil {
		return nil, err
	}

	argsSender := sender.ArgsSender{
		WithAcknowledge:        args.WithAcknowledge,
		RetryDurationInSeconds: args.RetryDurationInSeconds,
		PayloadConverter:       args.PayloadConverter,
		Log:                    args.Log,
	}
	webSocketsSender, err := sender.NewSender(argsSender)
	if err != nil {
		return nil, err
	}

	argsReceiver := receiver.ArgsReceiver{
		PayloadConverter:   args.PayloadConverter,
		Log:                args.Log,
		RetryDurationInSec: args.RetryDurationInSeconds,
		BlockingAckOnError: args.BlockingAckOnError,
	}
	webSocketsReceiver, err := receiver.NewReceiver(argsReceiver)
	if err != nil {
		return nil, err
	}

	conn := connection.NewWSConnClient()
	err = webSocketsSender.AddConnection(conn)
	if err != nil {
		return nil, err
	}

	wsUrl := url.URL{Scheme: "ws", Host: args.URL, Path: data.WSRoute}
	return &client{
		url:           wsUrl.String(),
		sender:        webSocketsSender,
		wsConn:        conn,
		retryDuration: time.Duration(args.RetryDurationInSeconds) * time.Second,
		safeCloser:    closing.NewSafeChanCloser(),
		receiver:      webSocketsReceiver,
		log:           args.Log,
	}, nil
}

func checkArgs(args ArgsWebSocketsClient) error {
	if check.IfNil(args.Log) {
		return core.ErrNilLogger
	}
	if check.IfNil(args.PayloadConverter) {
		return data.ErrNilPayloadConverter
	}
	if args.URL == "" {
		return data.ErrEmptyUrl
	}
	if args.RetryDurationInSeconds == 0 {
		return data.ErrZeroValueRetryDuration
	}
	return nil
}

// Start will start the websockets client (will initialize a connection with the ws server)
func (c *client) Start() {
	go func() {
		timer := time.NewTimer(c.retryDuration)
		defer timer.Stop()

		for {
			err := c.wsConn.OpenConnection(c.url)
			if err != nil && !errors.Is(err, data.ErrConnectionAlreadyOpened) {
				c.log.Warn(fmt.Sprintf("c.openConnection(), retrying in %v...", c.retryDuration), "error", err)
			}

			timer.Reset(c.retryDuration)

			select {
			case <-timer.C:
			case <-c.safeCloser.ChanClose():
				return
			}
		}
	}()
}

// Send will send the provided payload from args
func (c *client) Send(args data.WsSendArgs) error {
	return c.sender.Send(args)
}

// SetPayloadHandler set the payload handler
func (c *client) SetPayloadHandler(handler webSockets.PayloadHandler) error {
	return c.receiver.SetPayloadHandler(handler)
}

// Listen will listen from messages
func (c *client) Listen() {
	for {
		_ = c.receiver.Listen(c.wsConn)

		select {
		default:
		case <-c.safeCloser.ChanClose():
			return
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
