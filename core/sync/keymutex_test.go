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

	km := NewKeyRWMutex()
	require.NotNil(t, km)
	require.Equal(t, 0, len(km.managedMutexes))
}

func TestKeyMutex_Lock_Unlock(t *testing.T) {
	t.Parallel()

	km := NewKeyRWMutex()
	require.NotNil(t, km)
	require.Len(t, km.managedMutexes, 0)
	km.Lock("id1")
	require.Len(t, km.managedMutexes, 1)
	km.Lock("id2")
	require.Len(t, km.managedMutexes, 2)
	km.Unlock("id1")
	require.Len(t, km.managedMutexes, 1)
	km.Unlock("id2")
	require.Len(t, km.managedMutexes, 0)
}

func TestKeyMutex_RLock_RUnlock(t *testing.T) {
	t.Parallel()

	km := NewKeyRWMutex()
	require.NotNil(t, km)
	require.Len(t, km.managedMutexes, 0)
	km.RLock("id1")
	require.Len(t, km.managedMutexes, 1)
	km.RLock("id2")
	require.Len(t, km.managedMutexes, 2)
	km.RUnlock("id1")
	require.Len(t, km.managedMutexes, 1)
	km.RUnlock("id2")
	require.Len(t, km.managedMutexes, 0)
}

func TestKeyMutex_IsInterfaceNil(t *testing.T) {
	t.Parallel()

	km := NewKeyRWMutex()
	require.False(t, check.IfNil(km))

	km = nil
	require.True(t, check.IfNil(km))

	var km2 KeyRWMutexHandler
	require.True(t, check.IfNil(km2))
}

func TestKeyMutex_ConcurrencyMultipleCriticalSections(t *testing.T) {
	t.Parallel()

	wg := sync.WaitGroup{}
	km := NewKeyRWMutex()
	require.NotNil(t, km)

	f := func(wg *sync.WaitGroup, id string) {
		km.Lock(id)
		<-time.After(time.Millisecond * 10)
		km.Unlock(id)

		km.RLock(id)
		<-time.After(time.Millisecond * 10)
		km.RUnlock(id)

		wg.Done()
	}

	numConcurrentCalls := 200
	ids := []string{"id1", "id2", "id3", "id4", "id5", "id6", "id7", "id8", "id9", "id10"}
	wg.Add(numConcurrentCalls)

	for i := 1; i <= numConcurrentCalls; i++ {
		go f(&wg, ids[i%len(ids)])
	}
	wg.Wait()

	require.Len(t, km.managedMutexes, 0)
}

func TestKeyMutex_ConcurrencySameID(t *testing.T) {
	t.Parallel()

	wg := sync.WaitGroup{}
	km := NewKeyRWMutex()
	require.NotNil(t, km)

	f := func(wg *sync.WaitGroup, id string) {
		km.RLock(id)
		km.RUnlock(id)

		wg.Done()
	}

	numConcurrentCalls := 500
	wg.Add(numConcurrentCalls)

	for i := 1; i <= numConcurrentCalls; i++ {
		go f(&wg, "id")
	}
	wg.Wait()

	require.Len(t, km.managedMutexes, 0)
}
