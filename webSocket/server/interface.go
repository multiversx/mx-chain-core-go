package server

import (
	"github.com/multiversx/mx-chain-core-go/webSocket"
	"github.com/multiversx/mx-chain-core-go/webSocket/data"
)

// ReceiversHolder defines what a receivers holder should be able to do
type ReceiversHolder interface {
	AddReceiver(id string, rec Receiver)
	RemoveReceiver(id string)
	GetAll() map[string]Receiver
}

// Receiver defines what a WebSocket receiver should be able to do
type Receiver interface {
	Close() error
	SetPayloadHandler(handler webSocket.PayloadHandler) error
	Listen(connection webSocket.WSConClient) (closed bool)
}

// Sender defines what a WebSocket sender should be able to do
type Sender interface {
	AddConnection(client webSocket.WSConClient) error
	Send(args data.WsSendArgs) error
	Close() error
}
