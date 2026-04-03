package block_test

import (
	"testing"

	"github.com/multiversx/mx-chain-core-go/data"
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

	var nilSdp *block.ShardDataProposal
	sdp := &block.ShardDataProposal{}

	headerHash := []byte("header hash")
	round := uint64(100)
	nonce := uint64(42)
	shardID := uint32(1)
	epoch := uint32(5)
	numPending := uint32(10)

	err := sdp.SetHeaderHash(headerHash)
	require.NoError(t, err)
	require.Equal(t, headerHash, sdp.GetHeaderHash())
	err = nilSdp.SetHeaderHash(headerHash)
	require.Equal(t, data.ErrNilPointerReceiver, err)

	err = sdp.SetRound(round)
	require.NoError(t, err)
	require.Equal(t, round, sdp.GetRound())
	err = nilSdp.SetRound(round)
	require.Equal(t, data.ErrNilPointerReceiver, err)

	err = sdp.SetNonce(nonce)
	require.NoError(t, err)
	require.Equal(t, nonce, sdp.GetNonce())
	err = nilSdp.SetNonce(nonce)
	require.Equal(t, data.ErrNilPointerReceiver, err)

	err = sdp.SetShardID(shardID)
	require.NoError(t, err)
	require.Equal(t, shardID, sdp.GetShardID())
	err = nilSdp.SetShardID(shardID)
	require.Equal(t, data.ErrNilPointerReceiver, err)

	err = sdp.SetEpoch(epoch)
	require.NoError(t, err)
	require.Equal(t, epoch, sdp.GetEpoch())
	err = nilSdp.SetEpoch(epoch)
	require.Equal(t, data.ErrNilPointerReceiver, err)

	err = sdp.SetNumPendingMiniBlocks(numPending)
	require.NoError(t, err)
	require.Equal(t, numPending, sdp.GetNumPendingMiniBlocks())
	err = nilSdp.SetNumPendingMiniBlocks(numPending)
	require.Equal(t, data.ErrNilPointerReceiver, err)
}
