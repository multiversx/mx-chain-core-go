package transceiver

import (
	"errors"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"github.com/multiversx/mx-chain-core-go/core"
	"github.com/multiversx/mx-chain-core-go/core/mock"
	"github.com/multiversx/mx-chain-core-go/data/typeConverters/uint64ByteSlice"
	"github.com/multiversx/mx-chain-core-go/testscommon"
	"github.com/multiversx/mx-chain-core-go/webSocket"
	"github.com/multiversx/mx-chain-core-go/webSocket/data"
	"github.com/stretchr/testify/require"
)

func createArgs() ArgsTransceiver {
	payloadConverter, _ := webSocket.NewWebSocketPayloadConverter(uint64ByteSlice.NewBigEndianConverter())
	return ArgsTransceiver{
		BlockingAckOnError: false,
		PayloadConverter:   payloadConverter,
		Log:                &mock.LoggerMock{},
		RetryDurationInSec: 1,
		WithAcknowledge:    false,
	}
}

func TestNewReceiver(t *testing.T) {
	t.Parallel()

	t.Run("should work", func(t *testing.T) {
		args := createArgs()
		ws, err := NewTransceiver(args)
		require.NotNil(t, ws)
		require.Nil(t, err)
	})

	t.Run("empty logger, should return error", func(t *testing.T) {
		args := createArgs()
		args.Log = nil
		ws, err := NewTransceiver(args)
		require.Nil(t, ws)
		require.Equal(t, core.ErrNilLogger, err)
	})

	t.Run("nil payload converter, should return error", func(t *testing.T) {
		args := createArgs()
		args.PayloadConverter = nil
		ws, err := NewTransceiver(args)
		require.Nil(t, ws)
		require.Equal(t, data.ErrNilPayloadConverter, err)
	})

	t.Run("zero retry duration in seconds, should return error", func(t *testing.T) {
		args := createArgs()
		args.RetryDurationInSec = 0
		ws, err := NewTransceiver(args)
		require.Nil(t, ws)
		require.Equal(t, data.ErrZeroValueRetryDuration, err)
	})
}

func TestReceiver_ListenAndClose(t *testing.T) {
	args := createArgs()
	webSocketsReceiver, err := NewTransceiver(args)
	require.Nil(t, err)

	count := uint64(0)
	conn := &testscommon.WebsocketConnectionStub{
		ReadMessageCalled: func() (messageType int, payload []byte, err error) {
			time.Sleep(time.Second)
			if atomic.LoadUint64(&count) == 1 {
				return 0, nil, errors.New("closed")
			}
			return 0, nil, nil
		},
		CloseCalled: func() error {
			atomic.AddUint64(&count, 1)
			return nil
		},
	}
	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		webSocketsReceiver.Listen(conn)
		wg.Done()
		atomic.AddUint64(&count, 1)
	}()

	_ = webSocketsReceiver.Close()
	_ = conn.Close()
	wg.Wait()

	require.Equal(t, uint64(2), atomic.LoadUint64(&count))
}

func TestReceiver_ListenAndSendAck(t *testing.T) {
	args := createArgs()
	webSocketsReceiver, err := NewTransceiver(args)
	require.Nil(t, err)

	_ = webSocketsReceiver.SetPayloadHandler(&testscommon.PayloadHandlerStub{
		ProcessPayloadCalled: func(payload []byte) error {
			return nil
		},
	})

	wg := &sync.WaitGroup{}
	wg.Add(2)

	count := 0
	conn := &testscommon.WebsocketConnectionStub{
		ReadMessageCalled: func() (messageType int, payload []byte, err error) {
			time.Sleep(500 * time.Millisecond)
			if count >= 1 {
				wg.Done()
				return 0, nil, errors.New("closed")
			}
			count++
			preparedPayload := args.PayloadConverter.ConstructPayloadData(data.WsSendArgs{
				Payload: []byte("something"),
				OpType:  data.OperationSaveAccounts,
			}, 10, true)
			return websocket.BinaryMessage, preparedPayload, nil
		},
		CloseCalled: func() error {
			return nil
		},
		WriteMessageCalled: func(messageType int, data []byte) error {
			if count == 1 {
				count++
				return errors.New("local error")
			}
			return nil
		},
	}

	go func() {
		webSocketsReceiver.Listen(conn)
	}()

	wg.Wait()
	_ = webSocketsReceiver.Close()
	_ = conn.Close()

	require.Equal(t, 2, count)
}

func TestSender_AddConnectionSendAndClose(t *testing.T) {
	args := createArgs()
	args.WithAcknowledge = true
	webSocketTransceiver, _ := NewTransceiver(args)

	write := false
	readAck := false
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
				return websocket.TextMessage, counterBytes, nil
			}

			readAck = true
			return websocket.BinaryMessage, []byte("0"), nil

		},
	}

	go func() {
		webSocketTransceiver.Listen(conn1)
	}()

	err := webSocketTransceiver.Send(data.WsSendArgs{
		Payload: []byte("something"),
	}, conn1)
	require.Nil(t, err)
	require.True(t, write)
	require.True(t, readAck)

	err = webSocketTransceiver.Close()
	require.Nil(t, err)
}

func TestSender_AddConnectionSendAndWaitForAckClose(t *testing.T) {
	args := createArgs()
	args.WithAcknowledge = true
	webSocketTransceiver, _ := NewTransceiver(args)

	conn1 := &testscommon.WebsocketConnectionStub{
		GetIDCalled: func() string {
			return "conn1"
		},
		ReadMessageCalled: func() (messageType int, payload []byte, err error) {
			time.Sleep(50 * time.Millisecond)
			return websocket.BinaryMessage, []byte("0"), nil

		},
		CloseCalled: func() error {
			return nil
		},
	}

	called := false
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		err := webSocketTransceiver.Send(data.WsSendArgs{
			Payload: []byte("something"),
		}, conn1)
		require.Nil(t, err)
		called = true
		wg.Done()
	}()

	time.Sleep(100 * time.Millisecond)
	go func() {
		webSocketTransceiver.Listen(conn1)
	}()

	_ = webSocketTransceiver.Close()
	wg.Wait()
	require.True(t, called)
}

func TestWsTransceiverWaitForAck(t *testing.T) {
	args := createArgs()
	args.WithAcknowledge = true
	webSocketTransceiver, _ := NewTransceiver(args)

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		webSocketTransceiver.waitForAck()
		wg.Done()
	}()

	time.Sleep(time.Second)
	err := webSocketTransceiver.Close()
	require.Nil(t, err)

	wg.Wait()
}
