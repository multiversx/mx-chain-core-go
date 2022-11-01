package sender

import "sync"

type websocketClientAcknowledgesHolder struct {
	acks    map[uint64]struct{}
	mutAcks sync.Mutex
}

// NewWebsocketClientAcknowledgesHolder will return a new instance of websocketAcknowledgesHolder
func NewWebsocketClientAcknowledgesHolder() *websocketClientAcknowledgesHolder {
	return &websocketClientAcknowledgesHolder{
		acks: make(map[uint64]struct{}),
	}
}

// Add will add an element
func (wah *websocketClientAcknowledgesHolder) Add(counter uint64) {
	wah.mutAcks.Lock()
	wah.acks[counter] = struct{}{}
	wah.mutAcks.Unlock()
}

// ProcessAcknowledged will process the acknowledgment for the given counter. If found, the element will also be
// removed from the inner map
func (wah *websocketClientAcknowledgesHolder) ProcessAcknowledged(counter uint64) bool {
	wah.mutAcks.Lock()
	defer wah.mutAcks.Unlock()

	_, exists := wah.acks[counter]
	if !exists {
		return false
	}

	delete(wah.acks, counter)
	return true
}
