package block_test

import (
	"errors"
	"math/big"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/multiversx/mx-chain-core-go/core"
	"github.com/multiversx/mx-chain-core-go/data"
	"github.com/multiversx/mx-chain-core-go/data/block"
)

func TestMetaBlockV3_GetExecutionResultsHandlers(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()
		var mb2 *block.MetaBlockV3
		require.Nil(t, mb2.GetExecutionResultsHandlers())
	})

	t.Run("should work", func(t *testing.T) {
		t.Parallel()
		mb2 := &block.MetaBlockV3{
			ExecutionResults: []*block.MetaExecutionResult{
				{ExecutionResult: &block.BaseMetaExecutionResult{BaseExecutionResult: &block.BaseExecutionResult{HeaderHash: []byte("hash1")}}},
				{ExecutionResult: &block.BaseMetaExecutionResult{BaseExecutionResult: &block.BaseExecutionResult{HeaderHash: []byte("hash2")}}},
			},
		}
		expected := []data.BaseExecutionResultHandler{
			mb2.ExecutionResults[0], mb2.ExecutionResults[1],
		}
		result := mb2.GetExecutionResultsHandlers()
		require.Equal(t, expected, result)
	})
}

func TestMetaBlockV3_GetLastExecutionResultHandler(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()
		var mb2 *block.MetaBlockV3
		require.Nil(t, mb2.GetLastExecutionResultHandler())
	})

	t.Run("should work", func(t *testing.T) {
		t.Parallel()
		mb2 := &block.MetaBlockV3{
			LastExecutionResult: &block.MetaExecutionResultInfo{
				NotarizedInRound: 100,
				ExecutionResult: &block.BaseMetaExecutionResult{
					BaseExecutionResult: &block.BaseExecutionResult{HeaderHash: []byte("hash1")},
				},
			},
		}
		expected := mb2.LastExecutionResult
		result := mb2.GetLastExecutionResultHandler()
		require.Equal(t, expected, result)
	})
}

func TestMetaBlockV3_GetValidatorStatsRootHash(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()
		var mb2 *block.MetaBlockV3
		require.Nil(t, mb2.GetValidatorStatsRootHash())
	})

	t.Run("valid receiver, should return nil", func(t *testing.T) {
		t.Parallel()
		mb2 := &block.MetaBlockV3{}
		require.Nil(t, mb2.GetValidatorStatsRootHash())
	})
}

func TestMetaBlockV3_GetDevFeesInEpoch(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()
		var mb2 *block.MetaBlockV3
		require.Nil(t, mb2.GetDevFeesInEpoch())
	})

	t.Run("valid receiver, should return nil", func(t *testing.T) {
		t.Parallel()
		mb2 := &block.MetaBlockV3{}
		require.Nil(t, mb2.GetDevFeesInEpoch())
	})
}

func TestMetaBlockV3_GetEpochStartHandler(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()
		var mb2 *block.MetaBlockV3
		require.Nil(t, mb2.GetEpochStartHandler())
	})

	t.Run("valid receiver", func(t *testing.T) {
		t.Parallel()
		mb2 := &block.MetaBlockV3{
			EpochStart: block.EpochStart{
				LastFinalizedHeaders: nil,
				Economics:            block.Economics{},
			},
		}
		require.Equal(t, &mb2.EpochStart, mb2.GetEpochStartHandler())
	})
}

func TestMetaBlockV3_GetShardInfoHandlers(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()
		var mb2 *block.MetaBlockV3
		require.Nil(t, mb2.GetShardInfoHandlers())
	})

	t.Run("should work", func(t *testing.T) {
		t.Parallel()
		shardData1 := block.ShardData{ShardID: 0, HeaderHash: []byte("shard1")}
		shardData2 := block.ShardData{ShardID: 1, HeaderHash: []byte("shard2")}
		mb2 := &block.MetaBlockV3{
			ShardInfo: []block.ShardData{shardData1, shardData2},
		}
		expected := []data.ShardDataHandler{&mb2.ShardInfo[0], &mb2.ShardInfo[1]}
		result := mb2.GetShardInfoHandlers()
		require.Equal(t, expected, result)
	})
}

func TestMetaBlockV3_SetShardInfoHandlers(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()
		var mb2 *block.MetaBlockV3
		err := mb2.SetShardInfoHandlers(nil)
		require.Equal(t, data.ErrNilPointerReceiver, err)
	})

	t.Run("should work with nil", func(t *testing.T) {
		t.Parallel()
		mb2 := &block.MetaBlockV3{
			ShardInfo: make([]block.ShardData, 2),
		}
		err := mb2.SetShardInfoHandlers(nil)
		require.NoError(t, err)
		require.Nil(t, mb2.ShardInfo)
	})

	t.Run("should work with empty slice", func(t *testing.T) {
		t.Parallel()
		mb2 := &block.MetaBlockV3{}
		err := mb2.SetShardInfoHandlers([]data.ShardDataHandler{})
		require.NoError(t, err)
		require.Empty(t, mb2.ShardInfo)
	})

	t.Run("should error on list of nil shard data handlers", func(t *testing.T) {
		t.Parallel()
		mb2 := &block.MetaBlockV3{}
		err := mb2.SetShardInfoHandlers([]data.ShardDataHandler{nil})
		require.Equal(t, data.ErrInvalidTypeAssertion, err)
	})

	t.Run("should error on list of nil shard data", func(t *testing.T) {
		t.Parallel()
		mb2 := &block.MetaBlockV3{}
		var shardData *block.ShardData = nil
		err := mb2.SetShardInfoHandlers([]data.ShardDataHandler{shardData})
		require.Equal(t, data.ErrNilPointerDereference, err)
	})

	t.Run("should work with valid handlers", func(t *testing.T) {
		t.Parallel()
		shardData1 := &block.ShardData{ShardID: 0, HeaderHash: []byte("shard1")}
		shardData2 := &block.ShardData{ShardID: 1, HeaderHash: []byte("shard2")}
		handlers := []data.ShardDataHandler{shardData1, shardData2}

		mb2 := &block.MetaBlockV3{}
		err := mb2.SetShardInfoHandlers(handlers)
		require.NoError(t, err)
		require.Equal(t, 2, len(mb2.ShardInfo))
		assert.Equal(t, shardData1.GetShardID(), mb2.ShardInfo[0].ShardID)
		assert.Equal(t, shardData2.GetShardID(), mb2.ShardInfo[1].ShardID)
	})
}

func TestMetaBlockV3_SetValidatorStatsRootHash(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()
		var mb2 *block.MetaBlockV3
		err := mb2.SetValidatorStatsRootHash([]byte("root"))
		require.Equal(t, data.ErrFieldNotSupported, err)
	})

	t.Run("valid receiver should also error", func(t *testing.T) {
		t.Parallel()
		rootHash := []byte("validator stats root")
		mb2 := &block.MetaBlockV3{}
		require.Equal(t, data.ErrFieldNotSupported, mb2.SetValidatorStatsRootHash(rootHash))
	})
}

func TestMetaBlockV3_SetDevFeesInEpoch(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()
		var mb2 *block.MetaBlockV3
		err := mb2.SetDevFeesInEpoch(big.NewInt(100))
		require.Equal(t, data.ErrFieldNotSupported, err)
	})

	t.Run("valid receiver should also error", func(t *testing.T) {
		t.Parallel()
		devFees := big.NewInt(50)
		mb2 := &block.MetaBlockV3{}
		require.Equal(t, data.ErrFieldNotSupported, mb2.SetDevFeesInEpoch(devFees))
	})
}

func TestMetaBlockV3_SetAccumulatedFeesInEpoch(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()
		var mb2 *block.MetaBlockV3
		err := mb2.SetAccumulatedFeesInEpoch(big.NewInt(100))
		require.Equal(t, data.ErrFieldNotSupported, err)
	})

	t.Run("valid receiver should also error", func(t *testing.T) {
		t.Parallel()
		accumulatedFees := big.NewInt(50)
		mb2 := &block.MetaBlockV3{}
		require.Equal(t, data.ErrFieldNotSupported, mb2.SetAccumulatedFeesInEpoch(accumulatedFees))
	})
}

func TestMetaBlockV3_GetRootHash(t *testing.T) {
	t.Parallel()
	mb2 := &block.MetaBlockV3{}
	require.Nil(t, mb2.GetRootHash())
}

func TestMetaBlockV3_GetPubKeysBitmap(t *testing.T) {
	t.Parallel()
	mb2 := &block.MetaBlockV3{}
	require.Nil(t, mb2.GetPubKeysBitmap())
}

func TestMetaBlockV3_GetSignature(t *testing.T) {
	t.Parallel()
	mb2 := &block.MetaBlockV3{}
	require.Nil(t, mb2.GetSignature())
}

