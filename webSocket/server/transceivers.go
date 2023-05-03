package server

import (
	"sync"

	"github.com/multiversx/mx-chain-core-go/webSocket"
)

type TupleTransceiverAndConn struct {
	Transceiver Transceiver
	Conn        webSocket.WSConClient
}

type transceiversAndConnHolder struct {
	mutex              sync.RWMutex
	transceiverAndConn map[string]TupleTransceiverAndConn
}

// NewTransceiversAndConnHolder will create a new instance of transceiversHolder
func NewTransceiversAndConnHolder() *transceiversAndConnHolder {
	return &transceiversAndConnHolder{
		transceiverAndConn: map[string]TupleTransceiverAndConn{},
	}
}

// AddTransceiverAndConn will add the provided transceiver in the internal map
func (th *transceiversAndConnHolder) AddTransceiverAndConn(transceiver Transceiver, conn webSocket.WSConClient) {
	th.mutex.Lock()
	defer th.mutex.Unlock()

	th.transceiverAndConn[conn.GetID()] = TupleTransceiverAndConn{
		Transceiver: transceiver,
		Conn:        conn,
	}
}

// Remove will remove the provided transceiver from the internal map
func (th *transceiversAndConnHolder) Remove(id string) {
	th.mutex.Lock()
	defer th.mutex.Unlock()

	delete(th.transceiverAndConn, id)
}

// GetAll will return a map with all the stored transceivers
func (th *transceiversAndConnHolder) GetAll() map[string]TupleTransceiverAndConn {
	th.mutex.RLock()
	defer th.mutex.RUnlock()

	transceiversAndConn := make(map[string]TupleTransceiverAndConn)
	for id, tuple := range th.transceiverAndConn {
		transceiversAndConn[id] = tuple
	}

	return transceiversAndConn
}
