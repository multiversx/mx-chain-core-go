package factory

import (
	"fmt"
	"testing"

	"github.com/multiversx/mx-chain-core-go/core/mock"
	"github.com/multiversx/mx-chain-core-go/data/typeConverters/uint64ByteSlice"
	"github.com/multiversx/mx-chain-core-go/webSockets/data"
	"github.com/stretchr/testify/require"
)

func createArgs() ArgsWebSocketsDriverFactory {
	return ArgsWebSocketsDriverFactory{
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
}

func TestNewWebSocketsDriver(t *testing.T) {
	t.Parallel()

	args := createArgs()
	driver, err := NewWebSocketsDriver(args)
	require.Nil(t, err)
	require.NotNil(t, driver)
	require.Equal(t, "*webSockets.webSocketsDriver", fmt.Sprintf("%T", driver))

	err = driver.Close()
	require.Nil(t, err)
}

func TestCreateClient(t *testing.T) {
	t.Parallel()

	args := createArgs()
	webSocketsClient, err := createWebSocketsClient(args)
	require.Nil(t, err)
	require.Equal(t, "*client.client", fmt.Sprintf("%T", webSocketsClient))
}

func TestCreateServer(t *testing.T) {
	t.Parallel()

	args := createArgs()
	webSocketsClient, err := createWebSocketsServer(args)
	require.Nil(t, err)
	require.Equal(t, "*server.server", fmt.Sprintf("%T", webSocketsClient))
}
