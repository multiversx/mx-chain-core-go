package server

import (
	"sync"
)

type receiversHolder struct {
	mutex     sync.RWMutex
	receivers map[string]Receiver
}

func NewReceiversHolder() *receiversHolder {
	return &receiversHolder{
		receivers: map[string]Receiver{},
	}
}

func (rh *receiversHolder) AddReceiver(id string, rec Receiver) {
	rh.mutex.Lock()
	defer rh.mutex.Unlock()

	rh.receivers[id] = rec
}

func (rh *receiversHolder) RemoveReceiver(id string) {
	rh.mutex.Lock()
	defer rh.mutex.Unlock()

	delete(rh.receivers, id)
}

func (rh *receiversHolder) GetAll() map[string]Receiver {
	rh.mutex.RLock()
	defer rh.mutex.RUnlock()

	allReceivers := make(map[string]Receiver)
	for id, listener := range rh.receivers {
		allReceivers[id] = listener
	}

	return allReceivers
}
