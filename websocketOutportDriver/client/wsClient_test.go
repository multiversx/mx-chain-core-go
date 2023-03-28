package client

import (
	"errors"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"github.com/multiversx/mx-chain-core-go/core/atomic"
	"github.com/multiversx/mx-chain-core-go/testscommon"
	"github.com/multiversx/mx-chain-core-go/websocketOutportDriver/data"
	"github.com/stretchr/testify/require"
)

func createArgs() ArgsWsClient {
	return ArgsWsClient{
		Url:                      "url",
		RetryDurationInSec:       1,
		BlockingAckOnError:       false,
		PayloadProcessor:         &testscommon.PayloadProcessorStub{},
		PayloadParser:            &testscommon.PayloadParserStub{},
		Uint64ByteSliceConverter: &testscommon.Uint64ByteSliceConverterStub{},
		WSConnClient:             &testscommon.WebsocketConnectionStub{},
	}
}

func TestNewWsClientHandler(t *testing.T) {
	t.Parallel()

	t.Run("should work", func(t *testing.T) {
		args := createArgs()
		ws, err := NewWsClientHandler(args)
		require.NotNil(t, ws)
		require.Nil(t, err)
	})

	t.Run("nil payload parser, should return error", func(t *testing.T) {
		args := createArgs()
		args.PayloadParser = nil
		ws, err := NewWsClientHandler(args)
		require.Nil(t, ws)
		require.Equal(t, err, errNilPayloadParser)
	})

	t.Run("nil payload processor, should return error", func(t *testing.T) {
		args := createArgs()
		args.PayloadProcessor = nil
		ws, err := NewWsClientHandler(args)
		require.Nil(t, ws)
		require.Equal(t, err, errNilPayloadProcessor)
	})

	t.Run("nil ws conn, should return error", func(t *testing.T) {
		args := createArgs()
		args.WSConnClient = nil
		ws, err := NewWsClientHandler(args)
		require.Nil(t, ws)
		require.Equal(t, err, errNilWsConnReceiver)
	})

	t.Run("nil uint64 byte slice converter, should return error", func(t *testing.T) {
		args := createArgs()
		args.Uint64ByteSliceConverter = nil
		ws, err := NewWsClientHandler(args)
		require.Nil(t, ws)
		require.Equal(t, err, errNilUint64ByteSliceConverter)
	})

	t.Run("empty url, should return error", func(t *testing.T) {
		args := createArgs()
		args.Url = ""
		ws, err := NewWsClientHandler(args)
		require.Nil(t, ws)
		require.Equal(t, err, errEmptyUrl)
	})

	t.Run("zero value retry duration, should return error", func(t *testing.T) {
		args := createArgs()
		args.RetryDurationInSec = 0
		ws, err := NewWsClientHandler(args)
		require.Nil(t, ws)
		require.Equal(t, err, errZeroValueRetryDuration)
	})
}

func TestClient_EmptyPayload(t *testing.T) {
	t.Parallel()

	var errClosedConn error
	mutCloseConn := sync.RWMutex{}
	readMsgCalledCt := &atomic.Counter{}

	args := createArgs()
	args.WSConnClient = &testscommon.WebsocketConnectionStub{
		ReadMessageCalled: func() (messageType int, payload []byte, err error) {
			time.Sleep(100 * time.Millisecond)
			readMsgCalledCt.Increment()

			mutCloseConn.RLock()
			err = errClosedConn
			mutCloseConn.RUnlock()

			return 0, nil, err
		},
		CloseCalled: func() error {
			mutCloseConn.Lock()
			errClosedConn = errors.New("closed connection")
			mutCloseConn.Unlock()
			return nil
		},
	}

	wsClient, _ := NewWsClientHandler(args)

	go wsClient.Start()
	time.Sleep(500 * time.Millisecond)

	wsClient.Close()
	time.Sleep(100 * time.Millisecond)
	require.Equal(t, int64(5), readMsgCalledCt.Get())
}

