package clientServerReceiver

import (
	"github.com/multiversx/mx-chain-core-go/core"
	"github.com/multiversx/mx-chain-core-go/data/typeConverters/uint64ByteSlice"
	"github.com/multiversx/mx-chain-core-go/webSockets"
	"github.com/multiversx/mx-chain-core-go/webSockets/clientServerReceiver/client"
	"github.com/multiversx/mx-chain-core-go/webSockets/clientServerReceiver/server"
	"github.com/multiversx/mx-chain-core-go/webSockets/common"
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

// NewClientServerReceiver will create a new instance of receiver
func NewClientServerReceiver(args ArgsWsClientServerReceiver) (*receiver, error) {
	wsReceiver, err := createWsMessageReceiver(args)
	if err != nil {
		return nil, err
	}

	return &receiver{
		receiver: wsReceiver,
	}, nil
}

// Start will start the web-sockets receiver
func (r *receiver) Start() {
	r.receiver.Start()
}

// Close will close the web-sockets receiver
func (r *receiver) Close() {
	r.receiver.Close()
}

func createWsMessageReceiver(args ArgsWsClientServerReceiver) (WsMessagesReceiver, error) {
	uint64ByteSliceConverter := uint64ByteSlice.NewBigEndianConverter()
	payloadParser, err := webSockets.NewWebSocketPayloadParser(uint64ByteSliceConverter)
	if err != nil {
		return nil, err
	}

	if args.IsServer {
		return server.NewWsServer(server.ArgsWsServer{
			URL:                      args.Url,
			RetryDurationInSec:       args.RetryDurationInSec,
			BlockingAckOnError:       args.BlockingAckOnError,
			PayloadProcessor:         args.PayloadProcessor,
			Log:                      args.Log,
			PayloadParser:            payloadParser,
			Uint64ByteSliceConverter: uint64ByteSliceConverter,
		})
	}

	return client.NewWsClientHandler(client.ArgsWsClient{
		Url:                      args.Url,
		RetryDurationInSec:       args.RetryDurationInSec,
		BlockingAckOnError:       args.BlockingAckOnError,
		Log:                      args.Log,
		PayloadProcessor:         args.PayloadProcessor,
		PayloadParser:            payloadParser,
		Uint64ByteSliceConverter: uint64ByteSliceConverter,
		WSConnClient:             common.NewWSConnClient(),
	})
}
