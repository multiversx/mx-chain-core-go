package server

import (
	"errors"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	coreMock "github.com/multiversx/mx-chain-core-go/core/mock"
	"github.com/multiversx/mx-chain-core-go/testscommon"
	"github.com/multiversx/mx-chain-core-go/websocketOutportDriver/data"
	"github.com/stretchr/testify/require"
)

func TestNewWebSocketSender(t *testing.T) {
	t.Parallel()

	t.Run("nil uint64 byte slice converter", func(t *testing.T) {
		t.Parallel()

		args := getMockWebSocketSender()
		args.Uint64ByteSliceConverter = nil

		wss, err := NewServerSender(args)
		require.Nil(t, wss)
		require.Equal(t, data.ErrNilUint64ByteSliceConverter, err)
	})

	t.Run("nil logger", func(t *testing.T) {
		t.Parallel()

		args := getMockWebSocketSender()
		args.Log = nil

		wss, err := NewServerSender(args)
		require.Nil(t, wss)
		require.Equal(t, data.ErrNilLogger, err)
		require.True(t, wss.IsInterfaceNil())
	})

	t.Run("empty url", func(t *testing.T) {
		t.Parallel()

		args := getMockWebSocketSender()
		args.URL = ""

		wss, err := NewServerSender(args)
		require.Nil(t, wss)
		require.Equal(t, data.ErrEmptyUrl, err)
		require.True(t, wss.IsInterfaceNil())
	})

	t.Run("should work", func(t *testing.T) {
		t.Parallel()

		args := getMockWebSocketSender()

		wss, err := NewServerSender(args)
		require.NoError(t, err)
		require.NotNil(t, wss)
		require.False(t, wss.IsInterfaceNil())
	})
}

func TestWebSocketSender_AddClient(t *testing.T) {
	t.Parallel()

	t.Run("should work - without acknowledge", func(t *testing.T) {
		t.Parallel()

		wss, _ := NewServerSender(getMockWebSocketSender())

		wss.addClient(&testscommon.WebsocketConnectionStub{
			GetIDCalled: func() string {
				return "remote addr"
			},
		})

		clients := wss.clientsHolder.GetAll()
		require.NotNil(t, clients["remote addr"])

		require.False(t, wss.acknowledges.Exists("remote addr"))
	})

	t.Run("should work - with acknowledge", func(t *testing.T) {
		t.Parallel()

		args := getMockWebSocketSender()
		args.WithAcknowledge = true

		wss, _ := NewServerSender(args)

		wss.addClient(&testscommon.WebsocketConnectionStub{
			ReadMessageCalled: func() (_ int, _ []byte, err error) {
				err = errors.New("early exit - close the go routine")
				return
			},
			GetIDCalled: func() string {
				return "remote addr"
			},
		})

		clients := wss.clientsHolder.GetAll()
		require.NotNil(t, clients["remote addr"])

		require.True(t, wss.acknowledges.Exists("remote addr"))
	})
}

func TestWebSocketSender_Send(t *testing.T) {
	t.Parallel()

	t.Run("should error because no clients exist", func(t *testing.T) {
		t.Parallel()

		wss, _ := NewServerSender(getMockWebSocketSender())

		err := wss.Send(0, []byte("payload"))
		require.Equal(t, data.ErrNoClientToSendTo, err)
	})

	t.Run("should work - without acknowledge", func(t *testing.T) {
		t.Parallel()

		wss, _ := NewServerSender(getMockWebSocketSender())

		_ = wss.clientsHolder.AddClient(&testscommon.WebsocketConnectionStub{
			ReadMessageCalled: func() (_ int, _ []byte, err error) {
				err = errors.New("early exit - close the go routine")
				return
			},
			GetIDCalled: func() string {
				return "remove addr"
			},
		})

		err := wss.Send(0, []byte("payload"))
		require.NoError(t, err)
	})

	t.Run("should work - with acknowledge", func(t *testing.T) {
		t.Parallel()

		args := getMockWebSocketSender()
		args.WithAcknowledge = true
		wss, _ := NewServerSender(args)

		var ack []byte

		chClientAck := make(chan bool)
		wasMsgProcessed := false

		client1 := &testscommon.WebsocketConnectionStub{
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
			GetIDCalled: func() string {
				return "remote addr"
			},
		}

		wss.addClient(client1)

		err := wss.Send(0, []byte("payload"))
		require.NoError(t, err)
	})
}

func TestWebSocketSender_Close(t *testing.T) {
	t.Parallel()

	wss, _ := NewServerSender(getMockWebSocketSender())

	err := wss.Close()
	require.NoError(t, err)
}

func getMockWebSocketSender() ArgsServerSender {
	return ArgsServerSender{
		Uint64ByteSliceConverter: &testscommon.Uint64ByteSliceConverterStub{},
		Log:                      coreMock.LoggerMock{},
		URL:                      "local",
	}
}