func TestMetaBlockV3_GetTimeStamp(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()
		var mb2 *block.MetaBlockV3
		require.Equal(t, uint64(0), mb2.GetTimeStamp())
	})

	t.Run("should work", func(t *testing.T) {
		t.Parallel()
		timestamp := uint64(12345)
		mb2 := &block.MetaBlockV3{TimestampMs: timestamp}
		require.Equal(t, timestamp, mb2.GetTimeStamp())
	})
}

func TestMetaBlockV3_GetReceiptsHash(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()
		var mb2 *block.MetaBlockV3
		require.Nil(t, mb2.GetReceiptsHash())
	})
}

func TestMetaBlockV3_GetMiniBlockHeadersWithDst(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()
		var mb2 *block.MetaBlockV3
		require.Nil(t, mb2.GetMiniBlockHeadersWithDst(0))
	})
	t.Run("should return headers with correct destination", func(t *testing.T) {
		t.Parallel()

		metaHdr := &block.MetaBlockV3{Round: 15}
		metaHdr.ShardInfo = make([]block.ShardData, 0)

		shardMBHeader := make([]block.MiniBlockHeader, 0)
		shMBHdr1 := block.MiniBlockHeader{SenderShardID: 0, ReceiverShardID: 1, Hash: []byte("hash1")}
		shMBHdr2 := block.MiniBlockHeader{SenderShardID: 0, ReceiverShardID: 1, Hash: []byte("hash2")}
		shardMBHeader = append(shardMBHeader, shMBHdr1, shMBHdr2)

		shData1 := block.ShardData{ShardID: 0, HeaderHash: []byte("sh"), ShardMiniBlockHeaders: shardMBHeader}
		metaHdr.ShardInfo = append(metaHdr.ShardInfo, shData1)

		shData2 := block.ShardData{ShardID: 1, HeaderHash: []byte("sh"), ShardMiniBlockHeaders: shardMBHeader}
		metaHdr.ShardInfo = append(metaHdr.ShardInfo, shData2)

		mbsFromMetaToShard0 := []block.MiniBlockHeader{{Hash: []byte("hash3"), SenderShardID: core.MetachainShardId, ReceiverShardID: 0}}
		mbsFromMetaToShard1 := []block.MiniBlockHeader{{Hash: []byte("hash4"), SenderShardID: core.MetachainShardId, ReceiverShardID: 1}}
		metaHdr.ExecutionResults = make([]*block.MetaExecutionResult, 2)
		metaHdr.ExecutionResults[0] = &block.MetaExecutionResult{MiniBlockHeaders: mbsFromMetaToShard0}
		metaHdr.ExecutionResults[1] = &block.MetaExecutionResult{MiniBlockHeaders: mbsFromMetaToShard1}

		mbDst0 := metaHdr.GetMiniBlockHeadersWithDst(0)
		assert.Equal(t, len(mbsFromMetaToShard0), len(mbDst0))
		mbDst1 := metaHdr.GetMiniBlockHeadersWithDst(1)
		assert.Equal(t, len(shardMBHeader)+len(mbsFromMetaToShard1), len(mbDst1))
	})
}

func TestMetaBlockV3_GetProposedMiniBlockHeadersWithDst(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()
		var mb2 *block.MetaBlockV3
		require.Nil(t, mb2.GetProposedMiniBlockHeadersWithDst(0))
	})
	t.Run("should return headers with correct destination", func(t *testing.T) {
		t.Parallel()

		metaHdr := &block.MetaBlockV3{Round: 15}
		metaHdr.ShardInfo = make([]block.ShardData, 0)

		shardMBHeader := make([]block.MiniBlockHeader, 0)
		shMBHdr1 := block.MiniBlockHeader{SenderShardID: 0, ReceiverShardID: 1, Hash: []byte("hash1")}
		shMBHdr2 := block.MiniBlockHeader{SenderShardID: 0, ReceiverShardID: 1, Hash: []byte("hash2")}
		shardMBHeader = append(shardMBHeader, shMBHdr1, shMBHdr2)

		shData1 := block.ShardData{ShardID: 0, HeaderHash: []byte("sh"), ShardMiniBlockHeaders: shardMBHeader}
		metaHdr.ShardInfo = append(metaHdr.ShardInfo, shData1)

		shData2 := block.ShardData{ShardID: 1, HeaderHash: []byte("sh"), ShardMiniBlockHeaders: shardMBHeader}
		metaHdr.ShardInfo = append(metaHdr.ShardInfo, shData2)

		mbsFromMetaToShard0 := []block.MiniBlockHeader{{Hash: []byte("hash3"), SenderShardID: core.MetachainShardId, ReceiverShardID: 0}}
		mbsFromMetaToShard1 := []block.MiniBlockHeader{{Hash: []byte("hash4"), SenderShardID: core.MetachainShardId, ReceiverShardID: 1}}
		metaHdr.MiniBlockHeaders = append(metaHdr.MiniBlockHeaders, mbsFromMetaToShard0...)
		metaHdr.MiniBlockHeaders = append(metaHdr.MiniBlockHeaders, mbsFromMetaToShard1...)
		// should not include execution results
		metaHdr.ExecutionResults = make([]*block.MetaExecutionResult, 1)
		metaHdr.ExecutionResults[0] = &block.MetaExecutionResult{MiniBlockHeaders: mbsFromMetaToShard1}

		mbDst0 := metaHdr.GetProposedMiniBlockHeadersWithDst(0)
		require.Equal(t, len(mbsFromMetaToShard0), len(mbDst0))
		mbDst1 := metaHdr.GetProposedMiniBlockHeadersWithDst(1)
		require.Equal(t, len(mbsFromMetaToShard1), len(mbDst1)) // should not include shard info data
	})
}

func TestMetaBlockV3_GetOrderedCrossMiniblocksWithDst(t *testing.T) {
	t.Parallel()

	metaHdr := &block.MetaBlockV3{Round: 6}
	metaHdr.ShardInfo = make([]block.ShardData, 0)

	shardMBHeader1 := make([]block.MiniBlockHeader, 0)
	shMBHdr1 := block.MiniBlockHeader{SenderShardID: 0, ReceiverShardID: 1, Hash: []byte("hash1")}
	shardMBHeader1 = append(shardMBHeader1, shMBHdr1)
	shData1 := block.ShardData{Round: 11, ShardID: 0, HeaderHash: []byte("sh1"), ShardMiniBlockHeaders: shardMBHeader1}

	shardMBHeader2 := make([]block.MiniBlockHeader, 0)
	shMBHdr2 := block.MiniBlockHeader{SenderShardID: 0, ReceiverShardID: 1, Hash: []byte("hash2")}
	shardMBHeader2 = append(shardMBHeader2, shMBHdr2)
	shData2 := block.ShardData{Round: 9, ShardID: 0, HeaderHash: []byte("sh2"), ShardMiniBlockHeaders: shardMBHeader2}

	shardMBHeader3 := make([]block.MiniBlockHeader, 0)
	shMBHdr3 := block.MiniBlockHeader{SenderShardID: 2, ReceiverShardID: 1, Hash: []byte("hash3")}
	shardMBHeader3 = append(shardMBHeader3, shMBHdr3)
	shData3 := block.ShardData{Round: 10, ShardID: 2, HeaderHash: []byte("sh3"), ShardMiniBlockHeaders: shardMBHeader3}

	shardMBHeader4 := make([]block.MiniBlockHeader, 0)
	shMBHdr4 := block.MiniBlockHeader{SenderShardID: 2, ReceiverShardID: 1, Hash: []byte("hash4")}
	shardMBHeader4 = append(shardMBHeader4, shMBHdr4)
	shData4 := block.ShardData{Round: 8, ShardID: 2, HeaderHash: []byte("sh4"), ShardMiniBlockHeaders: shardMBHeader4}

	shardMBHeader5 := make([]block.MiniBlockHeader, 0)
	shMBHdr5 := block.MiniBlockHeader{SenderShardID: 1, ReceiverShardID: 2, Hash: []byte("hash5")}
	shardMBHeader5 = append(shardMBHeader5, shMBHdr5)
	shData5 := block.ShardData{Round: 7, ShardID: 1, HeaderHash: []byte("sh5"), ShardMiniBlockHeaders: shardMBHeader5}

	shardMBHeader6 := make([]block.MiniBlockHeader, 0)
	shMBHdr6 := block.MiniBlockHeader{SenderShardID: core.MetachainShardId, ReceiverShardID: core.AllShardId, Hash: []byte("hashAll")}
	shardMBHeader6 = append(shardMBHeader6, shMBHdr6)
	shData6 := block.ShardData{Round: 12, ShardID: 1, HeaderHash: []byte("sh6"), ShardMiniBlockHeaders: shardMBHeader6}

	metaHdr.ShardInfo = append(metaHdr.ShardInfo, shData1, shData2, shData3, shData4, shData5, shData6)

	mb1 := block.MiniBlockHeader{
		Hash:            []byte("hash6"),
		SenderShardID:   core.MetachainShardId,
		ReceiverShardID: 1,
	}

	mb2 := block.MiniBlockHeader{
		Hash:            []byte("hash7"),
		SenderShardID:   core.MetachainShardId,
		ReceiverShardID: core.AllShardId,
	}

	mb3 := block.MiniBlockHeader{
		Hash:            []byte("hash8"),
		SenderShardID:   core.MetachainShardId,
		ReceiverShardID: 2,
	}
	metaHdr.ExecutionResults = make([]*block.MetaExecutionResult, 1)
	metaHdr.ExecutionResults[0] = &block.MetaExecutionResult{
		MiniBlockHeaders: []block.MiniBlockHeader{mb1, mb2, mb3},
		ExecutionResult: &block.BaseMetaExecutionResult{
			BaseExecutionResult: &block.BaseExecutionResult{HeaderRound: 6},
		},
	}

	miniBlocksInfo := metaHdr.GetOrderedCrossMiniblocksWithDst(1)
	require.Equal(t, 6, len(miniBlocksInfo))
	assert.Equal(t, miniBlocksInfo[0].Hash, []byte("hash6"))
	assert.Equal(t, miniBlocksInfo[0].Round, uint64(6))
	assert.Equal(t, miniBlocksInfo[1].Hash, []byte("hash7"))
	assert.Equal(t, miniBlocksInfo[1].Round, uint64(6))
	assert.Equal(t, miniBlocksInfo[2].Hash, []byte("hash4"))
	assert.Equal(t, miniBlocksInfo[2].Round, uint64(8))
	assert.Equal(t, miniBlocksInfo[3].Hash, []byte("hash2"))
	assert.Equal(t, miniBlocksInfo[3].Round, uint64(9))
	assert.Equal(t, miniBlocksInfo[4].Hash, []byte("hash3"))
	assert.Equal(t, miniBlocksInfo[4].Round, uint64(10))
	assert.Equal(t, miniBlocksInfo[5].Hash, []byte("hash1"))
	assert.Equal(t, miniBlocksInfo[5].Round, uint64(11))

	miniBlocksInfo = metaHdr.GetOrderedCrossMiniblocksWithDst(2)
	require.Equal(t, 3, len(miniBlocksInfo))
	assert.Equal(t, miniBlocksInfo[0].Hash, []byte("hash7"))
	assert.Equal(t, miniBlocksInfo[0].Round, uint64(6))
	assert.Equal(t, miniBlocksInfo[1].Hash, []byte("hash8"))
	assert.Equal(t, miniBlocksInfo[1].Round, uint64(6))
	assert.Equal(t, miniBlocksInfo[2].Hash, []byte("hash5"))
	assert.Equal(t, miniBlocksInfo[2].Round, uint64(7))
}

