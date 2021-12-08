package throttler

import (
	"sync"

	"github.com/ElrondNetwork/elrond-go-core/core"
)

// NumGoRoutinesThrottler can limit the number of go routines launched
type NumGoRoutinesThrottler struct {
	max        int32
	mutCounter sync.RWMutex
	counter    int32
}

// NewNumGoRoutinesThrottler creates a new num go routine throttler instance
func NewNumGoRoutinesThrottler(max int32) (*NumGoRoutinesThrottler, error) {
	if max <= 0 {
		return nil, core.ErrNotPositiveValue
	}

	return &NumGoRoutinesThrottler{
		max: max,
	}, nil
}

// CanProcess returns true if current counter is less than max
func (ngrt *NumGoRoutinesThrottler) CanProcess() bool {
	ngrt.mutCounter.RLock()
	defer ngrt.mutCounter.RUnlock()

	return ngrt.counter < ngrt.max
}

// StartProcessingIfCanProcess returns true if current counter is less than max. It also increments the counter
func (ngrt *NumGoRoutinesThrottler) StartProcessingIfCanProcess() bool {
	ngrt.mutCounter.Lock()
	defer ngrt.mutCounter.Unlock()

	canIncrement := ngrt.counter < ngrt.max
	if canIncrement {
		ngrt.counter++
	}

	return canIncrement
}

// EndProcessing will decrement current counter
func (ngrt *NumGoRoutinesThrottler) EndProcessing() {
	ngrt.mutCounter.Lock()
	defer ngrt.mutCounter.Unlock()

	ngrt.counter--
}

// IsInterfaceNil returns true if there is no value under the interface
func (ngrt *NumGoRoutinesThrottler) IsInterfaceNil() bool {
	return ngrt == nil
}
