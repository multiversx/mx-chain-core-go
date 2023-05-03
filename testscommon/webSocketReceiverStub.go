package testscommon

import (
	"github.com/multiversx/mx-chain-core-go/webSocket"
	"github.com/multiversx/mx-chain-core-go/webSocket/data"
)

// WebSocketReceiverStub -
type WebSocketReceiverStub struct {
}

// Send -
func (w WebSocketReceiverStub) Send(_ data.WsSendArgs, _ webSocket.WSConClient) error {
	return nil
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