func TestMetaBlockV3_GetMiniBlockHeadersHashes(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()
		var mb2 *block.MetaBlockV3
		require.Nil(t, mb2.GetMiniBlockHeadersHashes())
	})

	t.Run("should work", func(t *testing.T) {
		t.Parallel()
		hash1 := []byte("hash1")
		hash2 := []byte("hash2")
		mb2 := &block.MetaBlockV3{
			MiniBlockHeaders: []block.MiniBlockHeader{
				{Hash: hash1},
				{Hash: hash2},
			},
		}
		expected := [][]byte{hash1, hash2}
		result := mb2.GetMiniBlockHeadersHashes()
		require.Equal(t, expected, result)
	})
}

func TestMetaBlockV3_GetMiniBlockHeaderHandlers(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()
		var mb2 *block.MetaBlockV3
		require.Nil(t, mb2.GetMiniBlockHeaderHandlers())
	})

	t.Run("should work", func(t *testing.T) {
		t.Parallel()
		mb2 := &block.MetaBlockV3{
			MiniBlockHeaders: []block.MiniBlockHeader{
				{Hash: []byte("hash1")},
				{Hash: []byte("hash2")},
			},
		}
		result := mb2.GetMiniBlockHeaderHandlers()
		require.Equal(t, 2, len(result))
		require.Equal(t, mb2.MiniBlockHeaders[0].Hash, result[0].GetHash())
		require.Equal(t, mb2.MiniBlockHeaders[1].Hash, result[1].GetHash())
	})
}

func TestMetaBlockV3_HasScheduledSupport(t *testing.T) {
	t.Parallel()
	mb2 := &block.MetaBlockV3{}
	require.False(t, mb2.HasScheduledSupport())
}

func TestMetaBlockV3_GetAdditionalData(t *testing.T) {
	t.Parallel()
	mb2 := &block.MetaBlockV3{}
	require.Nil(t, mb2.GetAdditionalData())
}

func TestMetaBlockV3_HasScheduledMiniBlocks(t *testing.T) {
	t.Parallel()
	mb2 := &block.MetaBlockV3{}
	require.False(t, mb2.HasScheduledMiniBlocks())
}

func TestMetaBlockV3_SetAccumulatedFees(t *testing.T) {
	t.Parallel()
	mb2 := &block.MetaBlockV3{}
	err := mb2.SetAccumulatedFees(big.NewInt(100))
	require.Equal(t, data.ErrFieldNotSupported, err)
}

func TestMetaBlockV3_SetDeveloperFees(t *testing.T) {
	t.Parallel()
	mb2 := &block.MetaBlockV3{}
	err := mb2.SetDeveloperFees(big.NewInt(50))
	require.Equal(t, data.ErrFieldNotSupported, err)
}

func TestMetaBlockV3_SetShardIDWillDoNothing(t *testing.T) {
	t.Parallel()

	mb2 := &block.MetaBlockV3{}
	err := mb2.SetShardID(0)
	require.NoError(t, err)

	shardID := mb2.GetShardID()
	require.Equal(t, core.MetachainShardId, shardID)
}

func TestMetaBlockV3_SetNonce(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()
		var mb2 *block.MetaBlockV3
		err := mb2.SetNonce(1)
		require.Equal(t, data.ErrNilPointerReceiver, err)
	})

	t.Run("should work", func(t *testing.T) {
		t.Parallel()
		mb2 := &block.MetaBlockV3{}
		require.NoError(t, mb2.SetNonce(42))
		require.Equal(t, uint64(42), mb2.Nonce)
	})
}

func TestMetaBlockV3_SetEpoch(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()
		var mb2 *block.MetaBlockV3
		err := mb2.SetEpoch(1)
		require.Equal(t, data.ErrNilPointerReceiver, err)
	})

	t.Run("should work", func(t *testing.T) {
		t.Parallel()
		mb2 := &block.MetaBlockV3{}
		require.NoError(t, mb2.SetEpoch(2))
		require.Equal(t, uint32(2), mb2.Epoch)
	})
}

func TestMetaBlockV3_SetRound(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()
		var mb2 *block.MetaBlockV3
		err := mb2.SetRound(1)
		require.Equal(t, data.ErrNilPointerReceiver, err)
	})

	t.Run("should work", func(t *testing.T) {
		t.Parallel()
		mb2 := &block.MetaBlockV3{}
		require.NoError(t, mb2.SetRound(42))
		require.Equal(t, uint64(42), mb2.Round)
	})
}

func TestMetaBlockV3_SetTimeStamp(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()
		var mb2 *block.MetaBlockV3
		err := mb2.SetTimeStamp(12345)
		require.Equal(t, data.ErrNilPointerReceiver, err)
	})

	t.Run("should work", func(t *testing.T) {
		t.Parallel()
		mb2 := &block.MetaBlockV3{}
		require.NoError(t, mb2.SetTimeStamp(12345))
		require.Equal(t, uint64(12345), mb2.TimestampMs)
	})
}

func TestMetaBlockV3_SetRootHash(t *testing.T) {
	t.Parallel()
	mb2 := &block.MetaBlockV3{}
	err := mb2.SetRootHash([]byte("root"))
	require.Equal(t, data.ErrFieldNotSupported, err)
}

func TestMetaBlockV3_SetPrevHash(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()
		var mb2 *block.MetaBlockV3
		err := mb2.SetPrevHash([]byte("prev"))
		require.Equal(t, data.ErrNilPointerReceiver, err)
	})

	t.Run("should work", func(t *testing.T) {
		t.Parallel()
		prevHash := []byte("prev hash")
		mb2 := &block.MetaBlockV3{}
		require.NoError(t, mb2.SetPrevHash(prevHash))
		require.Equal(t, prevHash, mb2.PrevHash)
	})
}

func TestMetaBlockV3_SetPrevRandSeed(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()
		var mb2 *block.MetaBlockV3
		err := mb2.SetPrevRandSeed([]byte("seed"))
		require.Equal(t, data.ErrNilPointerReceiver, err)
	})

	t.Run("should work", func(t *testing.T) {
		t.Parallel()
		seed := []byte("prev rand seed")
		mb2 := &block.MetaBlockV3{}
		require.NoError(t, mb2.SetPrevRandSeed(seed))
		require.Equal(t, seed, mb2.PrevRandSeed)
	})
}

