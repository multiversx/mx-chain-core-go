package block_test

import (
	"errors"
	"math/big"
	"strings"
	"testing"

	"github.com/multiversx/mx-chain-core-go/data"
	"github.com/multiversx/mx-chain-core-go/data/block"
	"github.com/stretchr/testify/require"
)

func TestHeaderV3_GetRootHash(t *testing.T) {
	t.Parallel()
	hv3 := &block.HeaderV3{}
	require.Nil(t, hv3.GetRootHash())
}

func TestHeaderV3_GetPubKeysBitmap(t *testing.T) {
	t.Parallel()
	hv3 := &block.HeaderV3{}
	require.Nil(t, hv3.GetPubKeysBitmap())
}

func TestHeaderV3_GetSignature(t *testing.T) {
	t.Parallel()
	hv3 := &block.HeaderV3{}
	require.Nil(t, hv3.GetSignature())
}

func TestHeaderV3_GetTimeStamp(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()
		var hv3 *block.HeaderV3
		require.Equal(t, uint64(0), hv3.GetTimeStamp())
	})

	t.Run("should work", func(t *testing.T) {
		t.Parallel()
		timestamp := uint64(12345)
		hv3 := &block.HeaderV3{TimeStampMs: timestamp}
		require.Equal(t, timestamp, hv3.GetTimeStamp())
	})
}

func TestHeaderV3_GetMiniBlockHeadersWithDst(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()
		var hv3 *block.HeaderV3
		require.Nil(t, hv3.GetMiniBlockHeadersWithDst(0))
	})

	t.Run("no mini blocks with dest", func(t *testing.T) {
		t.Parallel()
		hv3 := &block.HeaderV3{
			MiniBlockHeaders: []block.MiniBlockHeader{
				{ReceiverShardID: 1, SenderShardID: 0, Hash: []byte("hash1")},
			},
		}
		result := hv3.GetMiniBlockHeadersWithDst(2)
		require.Empty(t, result)
	})

	t.Run("should work", func(t *testing.T) {
		t.Parallel()
		hash1 := []byte("hash1")
		hash2 := []byte("hash2")
		hv3 := &block.HeaderV3{
			MiniBlockHeaders: []block.MiniBlockHeader{
				{ReceiverShardID: 1, SenderShardID: 0, Hash: hash1},
				{ReceiverShardID: 1, SenderShardID: 2, Hash: hash2},
			},
		}
		expected := map[string]uint32{
			string(hash1): 0,
			string(hash2): 2,
		}
		result := hv3.GetMiniBlockHeadersWithDst(1)
		require.Equal(t, expected, result)
	})
}

func TestHeaderV3_GetOrderedCrossMiniblocksWithDst(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()
		var hv3 *block.HeaderV3
		require.Nil(t, hv3.GetOrderedCrossMiniblocksWithDst(0))
	})

	t.Run("should work", func(t *testing.T) {
		t.Parallel()
		hash1 := []byte("hash1")
		hash2 := []byte("hash2")
		hv3 := &block.HeaderV3{
			Round: 42,
			MiniBlockHeaders: []block.MiniBlockHeader{
				{ReceiverShardID: 1, SenderShardID: 0, Hash: hash1},
				{ReceiverShardID: 2, SenderShardID: 1, Hash: hash2},
			},
		}
		expected := []*data.MiniBlockInfo{
			{Hash: hash2, SenderShardID: 1, Round: 42},
		}
		result := hv3.GetOrderedCrossMiniblocksWithDst(2)
		require.Equal(t, expected, result)
	})
}

func TestHeaderV3_GetMiniBlockHeadersHashes(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()
		var hv3 *block.HeaderV3
		require.Nil(t, hv3.GetMiniBlockHeadersHashes())
	})

	t.Run("should work", func(t *testing.T) {
		t.Parallel()
		hash1 := []byte("hash1")
		hash2 := []byte("hash2")
		hv3 := &block.HeaderV3{
			MiniBlockHeaders: []block.MiniBlockHeader{
				{Hash: hash1},
				{Hash: hash2},
			},
		}
		expected := [][]byte{hash1, hash2}
		result := hv3.GetMiniBlockHeadersHashes()
		require.Equal(t, expected, result)
	})
}

