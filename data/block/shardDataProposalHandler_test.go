package block_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/multiversx/mx-chain-core-go/data/block"
)

func TestShardDataProposal_IsInterfaceNil(t *testing.T) {
	t.Parallel()

	var sdp *block.ShardDataProposal
	require.True(t, sdp.IsInterfaceNil())

	sdp = &block.ShardDataProposal{}
	require.False(t, sdp.IsInterfaceNil())
}

func TestShardDataProposal_AllMethods(t *testing.T) {
	t.Parallel()

	sdp := &block.ShardDataProposal{}

	headerHash := []byte("header hash")
	round := uint64(100)
	nonce := uint64(42)
	shardID := uint32(1)
	epoch := uint32(5)

	sdp.SetHeaderHash(headerHash)
	require.Equal(t, headerHash, sdp.GetHeaderHash())

	sdp.SetRound(round)
	require.Equal(t, round, sdp.GetRound())

	sdp.SetNonce(nonce)
	require.Equal(t, nonce, sdp.GetNonce())

	sdp.SetShardID(shardID)
	require.Equal(t, shardID, sdp.GetShardID())

	sdp.SetEpoch(epoch)
	require.Equal(t, epoch, sdp.GetEpoch())
}
