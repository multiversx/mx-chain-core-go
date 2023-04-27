package testscommon

import (
	"github.com/multiversx/mx-chain-core-go/webSockets/data"
)

// PayloadParserStub -
type PayloadParserStub struct {
	ExtractPayloadDataCalled func(payload []byte) (*data.PayloadData, error)
}

// ExtractPayloadData -
func (pps *PayloadParserStub) ExtractPayloadData(payload []byte) (*data.PayloadData, error) {
	if pps.ExtractPayloadDataCalled != nil {
		return pps.ExtractPayloadDataCalled(payload)
	}

	return &data.PayloadData{}, nil
}

// IsInterfaceNil -
func (pps *PayloadParserStub) IsInterfaceNil() bool {
	return pps == nil
}