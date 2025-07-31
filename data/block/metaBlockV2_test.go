package block_test

import (
	"errors"
	"github.com/multiversx/mx-chain-core-go/core"
	"github.com/multiversx/mx-chain-core-go/data"
	"github.com/multiversx/mx-chain-core-go/data/block"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"math/big"
	"strings"
	"testing"
)

func TestMetaBlockV2_GetRootHash(t *testing.T) {
	t.Parallel()
	mb2 := &block.MetaBlockV2{}
	require.Nil(t, mb2.GetRootHash())
}

func TestMetaBlockV2_GetPubKeysBitmap(t *testing.T) {
	t.Parallel()
	mb2 := &block.MetaBlockV2{}
	require.Nil(t, mb2.GetPubKeysBitmap())
}

func TestMetaBlockV2_GetSignature(t *testing.T) {
	t.Parallel()
	mb2 := &block.MetaBlockV2{}
	require.Nil(t, mb2.GetSignature())
}

func TestMetaBlockV2_GetTimeStamp(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()
		var mb2 *block.MetaBlockV2
		require.Equal(t, uint64(0), mb2.GetTimeStamp())
	})

	t.Run("should work", func(t *testing.T) {
		t.Parallel()
		timestamp := uint64(12345)
		mb2 := &block.MetaBlockV2{TimestampMs: timestamp}
		require.Equal(t, timestamp, mb2.GetTimeStamp())
	})
}

func TestMetaBlockV2_GetMiniBlockHeadersWithDst(t *testing.T) {
	t.Parallel()

	metaHdr := &block.MetaBlockV2{Round: 15}
	metaHdr.ShardInfo = make([]block.ShardData, 0)

	shardMBHeader := make([]block.MiniBlockHeader, 0)
	shMBHdr1 := block.MiniBlockHeader{SenderShardID: 0, ReceiverShardID: 1, Hash: []byte("hash1")}
	shMBHdr2 := block.MiniBlockHeader{SenderShardID: 0, ReceiverShardID: 1, Hash: []byte("hash2")}
	shardMBHeader = append(shardMBHeader, shMBHdr1, shMBHdr2)

	shData1 := block.ShardData{ShardID: 0, HeaderHash: []byte("sh"), ShardMiniBlockHeaders: shardMBHeader}
	metaHdr.ShardInfo = append(metaHdr.ShardInfo, shData1)

	shData2 := block.ShardData{ShardID: 1, HeaderHash: []byte("sh"), ShardMiniBlockHeaders: shardMBHeader}
	metaHdr.ShardInfo = append(metaHdr.ShardInfo, shData2)

	mbDst0 := metaHdr.GetMiniBlockHeadersWithDst(0)
	assert.Equal(t, 0, len(mbDst0))
	mbDst1 := metaHdr.GetMiniBlockHeadersWithDst(1)
	assert.Equal(t, len(shardMBHeader), len(mbDst1))
}

