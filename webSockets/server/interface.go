package server

import (
	"github.com/multiversx/mx-chain-core-go/webSockets"
	"github.com/multiversx/mx-chain-core-go/webSockets/connection"
)

type ReceiversHolder interface {
	AddReceiver(id string, rec Receiver)
	RemoveReceiver(id string)
	GetAll() map[string]Receiver
}

type Receiver interface {
	Close() error
	SetPayloadHandler(handler webSockets.PayloadHandler)
	Listen(connection connection.WSConClient) (closed bool)
}

type Sender interface {
	AddConnection(client connection.WSConClient) error
	Send(payload []byte) error
	Close() error
}
