package client

import (
	"errors"
	"fmt"
	"net/url"
	"time"

	"github.com/multiversx/mx-chain-core-go/core"
	"github.com/multiversx/mx-chain-core-go/core/check"
	"github.com/multiversx/mx-chain-core-go/core/closing"
	"github.com/multiversx/mx-chain-core-go/webSocket"
	"github.com/multiversx/mx-chain-core-go/webSocket/connection"
	"github.com/multiversx/mx-chain-core-go/webSocket/data"
	"github.com/multiversx/mx-chain-core-go/webSocket/transceiver"
)

// ArgsWebSocketClient holds the arguments needed for creating a client
type ArgsWebSocketClient struct {
	RetryDurationInSeconds int
	WithAcknowledge        bool
	BlockingAckOnError     bool
	URL                    string
	PayloadConverter       webSocket.PayloadConverter
	Log                    core.Logger
}

type client struct {
	url           string
	retryDuration time.Duration
	safeCloser    core.SafeCloser
	log           core.Logger
	wsConn        webSocket.WSConClient
	transceiver   Receiver
}

// NewWebSocketClient will create a new instance of WebSocket client
func NewWebSocketClient(args ArgsWebSocketClient) (*client, error) {
	err := checkArgs(args)
	if err != nil {
		return nil, err
	}

	argsTransceiver := transceiver.ArgsTransceiver{
		PayloadConverter:   args.PayloadConverter,
		Log:                args.Log,
		RetryDurationInSec: args.RetryDurationInSeconds,
		BlockingAckOnError: args.BlockingAckOnError,
	}
	wsTransceiver, err := transceiver.NewReceiver(argsTransceiver)
	if err != nil {
		return nil, err
	}

	wsUrl := url.URL{Scheme: "ws", Host: args.URL, Path: data.WSRoute}
	return &client{
		url:           wsUrl.String(),
		wsConn:        connection.NewWSConnClient(),
		retryDuration: time.Duration(args.RetryDurationInSeconds) * time.Second,
		safeCloser:    closing.NewSafeChanCloser(),
		transceiver:   wsTransceiver,
		log:           args.Log,
	}, nil
}

func checkArgs(args ArgsWebSocketClient) error {
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

// Start will start the WebSocket client (will initialize a connection with the ws server) and listen for messages
func (c *client) Start() {
	go func() {
		timer := time.NewTimer(c.retryDuration)
		defer timer.Stop()

		for {
			err := c.wsConn.OpenConnection(c.url)
			if err != nil && !errors.Is(err, data.ErrConnectionAlreadyOpen) {
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

	go func() {
		for {
			_ = c.transceiver.Listen(c.wsConn)

			select {
			default:
			case <-c.safeCloser.ChanClose():
				return
			}
		}
	}()
}

// Send will send the provided payload from args
func (c *client) Send(args data.WsSendArgs) error {
	return c.transceiver.Send(args, c.wsConn)
}

// SetPayloadHandler set the payload handler
func (c *client) SetPayloadHandler(handler webSocket.PayloadHandler) error {
	return c.transceiver.SetPayloadHandler(handler)
}

// Close will close the component
func (c *client) Close() error {
	defer c.safeCloser.Close()

	c.log.Info("closing all components...")
	err := c.transceiver.Close()
	if err != nil {
		c.log.Warn("client.Close() sender", "error", err)
	}

	return nil
}

// IsInterfaceNil returns true if there is no value under the interface
func (c *client) IsInterfaceNil() bool {
	return c == nil
}
