package atomic

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFlag_SetReturningPrevious(t *testing.T) {
	t.Parallel()

	var flag Flag
	var wg sync.WaitGroup

	require.False(t, flag.IsSet())

	wg.Add(2)

	go func() {
		_ = flag.SetReturningPrevious()
		wg.Done()
	}()

	go func() {
		_ = flag.SetReturningPrevious()
		wg.Done()
	}()

	wg.Wait()
	require.True(t, flag.IsSet())
}

func TestFlag_Reset(t *testing.T) {
	t.Parallel()

	var flag Flag
	var wg sync.WaitGroup

	_ = flag.SetReturningPrevious()
	require.True(t, flag.IsSet())

	wg.Add(2)

	go func() {
		flag.Reset()
		wg.Done()
	}()

	go func() {
		flag.Reset()
		wg.Done()
	}()

	wg.Wait()
	require.False(t, flag.IsSet())
}

func TestFlag_SetValue(t *testing.T) {
	t.Parallel()

	var flag Flag
	var wg sync.WaitGroup

	// First, Toggle(true)
	wg.Add(2)

	go func() {
		flag.SetValue(true)
		wg.Done()
	}()

	go func() {
		flag.SetValue(true)
		wg.Done()
	}()

	wg.Wait()
	require.True(t, flag.IsSet())

	// Then, Toggle(false)
	wg.Add(2)

	go func() {
		flag.SetValue(false)
		wg.Done()
	}()

	go func() {
		flag.SetValue(false)
		wg.Done()
	}()

	wg.Wait()
	require.False(t, flag.IsSet())
}