func TestClient_CannotExtractPayload(t *testing.T) {
	t.Parallel()

	var errClosedConn error
	mutCloseConn := sync.RWMutex{}
	readMsgCalledCt := &atomic.Counter{}

	args := createArgs()
	args.WSConnClient = &testscommon.WebsocketConnectionStub{
		ReadMessageCalled: func() (messageType int, payload []byte, err error) {
			time.Sleep(100 * time.Millisecond)
			readMsgCalledCt.Increment()

			mutCloseConn.RLock()
			err = errClosedConn
			mutCloseConn.RUnlock()

			return 0, []byte("payload"), err
		},
		CloseCalled: func() error {
			mutCloseConn.Lock()
			errClosedConn = errors.New("closed connection")
			mutCloseConn.Unlock()

			return nil
		},
	}
	args.PayloadParser = &testscommon.PayloadParserStub{
		ExtractPayloadDataCalled: func(payload []byte) (*data.PayloadData, error) {
			require.Equal(t, []byte("payload"), payload)
			return nil, fmt.Errorf("error extracting payload")
		},
	}

	wsClient, _ := NewWsClientHandler(args)

	go wsClient.Start()
	time.Sleep(500 * time.Millisecond)

	wsClient.Close()
	time.Sleep(100 * time.Millisecond)
	require.Equal(t, int64(5), readMsgCalledCt.Get())
}

func TestClient_CannotWriteAckSignal(t *testing.T) {
	t.Parallel()

	var errClosedConn error
	mutCloseConn := sync.RWMutex{}
	writeMsgCalledCt := &atomic.Counter{}

	args := createArgs()
	args.WSConnClient = &testscommon.WebsocketConnectionStub{
		ReadMessageCalled: func() (messageType int, payload []byte, err error) {
			mutCloseConn.RLock()
			err = errClosedConn
			mutCloseConn.RUnlock()

			return 0, []byte("payload"), err
		},
		WriteMessageCalled: func(messageType int, data []byte) error {
			writeMsgCalledCt.Increment()
			return fmt.Errorf("cannot write message")
		},
		CloseCalled: func() error {
			mutCloseConn.Lock()
			errClosedConn = errors.New("closed connection")
			mutCloseConn.Unlock()

			return nil
		},
	}
	args.PayloadParser = &testscommon.PayloadParserStub{
		ExtractPayloadDataCalled: func(payload []byte) (*data.PayloadData, error) {
			require.Equal(t, []byte("payload"), payload)
			return &data.PayloadData{
				WithAcknowledge: true,
			}, nil
		},
	}

	wsClient, _ := NewWsClientHandler(args)

	go wsClient.Start()
	time.Sleep(1500 * time.Millisecond)

	wsClient.Close()
	time.Sleep(100 * time.Millisecond)
	require.Equal(t, int64(2), writeMsgCalledCt.Get())
}

func TestClient_ErrorProcessingPayloadBlockingAckOnError(t *testing.T) {
	t.Parallel()

	var errClosedConn error
	mutCloseConn := sync.RWMutex{}
	readMsgCalledCt := &atomic.Counter{}
	writeMsgCalledCt := &atomic.Counter{}

	args := createArgs()
	args.BlockingAckOnError = true
	args.WSConnClient = &testscommon.WebsocketConnectionStub{
		ReadMessageCalled: func() (messageType int, payload []byte, err error) {
			time.Sleep(100 * time.Millisecond)
			readMsgCalledCt.Increment()

			mutCloseConn.RLock()
			err = errClosedConn
			mutCloseConn.RUnlock()

			return 0, []byte("payload"), err
		},
		WriteMessageCalled: func(messageType int, data []byte) error {
			writeMsgCalledCt.Increment()
			return nil
		},
		CloseCalled: func() error {
			mutCloseConn.Lock()
			errClosedConn = errors.New("closed connection")
			mutCloseConn.Unlock()

			return nil
		},
	}
	args.PayloadParser = &testscommon.PayloadParserStub{
		ExtractPayloadDataCalled: func(payload []byte) (*data.PayloadData, error) {
			require.Equal(t, []byte("payload"), payload)
			return &data.PayloadData{
				WithAcknowledge: true,
			}, nil
		},
	}
	args.PayloadProcessor = &testscommon.PayloadProcessorStub{
		ProcessPayloadCalled: func(payload *data.PayloadData) error {
			return fmt.Errorf("cannot process payload")
		},
	}

	wsClient, _ := NewWsClientHandler(args)

	go wsClient.Start()
	time.Sleep(500 * time.Millisecond)

	wsClient.Close()
	time.Sleep(100 * time.Millisecond)
	require.Equal(t, int64(0), writeMsgCalledCt.Get())
	require.Equal(t, int64(5), readMsgCalledCt.Get())
}

