package webSockets

type nilPayloadHandler struct{}

// NewNilPayloadHandler will create a new instance of nilPayloadHandler
func NewNilPayloadHandler() PayloadHandler {
	return new(nilPayloadHandler)
}

// ProcessPayload will do nothing
func (n nilPayloadHandler) ProcessPayload(_ []byte) error {
	return nil
}

// Close will do nothing
func (n nilPayloadHandler) Close() error {
	return nil
}
