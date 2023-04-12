package client

import (
	"fmt"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
	"github.com/multiversx/mx-chain-core-go/core"
	"github.com/multiversx/mx-chain-core-go/core/check"
	"github.com/multiversx/mx-chain-core-go/core/closing"
	"github.com/multiversx/mx-chain-core-go/webSockets/common"
	"github.com/multiversx/mx-chain-core-go/webSockets/data"
)

// ArgsWsClientSender holds the arguments needed for creating a new instance of client
type ArgsWsClientSender struct {
	Uint64ByteSliceConverter common.Uint64ByteSliceConverter
	Log                      core.Logger
	RetryDurationInSec       int
	WithAcknowledge          bool
	URL                      string
}

type clientSender struct {
	log                      core.Logger
	uint64ByteSliceConverter common.Uint64ByteSliceConverter
	wsConn                   common.WSConClient
	retryDuration            time.Duration
	safeCloser               core.SafeCloser
	url                      string
	withAcknowledge          bool
}

// NewClientSender will create a new instance of *client
func NewClientSender(args ArgsWsClientSender) (*clientSender, error) {
	err := checkArgs(args)
	if err != nil {
		return nil, err
	}

	urlSendData := url.URL{Scheme: "ws", Host: args.URL, Path: data.WSRoute}
	return &clientSender{
		uint64ByteSliceConverter: args.Uint64ByteSliceConverter,
		retryDuration:            time.Duration(args.RetryDurationInSec) * time.Second,
		url:                      urlSendData.String(),
		wsConn:                   common.NewWSConnClient(),
		withAcknowledge:          args.WithAcknowledge,
		log:                      args.Log,
		safeCloser:               closing.NewSafeChanCloser(),
	}, nil
}

// Send will send the provided payload
func (c *clientSender) Send(counter uint64, payload []byte) error {
	err := c.writeMessage(payload)
	if err != nil {
		c.log.Warn("cannot write message", "error", err)
	}

	if !c.withAcknowledge {
		return nil
	}

	return c.waitForAckSignal(counter)
}

func (c *clientSender) waitForAckSignal(counter uint64) error {
	timer := time.NewTimer(c.retryDuration)
	defer timer.Stop()

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

		timer.Reset(c.retryDuration)
		c.log.Warn("clientSender.waitForAckSignal: different counter",
			"expected", counter, "actual", receivedCounter)

		select {
		case <-timer.C:
		case <-c.safeCloser.ChanClose():
			return nil
		}
	}
	return nil
}

func (c *clientSender) writeMessage(payload []byte) error {
	timer := time.NewTimer(c.retryDuration)
	defer timer.Stop()

	for {
		err := c.wsConn.WriteMessage(websocket.BinaryMessage, payload)
		if err == nil {
			return nil
		}

		c.log.Warn("client.writeMessage: connection problem", "error", err)
		c.openConnection()

		select {
		case <-timer.C:
		case <-c.safeCloser.ChanClose():
			return nil
		}
	}
}

func (c *clientSender) openConnection() {
	timer := time.NewTimer(c.retryDuration)
	defer timer.Stop()

	for {
		err := c.wsConn.OpenConnection(c.url)
		if err == nil {
			return
		} else {
			c.log.Warn(fmt.Sprintf("c.openConnection(), retrying in %v...", c.retryDuration), "error", err)
		}

		timer.Reset(c.retryDuration)

		select {
		case <-timer.C:
		case <-c.safeCloser.ChanClose():
			return
		}
	}
}

// Close will close the web-sockets connection
func (c *clientSender) Close() error {
	defer c.safeCloser.Close()

	c.log.Info("closing all components...")
	err := c.wsConn.Close()
	if err != nil {
		c.log.Warn("clientSender.Close()", "error", err)
	}

	return nil
}

func checkArgs(args ArgsWsClientSender) error {
	if check.IfNil(args.Uint64ByteSliceConverter) {
		return data.ErrNilUint64ByteSliceConverter
	}
	if len(args.URL) == 0 {
		return data.ErrEmptyUrl
	}
	if args.RetryDurationInSec == 0 {
		return data.ErrZeroValueRetryDuration
	}
	if check.IfNil(args.Log) {
		return data.ErrNilLogger
	}

	return nil
}