func TestMetaBlockV2_GetOrderedCrossMiniblocksWithDst(t *testing.T) {
	t.Parallel()

	metaHdr := &block.MetaBlock{Round: 6}
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

func TestMetaBlockV2_GetMiniBlockHeadersHashes(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()
		var mb2 *block.MetaBlockV2
		require.Nil(t, mb2.GetMiniBlockHeadersHashes())
	})

	t.Run("should work", func(t *testing.T) {
		t.Parallel()
		hash1 := []byte("hash1")
		hash2 := []byte("hash2")
		mb2 := &block.MetaBlockV2{
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

func TestMetaBlockV2_GetMiniBlockHeaderHandlers(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()
		var mb2 *block.MetaBlockV2
		require.Nil(t, mb2.GetMiniBlockHeaderHandlers())
	})

	t.Run("should work", func(t *testing.T) {
		t.Parallel()
		mb2 := &block.MetaBlockV2{
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

func TestMetaBlockV2_HasScheduledSupport(t *testing.T) {
	t.Parallel()
	mb2 := &block.MetaBlockV2{}
	require.False(t, mb2.HasScheduledSupport())
}

func TestMetaBlockV2_GetAdditionalData(t *testing.T) {
	t.Parallel()
	mb2 := &block.MetaBlockV2{}
	require.Nil(t, mb2.GetAdditionalData())
}

func TestMetaBlockV2_HasScheduledMiniBlocks(t *testing.T) {
	t.Parallel()
	mb2 := &block.MetaBlockV2{}
	require.False(t, mb2.HasScheduledMiniBlocks())
}

func TestMetaBlockV2_SetAccumulatedFees(t *testing.T) {
	t.Parallel()
	mb2 := &block.MetaBlockV2{}
	err := mb2.SetAccumulatedFees(big.NewInt(100))
	require.NoError(t, err)
}

func TestMetaBlockV2_SetDeveloperFees(t *testing.T) {
	t.Parallel()
	mb2 := &block.MetaBlockV2{}
	err := mb2.SetDeveloperFees(big.NewInt(50))
	require.NoError(t, err)
}

func TestMetaBlockV2_SetShardIDWillDoNothing(t *testing.T) {
	t.Parallel()

	mb2 := &block.MetaBlockV2{}
	err := mb2.SetShardID(0)
	require.NoError(t, err)

	shardID := mb2.GetShardID()
	require.Equal(t, core.MetachainShardId, shardID)
}

func TestMetaBlockV2_SetNonce(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()
		var mb2 *block.MetaBlockV2
		err := mb2.SetNonce(1)
		require.Equal(t, data.ErrNilPointerReceiver, err)
	})

	t.Run("should work", func(t *testing.T) {
		t.Parallel()
		mb2 := &block.MetaBlockV2{}
		require.NoError(t, mb2.SetNonce(42))
		require.Equal(t, uint64(42), mb2.Nonce)
	})
}

func TestMetaBlockV2_SetEpoch(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()
		var mb2 *block.MetaBlockV2
		err := mb2.SetEpoch(1)
		require.Equal(t, data.ErrNilPointerReceiver, err)
	})

	t.Run("should work", func(t *testing.T) {
		t.Parallel()
		mb2 := &block.MetaBlockV2{}
		require.NoError(t, mb2.SetEpoch(2))
		require.Equal(t, uint32(2), mb2.Epoch)
	})
}

func TestMetaBlockV2_SetRound(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()
		var mb2 *block.MetaBlockV2
		err := mb2.SetRound(1)
		require.Equal(t, data.ErrNilPointerReceiver, err)
	})

	t.Run("should work", func(t *testing.T) {
		t.Parallel()
		mb2 := &block.MetaBlockV2{}
		require.NoError(t, mb2.SetRound(42))
		require.Equal(t, uint64(42), mb2.Round)
	})
}

func TestMetaBlockV2_SetTimeStamp(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()
		var mb2 *block.MetaBlockV2
		err := mb2.SetTimeStamp(12345)
		require.Equal(t, data.ErrNilPointerReceiver, err)
	})

	t.Run("should work", func(t *testing.T) {
		t.Parallel()
		mb2 := &block.MetaBlockV2{}
		require.NoError(t, mb2.SetTimeStamp(12345))
		require.Equal(t, uint64(12345), mb2.TimestampMs)
	})
}

func TestMetaBlockV2_SetRootHash(t *testing.T) {
	t.Parallel()
	mb2 := &block.MetaBlockV2{}
	err := mb2.SetRootHash([]byte("root"))
	require.NoError(t, err)
}

func TestMetaBlockV2_SetPrevHash(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()
		var mb2 *block.MetaBlockV2
		err := mb2.SetPrevHash([]byte("prev"))
		require.Equal(t, data.ErrNilPointerReceiver, err)
	})

	t.Run("should work", func(t *testing.T) {
		t.Parallel()
		prevHash := []byte("prev hash")
		mb2 := &block.MetaBlockV2{}
		require.NoError(t, mb2.SetPrevHash(prevHash))
		require.Equal(t, prevHash, mb2.PrevHash)
	})
}

func TestMetaBlockV2_SetPrevRandSeed(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()
		var mb2 *block.MetaBlockV2
		err := mb2.SetPrevRandSeed([]byte("seed"))
		require.Equal(t, data.ErrNilPointerReceiver, err)
	})

	t.Run("should work", func(t *testing.T) {
		t.Parallel()
		seed := []byte("prev rand seed")
		mb2 := &block.MetaBlockV2{}
		require.NoError(t, mb2.SetPrevRandSeed(seed))
		require.Equal(t, seed, mb2.PrevRandSeed)
	})
}

