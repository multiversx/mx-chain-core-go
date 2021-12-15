package sender

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewWebsocketClientAcknowledgesHolder(t *testing.T) {
	t.Parallel()

	wcah := NewWebsocketClientAcknowledgesHolder()
	require.NotNil(t, wcah)
}

func TestWebsocketClientAcknowledgesHolder_Add(t *testing.T) {
	t.Parallel()

	counter := uint64(37)
	wcah := NewWebsocketClientAcknowledgesHolder()
	wcah.Add(counter)

	wcah.mutAcks.Lock()
	res, found := wcah.acks[counter]
	wcah.mutAcks.Unlock()

	require.True(t, found)
	require.NotNil(t, res)
}

func TestWebsocketClientAcknowledgesHolder_ProcessAcknowledged(t *testing.T) {
	t.Parallel()

	t.Run("ProcessAcknowledged: should not find", func(t *testing.T) {
		t.Parallel()

		wcah := NewWebsocketClientAcknowledgesHolder()
		res := wcah.ProcessAcknowledged(5)
		require.False(t, res)
	})

	t.Run("ProcessAcknowledged: should find and remove from inner map", func(t *testing.T) {
		t.Parallel()

		counter := uint64(37)
		wcah := NewWebsocketClientAcknowledgesHolder()
		wcah.Add(counter)

		res := wcah.ProcessAcknowledged(counter)
		require.True(t, res)

		wcah.mutAcks.Lock()
		require.Equal(t, 0, len(wcah.acks))
		wcah.mutAcks.Unlock()
	})
}

func TestWebsocketClientAcknowledgesHolder_ConcurrentOperations(t *testing.T) {
	t.Parallel()

	wcah := NewWebsocketClientAcknowledgesHolder()

	defer func() {
		r := recover()
		require.Nil(t, r)
	}()

	wg := sync.WaitGroup{}
	wg.Add(100)

	for i := uint64(0); i < 100; i++ {
		go func(index uint64) {
			switch index % 2 {
			case 0:
				wcah.Add(index)
			case 1:
				wcah.ProcessAcknowledged(index)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}
