package websocketOutportDriver

import (
	"errors"
	"testing"

	coreMock "github.com/ElrondNetwork/elrond-go-core/core/mock"
	"github.com/ElrondNetwork/elrond-go-core/data/indexer"
	"github.com/ElrondNetwork/elrond-go-core/data/typeConverters/uint64ByteSlice"
	"github.com/ElrondNetwork/elrond-go-core/marshal"
	"github.com/ElrondNetwork/elrond-go-core/websocketOutportDriver/data"
	"github.com/ElrondNetwork/elrond-go-core/websocketOutportDriver/mock"
	"github.com/stretchr/testify/require"
)

func TestNewWebsocketOutportDriverNodePart(t *testing.T) {
	t.Parallel()

	t.Run("nil marshalizer", func(t *testing.T) {
		args := getMockArgs()
		args.Marshaller = nil

		o, err := NewWebsocketOutportDriverNodePart(args)
		require.Nil(t, o)
		require.Equal(t, data.ErrNilMarshalizer, err)
	})

	t.Run("nil logger", func(t *testing.T) {
		args := getMockArgs()
		args.Log = nil

		o, err := NewWebsocketOutportDriverNodePart(args)
		require.Nil(t, o)
		require.Equal(t, data.ErrNilLogger, err)
	})

	t.Run("should work", func(t *testing.T) {
		args := getMockArgs()

		o, err := NewWebsocketOutportDriverNodePart(args)
		require.NotNil(t, o)
		require.NoError(t, err)
	})
}

func TestWebsocketOutportDriverNodePart_SaveBlock_ErrWhileSendingOnRoute(t *testing.T) {
	t.Skip("this test will run continuously so it should be run only when trying to reach that code area")

	expectedErr := errors.New("cannot send on route")
	defer func() {
		r := recover()
		require.Contains(t, r, expectedErr.Error())
	}()

	args := getMockArgs()
	args.WebsocketSender = &mock.WebSocketSenderStub{
		SendOnRouteCalled: func(_ data.WsSendArgs) error {
			return expectedErr
		},
	}
	o, err := NewWebsocketOutportDriverNodePart(args)
	require.NoError(t, err)

	err = o.SaveBlock(&indexer.ArgsSaveBlockData{})
	require.NoError(t, err)
}

func TestWebsocketOutportDriverNodePart_SaveBlock_ShouldWork(t *testing.T) {
	t.Parallel()

	defer func() {
		r := recover()
		require.Nil(t, r)
	}()
	args := getMockArgs()
	o, err := NewWebsocketOutportDriverNodePart(args)
	require.NoError(t, err)

	err = o.SaveBlock(&indexer.ArgsSaveBlockData{})
	require.NoError(t, err)
}

func TestWebsocketOutportDriverNodePart_SaveBlock_PayloadCheck(t *testing.T) {
	t.Parallel()

	args := getMockArgs()

	marshaledData, _ := args.Marshaller.Marshal(&indexer.ArgsSaveBlockData{})

	args.WebsocketSender = &mock.WebSocketSenderStub{
		SendOnRouteCalled: func(args data.WsSendArgs) error {
			expectedOpBytes := []byte{0, 0, 0, 0}
			expectedLengthBytes := []byte{0, 0, 0, 214} // json serialized empty ArgsSaveBlockData has 214 bytes
			expectedPayload := append(expectedOpBytes, expectedLengthBytes...)
			expectedPayload = append(expectedPayload, marshaledData...)

			require.Equal(t, expectedPayload, args.Payload)

			return nil
		},
	}
	o, err := NewWebsocketOutportDriverNodePart(args)
	require.NoError(t, err)

	err = o.SaveBlock(&indexer.ArgsSaveBlockData{})
	require.NoError(t, err)
}

func getMockArgs() WebsocketOutportDriverNodePartArgs {
	return WebsocketOutportDriverNodePartArgs{
		Enabled:    true,
		Marshaller: &marshal.JsonMarshalizer{},
		WebSocketConfig: data.WebSocketConfig{
			URL: "localhost:5555",
		},
		WebsocketSender:          &mock.WebSocketSenderStub{},
		Log:                      &coreMock.LoggerMock{},
		Uint64ByteSliceConverter: uint64ByteSlice.NewBigEndianConverter(),
	}
}