func TestMetaBlockV2_SetRandSeed(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()
		var mb2 *block.MetaBlockV2
		err := mb2.SetRandSeed([]byte("seed"))
		require.Equal(t, data.ErrNilPointerReceiver, err)
	})

	t.Run("should work", func(t *testing.T) {
		t.Parallel()
		seed := []byte("rand seed")
		mb2 := &block.MetaBlockV2{}
		require.NoError(t, mb2.SetRandSeed(seed))
		require.Equal(t, seed, mb2.RandSeed)
	})
}

func TestMetaBlockV2_SetPubKeysBitmap(t *testing.T) {
	t.Parallel()
	mb2 := &block.MetaBlockV2{}
	err := mb2.SetPubKeysBitmap([]byte("bitmap"))
	require.NoError(t, err)
}

func TestMetaBlockV2_SetSignature(t *testing.T) {
	t.Parallel()
	mb2 := &block.MetaBlockV2{}
	err := mb2.SetSignature([]byte("signature"))
	require.NoError(t, err)
}

func TestMetaBlockV2_SetLeaderSignature(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()
		var mb2 *block.MetaBlockV2
		err := mb2.SetLeaderSignature([]byte("sig"))
		require.Equal(t, data.ErrNilPointerReceiver, err)
	})

	t.Run("should work", func(t *testing.T) {
		t.Parallel()
		sig := []byte("leader signature")
		mb2 := &block.MetaBlockV2{}
		require.NoError(t, mb2.SetLeaderSignature(sig))
		require.Equal(t, sig, mb2.LeaderSignature)
	})
}

func TestMetaBlockV2_SetChainID(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()
		var mb2 *block.MetaBlockV2
		err := mb2.SetChainID([]byte("chain"))
		require.Equal(t, data.ErrNilPointerReceiver, err)
	})

	t.Run("should work", func(t *testing.T) {
		t.Parallel()
		chainID := []byte("chain ID")
		mb2 := &block.MetaBlockV2{}
		require.NoError(t, mb2.SetChainID(chainID))
		require.Equal(t, chainID, mb2.ChainID)
	})
}

func TestMetaBlockV2_SetSoftwareVersion(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()
		var mb2 *block.MetaBlockV2
		err := mb2.SetSoftwareVersion([]byte("v1.0.0"))
		require.Equal(t, data.ErrNilPointerReceiver, err)
	})

	t.Run("should work", func(t *testing.T) {
		t.Parallel()
		version := []byte("v1.2.3")
		mb2 := &block.MetaBlockV2{}
		require.NoError(t, mb2.SetSoftwareVersion(version))
		require.Equal(t, version, mb2.SoftwareVersion)
	})
}

func TestMetaBlockV2_SetTxCount(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()
		var mb2 *block.MetaBlockV2
		err := mb2.SetTxCount(10)
		require.Equal(t, data.ErrNilPointerReceiver, err)
	})

	t.Run("should work", func(t *testing.T) {
		t.Parallel()
		mb2 := &block.MetaBlockV2{}
		require.NoError(t, mb2.SetTxCount(42))
		require.Equal(t, uint32(42), mb2.TxCount)
	})
}

