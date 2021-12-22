package atomic

import "sync/atomic"

// Flag is an atomic flag
type Flag struct {
	value uint32
}

// SetReturningPrevious sets flag and returns its previous value
func (flag *Flag) SetReturningPrevious() bool {
	previousValue := atomic.SwapUint32(&flag.value, 1)
	return previousValue == 1
}

// Reset resets the flag, putting it in off position
func (flag *Flag) Reset() {
	atomic.StoreUint32(&flag.value, 0)
}

// IsSet checks whether flag is set
func (flag *Flag) IsSet() bool {
	value := atomic.LoadUint32(&flag.value)
	return value == 1
}

// Toggle toggles the flag
func (flag *Flag) Toggle(set bool) {
	if set {
		_ = flag.SetReturningPrevious()
	} else {
		flag.Reset()
	}
}
