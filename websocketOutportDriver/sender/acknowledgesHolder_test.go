package sender

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewAcknowledgesHolder(t *testing.T) {
	t.Parallel()

	ah := NewAcknowledgesHolder()
	require.NotNil(t, ah)
}

func TestAcknowledgesHolder_AddEntry(t *testing.T) {
	t.Parallel()

	remAddr := "test address"
	ah := NewAcknowledgesHolder()
	ah.AddEntry(remAddr)

	ah.mut.Lock()
	res, found := ah.acknowledges[remAddr]
	ah.mut.Unlock()

	require.True(t, found)
	require.NotNil(t, res)
}

func TestAcknowledgesHolder_AddReceivedAcknowledge(t *testing.T) {
	t.Parallel()

	remAddr := "test address"
	counter := uint64(37)
	ah := NewAcknowledgesHolder()
	ah.AddEntry(remAddr)

	ah.AddReceivedAcknowledge(remAddr, counter)

	ah.mut.Lock()
	found := ah.acknowledges[remAddr].ProcessAcknowledged(counter)
	ah.mut.Unlock()

	require.True(t, found)
}

func TestAcknowledgesHolder_GetAcknowledgesOfAddress(t *testing.T) {
	t.Parallel()

	t.Run("GetAcknowledgesOfAddress: not found", func(t *testing.T) {
		t.Parallel()

		ah := NewAcknowledgesHolder()

		res, found := ah.GetAcknowledgesOfAddress("new addr")
		require.False(t, found)
		require.Nil(t, res)
	})

	t.Run("GetAcknowledgesOfAddress: should work", func(t *testing.T) {
		t.Parallel()

		remAddr := "test address"
		counter0, counter1 := uint64(37), uint64(38)
		ah := NewAcknowledgesHolder()
		ah.AddEntry(remAddr)

		ah.AddReceivedAcknowledge(remAddr, counter0)
		ah.AddReceivedAcknowledge(remAddr, counter1)

		acks, found := ah.GetAcknowledgesOfAddress(remAddr)
		require.True(t, found)

		found0 := acks.ProcessAcknowledged(counter0)
		found1 := acks.ProcessAcknowledged(counter1)

		require.True(t, found0)
		require.True(t, found1)
	})
}

func TestAcknowledgesHolder_RemoveEntryForAddress(t *testing.T) {
	t.Parallel()

	remAddr := "remote addr"

	ah := NewAcknowledgesHolder()

	ah.AddEntry(remAddr)
	ah.RemoveEntryForAddress(remAddr)

	ah.mut.Lock()
	_, found := ah.acknowledges[remAddr]
	ah.mut.Unlock()

	require.False(t, found)
}

func TestAcknowledgesHolder_ConcurrentOperations(t *testing.T) {
	t.Parallel()

	ah := NewAcknowledgesHolder()

	defer func() {
		r := recover()
		require.Nil(t, r)
	}()

	wg := sync.WaitGroup{}
	wg.Add(100)

	for i := uint64(0); i < 100; i++ {
		go func(index uint64) {
			switch index % 4 {
			case 0:
				ah.AddReceivedAcknowledge("addr", index)
			case 1:
				_, _ = ah.GetAcknowledgesOfAddress("addr")
			case 2:
				ah.RemoveEntryForAddress("addr")
			case 3:
				ah.AddEntry("addr")
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}
