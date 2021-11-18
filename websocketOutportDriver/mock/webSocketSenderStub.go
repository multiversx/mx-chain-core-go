package mock

import (
	"github.com/ElrondNetwork/elrond-go-core/websocketOutportDriver/data"
	"github.com/ElrondNetwork/elrond-go-core/websocketOutportDriver/sender"
)

// WebSocketSenderStub -
type WebSocketSenderStub struct {
	SendOnRouteCalled func(args data.WsSendArgs) error
	AddClientCalled   func(wss sender.WSConn, remoteAddr string)
	CloseCalled       func() error
}

// AddClient -
func (w *WebSocketSenderStub) AddClient(wss sender.WSConn, remoteAddr string) {
	if w.AddClientCalled != nil {
		w.AddClientCalled(wss, remoteAddr)
	}
}

// Send -
func (w *WebSocketSenderStub) Send(args data.WsSendArgs) error {
	if w.SendOnRouteCalled != nil {
		return w.SendOnRouteCalled(args)
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
