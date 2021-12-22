package atomic

import (
	"sync"
)

// Flag is an atomic flag
type Flag struct {
	mut   sync.RWMutex
	value bool
}

// SetReturningPrevious sets flag and returns its previous value
func (flag *Flag) SetReturningPrevious() bool {
	flag.mut.Lock()
	previousValue := flag.value
	flag.value = true
	flag.mut.Unlock()

	return previousValue
}

// Reset resets the flag, putting it in off position
func (flag *Flag) Reset() {
	flag.mut.Lock()
	flag.value = false
	flag.mut.Unlock()
}

// IsSet checks whether flag is set
func (flag *Flag) IsSet() bool {
	flag.mut.RLock()
	defer flag.mut.RUnlock()

	return flag.value
}

// Toggle toggles the flag
func (flag *Flag) Toggle() {
	flag.mut.Lock()
	flag.value = !flag.value
	flag.mut.Unlock()
}

// SetValue sets the new value in the flag
func (flag *Flag) SetValue(newValue bool) {
	flag.mut.Lock()
	flag.value = newValue
	flag.mut.Unlock()
}