func TestMetaBlockV3_SetRandSeed(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()
		var mb2 *block.MetaBlockV3
		err := mb2.SetRandSeed([]byte("seed"))
		require.Equal(t, data.ErrNilPointerReceiver, err)
	})

	t.Run("should work", func(t *testing.T) {
		t.Parallel()
		seed := []byte("rand seed")
		mb2 := &block.MetaBlockV3{}
		require.NoError(t, mb2.SetRandSeed(seed))
		require.Equal(t, seed, mb2.RandSeed)
	})
}

func TestMetaBlockV3_SetPubKeysBitmap(t *testing.T) {
	t.Parallel()
	mb2 := &block.MetaBlockV3{}
	err := mb2.SetPubKeysBitmap([]byte("bitmap"))
	require.Equal(t, data.ErrFieldNotSupported, err)
}

func TestMetaBlockV3_SetSignature(t *testing.T) {
	t.Parallel()
	mb2 := &block.MetaBlockV3{}
	err := mb2.SetSignature([]byte("signature"))
	require.Equal(t, data.ErrFieldNotSupported, err)
}

func TestMetaBlockV3_SetLeaderSignature(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()
		var mb2 *block.MetaBlockV3
		err := mb2.SetLeaderSignature([]byte("sig"))
		require.Equal(t, data.ErrNilPointerReceiver, err)
	})

	t.Run("should work", func(t *testing.T) {
		t.Parallel()
		sig := []byte("leader signature")
		mb2 := &block.MetaBlockV3{}
		require.NoError(t, mb2.SetLeaderSignature(sig))
		require.Equal(t, sig, mb2.LeaderSignature)
	})
}

func TestMetaBlockV3_SetChainID(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()
		var mb2 *block.MetaBlockV3
		err := mb2.SetChainID([]byte("chain"))
		require.Equal(t, data.ErrNilPointerReceiver, err)
	})

	t.Run("should work", func(t *testing.T) {
		t.Parallel()
		chainID := []byte("chain ID")
		mb2 := &block.MetaBlockV3{}
		require.NoError(t, mb2.SetChainID(chainID))
		require.Equal(t, chainID, mb2.ChainID)
	})
}

func TestMetaBlockV3_SetSoftwareVersion(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()
		var mb2 *block.MetaBlockV3
		err := mb2.SetSoftwareVersion([]byte("v1.0.0"))
		require.Equal(t, data.ErrNilPointerReceiver, err)
	})

	t.Run("should work", func(t *testing.T) {
		t.Parallel()
		version := []byte("v1.2.3")
		mb2 := &block.MetaBlockV3{}
		require.NoError(t, mb2.SetSoftwareVersion(version))
		require.Equal(t, version, mb2.SoftwareVersion)
	})
}

func TestMetaBlockV3_SetTxCount(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()
		var mb2 *block.MetaBlockV3
		err := mb2.SetTxCount(10)
		require.Equal(t, data.ErrNilPointerReceiver, err)
	})

	t.Run("should work", func(t *testing.T) {
		t.Parallel()
		mb2 := &block.MetaBlockV3{}
		require.NoError(t, mb2.SetTxCount(42))
		require.Equal(t, uint32(42), mb2.TxCount)
	})
}

func TestMetaBlockV3_SetMiniBlockHeaderHandlers(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()
		var mb2 *block.MetaBlockV3
		err := mb2.SetMiniBlockHeaderHandlers(nil)
		require.Equal(t, data.ErrNilPointerReceiver, err)
	})

	t.Run("nil handlers", func(t *testing.T) {
		t.Parallel()
		mb2 := &block.MetaBlockV3{}
		err := mb2.SetMiniBlockHeaderHandlers(nil)
		require.NoError(t, err)
	})

	t.Run("should work", func(t *testing.T) {
		t.Parallel()
		mb2 := &block.MetaBlockV3{}
		mbh1 := &block.MiniBlockHeader{Hash: []byte("hash1")}
		mbh2 := &block.MiniBlockHeader{Hash: []byte("hash2")}
		handlers := []data.MiniBlockHeaderHandler{mbh1, mbh2}

		err := mb2.SetMiniBlockHeaderHandlers(handlers)
		require.NoError(t, err)
		require.Equal(t, 2, len(mb2.MiniBlockHeaders))
		require.Equal(t, mbh1.Hash, mb2.MiniBlockHeaders[0].Hash)
		require.Equal(t, mbh2.Hash, mb2.MiniBlockHeaders[1].Hash)
	})
}

func TestMetaBlockV3_SetScheduledRootHash(t *testing.T) {
	t.Parallel()
	mb2 := &block.MetaBlockV3{}
	err := mb2.SetScheduledRootHash([]byte("scheduled"))
	require.Equal(t, data.ErrScheduledRootHashNotSupported, err)
}

func TestMetaBlockV3_ValidateHeaderVersion(t *testing.T) {
	t.Parallel()
	mb2 := &block.MetaBlockV3{}
	err := mb2.ValidateHeaderVersion()
	require.NoError(t, err)
}

func TestMetaBlockV3_SetAdditionalData(t *testing.T) {
	t.Parallel()
	mb2 := &block.MetaBlockV3{}
	err := mb2.SetAdditionalData(nil)
	require.Equal(t, data.ErrFieldNotSupported, err)
}

func TestMetaBlockV3_SetReceiptsHash(t *testing.T) {
	t.Parallel()
	mb2 := &block.MetaBlockV3{}
	err := mb2.SetReceiptsHash([]byte("receipts"))
	require.Equal(t, data.ErrFieldNotSupported, err)
}

func TestMetaBlockV3_IsStartOfEpochBlock(t *testing.T) {
	t.Parallel()
	mb2 := &block.MetaBlockV3{}
	require.False(t, mb2.IsStartOfEpochBlock())

	mb2 = nil
	require.False(t, mb2.IsStartOfEpochBlock())
}

func TestMetaBlockV3_ShallowClone(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()
		var mb2 *block.MetaBlockV3
		clone := mb2.ShallowClone()
		require.Nil(t, clone)
	})

	t.Run("should work", func(t *testing.T) {
		t.Parallel()
		mb2 := &block.MetaBlockV3{
			Nonce:           42,
			Epoch:           2,
			Round:           100,
			TimestampMs:     12345,
			PrevHash:        []byte("prev hash"),
			PrevRandSeed:    []byte("prev rand seed"),
			RandSeed:        []byte("rand seed"),
			ChainID:         []byte("chain"),
			SoftwareVersion: []byte("v1.0.0"),
		}
		clone := mb2.ShallowClone()
		require.NotNil(t, clone)
		require.Equal(t, mb2, clone)
		require.False(t, &mb2.MiniBlockHeaders == &clone.(*block.MetaBlockV3).MiniBlockHeaders)
	})
}

func TestMetaBlockV3_CheckFieldsForNil(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()
		var mb2 *block.MetaBlockV3
		err := mb2.CheckFieldsForNil()
		require.Equal(t, data.ErrNilPointerReceiver, err)
	})

	t.Run("nil prev hash", func(t *testing.T) {
		t.Parallel()
		mb2 := &block.MetaBlockV3{
			PrevHash: nil,
		}
		err := mb2.CheckFieldsForNil()
		require.True(t, errors.Is(err, data.ErrNilValue))
		require.True(t, strings.Contains(err.Error(), "PrevHash"))
	})

	t.Run("nil prev rand seed", func(t *testing.T) {
		t.Parallel()
		mb2 := &block.MetaBlockV3{
			PrevHash:     []byte("prev hash"),
			PrevRandSeed: nil,
		}
		err := mb2.CheckFieldsForNil()
		require.True(t, errors.Is(err, data.ErrNilValue))
		require.True(t, strings.Contains(err.Error(), "PrevRandSeed"))
	})

	t.Run("nil rand seed", func(t *testing.T) {
		t.Parallel()
		mb2 := &block.MetaBlockV3{
			PrevHash:     []byte("prev hash"),
			PrevRandSeed: []byte("prev rand seed"),
			RandSeed:     nil,
		}
		err := mb2.CheckFieldsForNil()
		require.True(t, errors.Is(err, data.ErrNilValue))
		require.True(t, strings.Contains(err.Error(), "RandSeed"))
	})

	t.Run("nil leader sig", func(t *testing.T) {
		t.Parallel()
		mb2 := &block.MetaBlockV3{
			PrevHash:        []byte("prev hash"),
			PrevRandSeed:    []byte("prev rand seed"),
			RandSeed:        []byte("rand seed"),
			LeaderSignature: nil,
		}
		err := mb2.CheckFieldsForNil()
		require.True(t, errors.Is(err, data.ErrNilValue))
		require.True(t, strings.Contains(err.Error(), "LeaderSignature"))
	})

	t.Run("nil chain ID", func(t *testing.T) {
		t.Parallel()
		mb2 := &block.MetaBlockV3{
			PrevHash:        []byte("prev hash"),
			PrevRandSeed:    []byte("prev rand seed"),
			RandSeed:        []byte("rand seed"),
			LeaderSignature: []byte("leader signature"),
			SoftwareVersion: nil,
		}
		err := mb2.CheckFieldsForNil()
		require.True(t, errors.Is(err, data.ErrNilValue))
		require.True(t, strings.Contains(err.Error(), "ChainID"))
	})

	t.Run("nil software version", func(t *testing.T) {
		t.Parallel()
		mb2 := &block.MetaBlockV3{
			PrevHash:        []byte("prev hash"),
			PrevRandSeed:    []byte("prev rand seed"),
			RandSeed:        []byte("rand seed"),
			LeaderSignature: []byte("leader signature"),
			ChainID:         []byte("chain"),
			SoftwareVersion: nil,
		}
		err := mb2.CheckFieldsForNil()
		require.True(t, errors.Is(err, data.ErrNilValue))
		require.True(t, strings.Contains(err.Error(), "SoftwareVersion"))
	})

	t.Run("nil LastExecutionResult", func(t *testing.T) {
		t.Parallel()
		mb2 := &block.MetaBlockV3{
			PrevHash:        []byte("prev hash"),
			PrevRandSeed:    []byte("prev rand seed"),
			RandSeed:        []byte("rand seed"),
			LeaderSignature: []byte("leader signature"),
			SoftwareVersion: []byte("v1.0.0"),
			ChainID:         []byte("chain"),
		}
		err := mb2.CheckFieldsForNil()
		require.True(t, strings.Contains(err.Error(), "LastExecutionResult"))
	})

	t.Run("valid header", func(t *testing.T) {
		t.Parallel()
		mb2 := &block.MetaBlockV3{
			PrevHash:            []byte("prev hash"),
			PrevRandSeed:        []byte("prev rand seed"),
			RandSeed:            []byte("rand seed"),
			LeaderSignature:     []byte("leader sig"),
			SoftwareVersion:     []byte("v1.0.0"),
			ChainID:             []byte("chain"),
			LastExecutionResult: &block.MetaExecutionResultInfo{},
		}
		err := mb2.CheckFieldsForNil()
		require.NoError(t, err)
	})
}