func TestHeaderV3_GetMiniBlockHeaderHandlers(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()
		var hv3 *block.HeaderV3
		require.Nil(t, hv3.GetMiniBlockHeaderHandlers())
	})

	t.Run("should work", func(t *testing.T) {
		t.Parallel()
		hv3 := &block.HeaderV3{
			MiniBlockHeaders: []block.MiniBlockHeader{
				{Hash: []byte("hash1")},
				{Hash: []byte("hash2")},
			},
		}
		result := hv3.GetMiniBlockHeaderHandlers()
		require.Equal(t, 2, len(result))
		require.Equal(t, hv3.MiniBlockHeaders[0].Hash, result[0].GetHash())
		require.Equal(t, hv3.MiniBlockHeaders[1].Hash, result[1].GetHash())
	})
}

func TestHeaderV3_HasScheduledSupport(t *testing.T) {
	t.Parallel()
	hv3 := &block.HeaderV3{}
	require.False(t, hv3.HasScheduledSupport())
}

func TestHeaderV3_GetAdditionalData(t *testing.T) {
	t.Parallel()
	hv3 := &block.HeaderV3{}
	require.Nil(t, hv3.GetAdditionalData())
}

func TestHeaderV3_HasScheduledMiniBlocks(t *testing.T) {
	t.Parallel()
	hv3 := &block.HeaderV3{}
	require.False(t, hv3.HasScheduledMiniBlocks())
}

func TestHeaderV3_SetAccumulatedFees(t *testing.T) {
	t.Parallel()
	hv3 := &block.HeaderV3{}
	err := hv3.SetAccumulatedFees(big.NewInt(100))
	require.NoError(t, err)
}

func TestHeaderV3_SetDeveloperFees(t *testing.T) {
	t.Parallel()
	hv3 := &block.HeaderV3{}
	err := hv3.SetDeveloperFees(big.NewInt(50))
	require.NoError(t, err)
}

func TestHeaderV3_SetShardID(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()
		var hv3 *block.HeaderV3
		err := hv3.SetShardID(1)
		require.Equal(t, data.ErrNilPointerReceiver, err)
	})

	t.Run("should work", func(t *testing.T) {
		t.Parallel()
		hv3 := &block.HeaderV3{}
		require.NoError(t, hv3.SetShardID(1))
		require.Equal(t, uint32(1), hv3.ShardID)
	})
}

func TestHeaderV3_SetNonce(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()
		var hv3 *block.HeaderV3
		err := hv3.SetNonce(1)
		require.Equal(t, data.ErrNilPointerReceiver, err)
	})

	t.Run("should work", func(t *testing.T) {
		t.Parallel()
		hv3 := &block.HeaderV3{}
		require.NoError(t, hv3.SetNonce(42))
		require.Equal(t, uint64(42), hv3.Nonce)
	})
}

func TestHeaderV3_SetEpoch(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()
		var hv3 *block.HeaderV3
		err := hv3.SetEpoch(1)
		require.Equal(t, data.ErrNilPointerReceiver, err)
	})

	t.Run("should work", func(t *testing.T) {
		t.Parallel()
		hv3 := &block.HeaderV3{}
		require.NoError(t, hv3.SetEpoch(2))
		require.Equal(t, uint32(2), hv3.Epoch)
	})
}

func TestHeaderV3_SetRound(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()
		var hv3 *block.HeaderV3
		err := hv3.SetRound(1)
		require.Equal(t, data.ErrNilPointerReceiver, err)
	})

	t.Run("should work", func(t *testing.T) {
		t.Parallel()
		hv3 := &block.HeaderV3{}
		require.NoError(t, hv3.SetRound(42))
		require.Equal(t, uint64(42), hv3.Round)
	})
}

func TestHeaderV3_SetTimeStamp(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()
		var hv3 *block.HeaderV3
		err := hv3.SetTimeStamp(12345)
		require.Equal(t, data.ErrNilPointerReceiver, err)
	})

	t.Run("should work", func(t *testing.T) {
		t.Parallel()
		hv3 := &block.HeaderV3{}
		require.NoError(t, hv3.SetTimeStamp(12345))
		require.Equal(t, uint64(12345), hv3.TimeStampMs)
	})
}

