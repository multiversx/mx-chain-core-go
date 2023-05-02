package server

import (
	"errors"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/multiversx/mx-chain-core-go/core/mock"
	"github.com/multiversx/mx-chain-core-go/data/typeConverters/uint64ByteSlice"
	"github.com/multiversx/mx-chain-core-go/testscommon"
	"github.com/multiversx/mx-chain-core-go/webSockets/data"
	"github.com/stretchr/testify/require"
)

func createArgs() ArgsWebSocketsServer {
	return ArgsWebSocketsServer{
		RetryDurationInSeconds:   1,
		BlockingAckOnError:       false,
		WithAcknowledge:          false,
		URL:                      "url",
		Uint64ByteSliceConverter: uint64ByteSlice.NewBigEndianConverter(),
		Log:                      &mock.LoggerMock{},
	}
}

func TestNewWebSocketsServer(t *testing.T) {
	t.Parallel()

	t.Run("should work", func(t *testing.T) {
		args := createArgs()
		ws, err := NewWebSocketsServer(args)
		require.NotNil(t, ws)
		require.Nil(t, err)
		require.False(t, ws.IsInterfaceNil())
	})

	t.Run("empty url, should return error", func(t *testing.T) {
		args := createArgs()
		args.URL = ""
		ws, err := NewWebSocketsServer(args)
		require.Nil(t, ws)
		require.Equal(t, err, data.ErrEmptyUrl)
	})

	t.Run("nil uint64 byte slice converter, should return error", func(t *testing.T) {
		args := createArgs()
		args.Uint64ByteSliceConverter = nil
		ws, err := NewWebSocketsServer(args)
		require.Nil(t, ws)
		require.Equal(t, err, data.ErrNilUint64ByteSliceConverter)
	})

	t.Run("zero retry duration in seconds, should return error", func(t *testing.T) {
		args := createArgs()
		args.RetryDurationInSeconds = 0
		ws, err := NewWebSocketsServer(args)
		require.Nil(t, ws)
		require.Equal(t, err, data.ErrZeroValueRetryDuration)
	})
}

func TestServer_ListenAndClose(t *testing.T) {
	args := createArgs()
	args.URL = "localhost:9211"
	wsServer, _ := NewWebSocketsServer(args)

	count := uint64(0)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		wsServer.Start()
		wg.Done()
		atomic.AddUint64(&count, 1)
	}()

	_ = wsServer.Close()
	wg.Wait()
	require.Equal(t, uint64(1), atomic.LoadUint64(&count))
}

func TestServer_ListenAndRegisterPayloadHandlerAndClose(t *testing.T) {
	args := createArgs()
	args.URL = "localhost:9211"
	wsServer, _ := NewWebSocketsServer(args)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		wsServer.Start()
		wg.Done()
	}()

	_ = wsServer.SetPayloadHandler(&testscommon.PayloadHandlerStub{})
	wsServer.connectionHandler(&testscommon.WebsocketConnectionStub{
		ReadMessageCalled: func() (messageType int, payload []byte, err error) {
			return 0, nil, errors.New("local error")
		},
	})

	_ = wsServer.Close()
	wg.Wait()
}
