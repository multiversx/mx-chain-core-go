package sender

import (
	"sync"
	"testing"

	"github.com/gorilla/websocket"
	"github.com/multiversx/mx-chain-core-go/core"
	"github.com/multiversx/mx-chain-core-go/core/mock"
	"github.com/multiversx/mx-chain-core-go/data/typeConverters/uint64ByteSlice"
	"github.com/multiversx/mx-chain-core-go/testscommon"
	"github.com/multiversx/mx-chain-core-go/webSocket"
	"github.com/multiversx/mx-chain-core-go/webSocket/data"
	"github.com/stretchr/testify/require"
)

func createArgs() ArgsSender {
	payloadConverter, _ := webSocket.NewWebSocketPayloadConverter(uint64ByteSlice.NewBigEndianConverter())
	return ArgsSender{
		PayloadConverter:       payloadConverter,
		Log:                    &mock.LoggerStub{},
		RetryDurationInSeconds: 1,
	}
}

func TestNewSender(t *testing.T) {
	t.Parallel()

	t.Run("should work", func(t *testing.T) {
		args := createArgs()
		ws, err := NewSender(args)
		require.NotNil(t, ws)
		require.Nil(t, err)
	})

	t.Run("empty logger, should return error", func(t *testing.T) {
		args := createArgs()
		args.Log = nil
		ws, err := NewSender(args)
		require.Nil(t, ws)
		require.Equal(t, err, core.ErrNilLogger)
	})

	t.Run("nil payload converter should return error", func(t *testing.T) {
		args := createArgs()
		args.PayloadConverter = nil
		ws, err := NewSender(args)
		require.Nil(t, ws)
		require.Equal(t, err, data.ErrNilPayloadConverter)
	})

	t.Run("zero retry duration in seconds, should return error", func(t *testing.T) {
		args := createArgs()
		args.RetryDurationInSeconds = 0
		ws, err := NewSender(args)
		require.Nil(t, ws)
		require.Equal(t, err, data.ErrZeroValueRetryDuration)
	})
}

func TestSender_AddConnectionSendAndClose(t *testing.T) {
	args := createArgs()
	args.WithAcknowledge = true
	webSocketsSender, _ := NewSender(args)

	write := false
	readAck := false
	closeCalled := false
	conn1 := &testscommon.WebsocketConnectionStub{
		GetIDCalled: func() string {
			return "conn1"
		},
		WriteMessageCalled: func(messageType int, data []byte) error {
			write = true
			return nil
		},
		ReadMessageCalled: func() (messageType int, payload []byte, err error) {
			if readAck {
				counterBytes := args.PayloadConverter.EncodeUint64(1)
				return websocket.BinaryMessage, counterBytes, nil
			}

			readAck = true
			return websocket.BinaryMessage, []byte("0"), nil

		},
		CloseCalled: func() error {
			closeCalled = true
			return nil
		},
	}

	err := webSocketsSender.AddConnection(conn1)
	require.Nil(t, err)

	err = webSocketsSender.Send(data.WsSendArgs{
		Payload: []byte("something"),
	})
	require.Nil(t, err)
	require.True(t, write)
	require.True(t, readAck)

	err = webSocketsSender.Close()
	require.Nil(t, err)
	require.True(t, closeCalled)
}

func TestSender_AddConnectionSendAndWaitForAckClose(t *testing.T) {
	args := createArgs()
	args.WithAcknowledge = true
	webSocketsSender, _ := NewSender(args)

	conn1 := &testscommon.WebsocketConnectionStub{
		GetIDCalled: func() string {
			return "conn1"
		},
		ReadMessageCalled: func() (messageType int, payload []byte, err error) {
			return websocket.BinaryMessage, []byte("0"), nil

		},
		CloseCalled: func() error {
			return nil
		},
	}

	err := webSocketsSender.AddConnection(conn1)
	require.Nil(t, err)

	called := false
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		err = webSocketsSender.Send(data.WsSendArgs{
			Payload: []byte("something"),
		})
		require.Nil(t, err)
		called = true
		wg.Done()
	}()

	_ = webSocketsSender.Close()
	wg.Wait()
	require.True(t, called)
}
