package mock

import "github.com/multiversx/mx-chain-core-go/websocketOutportDriver"

type PayloadProcessorStub struct {
	ProcessPayloadCalled func(payload *websocketOutportDriver.PayloadData) error
	CloseCalled          func() error
}

func (pps *PayloadProcessorStub) ProcessPayload(payload *websocketOutportDriver.PayloadData) error {
	if pps.ProcessPayloadCalled != nil {
		return pps.ProcessPayloadCalled(payload)
	}

	return nil
}

func (pps *PayloadProcessorStub) Close() error {
	if pps.CloseCalled != nil {
		return pps.CloseCalled()
	}

	return nil
}
