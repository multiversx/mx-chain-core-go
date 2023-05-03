package server

import (
	"github.com/multiversx/mx-chain-core-go/webSocket"
	"github.com/multiversx/mx-chain-core-go/webSocket/data"
)

// TransceiversAndConnHolder defines what a transceivers holder should be able to do
type TransceiversAndConnHolder interface {
	AddTransceiverAndConn(transceiver Transceiver, conn webSocket.WSConClient)
	Remove(id string)
	GetAll() map[string]TupleTransceiverAndConn
}

// Transceiver defines what a WebSocket transceiver should be able to do
type Transceiver interface {
	Send(args data.WsSendArgs, connection webSocket.WSConClient) error
	SetPayloadHandler(handler webSocket.PayloadHandler) error
	Listen(connection webSocket.WSConClient) (closed bool)
	Close() error
}