func TestClient_ErrorProcessingPayloadNonBlockingAckOnError(t *testing.T) {
	t.Parallel()

	var errClosedConn error
	mutCloseConn := sync.RWMutex{}
	readMsgCalledCt := &atomic.Counter{}
	writeMsgCalledCt := &atomic.Counter{}

	args := createArgs()
	args.BlockingAckOnError = false
	args.WSConnClient = &testscommon.WebsocketConnectionStub{
		ReadMessageCalled: func() (messageType int, payload []byte, err error) {
			time.Sleep(100 * time.Millisecond)
			readMsgCalledCt.Increment()

			mutCloseConn.RLock()
			err = errClosedConn
			mutCloseConn.RUnlock()

			return 0, []byte("payload"), err
		},
		WriteMessageCalled: func(messageType int, data []byte) error {
			writeMsgCalledCt.Increment()
			return nil
		},
		CloseCalled: func() error {
			mutCloseConn.Lock()
			errClosedConn = errors.New("closed connection")
			mutCloseConn.Unlock()

			return nil
		},
	}
	args.PayloadParser = &testscommon.PayloadParserStub{
		ExtractPayloadDataCalled: func(payload []byte) (*data.PayloadData, error) {
			require.Equal(t, []byte("payload"), payload)
			return &data.PayloadData{
				WithAcknowledge: true,
			}, nil
		},
	}
	args.PayloadProcessor = &testscommon.PayloadProcessorStub{
		ProcessPayloadCalled: func(payload *data.PayloadData) error {
			return fmt.Errorf("cannot process payload")
		},
	}

	wsClient, _ := NewWsClientHandler(args)

	go wsClient.Start()
	time.Sleep(500 * time.Millisecond)

	wsClient.Close()
	time.Sleep(100 * time.Millisecond)
	require.Equal(t, int64(4), writeMsgCalledCt.Get())
	require.Equal(t, int64(5), readMsgCalledCt.Get())
}

func TestClient_NormalFlowWithAck(t *testing.T) {
	t.Parallel()

	var errClosedConn error
	mutCloseConn := sync.RWMutex{}
	mutPayload := sync.RWMutex{}
	readMsgCalledCt := &atomic.Counter{}
	writeMsgCalledCt := &atomic.Counter{}

	expectedPayload := &data.PayloadData{
		WithAcknowledge: true,
		Counter:         0,
	}

	args := createArgs()
	args.WSConnClient = &testscommon.WebsocketConnectionStub{
		ReadMessageCalled: func() (messageType int, payload []byte, err error) {
			time.Sleep(100 * time.Millisecond)
			readMsgCalledCt.Increment()

			mutCloseConn.RLock()
			err = errClosedConn
			mutCloseConn.RUnlock()

			mutPayload.Lock()
			defer mutPayload.Unlock()

			expectedPayload.Payload = []byte(fmt.Sprintf("payload%d", expectedPayload.Counter))
			return 0, expectedPayload.Payload, err
		},
		WriteMessageCalled: func(messageType int, data []byte) error {
			require.Equal(t, []byte{byte(expectedPayload.Counter)}, data)
			writeMsgCalledCt.Increment()
			return nil
		},
		CloseCalled: func() error {
			mutCloseConn.Lock()
			errClosedConn = errors.New("closed connection")
			mutCloseConn.Unlock()

			return nil
		},
	}
	args.Uint64ByteSliceConverter = &testscommon.Uint64ByteSliceConverterStub{
		ToByteSliceCalled: func(num uint64) []byte {
			mutPayload.RLock()
			defer mutPayload.RUnlock()

			require.Equal(t, expectedPayload.Counter, num)
			return []byte{byte(num)}
		},
	}

	args.PayloadParser = &testscommon.PayloadParserStub{
		ExtractPayloadDataCalled: func(payload []byte) (*data.PayloadData, error) {
			mutPayload.Lock()
			defer mutPayload.Unlock()

			require.Equal(t, []byte(fmt.Sprintf("payload%d", expectedPayload.Counter)), payload)
			expectedPayload.Counter++

			return expectedPayload, nil
		},
	}
	args.PayloadProcessor = &testscommon.PayloadProcessorStub{
		ProcessPayloadCalled: func(payload *data.PayloadData) error {
			require.Equal(t, expectedPayload, payload)
			return nil
		},
	}

	wsClient, _ := NewWsClientHandler(args)

	go wsClient.Start()
	time.Sleep(500 * time.Millisecond)

	wsClient.Close()
	time.Sleep(100 * time.Millisecond)
	require.Equal(t, int64(4), writeMsgCalledCt.Get())
	require.Equal(t, int64(5), readMsgCalledCt.Get())

	mutPayload.RLock()
	defer mutPayload.RUnlock()
	require.Equal(t, &data.PayloadData{
		WithAcknowledge: true,
		Payload:         []byte("payload4"),
		Counter:         4}, expectedPayload)
}

