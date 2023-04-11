package client

import (
	"fmt"
	"net/url"
	"time"

	"github.com/multiversx/mx-chain-core-go/core"
	"github.com/multiversx/mx-chain-core-go/core/check"
	"github.com/multiversx/mx-chain-core-go/core/closing"
	"github.com/multiversx/mx-chain-core-go/websocketOutportDriver/clientServerReceiver/payload"
	"github.com/multiversx/mx-chain-core-go/websocketOutportDriver/common"
	"github.com/multiversx/mx-chain-core-go/websocketOutportDriver/data"
)

const closedConnection = "use of closed network connection"

type client struct {
	url                string
	blockingAckOnError bool
	retryDuration      time.Duration
	wsConn             WSConnClient
	safeCloser         core.SafeCloser
	messagesListener   common.MessagesListener
	log                core.Logger
}

// ArgsWsClient holds the arguments required to create a new websocket client handler
type ArgsWsClient struct {
	Url                      string
	RetryDurationInSec       uint32
	BlockingAckOnError       bool
	Log                      core.Logger
	PayloadProcessor         common.PayloadProcessor
	PayloadParser            common.PayloadParser
	Uint64ByteSliceConverter common.Uint64ByteSliceConverter
	WSConnClient             WSConnClient
}

// NewWsClientHandler will create a ws client to receive data from an observer/light client
func NewWsClientHandler(args ArgsWsClient) (*client, error) {
	err := checkArgs(args)
	if err != nil {
		return nil, err
	}

	urlReceiveData := url.URL{Scheme: "ws", Host: args.Url, Path: data.WSRoute}

	messageListener, err := payload.NewMessagesListener(payload.ArgsMessagesProcessor{
		Log:                      args.Log,
		PayloadParser:            args.PayloadParser,
		PayloadProcessor:         args.PayloadProcessor,
		WsClient:                 args.WSConnClient,
		Uint64ByteSliceConverter: args.Uint64ByteSliceConverter,
		RetryDurationInSec:       args.RetryDurationInSec,
		BlockingAckOnError:       args.BlockingAckOnError,
	})
	if err != nil {
		return nil, err
	}

	return &client{
		url:              urlReceiveData.String(),
		wsConn:           args.WSConnClient,
		messagesListener: messageListener,
		log:              args.Log,
		safeCloser:       closing.NewSafeChanCloser(),
		retryDuration:    time.Duration(args.RetryDurationInSec) * time.Second,
	}, nil
}

func checkArgs(args ArgsWsClient) error {
	if check.IfNil(args.PayloadProcessor) {
		return data.ErrNilPayloadProcessor
	}
	if check.IfNil(args.PayloadParser) {
		return data.ErrNilPayloadParser
	}
	if check.IfNil(args.WSConnClient) {
		return data.ErrNilWsConnReceiver
	}
	if check.IfNil(args.Uint64ByteSliceConverter) {
		return data.ErrNilUint64ByteSliceConverter
	}
	if len(args.Url) == 0 {
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

// Start will start the client listening ws process
func (c *client) Start() {
	c.log.Info("connecting to", "url", c.url)

	timer := time.NewTimer(c.retryDuration)
	defer timer.Stop()

	closed := false
	for !closed {
		err := c.wsConn.OpenConnection(c.url)
		if err == nil {
			closed = c.messagesListener.Listen()
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

// Close will close the underlying ws connection
func (c *client) Close() {
	defer c.safeCloser.Close()

	c.messagesListener.Close()
}
