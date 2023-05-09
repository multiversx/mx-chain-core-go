package server

import (
	"github.com/multiversx/mx-chain-core-go/webSocket"
)

type transceiversAndConnHandler interface {
	addTransceiverAndConn(transceiver Transceiver, conn webSocket.WSConClient)
	remove(id string)
	getAll() map[string]tupleTransceiverAndConn
}

// Transceiver defines what a WebSocket transceiver should be able to do
type Transceiver interface {
	Send(payload []byte, topic string, connection webSocket.WSConClient) error
	SetPayloadHandler(handler webSocket.PayloadHandler) error
	Listen(connection webSocket.WSConClient) (closed bool)
	Close() error
}
