package server

import (
	"github.com/multiversx/mx-chain-core-go/webSockets"
	"github.com/multiversx/mx-chain-core-go/webSockets/connection"
)

// ReceiversHolder defines what a receivers holder should be able to do
type ReceiversHolder interface {
	AddReceiver(id string, rec Receiver)
	RemoveReceiver(id string)
	GetAll() map[string]Receiver
}

// Receiver defines what a web-sockets receiver should be able to do
type Receiver interface {
	Close() error
	SetPayloadHandler(handler webSockets.PayloadHandler)
	Listen(connection connection.WSConClient) (closed bool)
}

// Sender defines what a web-sockets sender should be able to do
type Sender interface {
	AddConnection(client connection.WSConClient) error
	Send(payload []byte) error
	Close() error
}
