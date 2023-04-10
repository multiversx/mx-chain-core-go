package clientServerReceiver

import (
	"github.com/multiversx/mx-chain-core-go/core"
	"github.com/multiversx/mx-chain-core-go/websocketOutportDriver/clientServerReceiver/client"
	"github.com/multiversx/mx-chain-core-go/websocketOutportDriver/clientServerReceiver/server"
	"github.com/multiversx/mx-chain-core-go/websocketOutportDriver/common"
)

type ArgsWsClientServerReceiver struct {
	IsServer           bool
	Url                string
	RetryDurationInSec uint32
	BlockingAckOnError bool
	PayloadProcessor   common.PayloadProcessor
	Log                core.Logger
}

type receiver struct {
	receiver WsMessagesReceiver
}

func NewClientServerReceiver(args ArgsWsClientServerReceiver) (*receiver, error) {
	wsReceiver, err := createWsMessageReceiver(args)
	if err != nil {
		return nil, err
	}

	return &receiver{
		receiver: wsReceiver,
	}, nil
}

func (r *receiver) Start() {
	r.receiver.Start()
}

func (r *receiver) Close() {
	r.receiver.Close()
}

func createWsMessageReceiver(args ArgsWsClientServerReceiver) (WsMessagesReceiver, error) {
	if args.IsServer {
		return server.NewWsServer(server.ArgsWsServer{
			URL:                args.Url,
			RetryDurationInSec: args.RetryDurationInSec,
			BlockingAckOnError: args.BlockingAckOnError,
			PayloadProcessor:   args.PayloadProcessor,
			Log:                args.Log,
		})
	}

	return client.CreateWsClient(client.ArgsCreateWsClient{
		Url:                args.Url,
		RetryDurationInSec: args.RetryDurationInSec,
		BlockingAckOnError: args.BlockingAckOnError,
		PayloadProcessor:   args.PayloadProcessor,
		Log:                args.Log,
	})
}
