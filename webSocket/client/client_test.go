package client

import (
	"errors"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/multiversx/mx-chain-core-go/core/mock"
	"github.com/multiversx/mx-chain-core-go/testscommon"
	"github.com/multiversx/mx-chain-core-go/webSocket"
	"github.com/multiversx/mx-chain-core-go/webSocket/data"
	"github.com/stretchr/testify/require"
)

func createArgs() ArgsWebSocketClient {
	payloadConverter, _ := webSocket.NewWebSocketPayloadConverter(&mock.MarshalizerMock{})
	return ArgsWebSocketClient{
		RetryDurationInSeconds: 1,
		BlockingAckOnError:     false,
		WithAcknowledge:        false,
		URL:                    "url",
		PayloadConverter:       payloadConverter,
		Log:                    &mock.LoggerMock{},
	}
}

func TestNewWebSocketServer(t *testing.T) {
	t.Parallel()

	t.Run("should work", func(t *testing.T) {
		args := createArgs()
		ws, err := NewWebSocketClient(args)
		require.NotNil(t, ws)
		require.Nil(t, err)
		require.False(t, ws.IsInterfaceNil())
	})

	t.Run("empty url, should return error", func(t *testing.T) {
		args := createArgs()
		args.URL = ""
		ws, err := NewWebSocketClient(args)
		require.Nil(t, ws)
		require.Equal(t, err, data.ErrEmptyUrl)
	})

	t.Run("nil payload converter, should return error", func(t *testing.T) {
		args := createArgs()
		args.PayloadConverter = nil
		ws, err := NewWebSocketClient(args)
		require.Nil(t, ws)
		require.Equal(t, data.ErrNilPayloadConverter, err)
	})

	t.Run("zero retry duration in seconds, should return error", func(t *testing.T) {
		args := createArgs()
		args.RetryDurationInSeconds = 0
		ws, err := NewWebSocketClient(args)
		require.Nil(t, ws)
		require.Equal(t, data.ErrZeroValueRetryDuration, err)
	})
}

func TestClient_SendAndClose(t *testing.T) {
	args := createArgs()
	ws, err := NewWebSocketClient(args)
	require.Nil(t, err)

	mockConn := &testscommon.WebsocketConnectionStub{
		WriteMessageCalled: func(messageType int, _ []byte) error {
			return errors.New(data.ClosedConnectionMessage)
		},
	}
	ws.wsConn = mockConn

	count := uint64(0)

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		err = ws.Send(data.WsMessage{
			Payload: []byte("send"),
		})
		require.Equal(t, "use of closed network connection", err.Error())
		atomic.AddUint64(&count, 1)
	}()

	_ = ws.Close()
	wg.Wait()
	require.Equal(t, uint64(1), atomic.LoadUint64(&count))
}

func TestClient_Send(t *testing.T) {
	args := createArgs()
	ws, err := NewWebSocketClient(args)
	require.Nil(t, err)

	mockConn := &testscommon.WebsocketConnectionStub{
		WriteMessageCalled: func(messageType int, _ []byte) error {
			return errors.New("local error")
		},
	}

	ws.wsConn = mockConn

	count := uint64(0)
	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		err = ws.Send(data.WsMessage{Payload: []byte("test")})
		require.Equal(t, "local error", err.Error())
		atomic.AddUint64(&count, 1)
	}()

	time.Sleep(2 * time.Second)
	_ = ws.Close()
	wg.Wait()

	require.Equal(t, uint64(1), atomic.LoadUint64(&count))
}
