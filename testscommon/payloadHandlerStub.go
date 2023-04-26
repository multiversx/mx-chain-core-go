package testscommon

import outportSenderData "github.com/multiversx/mx-chain-core-go/webSockets/data"

// PayloadHandlerStub -
type PayloadHandlerStub struct {
	HandlePayloadCalled func(payload []byte) (*outportSenderData.PayloadData, error)
	CloseCalled         func() error
}

// HandlePayload -
func (ph *PayloadHandlerStub) HandlePayload(payload []byte) (*outportSenderData.PayloadData, error) {
	if ph.HandlePayloadCalled != nil {
		return ph.HandlePayloadCalled(payload)
	}
	return nil, nil
}

// Close -
func (ph *PayloadHandlerStub) Close() error {
	if ph.CloseCalled != nil {
		return ph.CloseCalled()
	}
	return nil
}
