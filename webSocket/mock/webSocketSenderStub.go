package mock

import (
	"github.com/multiversx/mx-chain-core-go/webSocket/data"
)

// WebSocketSenderStub -
type WebSocketSenderStub struct {
	SendOnRouteCalled func(payload []byte, topic string) error
	AddClientCalled   func(wss data.WSConn, remoteAddr string)
	CloseCalled       func() error
}

// AddClient -
func (w *WebSocketSenderStub) AddClient(wss data.WSConn, remoteAddr string) {
	if w.AddClientCalled != nil {
		w.AddClientCalled(wss, remoteAddr)
	}
}

// Send -
func (w *WebSocketSenderStub) Send(payload []byte, topic string) error {
	if w.SendOnRouteCalled != nil {
		return w.SendOnRouteCalled(payload, topic)
	}

	return nil
}

// Close -
func (w *WebSocketSenderStub) Close() error {
	if w.CloseCalled != nil {
		return w.CloseCalled()
	}

	return nil
}

// IsInterfaceNil -
func (w *WebSocketSenderStub) IsInterfaceNil() bool {
	return w == nil
}
