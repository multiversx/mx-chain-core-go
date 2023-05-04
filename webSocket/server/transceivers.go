package server

import (
	"sync"

	"github.com/multiversx/mx-chain-core-go/webSocket"
)

type tupleTransceiverAndConn struct {
	transceiver Transceiver
	conn        webSocket.WSConClient
}

type transceiversAndConnHolder struct {
	mutex              sync.RWMutex
	transceiverAndConn map[string]tupleTransceiverAndConn
}

// NewTransceiversAndConnHolder will create a new instance of transceiversHolder
func newTransceiversAndConnHolder() *transceiversAndConnHolder {
	return &transceiversAndConnHolder{
		transceiverAndConn: map[string]tupleTransceiverAndConn{},
	}
}

// addTransceiverAndConn will add the provided transceiver in the internal map
func (th *transceiversAndConnHolder) addTransceiverAndConn(transceiver Transceiver, conn webSocket.WSConClient) {
	th.mutex.Lock()
	defer th.mutex.Unlock()

	th.transceiverAndConn[conn.GetID()] = tupleTransceiverAndConn{
		transceiver: transceiver,
		conn:        conn,
	}
}

// remove will remove the provided transceiver from the internal map
func (th *transceiversAndConnHolder) remove(id string) {
	th.mutex.Lock()
	defer th.mutex.Unlock()

	delete(th.transceiverAndConn, id)
}

// getAll will return a map with all the stored transceivers
func (th *transceiversAndConnHolder) getAll() map[string]tupleTransceiverAndConn {
	th.mutex.RLock()
	defer th.mutex.RUnlock()

	transceiversAndConn := make(map[string]tupleTransceiverAndConn)
	for id, tuple := range th.transceiverAndConn {
		transceiversAndConn[id] = tuple
	}

	return transceiversAndConn
}
