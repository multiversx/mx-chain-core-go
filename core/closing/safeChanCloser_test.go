package closing

import (
	"fmt"
	"testing"
	"time"

	"github.com/multiversx/mx-chain-core-go/core/check"
	"github.com/stretchr/testify/require"
)

const timeout = time.Second

func TestNewSafeChanCloser(t *testing.T) {
	t.Parallel()

	closer := NewSafeChanCloser()
	require.False(t, check.IfNil(closer))
}

func TestSafeChanCloser_CloseShouldWork(t *testing.T) {
	t.Parallel()

	chDone := make(chan struct{}, 1)
	closer := NewSafeChanCloser()
	go func() {
		<-closer.ChanClose()
		chDone <- struct{}{}
	}()

	closer.Close()

	select {
	case <-chDone:
	case <-time.After(timeout):
		require.Fail(t, "should have not timed out")
	}
}

func TestSafeChanCloser_Reset(t *testing.T) {
	t.Parallel()

	closer := NewSafeChanCloser()

	isClosed := func(closer *safeChanCloser) bool {
		select {
		case <-closer.ChanClose():
			return true
		default:
			return false
		}
	}

	// test multiple use-cases, such as resetting an already closed instance
	require.False(t, isClosed(closer))
	closer.Reset()
	require.False(t, isClosed(closer))
	closer.Close()
	require.True(t, isClosed(closer))
	closer.Reset()
	require.False(t, isClosed(closer))
}

func TestSafeChanCloser_ConcurrentOperationsShouldNotPanic(t *testing.T) {
	t.Parallel()

	defer panicHandler(t)

	closer := NewSafeChanCloser()
	for i := 0; i < 10; i++ {
		go func(idx int) {
			defer panicHandler(t)

			time.Sleep(time.Millisecond * 100)
			switch idx % 2 {
			case 0:
				closer.Close()
			case 1:
				closer.Reset()
			}

		}(i)
	}

	time.Sleep(timeout)
}

func panicHandler(t *testing.T) {
	r := recover()
	if r != nil {
		require.Fail(t, fmt.Sprintf("should have not panicked: %v", r))
	}
}
