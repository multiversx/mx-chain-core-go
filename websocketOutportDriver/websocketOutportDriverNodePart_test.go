package websocketOutportDriver

import (
	"errors"
	"testing"

	coreMock "github.com/multiversx/mx-chain-core-go/core/mock"
	"github.com/multiversx/mx-chain-core-go/data/block"
	"github.com/multiversx/mx-chain-core-go/data/outport"
	"github.com/multiversx/mx-chain-core-go/data/typeConverters/uint64ByteSlice"
	"github.com/multiversx/mx-chain-core-go/marshal"
	"github.com/multiversx/mx-chain-core-go/websocketOutportDriver/data"
	"github.com/multiversx/mx-chain-core-go/websocketOutportDriver/mock"
	"github.com/stretchr/testify/require"
)

var cannotSendOnRouteErr = errors.New("cannot send on route")

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

func TestNewWebsocketOutportDriverNodePart(t *testing.T) {
	t.Parallel()

	t.Run("nil marshaller", func(t *testing.T) {
		t.Parallel()

		args := getMockArgs()
		args.Marshaller = nil

		o, err := NewWebsocketOutportDriverNodePart(args)
		require.Nil(t, o)
		require.Equal(t, data.ErrNilMarshaller, err)
	})

	t.Run("nil uint64 byte slice converter", func(t *testing.T) {
		t.Parallel()

		args := getMockArgs()
		args.Uint64ByteSliceConverter = nil

		o, err := NewWebsocketOutportDriverNodePart(args)
		require.Nil(t, o)
		require.Equal(t, data.ErrNilUint64ByteSliceConverter, err)
	})

	t.Run("nil uint64 byte slice converter", func(t *testing.T) {
		t.Parallel()

		args := getMockArgs()
		args.Uint64ByteSliceConverter = nil

		o, err := NewWebsocketOutportDriverNodePart(args)
		require.Nil(t, o)
		require.Equal(t, data.ErrNilUint64ByteSliceConverter, err)
	})

	t.Run("nil logger", func(t *testing.T) {
		t.Parallel()

		args := getMockArgs()
		args.Log = nil

		o, err := NewWebsocketOutportDriverNodePart(args)
		require.Nil(t, o)
		require.Equal(t, data.ErrNilLogger, err)
	})

	t.Run("should work", func(t *testing.T) {
		t.Parallel()

		args := getMockArgs()

		o, err := NewWebsocketOutportDriverNodePart(args)
		require.NotNil(t, o)
		require.NoError(t, err)
		require.False(t, o.IsInterfaceNil())
	})
}

func TestWebsocketOutportDriverNodePart_SaveBlock(t *testing.T) {
	t.Parallel()

	t.Run("SaveBlock - should error", func(t *testing.T) {
		t.Parallel()

		args := getMockArgs()
		args.WebsocketSender = &mock.WebSocketSenderStub{
			SendOnRouteCalled: func(_ data.WsSendArgs) error {
				return cannotSendOnRouteErr
			},
		}
		o, err := NewWebsocketOutportDriverNodePart(args)
		require.NoError(t, err)

		err = o.SaveBlock(&outport.OutportBlock{})
		require.True(t, errors.Is(err, cannotSendOnRouteErr))
	})

	t.Run("SaveBlock - should work", func(t *testing.T) {
		t.Parallel()

		defer func() {
			r := recover()
			require.Nil(t, r)
		}()
		args := getMockArgs()
		o, err := NewWebsocketOutportDriverNodePart(args)
		require.NoError(t, err)

		err = o.SaveBlock(&outport.OutportBlock{})
		require.NoError(t, err)
	})
}

