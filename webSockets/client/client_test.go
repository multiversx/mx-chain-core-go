package client

import (
	"errors"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/multiversx/mx-chain-core-go/core/mock"
	"github.com/multiversx/mx-chain-core-go/data/typeConverters/uint64ByteSlice"
	"github.com/multiversx/mx-chain-core-go/testscommon"
	"github.com/multiversx/mx-chain-core-go/webSockets/data"
	"github.com/stretchr/testify/require"
)

func createArgs() ArgsWebSocketsClient {
	return ArgsWebSocketsClient{
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
		ws, err := NewWebSocketsClient(args)
		require.NotNil(t, ws)
		require.Nil(t, err)
		require.False(t, ws.IsInterfaceNil())
	})

	t.Run("empty url, should return error", func(t *testing.T) {
		args := createArgs()
		args.URL = ""
		ws, err := NewWebSocketsClient(args)
		require.Nil(t, ws)
		require.Equal(t, err, data.ErrEmptyUrl)
	})

	t.Run("nil uint64 byte slice converter, should return error", func(t *testing.T) {
		args := createArgs()
		args.Uint64ByteSliceConverter = nil
		ws, err := NewWebSocketsClient(args)
		require.Nil(t, ws)
		require.Equal(t, err, data.ErrNilUint64ByteSliceConverter)
	})

	t.Run("zero retry duration in seconds, should return error", func(t *testing.T) {
		args := createArgs()
		args.RetryDurationInSeconds = 0
		ws, err := NewWebSocketsClient(args)
		require.Nil(t, ws)
		require.Equal(t, err, data.ErrZeroValueRetryDuration)
	})
}

func TestClient_SendAndClose(t *testing.T) {
	args := createArgs()
	ws, err := NewWebSocketsClient(args)
	require.Nil(t, err)

	mockConn := &testscommon.WebsocketConnectionStub{
		WriteMessageCalled: func(messageType int, _ []byte) error {
			return errors.New(data.ClosedConnectionMessage)
		},
	}
	ws.wsConn = mockConn
	_ = ws.sender.AddConnection(mockConn)

	count := uint64(0)

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		err = ws.Send(data.WsSendArgs{
			Payload: []byte("send"),
		})
		require.Equal(t, "data wasn't sent to any client. last known error: use of closed network connection", err.Error())
		atomic.AddUint64(&count, 1)
	}()

	_ = ws.Close()
	wg.Wait()
	require.Equal(t, uint64(1), atomic.LoadUint64(&count))
}

func TestClient_Send(t *testing.T) {
	args := createArgs()
	ws, err := NewWebSocketsClient(args)
	require.Nil(t, err)

	mockConn := &testscommon.WebsocketConnectionStub{
		WriteMessageCalled: func(messageType int, _ []byte) error {
			return errors.New("local error")
		},
	}

	ws.wsConn = mockConn
	_ = ws.sender.AddConnection(mockConn)

	count := uint64(0)
	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		err = ws.Send(data.WsSendArgs{Payload: []byte("test")})
		require.Equal(t, "data wasn't sent to any client. last known error: local error", err.Error())
		atomic.AddUint64(&count, 1)
	}()

	time.Sleep(2 * time.Second)
	_ = ws.Close()
	wg.Wait()

	require.Equal(t, uint64(1), atomic.LoadUint64(&count))
}
