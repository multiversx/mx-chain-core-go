package indexer

import (
	"errors"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/ElrondNetwork/elrond-go/core/indexer/workItems"
	"github.com/ElrondNetwork/elrond-go/core/mock"
	"github.com/stretchr/testify/require"
)

func TestNewDataDispatcher_InvalidCacheSize(t *testing.T) {
	t.Parallel()

	dataDist, err := NewDataDispatcher(-1)

	require.Nil(t, dataDist)
	require.Equal(t, ErrNegativeCacheSize, err)
}

func TestNewDataDispatcher(t *testing.T) {
	t.Parallel()

	dispatcher, err := NewDataDispatcher(100)
	require.NoError(t, err)
	require.NotNil(t, dispatcher)
}

func TestDataDispatcher_StartIndexDataClose(t *testing.T) {
	t.Parallel()

	dispatcher, err := NewDataDispatcher(100)
	require.NoError(t, err)
	dispatcher.StartIndexData()

	called := false
	wg := sync.WaitGroup{}
	wg.Add(1)
	elasticProc := &mock.ElasticProcessorStub{
		SaveRoundsInfoCalled: func(infos []workItems.RoundInfo) error {
			called = true
			wg.Done()
			return nil
		},
	}
	dispatcher.Add(workItems.NewItemRounds(elasticProc, []workItems.RoundInfo{}))
	wg.Wait()

	require.True(t, called)

	err = dispatcher.Close()
	require.NoError(t, err)
}

func TestDataDispatcher_Add(t *testing.T) {
	t.Parallel()

	dispatcher, err := NewDataDispatcher(100)
	require.NoError(t, err)
	dispatcher.StartIndexData()

	calledCount := 0
	wg := sync.WaitGroup{}
	wg.Add(1)
	elasticProc := &mock.ElasticProcessorStub{
		SaveRoundsInfoCalled: func(infos []workItems.RoundInfo) error {
			if calledCount < 2 {
				calledCount++
				return fmt.Errorf("%w: wrapped error", ErrBackOff)
			}

			calledCount++
			wg.Done()
			return nil
		},
	}

	start := time.Now()
	dispatcher.Add(workItems.NewItemRounds(elasticProc, []workItems.RoundInfo{}))
	wg.Wait()

	timePassed := time.Since(start)
	require.Greater(t, 2*int64(timePassed), int64(backOffTime))

	require.Equal(t, 3, calledCount)

	err = dispatcher.Close()
	require.NoError(t, err)
}

func TestDataDispatcher_AddWithErrorShouldRetryTheReprocessing(t *testing.T) {
	t.Parallel()

	dispatcher, err := NewDataDispatcher(100)
	require.NoError(t, err)
	dispatcher.StartIndexData()

	calledCount := 0
	wg := sync.WaitGroup{}
	wg.Add(1)
	elasticProc := &mock.ElasticProcessorStub{
		SaveRoundsInfoCalled: func(infos []workItems.RoundInfo) error {
			if calledCount < 2 {
				calledCount++
				return errors.New("generic error")
			}

			calledCount++
			wg.Done()
			return nil
		},
	}

	start := time.Now()
	dispatcher.Add(workItems.NewItemRounds(elasticProc, []workItems.RoundInfo{}))
	wg.Wait()

	timePassed := time.Since(start)
	require.Greater(t, int64(timePassed), int64(2*durationBetweenErrorRetry))

	require.Equal(t, 3, calledCount)

	err = dispatcher.Close()
	require.NoError(t, err)
}
