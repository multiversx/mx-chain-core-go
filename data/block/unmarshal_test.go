package block

import (
	"testing"

	"github.com/multiversx/mx-chain-core-go/core/check"
	"github.com/multiversx/mx-chain-core-go/data"
	"github.com/multiversx/mx-chain-core-go/data/mock"
	"github.com/multiversx/mx-chain-core-go/marshal/factory"
	"github.com/stretchr/testify/require"
)

func TestGetHeaderFromBytes(t *testing.T) {
	t.Parallel()

	marshaller, _ := factory.NewMarshalizer(factory.GogoProtobuf)
	headerV1 := &Header{Nonce: 1}
	hBytes, _ := marshaller.Marshal(headerV1)
	emptyBlockCreator := &mock.EmptyBlockCreatorStub{
		CreateNewHeaderCalled: func() data.HeaderHandler {
			return &Header{}
		},
	}

	t.Run("should return error when nil marshaller", func(t *testing.T) {
		header, err := GetHeaderFromBytes(nil, emptyBlockCreator, hBytes)
		require.True(t, check.IfNil(header))
		require.Equal(t, data.ErrNilMarshalizer, err)
	})
	t.Run("should return error when nil empty block creator", func(t *testing.T) {
		header, err := GetHeaderFromBytes(marshaller, nil, hBytes)
		require.True(t, check.IfNil(header))
		require.Equal(t, data.ErrNilEmptyBlockCreator, err)
	})
	t.Run("wrong bytes should return error", func(t *testing.T) {
		header, err := GetHeaderFromBytes(marshaller, emptyBlockCreator, []byte("wrong bytes"))
		require.True(t, check.IfNil(header))
		require.NotNil(t, err)
	})
	t.Run("nil bytes should return empty header", func(t *testing.T) {
		header, err := GetHeaderFromBytes(marshaller, emptyBlockCreator, nil)
		require.False(t, check.IfNil(header))
		require.Nil(t, err)
		require.Equal(t, uint64(0), header.GetNonce())
	})
	t.Run("empty bytes should return empty header", func(t *testing.T) {
		header, err := GetHeaderFromBytes(marshaller, emptyBlockCreator, make([]byte, 0))
		require.False(t, check.IfNil(header))
		require.Nil(t, err)
		require.Equal(t, uint64(0), header.GetNonce())
	})
	t.Run("should work with correct bytes", func(t *testing.T) {
		header, err := GetHeaderFromBytes(marshaller, emptyBlockCreator, hBytes)
		require.False(t, check.IfNil(header))
		require.Nil(t, err)
		require.Equal(t, uint64(1), header.GetNonce())
	})
}
