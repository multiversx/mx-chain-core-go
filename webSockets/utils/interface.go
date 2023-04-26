package utils

import (
	"github.com/multiversx/mx-chain-core-go/webSockets/connection"
	"github.com/multiversx/mx-chain-core-go/webSockets/data"
)

// ConnectionsHandler defines what a clients handler should be able to do
type ConnectionsHandler interface {
	AddClient(client connection.WSConClient) error
	GetAll() map[string]connection.WSConClient
	CloseAndRemove(remoteAddr string) error
}

type PayloadHandler interface {
	HandlePayload(payload []byte) (*data.PayloadData, error)
	Close() error
}

type Sender interface {
	AddConnection(client connection.WSConClient) error
	Send(payload []byte) error
	Close() error
}

type Receiver interface {
	Close() error
	SetPayloadHandler(handler PayloadHandler)
	Listen(connection connection.WSConClient) (closed bool)
}
