package sync

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/multiversx/mx-chain-core-go/core/check"
)

func TestNewKeyMutex(t *testing.T) {
	t.Parallel()

	csa := NewKeyRWMutex()
	require.NotNil(t, csa)
	require.Equal(t, 0, len(csa.managedMutexes))
}

func TestKeyMutex_Lock_Unlock(t *testing.T) {
	t.Parallel()

	csa := NewKeyRWMutex()
	require.NotNil(t, csa)
	require.Len(t, csa.managedMutexes, 0)
	csa.Lock("id1")
	require.Len(t, csa.managedMutexes, 1)
	csa.Lock("id2")
	require.Len(t, csa.managedMutexes, 2)
	csa.Unlock("id1")
	require.Len(t, csa.managedMutexes, 1)
	csa.Unlock("id2")
	require.Len(t, csa.managedMutexes, 0)
}

func TestKeyMutex_RLock_RUnlock(t *testing.T) {
	t.Parallel()

	csa := NewKeyRWMutex()
	require.NotNil(t, csa)
	require.Len(t, csa.managedMutexes, 0)
	csa.RLock("id1")
	require.Len(t, csa.managedMutexes, 1)
	csa.RLock("id2")
	require.Len(t, csa.managedMutexes, 2)
	csa.RUnlock("id1")
	require.Len(t, csa.managedMutexes, 1)
	csa.RUnlock("id2")
	require.Len(t, csa.managedMutexes, 0)
}

func TestKeyMutex_IsInterfaceNil(t *testing.T) {
	t.Parallel()

	csa := NewKeyRWMutex()
	require.False(t, check.IfNil(csa))

	csa = nil
	require.True(t, check.IfNil(csa))

	var csa2 KeyRWMutexHandler
	require.True(t, check.IfNil(csa2))
}

func TestKeyMutex_ConcurrencyMultipleCriticalSections(t *testing.T) {
	t.Parallel()

	wg := sync.WaitGroup{}
	csa := NewKeyRWMutex()
	require.NotNil(t, csa)

	f := func(wg *sync.WaitGroup, id string) {
		csa.Lock(id)
		<-time.After(time.Millisecond * 10)
		csa.Unlock(id)

		csa.RLock(id)
		<-time.After(time.Millisecond * 10)
		csa.RUnlock(id)

		wg.Done()
	}

	numConcurrentCalls := 200
	ids := []string{"id1", "id2", "id3", "id4", "id5", "id6", "id7", "id8", "id9", "id10"}
	wg.Add(numConcurrentCalls)

	for i := 1; i <= numConcurrentCalls; i++ {
		go f(&wg, ids[i%len(ids)])
	}
	wg.Wait()

	require.Len(t, csa.managedMutexes, 0)
}

func TestKeyMutex_ConcurrencySameID(t *testing.T) {
	t.Parallel()

	wg := sync.WaitGroup{}
	csa := NewKeyRWMutex()
	require.NotNil(t, csa)

	f := func(wg *sync.WaitGroup, id string) {
		csa.RLock(id)
		csa.RUnlock(id)

		wg.Done()
	}

	numConcurrentCalls := 500
	wg.Add(numConcurrentCalls)

	for i := 1; i <= numConcurrentCalls; i++ {
		go f(&wg, "id")
	}
	wg.Wait()

	require.Len(t, csa.managedMutexes, 0)
}
