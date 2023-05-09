package testscommon

import "github.com/multiversx/mx-chain-core-go/webSocket/data"

// PayloadHandlerStub -
type PayloadHandlerStub struct {
	ProcessPayloadCalled func(payload []byte, payloadType data.PayloadType) error
	CloseCalled          func() error
}

// IsInterfaceNil -
func (ph *PayloadHandlerStub) IsInterfaceNil() bool {
	return ph == nil
}

// ProcessPayload -
func (ph *PayloadHandlerStub) ProcessPayload(payload []byte, payloadType data.PayloadType) error {
	if ph.ProcessPayloadCalled != nil {
		return ph.ProcessPayloadCalled(payload, payloadType)
	}
	return nil
}

// Close -
func (ph *PayloadHandlerStub) Close() error {
	if ph.CloseCalled != nil {
		return ph.CloseCalled()
	}
	return nil
}