func TestMetaBlockV2_SetMiniBlockHeaderHandlers(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()
		var mb2 *block.MetaBlockV2
		err := mb2.SetMiniBlockHeaderHandlers(nil)
		require.Equal(t, data.ErrNilPointerReceiver, err)
	})

	t.Run("nil handlers", func(t *testing.T) {
		t.Parallel()
		mb2 := &block.MetaBlockV2{}
		err := mb2.SetMiniBlockHeaderHandlers(nil)
		require.NoError(t, err)
	})

	t.Run("should work", func(t *testing.T) {
		t.Parallel()
		mb2 := &block.MetaBlockV2{}
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

func TestMetaBlockV2_SetScheduledRootHash(t *testing.T) {
	t.Parallel()
	mb2 := &block.MetaBlockV2{}
	err := mb2.SetScheduledRootHash([]byte("scheduled"))
	require.Equal(t, data.ErrScheduledRootHashNotSupported, err)
}

func TestMetaBlockV2_ValidateHeaderVersion(t *testing.T) {
	t.Parallel()
	mb2 := &block.MetaBlockV2{}
	err := mb2.ValidateHeaderVersion()
	require.NoError(t, err)
}

func TestMetaBlockV2_SetAdditionalData(t *testing.T) {
	t.Parallel()
	mb2 := &block.MetaBlockV2{}
	err := mb2.SetAdditionalData(nil)
	require.NoError(t, err)
}

func TestMetaBlockV2_IsStartOfEpochBlock(t *testing.T) {
	t.Parallel()
	mb2 := &block.MetaBlockV2{}
	require.False(t, mb2.IsStartOfEpochBlock())

	mb2 = nil
	require.False(t, mb2.IsStartOfEpochBlock())
}

func TestMetaBlockV2_ShallowClone(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()
		var mb2 *block.MetaBlockV2
		clone := mb2.ShallowClone()
		require.Nil(t, clone)
	})

	t.Run("should work", func(t *testing.T) {
		t.Parallel()
		mb2 := &block.HeaderV3{
			ShardID:         1,
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
		require.False(t, &mb2.MiniBlockHeaders == &clone.(*block.HeaderV3).MiniBlockHeaders)
	})
}

func TestMetaBlockV2_CheckFieldsForNil(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()
		var mb2 *block.MetaBlockV2
		err := mb2.CheckFieldsForNil()
		require.Equal(t, data.ErrNilPointerReceiver, err)
	})

	t.Run("nil prev hash", func(t *testing.T) {
		t.Parallel()
		mb2 := &block.HeaderV3{
			PrevHash: nil,
		}
		err := mb2.CheckFieldsForNil()
		require.True(t, errors.Is(err, data.ErrNilValue))
		require.True(t, strings.Contains(err.Error(), "PrevHash"))
	})

	t.Run("nil prev rand seed", func(t *testing.T) {
		t.Parallel()
		mb2 := &block.HeaderV3{
			PrevHash:     []byte("prev hash"),
			PrevRandSeed: nil,
		}
		err := mb2.CheckFieldsForNil()
		require.True(t, errors.Is(err, data.ErrNilValue))
		require.True(t, strings.Contains(err.Error(), "PrevRandSeed"))
	})

	t.Run("nil rand seed", func(t *testing.T) {
		t.Parallel()
		mb2 := &block.HeaderV3{
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
		mb2 := &block.HeaderV3{
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
		mb2 := &block.HeaderV3{
			PrevHash:        []byte("prev hash"),
			PrevRandSeed:    []byte("prev rand seed"),
			RandSeed:        []byte("rand seed"),
			LeaderSignature: []byte("leader signature"),
			SoftwareVersion: nil,
		}
		err := mb2.CheckFieldsForNil()
		require.True(t, errors.Is(err, data.ErrNilValue))
		require.True(t, strings.Contains(err.Error(), "SoftwareVersion"))
	})

	t.Run("nil last exec result", func(t *testing.T) {
		t.Parallel()
		mb2 := &block.HeaderV3{
			PrevHash:            []byte("prev hash"),
			PrevRandSeed:        []byte("prev rand seed"),
			RandSeed:            []byte("rand seed"),
			LeaderSignature:     []byte("leader signature"),
			SoftwareVersion:     []byte("v1.0.0"),
			LastExecutionResult: nil,
		}
		err := mb2.CheckFieldsForNil()
		require.True(t, errors.Is(err, data.ErrNilValue))
		require.True(t, strings.Contains(err.Error(), "LastExecutionResult"))
	})

	t.Run("valid header", func(t *testing.T) {
		t.Parallel()
		mb2 := &block.HeaderV3{
			PrevHash:            []byte("prev hash"),
			PrevRandSeed:        []byte("prev rand seed"),
			RandSeed:            []byte("rand seed"),
			LeaderSignature:     []byte("leader sig"),
			SoftwareVersion:     []byte("v1.0.0"),
			LastExecutionResult: &block.ExecutionResultInfo{},
		}
		err := mb2.CheckFieldsForNil()
		require.NoError(t, err)
	})
}

func TestMetaBlockV2_GetAccumulatedFees(t *testing.T) {
	t.Parallel()
	mb2 := &block.MetaBlockV2{}
	require.Equal(t, big.NewInt(0), mb2.GetAccumulatedFees())
}

func TestMetaBlockV2_GetDeveloperFees(t *testing.T) {
	t.Parallel()
	mb2 := &block.MetaBlockV2{}
	require.Equal(t, big.NewInt(0), mb2.GetDeveloperFees())
}

func TestMetaBlockV2_IsInterfaceNil(t *testing.T) {
	t.Parallel()

	var mb2 *block.MetaBlockV2
	require.True(t, mb2.IsInterfaceNil())

	mb2 = &block.MetaBlockV2{}
	require.False(t, mb2.IsInterfaceNil())
}
