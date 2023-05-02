package factory

import (
	"testing"

	"github.com/multiversx/mx-chain-core-go/core/mock"
	"github.com/multiversx/mx-chain-core-go/data/typeConverters/uint64ByteSlice"
	"github.com/multiversx/mx-chain-core-go/webSockets/data"
	"github.com/stretchr/testify/require"
)

func TestNewWebSocketsDriver(t *testing.T) {
	t.Parallel()

	args := ArgsWebSocketsDriverFactory{
		WebSocketConfig: data.WebSocketConfig{
			URL:                "localhost:1234",
			WithAcknowledge:    false,
			IsServer:           false,
			RetryDurationInSec: 1,
		},
		Marshaller:               &mock.MarshalizerMock{},
		Uint64ByteSliceConverter: uint64ByteSlice.NewBigEndianConverter(),
		Log:                      &mock.LoggerMock{},
		WithAcknowledge:          false,
	}

	driver, err := NewWebSocketsDriver(args)
	require.Nil(t, err)
	require.NotNil(t, driver)

	err = driver.Close()
	require.Nil(t, err)
}
