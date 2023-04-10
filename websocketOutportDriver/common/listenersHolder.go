package common

import "sync"

type listenerHolder struct {
	mutex     sync.RWMutex
	listeners map[string]MessagesListener
}

// NewListenersHolder will create a new instance of listenerHolder
func NewListenersHolder() *listenerHolder {
	return &listenerHolder{
		mutex:     sync.RWMutex{},
		listeners: make(map[string]MessagesListener),
	}
}

// Add will add a new listener in the map
func (lh *listenerHolder) Add(id string, listener MessagesListener) {
	lh.mutex.Lock()
	defer lh.mutex.Unlock()

	lh.listeners[id] = listener
}

// Remove will remove a listener from map based on the provided id
func (lh *listenerHolder) Remove(id string) {
	lh.mutex.Lock()
	defer lh.mutex.Unlock()

	delete(lh.listeners, id)
}

// GetAll will return a new map with all the listeners
func (lh *listenerHolder) GetAll() map[string]MessagesListener {
	lh.mutex.RLock()
	defer lh.mutex.RUnlock()

	allListeners := make(map[string]MessagesListener)
	for id, listener := range lh.listeners {
		allListeners[id] = listener
	}

	return allListeners
}
