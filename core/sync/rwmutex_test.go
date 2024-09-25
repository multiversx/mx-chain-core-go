package sync

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewRWMutex(t *testing.T) {
	t.Parallel()

	rwm := newRWMutex()
	require.NotNil(t, rwm)
	require.Equal(t, int32(0), rwm.cntLocks)
	require.Equal(t, int32(0), rwm.cntRLocks)
}

func TestRWMutex_Lock_Unlock_IsLocked_NumLocks(t *testing.T) {
	t.Parallel()

	rwm := &rwMutex{}
	rwm.lock()
	rwm.updateCounterLock()
	require.Equal(t, int32(1), rwm.numLocks())

	rwm.unlock()
	rwm.updateCounterUnlock()
	require.Equal(t, int32(0), rwm.numLocks())

	rwm.rLock()
	rwm.updateCounterRLock()
	require.Equal(t, int32(1), rwm.numLocks())

	rwm.rUnlock()
	rwm.updateCounterRUnlock()
	require.Equal(t, int32(0), rwm.numLocks())
}