func TestHeaderV3_SetRootHash(t *testing.T) {
	t.Parallel()
	hv3 := &block.HeaderV3{}
	err := hv3.SetRootHash([]byte("root"))
	require.NoError(t, err)
}

func TestHeaderV3_SetPrevHash(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()
		var hv3 *block.HeaderV3
		err := hv3.SetPrevHash([]byte("prev"))
		require.Equal(t, data.ErrNilPointerReceiver, err)
	})

	t.Run("should work", func(t *testing.T) {
		t.Parallel()
		prevHash := []byte("prev hash")
		hv3 := &block.HeaderV3{}
		require.NoError(t, hv3.SetPrevHash(prevHash))
		require.Equal(t, prevHash, hv3.PrevHash)
	})
}

func TestHeaderV3_SetPrevRandSeed(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()
		var hv3 *block.HeaderV3
		err := hv3.SetPrevRandSeed([]byte("seed"))
		require.Equal(t, data.ErrNilPointerReceiver, err)
	})

	t.Run("should work", func(t *testing.T) {
		t.Parallel()
		seed := []byte("prev rand seed")
		hv3 := &block.HeaderV3{}
		require.NoError(t, hv3.SetPrevRandSeed(seed))
		require.Equal(t, seed, hv3.PrevRandSeed)
	})
}

func TestHeaderV3_SetRandSeed(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()
		var hv3 *block.HeaderV3
		err := hv3.SetRandSeed([]byte("seed"))
		require.Equal(t, data.ErrNilPointerReceiver, err)
	})

	t.Run("should work", func(t *testing.T) {
		t.Parallel()
		seed := []byte("rand seed")
		hv3 := &block.HeaderV3{}
		require.NoError(t, hv3.SetRandSeed(seed))
		require.Equal(t, seed, hv3.RandSeed)
	})
}

func TestHeaderV3_SetPubKeysBitmap(t *testing.T) {
	t.Parallel()
	hv3 := &block.HeaderV3{}
	err := hv3.SetPubKeysBitmap([]byte("bitmap"))
	require.NoError(t, err)
}

func TestHeaderV3_SetSignature(t *testing.T) {
	t.Parallel()
	hv3 := &block.HeaderV3{}
	err := hv3.SetSignature([]byte("signature"))
	require.NoError(t, err)
}

func TestHeaderV3_SetLeaderSignature(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()
		var hv3 *block.HeaderV3
		err := hv3.SetLeaderSignature([]byte("sig"))
		require.Equal(t, data.ErrNilPointerReceiver, err)
	})

	t.Run("should work", func(t *testing.T) {
		t.Parallel()
		sig := []byte("leader signature")
		hv3 := &block.HeaderV3{}
		require.NoError(t, hv3.SetLeaderSignature(sig))
		require.Equal(t, sig, hv3.LeaderSignature)
	})
}

func TestHeaderV3_SetChainID(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()
		var hv3 *block.HeaderV3
		err := hv3.SetChainID([]byte("chain"))
		require.Equal(t, data.ErrNilPointerReceiver, err)
	})

	t.Run("should work", func(t *testing.T) {
		t.Parallel()
		chainID := []byte("chain ID")
		hv3 := &block.HeaderV3{}
		require.NoError(t, hv3.SetChainID(chainID))
		require.Equal(t, chainID, hv3.ChainID)
	})
}

func TestHeaderV3_SetSoftwareVersion(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()
		var hv3 *block.HeaderV3
		err := hv3.SetSoftwareVersion([]byte("v1.0.0"))
		require.Equal(t, data.ErrNilPointerReceiver, err)
	})

	t.Run("should work", func(t *testing.T) {
		t.Parallel()
		version := []byte("v1.2.3")
		hv3 := &block.HeaderV3{}
		require.NoError(t, hv3.SetSoftwareVersion(version))
		require.Equal(t, version, hv3.SoftwareVersion)
	})
}

func TestHeaderV3_SetTxCount(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()
		var hv3 *block.HeaderV3
		err := hv3.SetTxCount(10)
		require.Equal(t, data.ErrNilPointerReceiver, err)
	})

	t.Run("should work", func(t *testing.T) {
		t.Parallel()
		hv3 := &block.HeaderV3{}
		require.NoError(t, hv3.SetTxCount(42))
		require.Equal(t, uint32(42), hv3.TxCount)
	})
}

