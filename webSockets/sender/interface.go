package sender

import (
	"github.com/multiversx/mx-chain-core-go/webSockets"
)

// ConnectionsHandler defines what a clients handler should be able to do
type ConnectionsHandler interface {
	AddClient(client webSockets.WSConClient) error
	GetAll() map[string]webSockets.WSConClient
	CloseAndRemove(remoteAddr string) error
}
