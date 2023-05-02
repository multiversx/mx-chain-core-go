package server

import (
	"github.com/multiversx/mx-chain-core-go/webSockets"
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
	SetPayloadHandler(handler webSockets.PayloadHandler) error
	Listen(connection webSockets.WSConClient) (closed bool)
}

// Sender defines what a web-sockets sender should be able to do
type Sender interface {
	AddConnection(client webSockets.WSConClient) error
	Send(payload []byte) error
	Close() error
}
