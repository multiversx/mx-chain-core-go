package client

import (
	"github.com/multiversx/mx-chain-core-go/webSockets"
	"github.com/multiversx/mx-chain-core-go/webSockets/connection"
)

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
