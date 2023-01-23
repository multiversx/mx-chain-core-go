package appStatusPolling_test

import (
	"context"
	"testing"
	"time"

	"github.com/multiversx/mx-chain-core-go/core"
	"github.com/multiversx/mx-chain-core-go/core/appStatusPolling"
	"github.com/multiversx/mx-chain-core-go/core/mock"
	"github.com/stretchr/testify/assert"
)

func TestNewAppStatusPooling_NilAppStatusHandlerShouldErr(t *testing.T) {
	t.Parallel()

	_, err := appStatusPolling.NewAppStatusPolling(nil, time.Second, &mock.LoggerMock{})
	assert.Equal(t, err, appStatusPolling.ErrNilAppStatusHandler)
}

func TestNewAppStatusPooling_NilLoggerShouldErr(t *testing.T) {
	t.Parallel()

	_, err := appStatusPolling.NewAppStatusPolling(&mock.StatusHandlerMock{}, time.Second, nil)
	assert.Equal(t, err, core.ErrNilLogger)
}

func TestNewAppStatusPooling_NegativePollingDurationShouldErr(t *testing.T) {
	t.Parallel()

	_, err := appStatusPolling.NewAppStatusPolling(&mock.StatusHandlerMock{}, time.Duration(-1), &mock.LoggerMock{})
	assert.Equal(t, err, appStatusPolling.ErrPollingDurationToSmall)
}

func TestNewAppStatusPooling_ZeroPollingDurationShouldErr(t *testing.T) {
	t.Parallel()

	_, err := appStatusPolling.NewAppStatusPolling(&mock.StatusHandlerMock{}, 0, &mock.LoggerMock{})
	assert.Equal(t, err, appStatusPolling.ErrPollingDurationToSmall)
}

func TestNewAppStatusPooling_OkValsShouldPass(t *testing.T) {
	t.Parallel()

	_, err := appStatusPolling.NewAppStatusPolling(&mock.StatusHandlerMock{}, time.Second, &mock.LoggerMock{})
	assert.Nil(t, err)
}

func TestNewAppStatusPolling_RegisterHandlerFuncShouldErr(t *testing.T) {
	t.Parallel()

	asp, err := appStatusPolling.NewAppStatusPolling(&mock.StatusHandlerMock{}, time.Second, &mock.LoggerMock{})
	assert.Nil(t, err)

	err = asp.RegisterPollingFunc(nil)
	assert.Equal(t, appStatusPolling.ErrNilHandlerFunc, err)
}

func TestAppStatusPolling_Poll_TestNumOfConnectedAddressesCalled(t *testing.T) {
	t.Parallel()

	pollingDuration := time.Second
	chDone := make(chan struct{})
	ash := mock.AppStatusHandlerStub{
		SetInt64ValueHandler: func(key string, value int64) {
			chDone <- struct{}{}
		},
	}
	asp, err := appStatusPolling.NewAppStatusPolling(&ash, pollingDuration, &mock.LoggerMock{})
	assert.Nil(t, err)

	err = asp.RegisterPollingFunc(func(appStatusHandler core.AppStatusHandler) {
		appStatusHandler.SetInt64Value("metric", int64(10))
	})
	assert.Nil(t, err)

	ctx := context.Background()
	asp.Poll(ctx)

	select {
	case <-chDone:
	case <-time.After(pollingDuration * 2 * time.Second):
		assert.Fail(t, "timeout calling SetInt64Value")
	}
}
