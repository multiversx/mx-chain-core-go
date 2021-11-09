package mock

import (
	"github.com/ElrondNetwork/elrond-go-core/websocketOutportDriver/data"
)

// WebSocketSenderStub -
type WebSocketSenderStub struct {
	SendOnRouteCalled func(args data.WsSendArgs) error
}

// SendOnRoute -
func (w *WebSocketSenderStub) SendOnRoute(args data.WsSendArgs) error {
	if w.SendOnRouteCalled != nil {
		return w.SendOnRouteCalled(args)
	}

	return nil
}

// IsInterfaceNil -
func (w *WebSocketSenderStub) IsInterfaceNil() bool {
	return w == nil
}