func TestWebsocketOutportDriverNodePart_FinalizedBlock(t *testing.T) {
	t.Parallel()

	t.Run("Finalized block - should error", func(t *testing.T) {
		args := getMockArgs()
		args.WebsocketSender = &mock.WebSocketSenderStub{
			SendOnRouteCalled: func(_ data.WsSendArgs) error {
				return cannotSendOnRouteErr
			},
		}
		o, err := NewWebsocketOutportDriverNodePart(args)
		require.NoError(t, err)

		err = o.FinalizedBlock(&outport.FinalizedBlock{HeaderHash: []byte("header hash")})
		require.True(t, errors.Is(err, cannotSendOnRouteErr))
	})

	t.Run("Finalized block - should work", func(t *testing.T) {
		args := getMockArgs()
		args.WebsocketSender = &mock.WebSocketSenderStub{
			SendOnRouteCalled: func(_ data.WsSendArgs) error {
				return nil
			},
		}
		o, err := NewWebsocketOutportDriverNodePart(args)
		require.NoError(t, err)

		err = o.FinalizedBlock(&outport.FinalizedBlock{HeaderHash: []byte("header hash")})
		require.NoError(t, err)
	})
}

func TestWebsocketOutportDriverNodePart_RevertIndexedBlock(t *testing.T) {
	t.Parallel()

	t.Run("RevertIndexedBlock - should error", func(t *testing.T) {
		args := getMockArgs()
		args.WebsocketSender = &mock.WebSocketSenderStub{
			SendOnRouteCalled: func(_ data.WsSendArgs) error {
				return cannotSendOnRouteErr
			},
		}
		o, err := NewWebsocketOutportDriverNodePart(args)
		require.NoError(t, err)

		err = o.RevertIndexedBlock(nil)
		require.True(t, errors.Is(err, cannotSendOnRouteErr))
	})

	t.Run("RevertIndexedBlock block - should work", func(t *testing.T) {
		args := getMockArgs()
		args.WebsocketSender = &mock.WebSocketSenderStub{
			SendOnRouteCalled: func(_ data.WsSendArgs) error {
				return nil
			},
		}
		o, err := NewWebsocketOutportDriverNodePart(args)
		require.NoError(t, err)

		err = o.RevertIndexedBlock(nil)
		require.NoError(t, err)
	})
}

func TestWebsocketOutportDriverNodePart_SaveAccounts(t *testing.T) {
	t.Parallel()

	t.Run("SaveAccounts - should error", func(t *testing.T) {
		args := getMockArgs()
		args.WebsocketSender = &mock.WebSocketSenderStub{
			SendOnRouteCalled: func(_ data.WsSendArgs) error {
				return cannotSendOnRouteErr
			},
		}
		o, err := NewWebsocketOutportDriverNodePart(args)
		require.NoError(t, err)

		err = o.SaveAccounts(nil)
		require.True(t, errors.Is(err, cannotSendOnRouteErr))
	})

	t.Run("SaveAccounts block - should work", func(t *testing.T) {
		args := getMockArgs()
		args.WebsocketSender = &mock.WebSocketSenderStub{
			SendOnRouteCalled: func(_ data.WsSendArgs) error {
				return nil
			},
		}
		o, err := NewWebsocketOutportDriverNodePart(args)
		require.NoError(t, err)

		err = o.SaveAccounts(nil)
		require.NoError(t, err)
	})
}

func TestWebsocketOutportDriverNodePart_SaveRoundsInfo(t *testing.T) {
	t.Parallel()

	t.Run("SaveRoundsInfo - should error", func(t *testing.T) {
		args := getMockArgs()
		args.WebsocketSender = &mock.WebSocketSenderStub{
			SendOnRouteCalled: func(_ data.WsSendArgs) error {
				return cannotSendOnRouteErr
			},
		}
		o, err := NewWebsocketOutportDriverNodePart(args)
		require.NoError(t, err)

		err = o.SaveRoundsInfo(nil)
		require.True(t, errors.Is(err, cannotSendOnRouteErr))
	})

	t.Run("SaveRoundsInfo block - should work", func(t *testing.T) {
		args := getMockArgs()
		args.WebsocketSender = &mock.WebSocketSenderStub{
			SendOnRouteCalled: func(_ data.WsSendArgs) error {
				return nil
			},
		}
		o, err := NewWebsocketOutportDriverNodePart(args)
		require.NoError(t, err)

		err = o.SaveRoundsInfo(nil)
		require.NoError(t, err)
	})
}

