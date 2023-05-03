package server

import (
	"sync"
)

type receiversHolder struct {
	mutex     sync.RWMutex
	receivers map[string]Receiver
}

// NewReceiversHolder will create a new instance of receiversHolder
func NewReceiversHolder() *receiversHolder {
	return &receiversHolder{
		receivers: map[string]Receiver{},
	}
}

// AddReceiver will add the provided receiver in the internal map
func (rh *receiversHolder) AddReceiver(id string, rec Receiver) {
	rh.mutex.Lock()
	defer rh.mutex.Unlock()

	rh.receivers[id] = rec
}

// RemoveReceiver will remove the provided receiver from the internal map
func (rh *receiversHolder) RemoveReceiver(id string) {
	rh.mutex.Lock()
	defer rh.mutex.Unlock()

	delete(rh.receivers, id)
}

// GetAll will return a map with all the stored receivers
func (rh *receiversHolder) GetAll() map[string]Receiver {
	rh.mutex.RLock()
	defer rh.mutex.RUnlock()

	allReceivers := make(map[string]Receiver)
	for id, listener := range rh.receivers {
		allReceivers[id] = listener
	}

	return allReceivers
}
