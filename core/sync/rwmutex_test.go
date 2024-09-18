package sync

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewRWMutex(t *testing.T) {
	t.Parallel()

	cs := newRWMutex()
	require.NotNil(t, cs)
	require.Equal(t, int32(0), cs.cntLocks)
	require.Equal(t, int32(0), cs.cntRLocks)
}

func TestRWMutex_Lock_Unlock_IsLocked_NumLocks(t *testing.T) {
	t.Parallel()

	cs := &rwMutex{}
	cs.lock()
	cs.updateCounterLock()
	require.Equal(t, int32(1), cs.numLocks())

	cs.unlock()
	cs.updateCounterUnlock()
	require.Equal(t, int32(0), cs.numLocks())

	cs.rLock()
	cs.updateCounterRLock()
	require.Equal(t, int32(1), cs.numLocks())

	cs.rUnlock()
	cs.updateCounterRUnlock()
	require.Equal(t, int32(0), cs.numLocks())
}