func TestMetaBlockV3_GetAccumulatedFees(t *testing.T) {
	t.Parallel()
	mb2 := &block.MetaBlockV3{}
	require.Nil(t, mb2.GetAccumulatedFees())
}

func TestMetaBlockV3_GetDeveloperFees(t *testing.T) {
	t.Parallel()
	mb2 := &block.MetaBlockV3{}
	require.Nil(t, mb2.GetDeveloperFees())
}

func TestMetaBlockV3_IsHeaderV3(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()
		var mb2 *block.MetaBlockV3
		require.False(t, mb2.IsHeaderV3())
	})

	t.Run("valid receiver", func(t *testing.T) {
		t.Parallel()
		mb2 := &block.MetaBlockV3{}
		require.True(t, mb2.IsHeaderV3())
	})
}

func TestMetaBlockV3_IsInterfaceNil(t *testing.T) {
	t.Parallel()

	var mb2 *block.MetaBlockV3
	require.True(t, mb2.IsInterfaceNil())

	mb2 = &block.MetaBlockV3{}
	require.False(t, mb2.IsInterfaceNil())
}

func TestMetaBlockV3_SetLastExecutionResultHandler(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()

		var header *block.MetaBlockV3
		require.Equal(t, data.ErrNilPointerReceiver, header.SetLastExecutionResultHandler(nil))
	})

	t.Run("nil exec result", func(t *testing.T) {
		t.Parallel()

		header := &block.MetaBlockV3{}
		require.Equal(t, data.ErrNilPointerDereference, header.SetLastExecutionResultHandler(nil))
	})

	t.Run("cast fails", func(t *testing.T) {
		t.Parallel()

		header := &block.MetaBlockV3{}
		require.Equal(t, data.ErrInvalidTypeAssertion, header.SetLastExecutionResultHandler(&block.MetaBlockV3{}))
	})

	t.Run("should work", func(t *testing.T) {
		t.Parallel()

		header := &block.MetaBlockV3{}
		execResult := &block.MetaExecutionResultInfo{}
		require.NoError(t, header.SetLastExecutionResultHandler(execResult))
		require.Equal(t, execResult, header.GetLastExecutionResult())
	})
}

func TestMetaBlockV3_SetExecutionResultsHandlers(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()

		var header *block.MetaBlockV3
		require.Equal(t, data.ErrNilPointerReceiver, header.SetExecutionResultsHandlers(nil))
	})

	t.Run("empty slice", func(t *testing.T) {
		t.Parallel()

		header := &block.MetaBlockV3{}
		require.NoError(t, header.SetExecutionResultsHandlers(nil))
		require.Nil(t, header.GetExecutionResults())
	})

	t.Run("invalid cast", func(t *testing.T) {
		t.Parallel()

		header := &block.MetaBlockV3{}
		handlers := []data.BaseExecutionResultHandler{
			&block.MetaExecutionResult{}, // ok
			&block.ExecutionResult{},     // cast fails
		}
		require.Equal(t, data.ErrInvalidTypeAssertion, header.SetExecutionResultsHandlers(handlers))
	})

	t.Run("nil handler", func(t *testing.T) {
		t.Parallel()

		header := &block.MetaBlockV3{}
		nilHandler := (*block.MetaExecutionResult)(nil)
		handlers := []data.BaseExecutionResultHandler{
			&block.MetaExecutionResult{}, // ok
			nilHandler,                   // nil
		}
		require.Equal(t, data.ErrNilPointerDereference, header.SetExecutionResultsHandlers(handlers))
	})

	t.Run("should work", func(t *testing.T) {
		t.Parallel()

		header := &block.MetaBlockV3{}
		handlers := []data.BaseExecutionResultHandler{
			&block.MetaExecutionResult{
				ReceiptsHash: []byte("receiptsHash1"),
			},
			&block.MetaExecutionResult{
				ReceiptsHash: []byte("receiptsHash2"),
			},
		}
		require.NoError(t, header.SetExecutionResultsHandlers(handlers))
		require.Equal(t, handlers, header.GetExecutionResultsHandlers())
	})
}

func TestMetaBlockV3_SetEpochStartHandler(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()

		var header *block.MetaBlockV3
		require.Equal(t, data.ErrNilPointerReceiver, header.SetEpochStartHandler(nil))
	})
	t.Run("valid receiver, nil epochStartHandler should return nil", func(t *testing.T) {
		t.Parallel()
		header := &block.MetaBlockV3{}
		require.Nil(t, header.SetEpochStartHandler(nil))
		require.Len(t, header.EpochStart.LastFinalizedHeaders, 0)
	})
	t.Run("valid receiver, nil EpochStart for epochStartHandler should return error", func(t *testing.T) {
		t.Parallel()

		header := &block.MetaBlockV3{}
		var nilValue *block.EpochStart
		err := header.SetEpochStartHandler(nilValue)
		require.Equal(t, data.ErrNilPointerDereference, err)
	})
	t.Run("valid receiver, non-nil epochStartHandler should set the field", func(t *testing.T) {
		t.Parallel()

		header := &block.MetaBlockV3{}
		epochStartHandler := &block.EpochStart{}
		err := header.SetEpochStartHandler(epochStartHandler)
		require.Nil(t, err)
		require.Equal(t, epochStartHandler, header.GetEpochStartHandler())
	})
}

func TestMetaBlockV3_GetShardInfoProposalHandlers(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()

		var mb3 *block.MetaBlockV3
		require.Nil(t, mb3.GetShardInfoProposalHandlers())
	})

	t.Run("should work", func(t *testing.T) {
		t.Parallel()

		shardDataProposal1 := block.ShardDataProposal{ShardID: 0, HeaderHash: []byte("shard1"), Nonce: 10, Round: 100, Epoch: 1}
		shardDataProposal2 := block.ShardDataProposal{ShardID: 1, HeaderHash: []byte("shard2"), Nonce: 20, Round: 200, Epoch: 2}
		mb3 := &block.MetaBlockV3{
			ShardInfoProposal: []block.ShardDataProposal{shardDataProposal1, shardDataProposal2},
		}
		expected := []data.ShardDataProposalHandler{&mb3.ShardInfoProposal[0], &mb3.ShardInfoProposal[1]}
		result := mb3.GetShardInfoProposalHandlers()
		require.Equal(t, expected, result)
	})
}

