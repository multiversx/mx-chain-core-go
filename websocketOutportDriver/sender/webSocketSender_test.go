package sender

import (
	"errors"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	coreMock "github.com/multiversx/mx-chain-core-go/core/mock"
	"github.com/multiversx/mx-chain-core-go/testscommon"
	"github.com/multiversx/mx-chain-core-go/websocketOutportDriver/data"
	"github.com/multiversx/mx-chain-core-go/websocketOutportDriver/mock"
	"github.com/stretchr/testify/require"
)

func TestNewWebSocketSender(t *testing.T) {
	t.Parallel()

	t.Run("nil server", func(t *testing.T) {
		t.Parallel()

		args := getMockWebSocketSender()
		args.Server = nil

		wss, err := NewWebSocketSender(args)
		require.Nil(t, wss)
		require.Equal(t, data.ErrNilHttpServer, err)
	})

	t.Run("nil uint64 byte slice converter", func(t *testing.T) {
		t.Parallel()

		args := getMockWebSocketSender()
		args.Uint64ByteSliceConverter = nil

		wss, err := NewWebSocketSender(args)
		require.Nil(t, wss)
		require.Equal(t, data.ErrNilUint64ByteSliceConverter, err)
	})

	t.Run("nil logger", func(t *testing.T) {
		t.Parallel()

		args := getMockWebSocketSender()
		args.Log = nil

		wss, err := NewWebSocketSender(args)
		require.Nil(t, wss)
		require.Equal(t, data.ErrNilLogger, err)
		require.True(t, wss.IsInterfaceNil())
	})

	t.Run("should work", func(t *testing.T) {
		t.Parallel()

		args := getMockWebSocketSender()

		wss, err := NewWebSocketSender(args)
		require.NoError(t, err)
		require.NotNil(t, wss)
		require.False(t, wss.IsInterfaceNil())
	})
}

func TestWebSocketSender_AddClient(t *testing.T) {
	t.Parallel()

	t.Run("nil client", func(t *testing.T) {
		t.Parallel()

		wss, _ := NewWebSocketSender(getMockWebSocketSender())

		wss.AddClient(nil, "remote addr")
		require.Equal(t, 0, len(wss.clientsHolder.GetAll()))
	})

	t.Run("should work - without acknowledge", func(t *testing.T) {
		t.Parallel()

		wss, _ := NewWebSocketSender(getMockWebSocketSender())

		wss.AddClient(&testscommon.WebsocketConnectionStub{}, "remote addr")

		clients := wss.clientsHolder.GetAll()
		require.NotNil(t, clients["remote addr"])

		wss.acknowledges.mut.Lock()
		acksForAddress := wss.acknowledges.acknowledges["remote addr"]
		wss.acknowledges.mut.Unlock()

		require.Nil(t, acksForAddress)
	})

	t.Run("should work - with acknowledge", func(t *testing.T) {
		t.Parallel()

		args := getMockWebSocketSender()
		args.WithAcknowledge = true

		wss, _ := NewWebSocketSender(args)

		wss.AddClient(&testscommon.WebsocketConnectionStub{
			ReadMessageCalled: func() (_ int, _ []byte, err error) {
				err = errors.New("early exit - close the go routine")
				return
			},
		}, "remote addr")

		clients := wss.clientsHolder.GetAll()
		require.NotNil(t, clients["remote addr"])

		wss.acknowledges.mut.Lock()
		acksForAddress := wss.acknowledges.acknowledges["remote addr"]
		wss.acknowledges.mut.Unlock()

		require.NotNil(t, acksForAddress)
	})
}

func TestWebSocketSender_Send(t *testing.T) {
	t.Parallel()

	t.Run("should error because no clients exist", func(t *testing.T) {
		t.Parallel()

		wss, _ := NewWebSocketSender(getMockWebSocketSender())

		err := wss.Send(data.WsSendArgs{
			Payload: []byte("payload"),
		})
		require.Equal(t, data.ErrNoClientToSendTo, err)
	})

	t.Run("should work - without acknowledge", func(t *testing.T) {
		t.Parallel()

		wss, _ := NewWebSocketSender(getMockWebSocketSender())

		wss.AddClient(&testscommon.WebsocketConnectionStub{
			ReadMessageCalled: func() (_ int, _ []byte, err error) {
				err = errors.New("early exit - close the go routine")
				return
			},
		}, "remote addr")

		err := wss.Send(data.WsSendArgs{
			Payload: []byte("payload"),
		})
		require.NoError(t, err)
	})

	t.Run("should work - with acknowledge", func(t *testing.T) {
		t.Parallel()

		args := getMockWebSocketSender()
		args.WithAcknowledge = true
		wss, _ := NewWebSocketSender(args)

		var ack []byte

		chClientAck := make(chan bool)
		wasMsgProcessed := false

		wss.AddClient(&testscommon.WebsocketConnectionStub{
			ReadMessageCalled: func() (msgType int, payload []byte, err error) {
				if wasMsgProcessed {
					time.Sleep(100 * time.Millisecond)
					msgType = websocket.BinaryMessage
					err = errors.New("end")
					return
				}

				<-chClientAck

				time.Sleep(100 * time.Millisecond)

				msgType = websocket.BinaryMessage
				payload = ack
				err = nil
				wasMsgProcessed = true

				return
			},
			WriteMessageCalled: func(_ int, data []byte) error {
				ack = data[1:3]
				chClientAck <- true

				return nil
			},
		}, "remote addr")

		err := wss.Send(data.WsSendArgs{
			Payload: []byte("payload"),
		})
		require.NoError(t, err)
	})
}

func TestWebSocketSender_Close(t *testing.T) {
	t.Parallel()

	wss, _ := NewWebSocketSender(getMockWebSocketSender())

	err := wss.Close()
	require.NoError(t, err)
}

func getMockWebSocketSender() WebSocketSenderArgs {
	return WebSocketSenderArgs{
		Server:                   &mock.HttpServerStub{},
		Uint64ByteSliceConverter: &testscommon.Uint64ByteSliceConverterStub{},
		Log:                      coreMock.LoggerMock{},
	}
}