func TestHeaderV3_SetMiniBlockHeaderHandlers(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()
		var hv3 *block.HeaderV3
		err := hv3.SetMiniBlockHeaderHandlers(nil)
		require.Equal(t, data.ErrNilPointerReceiver, err)
	})

	t.Run("nil handlers", func(t *testing.T) {
		t.Parallel()
		hv3 := &block.HeaderV3{}
		err := hv3.SetMiniBlockHeaderHandlers(nil)
		require.NoError(t, err)
	})

	t.Run("should work", func(t *testing.T) {
		t.Parallel()
		hv3 := &block.HeaderV3{}
		mbh1 := &block.MiniBlockHeader{Hash: []byte("hash1")}
		mbh2 := &block.MiniBlockHeader{Hash: []byte("hash2")}
		handlers := []data.MiniBlockHeaderHandler{mbh1, mbh2}

		err := hv3.SetMiniBlockHeaderHandlers(handlers)
		require.NoError(t, err)
		require.Equal(t, 2, len(hv3.MiniBlockHeaders))
		require.Equal(t, mbh1.Hash, hv3.MiniBlockHeaders[0].Hash)
		require.Equal(t, mbh2.Hash, hv3.MiniBlockHeaders[1].Hash)
	})
}

func TestHeaderV3_SetReceiptsHash(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()
		var hv3 *block.HeaderV3
		err := hv3.SetReceiptsHash([]byte("receipts"))
		require.Equal(t, data.ErrNilPointerReceiver, err)
	})

	t.Run("should work", func(t *testing.T) {
		t.Parallel()
		receiptsHash := []byte("receipts hash")
		hv3 := &block.HeaderV3{}
		require.NoError(t, hv3.SetReceiptsHash(receiptsHash))
		require.Equal(t, receiptsHash, hv3.ReceiptsHash)
	})
}

func TestHeaderV3_SetScheduledRootHash(t *testing.T) {
	t.Parallel()
	hv3 := &block.HeaderV3{}
	err := hv3.SetScheduledRootHash([]byte("scheduled"))
	require.NoError(t, err)
}

func TestHeaderV3_ValidateHeaderVersion(t *testing.T) {
	t.Parallel()
	hv3 := &block.HeaderV3{}
	err := hv3.ValidateHeaderVersion()
	require.NoError(t, err)
}

func TestHeaderV3_SetAdditionalData(t *testing.T) {
	t.Parallel()
	hv3 := &block.HeaderV3{}
	err := hv3.SetAdditionalData(nil)
	require.NoError(t, err)
}

func TestHeaderV3_IsStartOfEpochBlock(t *testing.T) {
	t.Parallel()
	hv3 := &block.HeaderV3{}
	require.False(t, hv3.IsStartOfEpochBlock())

	hv3 = nil
	require.False(t, hv3.IsStartOfEpochBlock())
}

func TestHeaderV3_ShallowClone(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()
		var hv3 *block.HeaderV3
		clone := hv3.ShallowClone()
		require.Nil(t, clone)
	})

	t.Run("should work", func(t *testing.T) {
		t.Parallel()
		hv3 := &block.HeaderV3{
			ShardID:         1,
			Nonce:           42,
			Epoch:           2,
			Round:           100,
			TimeStampMs:     12345,
			PrevHash:        []byte("prev hash"),
			PrevRandSeed:    []byte("prev rand seed"),
			RandSeed:        []byte("rand seed"),
			ChainID:         []byte("chain"),
			SoftwareVersion: []byte("v1.0.0"),
		}
		clone := hv3.ShallowClone()
		require.NotNil(t, clone)
		require.Equal(t, hv3, clone)
		require.False(t, &hv3.MiniBlockHeaders == &clone.(*block.HeaderV3).MiniBlockHeaders)
	})
}

