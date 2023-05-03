package testscommon

import (
	"github.com/multiversx/mx-chain-core-go/webSocket"
)

// WebSocketReceiverStub -
type WebSocketReceiverStub struct {
}

// Close -
func (w WebSocketReceiverStub) Close() error {
	return nil
}

// SetPayloadHandler -
func (w WebSocketReceiverStub) SetPayloadHandler(_ webSocket.PayloadHandler) error {
	return nil
}

// Listen -
func (w WebSocketReceiverStub) Listen(_ webSocket.WSConClient) (closed bool) {
	return false
}
