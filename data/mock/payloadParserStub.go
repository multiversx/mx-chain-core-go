package mock

import "github.com/multiversx/mx-chain-core-go/websocketOutportDriver"

// PayloadParserStub -
type PayloadParserStub struct {
	ExtractPayloadDataCalled func(payload []byte) (*websocketOutportDriver.PayloadData, error)
}

// ExtractPayloadData -
func (pps *PayloadParserStub) ExtractPayloadData(payload []byte) (*websocketOutportDriver.PayloadData, error) {
	if pps.ExtractPayloadDataCalled != nil {
		return pps.ExtractPayloadData(payload)
	}

	return &websocketOutportDriver.PayloadData{}, nil
}