func TestMetaBlockV3_SetShardInfoProposalHandlers(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()

		var mb3 *block.MetaBlockV3
		err := mb3.SetShardInfoProposalHandlers(nil)
		require.Equal(t, data.ErrNilPointerReceiver, err)
	})

	t.Run("should work with nil", func(t *testing.T) {
		t.Parallel()

		mb3 := &block.MetaBlockV3{
			ShardInfoProposal: make([]block.ShardDataProposal, 2),
		}
		err := mb3.SetShardInfoProposalHandlers(nil)
		require.NoError(t, err)
		require.Nil(t, mb3.ShardInfoProposal)
	})

	t.Run("should error on list of nil shard data proposal handlers", func(t *testing.T) {
		t.Parallel()

		mb3 := &block.MetaBlockV3{}
		err := mb3.SetShardInfoProposalHandlers([]data.ShardDataProposalHandler{nil})
		require.Equal(t, data.ErrInvalidTypeAssertion, err)
	})

	t.Run("should error on list of nil shard data proposal", func(t *testing.T) {
		t.Parallel()

		mb3 := &block.MetaBlockV3{}
		var shardDataProposal *block.ShardDataProposal = nil
		err := mb3.SetShardInfoProposalHandlers([]data.ShardDataProposalHandler{shardDataProposal})
		require.Equal(t, data.ErrNilPointerDereference, err)
	})

	t.Run("should work with valid handlers", func(t *testing.T) {
		t.Parallel()

		shardDataProposal1 := &block.ShardDataProposal{ShardID: 0, HeaderHash: []byte("shard1"), Nonce: 10, Round: 100, Epoch: 1}
		shardDataProposal2 := &block.ShardDataProposal{ShardID: 1, HeaderHash: []byte("shard2"), Nonce: 20, Round: 200, Epoch: 2}
		handlers := []data.ShardDataProposalHandler{shardDataProposal1, shardDataProposal2}

		mb3 := &block.MetaBlockV3{}
		err := mb3.SetShardInfoProposalHandlers(handlers)
		require.NoError(t, err)
		require.Equal(t, 2, len(mb3.ShardInfoProposal))
		assert.True(t, shardDataProposal1.Equal(mb3.ShardInfoProposal[0]))
		assert.True(t, shardDataProposal2.Equal(mb3.ShardInfoProposal[1]))
	})
}

func TestMetaHeaderV3_checkBaseExecutionResultsIntegrity(t *testing.T) {
	t.Parallel()

	t.Run("nil base exec result", func(t *testing.T) {
		t.Parallel()

		metaV3 := &block.MetaBlockV3{}
		err := metaV3.CheckBaseExecutionResultIntegrity(nil)
		require.Equal(t, data.ErrNilValue, err)
	})
	t.Run("nil base exec result with reflect", func(t *testing.T) {
		t.Parallel()

		metaV3 := &block.MetaBlockV3{
			Round: 2,
			LastExecutionResult: &block.MetaExecutionResultInfo{
				NotarizedInRound: 1,
			},
			ExecutionResults: []*block.MetaExecutionResult{
				&block.MetaExecutionResult{},
			},
		}
		err := metaV3.CheckBaseExecutionResultIntegrity(metaV3.LastExecutionResult.ExecutionResult)
		require.Equal(t, data.ErrNilValue, err)
		err = metaV3.CheckBaseExecutionResultIntegrity(metaV3.ExecutionResults[0].ExecutionResult)
		require.Equal(t, data.ErrNilValue, err)
	})
	t.Run("invalid base execution result header hash", func(t *testing.T) {
		t.Parallel()

		metaV3 := &block.MetaBlockV3{
			Nonce: 1,
			Round: 1,
			Epoch: 1,
		}
		baseExecResult := &block.BaseExecutionResult{
			HeaderHash: []byte{},
		}
		err := metaV3.CheckBaseExecutionResultIntegrity(baseExecResult)
		require.True(t, errors.Is(err, data.ErrNilValue))
		require.True(t, strings.Contains(err.Error(), "HeaderHash"))
	})
	t.Run("invalid base execution result header nonce", func(t *testing.T) {
		t.Parallel()

		metaV3 := &block.MetaBlockV3{
			Nonce: 1,
			Round: 1,
			Epoch: 1,
		}
		baseExecResult := &block.BaseExecutionResult{
			HeaderHash:  []byte("header hash"),
			HeaderNonce: 1,
		}
		err := metaV3.CheckBaseExecutionResultIntegrity(baseExecResult)
		require.NotNil(t, err)
		require.True(t, strings.Contains(err.Error(), "HeaderNonce"))

		baseExecResult.HeaderNonce = 2
		err = metaV3.CheckBaseExecutionResultIntegrity(baseExecResult)
		require.NotNil(t, err)
		require.True(t, strings.Contains(err.Error(), "HeaderNonce"))
	})
	t.Run("invalid base execution result header round", func(t *testing.T) {
		t.Parallel()

		metaV3 := &block.MetaBlockV3{
			Nonce: 2,
			Round: 1,
			Epoch: 1,
		}
		baseExecResult := &block.BaseExecutionResult{
			HeaderHash:  []byte("header hash"),
			HeaderNonce: 1,
			HeaderRound: 1,
		}
		err := metaV3.CheckBaseExecutionResultIntegrity(baseExecResult)
		require.NotNil(t, err)
		require.True(t, strings.Contains(err.Error(), "HeaderRound"))

		baseExecResult.HeaderRound = 2
		err = metaV3.CheckBaseExecutionResultIntegrity(baseExecResult)
		require.NotNil(t, err)
		require.True(t, strings.Contains(err.Error(), "HeaderRound"))
	})
	t.Run("invalid base execution result header epoch", func(t *testing.T) {
		t.Parallel()

		metaV3 := &block.MetaBlockV3{
			Nonce: 2,
			Round: 2,
			Epoch: 1,
		}
		baseExecResult := &block.BaseExecutionResult{
			HeaderHash:  []byte("header hash"),
			HeaderNonce: 1,
			HeaderRound: 1,
			HeaderEpoch: 2,
		}
		err := metaV3.CheckBaseExecutionResultIntegrity(baseExecResult)
		require.NotNil(t, err)
		require.True(t, strings.Contains(err.Error(), "HeaderEpoch"))
	})
	t.Run("invalid base execution result root hash", func(t *testing.T) {
		t.Parallel()

		metaV3 := &block.MetaBlockV3{
			Nonce: 2,
			Round: 2,
			Epoch: 2,
		}
		baseExecResult := &block.BaseExecutionResult{
			HeaderHash:  []byte("header hash"),
			HeaderNonce: 1,
			HeaderRound: 1,
			HeaderEpoch: 1,
		}
		err := metaV3.CheckBaseExecutionResultIntegrity(baseExecResult)
		require.True(t, errors.Is(err, data.ErrNilValue))
		require.True(t, strings.Contains(err.Error(), "RootHash"))

		baseExecResult.RootHash = []byte{}
		err = metaV3.CheckBaseExecutionResultIntegrity(baseExecResult)
		require.True(t, errors.Is(err, data.ErrNilValue))
		require.True(t, strings.Contains(err.Error(), "RootHash"))
	})
	t.Run("should work", func(t *testing.T) {
		t.Parallel()

		metaV3 := &block.MetaBlockV3{
			Nonce: 2,
			Round: 2,
			Epoch: 2,
		}
		baseExecResult := &block.BaseExecutionResult{
			HeaderHash:  []byte("header hash"),
			HeaderNonce: 1,
			HeaderRound: 1,
			HeaderEpoch: 1,
			RootHash:    []byte("root hash"),
		}
		err := metaV3.CheckBaseExecutionResultIntegrity(baseExecResult)
		require.NoError(t, err)
	})
}

