package testscommon

import (
	"github.com/multiversx/mx-chain-core-go/websocketOutportDriver/data"
)

// PayloadProcessorStub -
type PayloadProcessorStub struct {
	ProcessPayloadCalled func(payload *data.PayloadData) error
	CloseCalled          func() error
}

// ProcessPayload -
func (pps *PayloadProcessorStub) ProcessPayload(payload *data.PayloadData) error {
	if pps.ProcessPayloadCalled != nil {
		return pps.ProcessPayloadCalled(payload)
	}

	return nil
}

// Close -
func (pps *PayloadProcessorStub) Close() error {
	if pps.CloseCalled != nil {
		return pps.CloseCalled()
	}

	return nil
}

// IsInterfaceNil -
func (pps *PayloadProcessorStub) IsInterfaceNil() bool {
	return pps == nil
}
