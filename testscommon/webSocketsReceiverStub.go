package testscommon

import (
	"github.com/multiversx/mx-chain-core-go/webSockets"
)

// WebSocketsReceiverStub -
type WebSocketsReceiverStub struct {
}

// Close -
func (w WebSocketsReceiverStub) Close() error {
	return nil
}

// SetPayloadHandler -
func (w WebSocketsReceiverStub) SetPayloadHandler(_ webSockets.PayloadHandler) error {
	return nil
}

// Listen -
func (w WebSocketsReceiverStub) Listen(_ webSockets.WSConClient) (closed bool) {
	return false
}