func TestMetaHeaderV3_CheckBaseMetaExecutionResultIntegrity(t *testing.T) {
	t.Parallel()

	t.Run("nil own base exec result", func(t *testing.T) {
		t.Parallel()

		metaV3 := &block.MetaBlockV3{}
		err := metaV3.CheckBaseMetaExecutionResultIntegrity(nil)
		require.Equal(t, data.ErrNilValue, err)
	})
	t.Run("with nil value in ValidatorStatsRootHash", func(t *testing.T) {
		t.Parallel()

		metaV3 := &block.MetaBlockV3{
			Nonce: 2,
			Round: 2,
			Epoch: 2,
		}
		metaExecResult := &block.BaseMetaExecutionResult{
			BaseExecutionResult: &block.BaseExecutionResult{
				HeaderHash:  []byte("header hash"),
				HeaderNonce: 1,
				HeaderRound: 1,
				HeaderEpoch: 1,
				RootHash:    []byte("root hash"),
			},
		}
		err := metaV3.CheckBaseMetaExecutionResultIntegrity(metaExecResult)
		require.ErrorIs(t, err, data.ErrNilValue)
		require.Contains(t, err.Error(), "ValidatorStatsRootHash")
	})
	t.Run("with nil value in AccumulatedFeesInEpoch", func(t *testing.T) {
		t.Parallel()

		metaV3 := &block.MetaBlockV3{
			Nonce: 2,
			Round: 2,
			Epoch: 2,
		}
		metaExecResult := &block.BaseMetaExecutionResult{
			BaseExecutionResult: &block.BaseExecutionResult{
				HeaderHash:  []byte("header hash"),
				HeaderNonce: 1,
				HeaderRound: 1,
				HeaderEpoch: 1,
				RootHash:    []byte("root hash"),
			},
			ValidatorStatsRootHash: []byte("validator stats root hash"),
		}
		err := metaV3.CheckBaseMetaExecutionResultIntegrity(metaExecResult)
		require.ErrorIs(t, err, data.ErrNilValue)
		require.Contains(t, err.Error(), "AccumulatedFeesInEpoch")
	})
	t.Run("with negative accumulated fees", func(t *testing.T) {
		t.Parallel()

		metaV3 := &block.MetaBlockV3{
			Nonce: 2,
			Round: 2,
			Epoch: 2,
		}
		metaExecResult := &block.BaseMetaExecutionResult{
			BaseExecutionResult: &block.BaseExecutionResult{
				HeaderHash:  []byte("header hash"),
				HeaderNonce: 1,
				HeaderRound: 1,
				HeaderEpoch: 1,
				RootHash:    []byte("root hash"),
			},
			ValidatorStatsRootHash: []byte("validator stats root hash"),
			AccumulatedFeesInEpoch: big.NewInt(-100),
		}
		err := metaV3.CheckBaseMetaExecutionResultIntegrity(metaExecResult)
		require.ErrorIs(t, err, data.ErrInvalidValue)
		require.Contains(t, err.Error(), "AccumulatedFeesInEpoch")
	})
	t.Run("with nil value in DevFeesInEpoch", func(t *testing.T) {
		t.Parallel()

		metaV3 := &block.MetaBlockV3{
			Nonce: 2,
			Round: 2,
			Epoch: 2,
		}
		metaExecResult := &block.BaseMetaExecutionResult{
			BaseExecutionResult: &block.BaseExecutionResult{
				HeaderHash:  []byte("header hash"),
				HeaderNonce: 1,
				HeaderRound: 1,
				HeaderEpoch: 1,
				RootHash:    []byte("root hash"),
			},
			ValidatorStatsRootHash: []byte("validator stats root hash"),
			AccumulatedFeesInEpoch: big.NewInt(100),
		}
		err := metaV3.CheckBaseMetaExecutionResultIntegrity(metaExecResult)
		require.ErrorIs(t, err, data.ErrNilValue)
		require.Contains(t, err.Error(), "DevFeesInEpoch")
	})
	t.Run("with negative developers fees", func(t *testing.T) {
		t.Parallel()

		metaV3 := &block.MetaBlockV3{
			Nonce: 2,
			Round: 2,
			Epoch: 2,
		}
		metaExecResult := &block.BaseMetaExecutionResult{
			BaseExecutionResult: &block.BaseExecutionResult{
				HeaderHash:  []byte("header hash"),
				HeaderNonce: 1,
				HeaderRound: 1,
				HeaderEpoch: 1,
				RootHash:    []byte("root hash"),
			},
			ValidatorStatsRootHash: []byte("validator stats root hash"),
			AccumulatedFeesInEpoch: big.NewInt(100),
			DevFeesInEpoch:         big.NewInt(-50),
		}
		err := metaV3.CheckBaseMetaExecutionResultIntegrity(metaExecResult)
		require.ErrorIs(t, err, data.ErrInvalidValue)
		require.Contains(t, err.Error(), "DevFeesInEpoch")
	})
	t.Run("with error in base execution result", func(t *testing.T) {
		t.Parallel()
		metaV3 := &block.MetaBlockV3{
			Nonce: 2,
			Round: 2,
			Epoch: 2,
		}
		metaExecResult := &block.BaseMetaExecutionResult{
			BaseExecutionResult: &block.BaseExecutionResult{
				HeaderHash: []byte{},
			},
			ValidatorStatsRootHash: []byte("validator stats root hash"),
			AccumulatedFeesInEpoch: big.NewInt(100),
			DevFeesInEpoch:         big.NewInt(50),
		}
		err := metaV3.CheckBaseMetaExecutionResultIntegrity(metaExecResult)
		require.ErrorIs(t, err, data.ErrNilValue)
		require.Contains(t, err.Error(), "BaseExecutionResult.HeaderHash")
	})
	t.Run("should work", func(t *testing.T) {
		t.Parallel()

		metaV3 := &block.MetaBlockV3{
			Nonce: 2,
			Round: 2,
			Epoch: 2,
		}
		metaExecResult := &block.BaseMetaExecutionResult{
			BaseExecutionResult: &block.BaseExecutionResult{
				HeaderHash:  []byte("header hash"),
				HeaderNonce: 1,
				HeaderRound: 1,
				HeaderEpoch: 1,
				RootHash:    []byte("root hash"),
			},
			ValidatorStatsRootHash: []byte("validator stats root hash"),
			AccumulatedFeesInEpoch: big.NewInt(100),
			DevFeesInEpoch:         big.NewInt(50),
		}
		err := metaV3.CheckBaseMetaExecutionResultIntegrity(metaExecResult)
		require.NoError(t, err)
	})
}

func TestMetaHeaderV3_checkExecutionResultsIntegrity(t *testing.T) {
	t.Parallel()
	t.Run("nil execution result", func(t *testing.T) {
		t.Parallel()

		metaV3 := &block.MetaBlockV3{}
		metaV3.ExecutionResults = make([]*block.MetaExecutionResult, 1)
		assert.Equal(t, metaV3.ExecutionResults[0].IsInterfaceNil(), true)
		err := metaV3.CheckExecutionResultsIntegrity()
		require.Error(t, err)
		require.ErrorIs(t, err, data.ErrNilValue)
	})
	t.Run("with nil receipts hash", func(t *testing.T) {
		t.Parallel()
		metaV3 := createValidMetaHeaderV3ToTest()
		metaV3.ExecutionResults[0].ReceiptsHash = nil
		err := metaV3.CheckExecutionResultsIntegrity()
		require.ErrorIs(t, err, data.ErrNilValue)
		require.Contains(t, err.Error(), "ReceiptsHash")
	})
	t.Run("with nil accumulated fees", func(t *testing.T) {
		t.Parallel()
		metaV3 := createValidMetaHeaderV3ToTest()
		metaV3.ExecutionResults[0].AccumulatedFees = nil
		err := metaV3.CheckExecutionResultsIntegrity()
		require.ErrorIs(t, err, data.ErrNilValue)
		require.Contains(t, err.Error(), "AccumulatedFees")
	})
	t.Run("with negative accumulated fees", func(t *testing.T) {
		t.Parallel()

		metaV3 := createValidMetaHeaderV3ToTest()
		metaV3.ExecutionResults[0].AccumulatedFees = big.NewInt(-100)
		err := metaV3.CheckExecutionResultsIntegrity()
		require.ErrorIs(t, err, data.ErrInvalidValue)
		require.Contains(t, err.Error(), "AccumulatedFees")
	})
	t.Run("with nil developers fees", func(t *testing.T) {
		t.Parallel()

		metaV3 := createValidMetaHeaderV3ToTest()
		metaV3.ExecutionResults[0].DeveloperFees = nil
		err := metaV3.CheckExecutionResultsIntegrity()
		require.ErrorIs(t, err, data.ErrNilValue)
		require.Contains(t, err.Error(), "DeveloperFees")
	})
	t.Run("with negative developers fees", func(t *testing.T) {
		t.Parallel()

		metaV3 := createValidMetaHeaderV3ToTest()
		metaV3.ExecutionResults[0].DeveloperFees = big.NewInt(-50)
		err := metaV3.CheckExecutionResultsIntegrity()
		require.ErrorIs(t, err, data.ErrInvalidValue)
		require.Contains(t, err.Error(), "DeveloperFees")
	})
	t.Run("invalid execution result base exec result", func(t *testing.T) {
		t.Parallel()

		metaV3 := createValidMetaHeaderV3ToTest()
		metaV3.ExecutionResults[0].ExecutionResult = nil
		err := metaV3.CheckExecutionResultsIntegrity()
		require.ErrorIs(t, err, data.ErrNilValue)
	})
	t.Run("invalid execution result base exec result first one good", func(t *testing.T) {
		t.Parallel()

		metaV3 := createValidMetaHeaderV3ToTest()
		metaV3.ExecutionResults[1].ExecutionResult = nil
		err := metaV3.CheckExecutionResultsIntegrity()
		require.ErrorIs(t, err, data.ErrNilValue)
	})

	t.Run("should work", func(t *testing.T) {
		t.Parallel()

		metaV3 := createValidMetaHeaderV3ToTest()
		err := metaV3.CheckExecutionResultsIntegrity()
		require.NoError(t, err)
	})
}

