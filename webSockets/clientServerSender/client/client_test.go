package client

import (
	"errors"
	"sync"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"github.com/multiversx/mx-chain-core-go/core/mock"
	"github.com/multiversx/mx-chain-core-go/data/typeConverters/uint64ByteSlice"
	"github.com/multiversx/mx-chain-core-go/testscommon"
	"github.com/multiversx/mx-chain-core-go/webSockets/data"
	"github.com/stretchr/testify/require"
)

func createArgs() ArgsWsClientSender {
	return ArgsWsClientSender{
		URL:                      "url",
		RetryDurationInSec:       1,
		Uint64ByteSliceConverter: &testscommon.Uint64ByteSliceConverterStub{},
		Log:                      &mock.LoggerMock{},
	}
}

func TestNewWsClient(t *testing.T) {
	t.Parallel()

	t.Run("should work", func(t *testing.T) {
		args := createArgs()
		ws, err := NewClientSender(args)
		require.NotNil(t, ws)
		require.Nil(t, err)
	})

	t.Run("nil uint64 byte slice converter, should return error", func(t *testing.T) {
		args := createArgs()
		args.Uint64ByteSliceConverter = nil
		ws, err := NewClientSender(args)
		require.Nil(t, ws)
		require.Equal(t, err, data.ErrNilUint64ByteSliceConverter)
	})
	t.Run("nil logger, should return error", func(t *testing.T) {
		args := createArgs()
		args.Log = nil
		ws, err := NewClientSender(args)
		require.Nil(t, ws)
		require.Equal(t, err, data.ErrNilLogger)
	})

	t.Run("empty url, should return error", func(t *testing.T) {
		args := createArgs()
		args.URL = ""
		ws, err := NewClientSender(args)
		require.Nil(t, ws)
		require.Equal(t, err, data.ErrEmptyUrl)
	})

	t.Run("zero value retry duration, should return error", func(t *testing.T) {
		args := createArgs()
		args.RetryDurationInSec = 0
		ws, err := NewClientSender(args)
		require.Nil(t, ws)
		require.Equal(t, err, data.ErrZeroValueRetryDuration)
	})
}

func TestClientSender_SendNotConnectedToServer(t *testing.T) {
	args := createArgs()
	args.URL = "127.0.0.1:21113"
	cSender, _ := NewClientSender(args)

	called := false
	wg := sync.WaitGroup{}
	wg.Add(1)
	cSender.wsConn = &testscommon.WebsocketConnectionStub{
		OpenConnectionCalled: func(url string) error {
			return nil
		},
		WriteMessageCalled: func(messageType int, data []byte) error {
			called = true
			wg.Done()
			return nil
		},
	}

	go func() {
		_ = cSender.Send(0, []byte("message"))
	}()

	time.Sleep(100 * time.Millisecond)

	wg.Wait()

	require.True(t, called)
}

func TestClientSender_SendMessageNoConnectionAndClose(t *testing.T) {
	args := createArgs()
	args.URL = "127.0.0.1:21112"
	cSender, _ := NewClientSender(args)

	called := false
	cSender.wsConn = &testscommon.WebsocketConnectionStub{
		WriteMessageCalled: func(_ int, _ []byte) error {
			return errors.New("error")
		},
		CloseCalled: func() error {
			called = true
			return nil
		},
	}

	go func() {
		time.Sleep(1*time.Second + 500*time.Millisecond)
		err := cSender.Close()
		require.Nil(t, err)
	}()

	err := cSender.Send(0, []byte("message"))
	require.Nil(t, err)
	require.True(t, called)
}

func TestClientSender_SendMessageAndWaitForAck(t *testing.T) {
	args := createArgs()
	args.Uint64ByteSliceConverter = uint64ByteSlice.NewBigEndianConverter()
	args.URL = "127.0.0.1:21112"
	args.WithAcknowledge = true
	cSender, _ := NewClientSender(args)

	counter := uint64(3)
	newCounter := uint64(0)
	cSender.wsConn = &testscommon.WebsocketConnectionStub{
		WriteMessageCalled: func(messageType int, data []byte) error {
			return nil
		},
		ReadMessageCalled: func() (messageType int, payload []byte, err error) {
			newCounter++
			return websocket.BinaryMessage, args.Uint64ByteSliceConverter.ToByteSlice(newCounter), nil
		},
	}

	err := cSender.Send(counter, []byte("something"))
	require.Nil(t, err)
	require.Equal(t, counter, newCounter)
}
