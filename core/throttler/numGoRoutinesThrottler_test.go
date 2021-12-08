package throttler_test

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/ElrondNetwork/elrond-go-core/core"
	"github.com/ElrondNetwork/elrond-go-core/core/check"
	"github.com/ElrondNetwork/elrond-go-core/core/throttler"
	"github.com/stretchr/testify/assert"
)

func TestNewNumGoRoutinesThrottler_WithNegativeShouldError(t *testing.T) {
	t.Parallel()

	nt, err := throttler.NewNumGoRoutinesThrottler(-1)

	assert.Nil(t, nt)
	assert.Equal(t, core.ErrNotPositiveValue, err)
}

func TestNewNumGoRoutinesThrottler_WithZeroShouldError(t *testing.T) {
	t.Parallel()

	nt, err := throttler.NewNumGoRoutinesThrottler(0)

	assert.Nil(t, nt)
	assert.Equal(t, core.ErrNotPositiveValue, err)
}

func TestNewNumGoRoutinesThrottler_ShouldWork(t *testing.T) {
	t.Parallel()

	nt, err := throttler.NewNumGoRoutinesThrottler(1)

	assert.Nil(t, err)
	assert.False(t, check.IfNil(nt))
}

func TestNumGoRoutinesThrottler_CanProcessMessageWithZeroCounter(t *testing.T) {
	t.Parallel()

	nt, _ := throttler.NewNumGoRoutinesThrottler(1)

	assert.True(t, nt.CanProcess())
}

func TestNumGoRoutinesThrottler_StartProcessingIfCanProcessEqualsMax(t *testing.T) {
	t.Parallel()

	nt, _ := throttler.NewNumGoRoutinesThrottler(1)
	result := nt.StartProcessingIfCanProcess()

	assert.False(t, nt.CanProcess())
	assert.True(t, result)

	result = nt.StartProcessingIfCanProcess()
	assert.False(t, result)
}

func TestNumGoRoutinesThrottler_StartProcessingIfCanProcessCounterIsMaxLessThanOne(t *testing.T) {
	t.Parallel()

	max := int32(45)
	nt, _ := throttler.NewNumGoRoutinesThrottler(max)

	for i := int32(0); i < max-1; i++ {
		result := nt.StartProcessingIfCanProcess()
		assert.True(t, result)
	}

	assert.True(t, nt.CanProcess())
}

func TestNumGoRoutinesThrottler_StartProcessingIfCanProcessCounterIsMax(t *testing.T) {
	t.Parallel()

	max := int32(45)
	nt, _ := throttler.NewNumGoRoutinesThrottler(max)

	for i := int32(0); i < max; i++ {
		result := nt.StartProcessingIfCanProcess()
		assert.True(t, result)
	}

	assert.False(t, nt.CanProcess())
}

func TestNumGoRoutinesThrottler_StartProcessingIfCanProcessCounterIsMaxLessOneFromEndProcessMessage(t *testing.T) {
	t.Parallel()

	max := int32(45)
	nt, _ := throttler.NewNumGoRoutinesThrottler(max)

	for i := int32(0); i < max; i++ {
		result := nt.StartProcessingIfCanProcess()
		assert.True(t, result)
	}
	nt.EndProcessing()

	assert.True(t, nt.CanProcess())
}

func TestNumGoRoutinesThrottler_StartProcessingIfCanProcessSimultaneousCallOnLastEntry(t *testing.T) {
	t.Parallel()

	max := int32(45)
	nt, _ := throttler.NewNumGoRoutinesThrottler(max)

	for i := int32(0); i < max-1; i++ {
		_ = nt.StartProcessingIfCanProcess()
	}

	numGoRoutines := 100
	wg := &sync.WaitGroup{}
	wg.Add(numGoRoutines)
	numStarted := uint32(0)
	numNotStarted := uint32(0)

	for i := 0; i < numGoRoutines; i++ {
		go func() {
			time.Sleep(time.Millisecond * 10)

			result := nt.StartProcessingIfCanProcess()
			if result {
				atomic.AddUint32(&numStarted, 1)
			} else {
				atomic.AddUint32(&numNotStarted, 1)
			}

			wg.Done()
		}()
	}

	wg.Wait()

	assert.Equal(t, uint32(1), atomic.LoadUint32(&numStarted))
	assert.Equal(t, uint32(numGoRoutines-1), atomic.LoadUint32(&numNotStarted))
}
