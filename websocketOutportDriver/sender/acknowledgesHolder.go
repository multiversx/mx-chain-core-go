package sender

import "sync"

type acknowledgesHolder struct {
	acknowledges map[string]*websocketClientAcknowledgesHolder
	mut          sync.Mutex
}

// NewAcknowledgesHolder returns a new instance of acknowledgesHolder
func NewAcknowledgesHolder() *acknowledgesHolder {
	return &acknowledgesHolder{
		acknowledges: make(map[string]*websocketClientAcknowledgesHolder),
	}
}

// AddEntryForClient will add the client to the inner map
func (ah *acknowledgesHolder) AddEntryForClient(remoteAddr string) {
	ah.mut.Lock()
	ah.acknowledges[remoteAddr] = NewWebsocketClientAcknowledgesHolder()
	ah.mut.Unlock()
}

// GetAcknowledgesOfAddress will return the acknowledges for the specified address, if any
func (ah *acknowledgesHolder) GetAcknowledgesOfAddress(remoteAddr string) (*websocketClientAcknowledgesHolder, bool) {
	ah.mut.Lock()
	defer ah.mut.Unlock()

	acks, found := ah.acknowledges[remoteAddr]
	return acks, found
}

// AddReceivedAcknowledge will add the received acknowledge as a counter for the given address
func (ah *acknowledgesHolder) AddReceivedAcknowledge(remoteAddr string, counter uint64) {
	ah.mut.Lock()
	ah.acknowledges[remoteAddr].Add(counter)
	ah.mut.Unlock()
}
