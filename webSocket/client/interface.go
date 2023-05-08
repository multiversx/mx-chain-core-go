package client

import (
	"github.com/multiversx/mx-chain-core-go/webSocket"
	"github.com/multiversx/mx-chain-core-go/webSocket/data"
)

// Receiver defines what a WebSocket transceiver should be able to do
type Receiver interface {
	Send(args data.PayloadData, connection webSocket.WSConClient) error
	SetPayloadHandler(handler webSocket.PayloadHandler) error
	Listen(connection webSocket.WSConClient) (closed bool)
	Close() error
}