func TestHeaderV3_CheckFieldsForNil(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()
		var hv3 *block.HeaderV3
		err := hv3.CheckFieldsForNil()
		require.Equal(t, data.ErrNilPointerReceiver, err)
	})

	t.Run("nil prev hash", func(t *testing.T) {
		t.Parallel()
		hv3 := &block.HeaderV3{
			PrevHash: nil,
		}
		err := hv3.CheckFieldsForNil()
		require.True(t, errors.Is(err, data.ErrNilValue))
		require.True(t, strings.Contains(err.Error(), "PrevHash"))
	})

	t.Run("nil prev rand seed", func(t *testing.T) {
		t.Parallel()
		hv3 := &block.HeaderV3{
			PrevHash:     []byte("prev hash"),
			PrevRandSeed: nil,
		}
		err := hv3.CheckFieldsForNil()
		require.True(t, errors.Is(err, data.ErrNilValue))
		require.True(t, strings.Contains(err.Error(), "PrevRandSeed"))
	})

	t.Run("nil rand seed", func(t *testing.T) {
		t.Parallel()
		hv3 := &block.HeaderV3{
			PrevHash:     []byte("prev hash"),
			PrevRandSeed: []byte("prev rand seed"),
			RandSeed:     nil,
		}
		err := hv3.CheckFieldsForNil()
		require.True(t, errors.Is(err, data.ErrNilValue))
		require.True(t, strings.Contains(err.Error(), "RandSeed"))
	})

	t.Run("nil leader sig", func(t *testing.T) {
		t.Parallel()
		hv3 := &block.HeaderV3{
			PrevHash:        []byte("prev hash"),
			PrevRandSeed:    []byte("prev rand seed"),
			RandSeed:        []byte("rand seed"),
			LeaderSignature: nil,
		}
		err := hv3.CheckFieldsForNil()
		require.True(t, errors.Is(err, data.ErrNilValue))
		require.True(t, strings.Contains(err.Error(), "LeaderSignature"))
	})

	t.Run("nil software version", func(t *testing.T) {
		t.Parallel()
		hv3 := &block.HeaderV3{
			PrevHash:        []byte("prev hash"),
			PrevRandSeed:    []byte("prev rand seed"),
			RandSeed:        []byte("rand seed"),
			LeaderSignature: []byte("leader signature"),
			SoftwareVersion: nil,
		}
		err := hv3.CheckFieldsForNil()
		require.True(t, errors.Is(err, data.ErrNilValue))
		require.True(t, strings.Contains(err.Error(), "SoftwareVersion"))
	})

	t.Run("nil last exec result", func(t *testing.T) {
		t.Parallel()
		hv3 := &block.HeaderV3{
			PrevHash:            []byte("prev hash"),
			PrevRandSeed:        []byte("prev rand seed"),
			RandSeed:            []byte("rand seed"),
			LeaderSignature:     []byte("leader signature"),
			SoftwareVersion:     []byte("v1.0.0"),
			LastExecutionResult: nil,
		}
		err := hv3.CheckFieldsForNil()
		require.True(t, errors.Is(err, data.ErrNilValue))
		require.True(t, strings.Contains(err.Error(), "LastExecutionResult"))
	})

	t.Run("valid header", func(t *testing.T) {
		t.Parallel()
		hv3 := &block.HeaderV3{
			PrevHash:            []byte("prev hash"),
			PrevRandSeed:        []byte("prev rand seed"),
			RandSeed:            []byte("rand seed"),
			LeaderSignature:     []byte("leader sig"),
			SoftwareVersion:     []byte("v1.0.0"),
			LastExecutionResult: &block.ExecutionResultInfo{},
		}
		err := hv3.CheckFieldsForNil()
		require.NoError(t, err)
	})
}

func TestHeaderV3_GetAccumulatedFees(t *testing.T) {
	t.Parallel()
	hv3 := &block.HeaderV3{}
	require.Equal(t, big.NewInt(0), hv3.GetAccumulatedFees())
}

func TestHeaderV3_GetDeveloperFees(t *testing.T) {
	t.Parallel()
	hv3 := &block.HeaderV3{}
	require.Equal(t, big.NewInt(0), hv3.GetDeveloperFees())
}

func TestHeaderV3_IsInterfaceNil(t *testing.T) {
	t.Parallel()

	var hv3 *block.HeaderV3
	require.True(t, hv3.IsInterfaceNil())

	hv3 = &block.HeaderV3{}
	require.False(t, hv3.IsInterfaceNil())
}
