package webSockets

import "github.com/multiversx/mx-chain-core-go/webSockets/data"

type nilPayloadHandler struct{}

// NewNilPayloadHandler will create a new instance of nilPayloadHandler
func NewNilPayloadHandler() PayloadHandler {
	return new(nilPayloadHandler)
}

// HandlePayload will do nothing
func (n nilPayloadHandler) HandlePayload(_ []byte) (*data.PayloadData, error) {
	return nil, nil
}

// Close will do nothing
func (n nilPayloadHandler) Close() error {
	return nil
}
