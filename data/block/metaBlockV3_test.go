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
		expected := []data.MetaExecutionResultHandler{
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
				NotarizedAtHeaderHash: []byte("notarizedHash"),
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
		metaHdr.MiniBlockHeaders = append(metaHdr.MiniBlockHeaders, mbsFromMetaToShard0...)
		metaHdr.MiniBlockHeaders = append(metaHdr.MiniBlockHeaders, mbsFromMetaToShard1...)

		mbDst0 := metaHdr.GetMiniBlockHeadersWithDst(0)
		assert.Equal(t, len(mbsFromMetaToShard0), len(mbDst0))
		mbDst1 := metaHdr.GetMiniBlockHeadersWithDst(1)
		assert.Equal(t, len(shardMBHeader)+len(mbsFromMetaToShard1), len(mbDst1))
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

	metaHdr.ShardInfo = append(metaHdr.ShardInfo, shData1, shData2, shData3, shData4, shData5)

	metaHdr.MiniBlockHeaders = append(metaHdr.MiniBlockHeaders, block.MiniBlockHeader{
		Hash:            []byte("hash6"),
		SenderShardID:   core.MetachainShardId,
		ReceiverShardID: 1,
	})

	metaHdr.MiniBlockHeaders = append(metaHdr.MiniBlockHeaders, block.MiniBlockHeader{
		Hash:            []byte("hash7"),
		SenderShardID:   core.MetachainShardId,
		ReceiverShardID: core.AllShardId,
	})

	metaHdr.MiniBlockHeaders = append(metaHdr.MiniBlockHeaders, block.MiniBlockHeader{
		Hash:            []byte("hash8"),
		SenderShardID:   core.MetachainShardId,
		ReceiverShardID: 2,
	})

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

	t.Run("valid header", func(t *testing.T) {
		t.Parallel()
		mb2 := &block.MetaBlockV3{
			PrevHash:        []byte("prev hash"),
			PrevRandSeed:    []byte("prev rand seed"),
			RandSeed:        []byte("rand seed"),
			LeaderSignature: []byte("leader sig"),
			SoftwareVersion: []byte("v1.0.0"),
			ChainID:         []byte("chain"),
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
