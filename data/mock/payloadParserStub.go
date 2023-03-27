package mock

import "github.com/multiversx/mx-chain-core-go/websocketOutportDriver"

type PayloadParserStub struct {
	ExtractPayloadDataCalled func(payload []byte) (*websocketOutportDriver.PayloadData, error)
}

func (pps *PayloadParserStub) ExtractPayloadData(payload []byte) (*websocketOutportDriver.PayloadData, error) {
	if pps.ExtractPayloadDataCalled != nil {
		return pps.ExtractPayloadData(payload)
	}

	return &websocketOutportDriver.PayloadData{}, nil
}
