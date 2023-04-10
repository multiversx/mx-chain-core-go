package client

import (
	"github.com/multiversx/mx-chain-core-go/core"
	"github.com/multiversx/mx-chain-core-go/data/typeConverters/uint64ByteSlice"
	"github.com/multiversx/mx-chain-core-go/websocketOutportDriver"
	"github.com/multiversx/mx-chain-core-go/websocketOutportDriver/common"
)

// ArgsCreateWsClient is a placeholder struct for args required to create a WSClient
type ArgsCreateWsClient struct {
	Url                string
	RetryDurationInSec uint32
	BlockingAckOnError bool
	PayloadProcessor   common.PayloadProcessor
	Log                core.Logger
}

// CreateWsClient creates a WSClient
func CreateWsClient(args ArgsCreateWsClient) (WSClient, error) {
	uint64ByteSliceConverter := uint64ByteSlice.NewBigEndianConverter()
	payloadParser, err := websocketOutportDriver.NewWebSocketPayloadParser(uint64ByteSliceConverter)
	if err != nil {
		return nil, err
	}

	argsWsClient := ArgsWsClient{
		Log:                      args.Log,
		Url:                      args.Url,
		RetryDurationInSec:       args.RetryDurationInSec,
		BlockingAckOnError:       args.BlockingAckOnError,
		PayloadProcessor:         args.PayloadProcessor,
		PayloadParser:            payloadParser,
		Uint64ByteSliceConverter: uint64ByteSliceConverter,
		WSConnClient:             common.NewWSConnClient(),
	}

	return NewWsClientHandler(argsWsClient)
}
