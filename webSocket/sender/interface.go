package sender

import (
	"github.com/multiversx/mx-chain-core-go/webSocket"
)

// ConnectionsHandler defines what a clients handler should be able to do
type ConnectionsHandler interface {
	AddClient(client webSocket.WSConClient) error
	GetAll() map[string]webSocket.WSConClient
	CloseAndRemove(remoteAddr string) error
}