func TestClient_NormalFlowWithoutAck(t *testing.T) {
	t.Parallel()

	var errClosedConn error
	mutCloseConn := sync.RWMutex{}
	readMsgCalledCt := &atomic.Counter{}
	writeMsgCalledCt := &atomic.Counter{}

	args := createArgs()
	args.WSConnClient = &testscommon.WebsocketConnectionStub{
		ReadMessageCalled: func() (messageType int, payload []byte, err error) {
			time.Sleep(100 * time.Millisecond)
			readMsgCalledCt.Increment()

			mutCloseConn.RLock()
			err = errClosedConn
			mutCloseConn.RUnlock()

			return 0, []byte("payload"), err
		},
		WriteMessageCalled: func(messageType int, data []byte) error {
			writeMsgCalledCt.Increment()
			return nil
		},
		CloseCalled: func() error {
			mutCloseConn.Lock()
			errClosedConn = errors.New("closed connection")
			mutCloseConn.Unlock()

			return nil
		},
	}
	expectedPayload := &data.PayloadData{WithAcknowledge: false}
	args.PayloadParser = &testscommon.PayloadParserStub{
		ExtractPayloadDataCalled: func(payload []byte) (*data.PayloadData, error) {
			require.Equal(t, []byte("payload"), payload)
			return expectedPayload, nil
		},
	}
	args.PayloadProcessor = &testscommon.PayloadProcessorStub{
		ProcessPayloadCalled: func(payload *data.PayloadData) error {
			require.Equal(t, expectedPayload, payload)
			return nil
		},
	}

	wsClient, _ := NewWsClientHandler(args)

	go wsClient.Start()
	time.Sleep(500 * time.Millisecond)

	wsClient.Close()
	time.Sleep(100 * time.Millisecond)
	require.Equal(t, int64(0), writeMsgCalledCt.Get())
	require.Equal(t, int64(5), readMsgCalledCt.Get())
}

func TestClient_ConnectionClosedFromServer(t *testing.T) {
	t.Parallel()

	readMsgCalledCt := &atomic.Counter{}
	writeMsgCalledCt := &atomic.Counter{}
	openConnectionCalledCt := &atomic.Counter{}

	args := createArgs()
	args.WSConnClient = &testscommon.WebsocketConnectionStub{
		ReadMessageCalled: func() (messageType int, payload []byte, err error) {
			time.Sleep(100 * time.Millisecond)

			if readMsgCalledCt.Get() == 3 {
				return 0, nil, errors.New(closedConnection)
			}

			readMsgCalledCt.Increment()
			return 0, []byte("payload"), err
		},
		WriteMessageCalled: func(messageType int, data []byte) error {
			writeMsgCalledCt.Increment()
			return nil
		},
		OpenConnectionCalled: func(url string) error {
			openConnectionCalledCt.Increment()
			return nil
		},
	}
	expectedPayload := &data.PayloadData{WithAcknowledge: true}
	args.PayloadParser = &testscommon.PayloadParserStub{
		ExtractPayloadDataCalled: func(payload []byte) (*data.PayloadData, error) {
			require.Equal(t, []byte("payload"), payload)
			return expectedPayload, nil
		},
	}
	args.PayloadProcessor = &testscommon.PayloadProcessorStub{
		ProcessPayloadCalled: func(payload *data.PayloadData) error {
			require.Equal(t, expectedPayload, payload)
			return nil
		},
	}

	wsClient, _ := NewWsClientHandler(args)

	go wsClient.Start()
	time.Sleep(500 * time.Millisecond)

	require.Equal(t, int64(3), readMsgCalledCt.Get())
	require.Equal(t, int64(3), writeMsgCalledCt.Get())
	require.Equal(t, int64(1), openConnectionCalledCt.Get())
}

