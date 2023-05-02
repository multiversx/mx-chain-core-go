package client

import (
	"github.com/multiversx/mx-chain-core-go/webSockets"
)

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
