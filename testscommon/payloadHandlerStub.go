package testscommon

// PayloadHandlerStub -
type PayloadHandlerStub struct {
	ProcessPayloadCalled func(payload []byte) error
	CloseCalled          func() error
}

// ProcessPayload -
func (ph *PayloadHandlerStub) ProcessPayload(payload []byte) error {
	if ph.ProcessPayloadCalled != nil {
		return ph.ProcessPayloadCalled(payload)
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