func TestClient_CloseErrorFromServerShouldRetryOpenConnection(t *testing.T) {
	t.Parallel()

	readMsgCalledCt := &atomic.Counter{}
	writeMsgCalledCt := &atomic.Counter{}
	openConnectionCalledCt := &atomic.Counter{}

	args := createArgs()
	args.WSConnClient = &testscommon.WebsocketConnectionStub{
		ReadMessageCalled: func() (messageType int, payload []byte, err error) {
			if readMsgCalledCt.Get() == 3 {
				return 0, nil, &websocket.CloseError{}
			}

			readMsgCalledCt.Increment()
			return 0, []byte("payload"), err
		},
		WriteMessageCalled: func(messageType int, data []byte) error {
			writeMsgCalledCt.Increment()
			return nil
		},
		OpenConnectionCalled: func(url string) error {
			openConnectionCalledCt.Increment()
			return nil
		},
	}
	expectedPayload := &data.PayloadData{WithAcknowledge: true}
	args.PayloadParser = &testscommon.PayloadParserStub{
		ExtractPayloadDataCalled: func(payload []byte) (*data.PayloadData, error) {
			require.Equal(t, []byte("payload"), payload)
			return expectedPayload, nil
		},
	}
	args.PayloadProcessor = &testscommon.PayloadProcessorStub{
		ProcessPayloadCalled: func(payload *data.PayloadData) error {
			require.Equal(t, expectedPayload, payload)
			return nil
		},
	}

	wsClient, _ := NewWsClientHandler(args)

	go wsClient.Start()
	time.Sleep(1100 * time.Millisecond)

	wsClient.Close()
	time.Sleep(100 * time.Millisecond)

	require.Equal(t, int64(3), readMsgCalledCt.Get())
	require.Equal(t, int64(3), writeMsgCalledCt.Get())
	require.Equal(t, int64(2), openConnectionCalledCt.Get())
}

func TestClient_StartWaitForServerConnection(t *testing.T) {
	t.Parallel()

	readMsgCalledCt := &atomic.Counter{}
	writeMsgCalledCt := &atomic.Counter{}
	openConnectionCalledCt := &atomic.Counter{}

	args := createArgs()
	args.WSConnClient = &testscommon.WebsocketConnectionStub{
		ReadMessageCalled: func() (messageType int, payload []byte, err error) {
			if readMsgCalledCt.Get() == 5 {
				return 0, nil, fmt.Errorf(closedConnection)
			}

			readMsgCalledCt.Increment()
			return 0, []byte("payload"), err
		},
		WriteMessageCalled: func(messageType int, data []byte) error {
			writeMsgCalledCt.Increment()
			return nil
		},
		OpenConnectionCalled: func(url string) error {
			openConnectionCalledCt.Increment()
			if openConnectionCalledCt.Get() <= 2 {
				return fmt.Errorf("cannot open connection")
			}

			return nil
		},
	}
	expectedPayload := &data.PayloadData{WithAcknowledge: true}
	args.PayloadParser = &testscommon.PayloadParserStub{
		ExtractPayloadDataCalled: func(payload []byte) (*data.PayloadData, error) {
			require.Equal(t, []byte("payload"), payload)
			return expectedPayload, nil
		},
	}
	args.PayloadProcessor = &testscommon.PayloadProcessorStub{
		ProcessPayloadCalled: func(payload *data.PayloadData) error {
			require.Equal(t, expectedPayload, payload)
			return nil
		},
	}

	wsClient, _ := NewWsClientHandler(args)

	go wsClient.Start()
	time.Sleep(2050 * time.Millisecond)

	require.Equal(t, int64(5), readMsgCalledCt.Get())
	require.Equal(t, int64(5), writeMsgCalledCt.Get())
	require.Equal(t, int64(3), openConnectionCalledCt.Get())
}
