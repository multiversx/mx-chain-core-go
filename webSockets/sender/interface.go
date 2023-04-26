package sender

import "github.com/multiversx/mx-chain-core-go/webSockets/connection"

// ConnectionsHandler defines what a clients handler should be able to do
type ConnectionsHandler interface {
	AddClient(client connection.WSConClient) error
	GetAll() map[string]connection.WSConClient
	CloseAndRemove(remoteAddr string) error
}
