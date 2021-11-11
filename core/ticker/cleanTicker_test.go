package ticker

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/ElrondNetwork/elrond-go-core/core/check"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewCleanTicker(t *testing.T) {
	t.Parallel()

	ct := NewCleanTicker(time.Second)
	defer ct.Stop()

	assert.False(t, check.IfNil(ct))
}

func TestCleanTicker_ResetAndCleanShouldNotCallTwice(t *testing.T) {
	t.Parallel()

	for i := 0; i < 100; i++ {
		testCleanTickerResetAndCleanShouldNotCallTwice(t)
	}
}

func testCleanTickerResetAndCleanShouldNotCallTwice(t *testing.T) {
	interval := time.Millisecond
	ct := NewCleanTicker(interval)
	defer ct.Stop()

	numCalls := uint32(0)
	wg := sync.WaitGroup{}
	wg.Add(1)

	chFinished := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		for {
			ct.ResetAndClean(interval)

			select {
			case <-ct.Chan():
				atomic.AddUint32(&numCalls, 1)
				wg.Wait()
			case <-ctx.Done():
				close(chFinished)
				return
			}
		}
	}()

	// this delay will cause a new time.Time to be available in the ticker's chan
	time.Sleep(time.Millisecond * 10)

	cancel()
	wg.Done()

	select {
	case <-chFinished:
	case <-time.After(time.Second):
		require.Fail(t, "called multiple times even if the context was done")
	}

	assert.Equal(t, uint32(1), atomic.LoadUint32(&numCalls))
}

func TestCleanTicker_ResetMightCallTwice(t *testing.T) {
	t.Parallel()

	for i := 0; i < 100; i++ {
		testCleanTickerResetMightCallTwice(t)
	}
}

func testCleanTickerResetMightCallTwice(t *testing.T) {
	interval := time.Millisecond
	ct := NewCleanTicker(interval)
	defer ct.Stop()

	numCalls := uint32(0)
	wg := sync.WaitGroup{}
	wg.Add(1)

	chFinished := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		for {
			ct.Reset(interval)

			select {
			case <-ct.Chan():
				atomic.AddUint32(&numCalls, 1)
				wg.Wait()
			case <-ctx.Done():
				close(chFinished)
				return
			}
		}
	}()

	// this delay will cause a new time.Time to be available in the ticker's chan
	time.Sleep(time.Millisecond * 10)

	cancel()
	wg.Done()

	time.Sleep(time.Millisecond * 10)

	assert.True(t, uint32(1) <= atomic.LoadUint32(&numCalls))
	if atomic.LoadUint32(&numCalls) > 1 {
		fmt.Println("found one extra call")
	}
}