func TestMetaHeaderV3_CheckLastExecutionResultIntegrity(t *testing.T) {
	t.Parallel()

	t.Run("nil last execution result", func(t *testing.T) {
		t.Parallel()
		metaV3 := &block.MetaBlockV3{}
		err := metaV3.CheckLastExecutionResultIntegrity()
		require.True(t, errors.Is(err, data.ErrNilValue))
		require.True(t, strings.Contains(err.Error(), "LastExecutionResult"))
	})
	t.Run("invalid last execution result base exec result", func(t *testing.T) {
		t.Parallel()

		metaV3 := &block.MetaBlockV3{
			Nonce: 2,
			Round: 2,
			Epoch: 2,
			LastExecutionResult: &block.MetaExecutionResultInfo{
				NotarizedInRound: 1,
			},
		}
		err := metaV3.CheckLastExecutionResultIntegrity()
		require.ErrorIs(t, err, data.ErrNilValue)
	})

	t.Run("should work", func(t *testing.T) {
		t.Parallel()
		metaV3 := createValidMetaHeaderV3ToTest()
		err := metaV3.CheckLastExecutionResultIntegrity()
		require.NoError(t, err)
	})
}

func TestMetaHeaderV3_CheckFieldsIntegrity(t *testing.T) {
	t.Parallel()
	t.Run("nil header", func(t *testing.T) {
		t.Parallel()
		var metaV3 *block.MetaBlockV3
		err := metaV3.CheckFieldsIntegrity()
		require.Equal(t, data.ErrNilPointerReceiver, err)
	})

	t.Run("not nil reserved field", func(t *testing.T) {
		t.Parallel()
		metaV3 := createValidMetaHeaderV3ToTest()
		metaV3.Reserved = []byte("not nil reserved")
		err := metaV3.CheckFieldsIntegrity()
		require.Error(t, err)
		require.ErrorIs(t, err, data.ErrNotNilValue)
	})

	t.Run("invalid shard info proposal", func(t *testing.T) {
		t.Parallel()
		metaV3 := createValidMetaHeaderV3ToTest()
		metaV3.ShardInfoProposal = make([]block.ShardDataProposal, 0)
		err := metaV3.CheckFieldsIntegrity()
		require.Error(t, err)
		require.Contains(t, err.Error(), "ShardInfoProposal")
	})

	t.Run("genesis round should work", func(t *testing.T) {
		t.Parallel()
		metaV3 := createValidMetaHeaderV3ToTestForGenesisRound()
		err := metaV3.CheckFieldsIntegrity()
		require.NoError(t, err)
	})
	t.Run("nil last execution result", func(t *testing.T) {
		t.Parallel()
		metaV3 := &block.MetaBlockV3{
			Nonce: 2,
			Round: 2,
			Epoch: 2,
		}
		err := metaV3.CheckFieldsIntegrity()
		require.Error(t, err)
		require.ErrorIs(t, err, data.ErrNilValue)
	})

	t.Run("invalid execution results", func(t *testing.T) {
		t.Parallel()
		metaV3 := createValidMetaHeaderV3ToTest()
		metaV3.ExecutionResults[0].ExecutionResult = nil
		err := metaV3.CheckFieldsIntegrity()
		require.Error(t, err)
		require.ErrorIs(t, err, data.ErrNilValue)
	})
	t.Run("invalid last execution result", func(t *testing.T) {
		t.Parallel()
		metaV3 := createValidMetaHeaderV3ToTest()
		metaV3.LastExecutionResult.ExecutionResult = nil
		err := metaV3.CheckFieldsIntegrity()
		require.Error(t, err)
		require.ErrorIs(t, err, data.ErrNilValue)
	})
	t.Run("invalid round in last execution result", func(t *testing.T) {
		t.Parallel()
		metaV3 := createValidMetaHeaderV3ToTest()
		metaV3.LastExecutionResult.NotarizedInRound = 1300
		err := metaV3.CheckFieldsIntegrity()
		require.Error(t, err)
		require.Contains(t, err.Error(), "LastExecutionResult.NotarizedInRound")
	})
	t.Run("should work", func(t *testing.T) {
		t.Parallel()
		metaV3 := createValidMetaHeaderV3ToTest()
		err := metaV3.CheckFieldsIntegrity()
		require.NoError(t, err)
	})
}

func createValidMetaHeaderV3ToTest() *block.MetaBlockV3 {
	return &block.MetaBlockV3{
		Nonce:           42,
		Epoch:           2,
		Round:           15,
		TimestampMs:     123456789,
		PrevHash:        []byte("prev_hash"),
		PrevRandSeed:    []byte("prev_seed"),
		RandSeed:        []byte("new_seed"),
		ChainID:         []byte("chain-id"),
		SoftwareVersion: []byte("v1.0.0"),

		MiniBlockHeaders: []block.MiniBlockHeader{
			{Hash: []byte("meta-to-s0"), SenderShardID: core.MetachainShardId, ReceiverShardID: 0},
			{Hash: []byte("meta-to-s1"), SenderShardID: core.MetachainShardId, ReceiverShardID: 1},
		},

		ShardInfo: []block.ShardData{
			{
				ShardID:    0,
				Round:      10,
				Nonce:      41,
				Epoch:      1,
				HeaderHash: []byte("shard0-hash"),
				ShardMiniBlockHeaders: []block.MiniBlockHeader{
					{SenderShardID: 0, ReceiverShardID: 1, Hash: []byte("s0-to-s1")},
				},
			},
			{
				ShardID:    1,
				Round:      11,
				Nonce:      40,
				Epoch:      1,
				HeaderHash: []byte("shard1-hash"),
				ShardMiniBlockHeaders: []block.MiniBlockHeader{
					{SenderShardID: 1, ReceiverShardID: 0, Hash: []byte("s1-to-s0")},
				},
			},
		},
		ShardInfoProposal: []block.ShardDataProposal{
			{ShardID: 0, HeaderHash: []byte("shard-0-hash"), Nonce: 41, Round: 10, Epoch: 1},
			{ShardID: 1, HeaderHash: []byte("shard-1-hash"), Nonce: 40, Round: 11, Epoch: 1},
		},
		ExecutionResults: []*block.MetaExecutionResult{
			{
				ExecutionResult: &block.BaseMetaExecutionResult{
					BaseExecutionResult: &block.BaseExecutionResult{
						HeaderHash:  []byte("hdr-hash-10"),
						HeaderNonce: 39,
						HeaderRound: 10,
						HeaderEpoch: 1,
						RootHash:    []byte("root-hash-10"),
					},
					AccumulatedFeesInEpoch: big.NewInt(1000),
					DevFeesInEpoch:         big.NewInt(100),
					ValidatorStatsRootHash: []byte("validator-stats-root-hash-10"),
				},
				ReceiptsHash:    []byte("receipts-hash-10"),
				AccumulatedFees: big.NewInt(1000),
				DeveloperFees:   big.NewInt(100),
			},
			{
				ExecutionResult: &block.BaseMetaExecutionResult{
					BaseExecutionResult: &block.BaseExecutionResult{
						HeaderHash:  []byte("hdr-hash-11"),
						HeaderNonce: 40,
						HeaderRound: 11,
						HeaderEpoch: 1,
						RootHash:    []byte("root-hash-11"),
					},
					AccumulatedFeesInEpoch: big.NewInt(2000),
					DevFeesInEpoch:         big.NewInt(200),
					ValidatorStatsRootHash: []byte("validator-stats-root-hash-11"),
				},
				ReceiptsHash:    []byte("receipts-hash-11"),
				AccumulatedFees: big.NewInt(2000),
				DeveloperFees:   big.NewInt(200),
			},
			{
				ExecutionResult: &block.BaseMetaExecutionResult{
					BaseExecutionResult: &block.BaseExecutionResult{
						HeaderHash:  []byte("hdr-hash-last"),
						HeaderNonce: 41,
						HeaderRound: 12,
						HeaderEpoch: 2,
						RootHash:    []byte("root-hash-last"),
					},
					AccumulatedFeesInEpoch: big.NewInt(3000),
					DevFeesInEpoch:         big.NewInt(300),
					ValidatorStatsRootHash: []byte("validator-stats-root-hash-last"),
				},
				ReceiptsHash:    []byte("receipts-hash-last"),
				AccumulatedFees: big.NewInt(3000),
				DeveloperFees:   big.NewInt(300),
			},
		},

		LastExecutionResult: &block.MetaExecutionResultInfo{
			NotarizedInRound: 14,
			ExecutionResult: &block.BaseMetaExecutionResult{
				BaseExecutionResult: &block.BaseExecutionResult{
					HeaderHash:  []byte("hdr-hash-last"),
					HeaderNonce: 41,
					HeaderRound: 12,
					HeaderEpoch: 2,
					RootHash:    []byte("root-hash-last"),
				},
				AccumulatedFeesInEpoch: big.NewInt(3000),
				DevFeesInEpoch:         big.NewInt(300),
				ValidatorStatsRootHash: []byte("validator-stats-root-hash-last"),
			},
		},
	}
}

func createValidMetaHeaderV3ToTestForGenesisRound() *block.MetaBlockV3 {
	return &block.MetaBlockV3{
		Nonce:           0,
		Round:           0,
		Epoch:           0,
		RandSeed:        []byte("rand seed for genesis"),
		LeaderSignature: []byte("leader signature for genesis"),
		SoftwareVersion: []byte("v1.0.0"),
	}
}
