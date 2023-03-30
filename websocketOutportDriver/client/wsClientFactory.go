package client

import (
	"github.com/multiversx/mx-chain-core-go/data/typeConverters/uint64ByteSlice"
	"github.com/multiversx/mx-chain-core-go/websocketOutportDriver"
)

// ArgsCreateWsClient is a placeholder struct for args required to create a WSClient
type ArgsCreateWsClient struct {
	Url                string
	RetryDurationInSec uint32
	BlockingAckOnError bool
	PayloadProcessor   PayloadProcessor
}

// CreateWsClient creates a WSClient
func CreateWsClient(args ArgsCreateWsClient) (WSClient, error) {
	uint64ByteSliceConverter := uint64ByteSlice.NewBigEndianConverter()
	payloadParser, err := websocketOutportDriver.NewWebSocketPayloadParser(uint64ByteSliceConverter)
	if err != nil {
		return nil, err
	}

	argsWsClient := ArgsWsClient{
		Url:                      args.Url,
		RetryDurationInSec:       args.RetryDurationInSec,
		BlockingAckOnError:       args.BlockingAckOnError,
		PayloadProcessor:         args.PayloadProcessor,
		PayloadParser:            payloadParser,
		Uint64ByteSliceConverter: uint64ByteSliceConverter,
		WSConnClient:             NewWSConnClient(),
	}

	return NewWsClientHandler(argsWsClient)
}
