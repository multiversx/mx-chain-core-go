package unmarshal

import (
	"testing"

	"github.com/multiversx/mx-chain-core-go/core"
	"github.com/multiversx/mx-chain-core-go/data/block"
	"github.com/multiversx/mx-chain-core-go/marshal/factory"
	"github.com/stretchr/testify/require"
)

func TestGetHeaderFromBytes(t *testing.T) {
	t.Parallel()

	marshaller, _ := factory.NewMarshalizer(factory.GogoProtobuf)

	header, err := GetHeaderFromBytes(marshaller, "wrong", nil)
	require.Nil(t, header)
	require.Equal(t, errInvalidHeaderType, err)

	// header v1
	headerV1 := &block.Header{Nonce: 1}
	hBytes, _ := marshaller.Marshal(headerV1)
	header, err = GetHeaderFromBytes(marshaller, core.ShardHeaderV1, hBytes)
	require.Nil(t, err)
	require.NotNil(t, header)
	require.Equal(t, uint64(1), header.GetNonce())

	// header v2
	headerV2 := &block.HeaderV2{ScheduledRootHash: []byte("aaaaaa")}
	hBytes, _ = marshaller.Marshal(headerV2)
	header, err = GetHeaderFromBytes(marshaller, core.ShardHeaderV2, hBytes)
	require.Nil(t, err)
	require.NotNil(t, header)
	require.Equal(t, []byte("aaaaaa"), header.GetAdditionalData().GetScheduledRootHash())

	// meta
	metaHeader := &block.MetaBlock{
		Nonce: 1,
	}
	hBytes, _ = marshaller.Marshal(metaHeader)
	header, err = GetHeaderFromBytes(marshaller, core.MetaHeader, hBytes)
	require.Nil(t, err)
	require.NotNil(t, header)
	require.Equal(t, uint64(1), header.GetNonce())
}
