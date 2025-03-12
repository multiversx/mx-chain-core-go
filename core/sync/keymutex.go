package sync

import "sync"

// keyRWMutex is a mutex that can be used to lock/unlock a resource identified by a key
type keyRWMutex struct {
	mut            sync.RWMutex
	managedMutexes map[string]*rwMutex
}

// NewKeyRWMutex returns a new instance of keyRWMutex
func NewKeyRWMutex() *keyRWMutex {
	return &keyRWMutex{
		managedMutexes: make(map[string]*rwMutex),
	}
}

// RLock locks for read the Mutex for the given key
func (km *keyRWMutex) RLock(key string) {
	km.getForRLock(key).rLock()
}

// RUnlock unlocks for read the Mutex for the given key
func (km *keyRWMutex) RUnlock(key string) {
	km.getForRUnlock(key).rUnlock()
	km.cleanupMutex(key)
}

// Lock locks the Mutex for the given key
func (km *keyRWMutex) Lock(key string) {
	km.getForLock(key).lock()
}

// Unlock unlocks the Mutex for the given key
func (km *keyRWMutex) Unlock(key string) {
	km.getForUnlock(key).unlock()
	km.cleanupMutex(key)
}

// getForLock returns the Mutex for the given key, updating the Lock counter
func (km *keyRWMutex) getForLock(key string) *rwMutex {
	km.mut.Lock()
	defer km.mut.Unlock()

	mutex, ok := km.managedMutexes[key]
	if !ok {
		mutex = km.newInternalMutex(key)
	}
	mutex.updateCounterLock()

	return mutex
}

// getForRLock returns the Mutex for the given key, updating the RLock counter
func (km *keyRWMutex) getForRLock(key string) *rwMutex {
	km.mut.Lock()
	defer km.mut.Unlock()

	mutex, ok := km.managedMutexes[key]
	if !ok {
		mutex = km.newInternalMutex(key)
	}
	mutex.updateCounterRLock()

	return mutex
}

// getForUnlock returns the Mutex for the given key, updating the Unlock counter
func (km *keyRWMutex) getForUnlock(key string) *rwMutex {
	km.mut.Lock()
	defer km.mut.Unlock()

	mutex, ok := km.managedMutexes[key]
	if !ok {
		mutex = km.newInternalMutex(key)
	}
	mutex.updateCounterUnlock()

	return mutex
}

// getForRUnlock returns the Mutex for the given key, updating the RUnlock counter
func (km *keyRWMutex) getForRUnlock(key string) *rwMutex {
	km.mut.Lock()
	defer km.mut.Unlock()

	mutex, ok := km.managedMutexes[key]
	if !ok {
		mutex = km.newInternalMutex(key)
	}
	mutex.updateCounterRUnlock()

	return mutex
}

// newInternalMutex creates a new mutex for the given key and adds it to the map
func (km *keyRWMutex) newInternalMutex(key string) *rwMutex {
	mutex, ok := km.managedMutexes[key]
	if !ok {
		mutex = newRWMutex()
		km.managedMutexes[key] = mutex
	}
	return mutex
}

// cleanupMutex removes the mutex from the map if it is not used anymore
func (km *keyRWMutex) cleanupMutex(key string) {
	km.mut.Lock()
	defer km.mut.Unlock()

	mut, ok := km.managedMutexes[key]
	if ok && mut.numLocks() == 0 {
		delete(km.managedMutexes, key)
	}
}

// IsInterfaceNil returns true if there is no value under the interface
func (km *keyRWMutex) IsInterfaceNil() bool {
	return km == nil
}