func TestWebsocketOutportDriverNodePart_SaveValidatorsPubKeys(t *testing.T) {
	t.Parallel()

	t.Run("SaveValidatorsPubKeys - should error", func(t *testing.T) {
		args := getMockArgs()
		args.WebsocketSender = &mock.WebSocketSenderStub{
			SendOnRouteCalled: func(_ data.WsSendArgs) error {
				return cannotSendOnRouteErr
			},
		}
		o, err := NewWebsocketOutportDriverNodePart(args)
		require.NoError(t, err)

		err = o.SaveValidatorsPubKeys(nil)
		require.True(t, errors.Is(err, cannotSendOnRouteErr))
	})

	t.Run("SaveValidatorsPubKeys block - should work", func(t *testing.T) {
		args := getMockArgs()
		args.WebsocketSender = &mock.WebSocketSenderStub{
			SendOnRouteCalled: func(_ data.WsSendArgs) error {
				return nil
			},
		}
		o, err := NewWebsocketOutportDriverNodePart(args)
		require.NoError(t, err)

		err = o.SaveValidatorsPubKeys(nil)
		require.NoError(t, err)
	})
}

func TestWebsocketOutportDriverNodePart_SaveValidatorsRating(t *testing.T) {
	t.Parallel()

	t.Run("SaveValidatorsRating - should error", func(t *testing.T) {
		args := getMockArgs()
		args.WebsocketSender = &mock.WebSocketSenderStub{
			SendOnRouteCalled: func(_ data.WsSendArgs) error {
				return cannotSendOnRouteErr
			},
		}
		o, err := NewWebsocketOutportDriverNodePart(args)
		require.NoError(t, err)

		err = o.SaveValidatorsRating(nil)
		require.True(t, errors.Is(err, cannotSendOnRouteErr))
	})

	t.Run("SaveValidatorsRating block - should work", func(t *testing.T) {
		args := getMockArgs()
		args.WebsocketSender = &mock.WebSocketSenderStub{
			SendOnRouteCalled: func(_ data.WsSendArgs) error {
				return nil
			},
		}
		o, err := NewWebsocketOutportDriverNodePart(args)
		require.NoError(t, err)

		err = o.SaveValidatorsRating(nil)
		require.NoError(t, err)
	})
}

func TestWebsocketOutportDriverNodePart_SaveBlock_PayloadCheck(t *testing.T) {
	t.Parallel()

	mockArgs := getMockArgs()

	outportBlock := &outport.OutportBlock{BlockData: &outport.BlockData{Body: &block.Body{}}}
	marshaledData, err := mockArgs.Marshaller.Marshal(outportBlock)
	require.Nil(t, err)

	mockArgs.WebsocketSender = &mock.WebSocketSenderStub{
		SendOnRouteCalled: func(args data.WsSendArgs) error {
			expectedOpBytes := []byte{0, 0, 0, 0}

			messageLength := uint64(len(marshaledData))
			messageLengthBytes := mockArgs.Uint64ByteSliceConverter.ToByteSlice(messageLength)
			expectedLengthBytes := messageLengthBytes[uint32NumBytes:]

			expectedPayload := append(expectedOpBytes, expectedLengthBytes...)
			expectedPayload = append(expectedPayload, marshaledData...)

			require.Equal(t, expectedPayload, args.Payload)

			return nil
		},
	}
	o, err := NewWebsocketOutportDriverNodePart(mockArgs)
	require.NoError(t, err)

	err = o.SaveBlock(outportBlock)
	require.NoError(t, err)
}

func TestWebsocketOutportDriverNodePart_Close(t *testing.T) {
	t.Parallel()

	closedWasCalled := false
	args := getMockArgs()
	args.WebsocketSender = &mock.WebSocketSenderStub{
		CloseCalled: func() error {
			closedWasCalled = true
			return nil
		},
	}

	o, err := NewWebsocketOutportDriverNodePart(args)
	require.NoError(t, err)

	err = o.Close()
	require.NoError(t, err)
	require.True(t, closedWasCalled)
}
