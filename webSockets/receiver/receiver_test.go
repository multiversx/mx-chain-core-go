package receiver

import (
	"errors"
	"sync"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"github.com/multiversx/mx-chain-core-go/core"
	"github.com/multiversx/mx-chain-core-go/core/mock"
	"github.com/multiversx/mx-chain-core-go/data/typeConverters/uint64ByteSlice"
	"github.com/multiversx/mx-chain-core-go/testscommon"
	"github.com/multiversx/mx-chain-core-go/webSockets/data"
	"github.com/stretchr/testify/require"
)

func createArgs() ArgsReceiver {
	return ArgsReceiver{
		BlockingAckOnError:       false,
		Uint64ByteSliceConverter: uint64ByteSlice.NewBigEndianConverter(),
		Log:                      &mock.LoggerMock{},
		RetryDurationInSec:       1,
	}
}

func TestNewReceiver(t *testing.T) {
	t.Parallel()

	t.Run("should work", func(t *testing.T) {
		args := createArgs()
		ws, err := NewReceiver(args)
		require.NotNil(t, ws)
		require.Nil(t, err)
	})

	t.Run("empty logger, should return error", func(t *testing.T) {
		args := createArgs()
		args.Log = nil
		ws, err := NewReceiver(args)
		require.Nil(t, ws)
		require.Equal(t, err, core.ErrNilLogger)
	})

	t.Run("nil uint64 byte slice converter, should return error", func(t *testing.T) {
		args := createArgs()
		args.Uint64ByteSliceConverter = nil
		ws, err := NewReceiver(args)
		require.Nil(t, ws)
		require.Equal(t, err, data.ErrNilUint64ByteSliceConverter)
	})

	t.Run("zero retry duration in seconds, should return error", func(t *testing.T) {
		args := createArgs()
		args.RetryDurationInSec = 0
		ws, err := NewReceiver(args)
		require.Nil(t, ws)
		require.Equal(t, err, data.ErrZeroValueRetryDuration)
	})
}

func TestReceiver_ListenAndClose(t *testing.T) {
	args := createArgs()
	webSocketsReceiver, err := NewReceiver(args)
	require.Nil(t, err)

	count := 0
	conn := &testscommon.WebsocketConnectionStub{
		ReadMessageCalled: func() (messageType int, payload []byte, err error) {
			time.Sleep(time.Second)
			if count == 1 {
				return 0, nil, errors.New("closed")
			}
			return 0, nil, nil
		},
		CloseCalled: func() error {
			count++
			return nil
		},
	}
	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		webSocketsReceiver.Listen(conn)
		wg.Done()
		count++
	}()

	_ = webSocketsReceiver.Close()
	_ = conn.Close()
	wg.Wait()

	require.Equal(t, 2, count)
}

func TestReceiver_ListenAndSendAck(t *testing.T) {
	args := createArgs()
	webSocketsReceiver, err := NewReceiver(args)
	require.Nil(t, err)

	webSocketsReceiver.SetPayloadHandler(&testscommon.PayloadHandlerStub{
		HandlePayloadCalled: func(payload []byte) (*data.PayloadData, error) {
			return &data.PayloadData{
				WithAcknowledge: true,
				Counter:         10,
			}, nil
		},
	})

	wg := &sync.WaitGroup{}
	wg.Add(1)

	count := 0
	conn := &testscommon.WebsocketConnectionStub{
		ReadMessageCalled: func() (messageType int, payload []byte, err error) {
			time.Sleep(500 * time.Millisecond)
			if count >= 1 {
				wg.Done()
				return 0, nil, errors.New("closed")
			}
			count++
			return websocket.BinaryMessage, []byte("something"), nil
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