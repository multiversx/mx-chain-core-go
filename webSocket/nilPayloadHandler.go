package webSocket

type nilPayloadHandler struct{}

// NewNilPayloadHandler will create a new instance of nilPayloadHandler
func NewNilPayloadHandler() PayloadHandler {
	return new(nilPayloadHandler)
}

// ProcessPayload will do nothing
func (n nilPayloadHandler) ProcessPayload(_ []byte, _ string) error {
	return nil
}

// Close will do nothing
func (n nilPayloadHandler) Close() error {
	return nil
}

// IsInterfaceNil returns true if there is no value under the interface
func (n nilPayloadHandler) IsInterfaceNil() bool {
	return false
}
