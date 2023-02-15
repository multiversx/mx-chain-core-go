package block_test

import (
	"math/big"
	"reflect"
	"testing"

	"github.com/multiversx/mx-chain-core-go/core"
	"github.com/multiversx/mx-chain-core-go/data"
	"github.com/multiversx/mx-chain-core-go/data/block"
	"github.com/multiversx/mx-chain-core-go/data/headerVersionData"
	"github.com/stretchr/testify/require"
)

func TestHeaderV2_GetEpochNilPointerReceiverOrInnerHeader(t *testing.T) {
	t.Parallel()

	var h *block.HeaderV2
	require.Equal(t, uint32(0), h.GetEpoch())

	h = &block.HeaderV2{
		Header: nil,
	}
	require.Equal(t, uint32(0), h.GetEpoch())
}

func TestHeaderV2_GetEpoch(t *testing.T) {
	t.Parallel()

	epoch := uint32(1)
	h := &block.HeaderV2{
		Header: &block.Header{
			Epoch: epoch,
		},
	}

	require.Equal(t, epoch, h.GetEpoch())
}

func TestHeaderV2_GetShardNilPointerReceiverOrInnerHeader(t *testing.T) {
	t.Parallel()

	var h *block.HeaderV2
	require.Equal(t, uint32(0), h.GetShardID())

	h = &block.HeaderV2{
		Header: nil,
	}
	require.Equal(t, uint32(0), h.GetShardID())
}

func TestHeaderV2_GetShard(t *testing.T) {
	t.Parallel()

	shardId := uint32(2)
	h := &block.HeaderV2{
		Header: &block.Header{
			ShardID: shardId,
		},
	}

	require.Equal(t, shardId, h.GetShardID())
}
func TestHeaderV2_GetNonceNilPointerReceiverOrInnerHeader(t *testing.T) {
	t.Parallel()

	var h *block.HeaderV2
	require.Equal(t, uint64(0), h.GetNonce())

	h = &block.HeaderV2{
		Header: nil,
	}
	require.Equal(t, uint64(0), h.GetNonce())
}

func TestHeaderV2_GetNonce(t *testing.T) {
	t.Parallel()

	nonce := uint64(2)
	h := &block.HeaderV2{
		Header: &block.Header{
			Nonce: nonce,
		},
	}

	require.Equal(t, nonce, h.GetNonce())
}

func TestHeaderV2_GetPrevHashNilPointerReceiverOrInnerHeader(t *testing.T) {
	t.Parallel()

	var h *block.HeaderV2
	require.Equal(t, []byte(nil), h.GetPrevHash())

	h = &block.HeaderV2{
		Header: nil,
	}
	require.Equal(t, []byte(nil), h.GetPrevHash())
}

func TestHeaderV2_GetPrevHash(t *testing.T) {
	t.Parallel()

	prevHash := []byte("prev hash")
	h := &block.HeaderV2{
		Header: &block.Header{
			PrevHash: prevHash,
		},
	}

	require.Equal(t, prevHash, h.GetPrevHash())
}

func TestHeaderV2_GetPrevRandSeedNilPointerReceiverOrInnerHeader(t *testing.T) {
	t.Parallel()

	var h *block.HeaderV2
	require.Equal(t, []byte(nil), h.GetPrevRandSeed())

	h = &block.HeaderV2{
		Header: nil,
	}
	require.Equal(t, []byte(nil), h.GetPrevRandSeed())
}

func TestHeaderV2_GetPrevRandSeed(t *testing.T) {
	t.Parallel()

	prevRandSeed := []byte("prev random seed")
	h := &block.HeaderV2{
		Header: &block.Header{
			PrevRandSeed: prevRandSeed,
		},
	}

	require.Equal(t, prevRandSeed, h.GetPrevRandSeed())
}

func TestHeaderV2_GetRandSeedNilPointerReceiverOrInnerHeader(t *testing.T) {
	t.Parallel()

	var h *block.HeaderV2
	require.Equal(t, []byte(nil), h.GetRandSeed())

	h = &block.HeaderV2{
		Header: nil,
	}
	require.Equal(t, []byte(nil), h.GetRandSeed())
}

func TestHeaderV2_GetRandSeed(t *testing.T) {
	t.Parallel()

	randSeed := []byte("random seed")
	h := &block.HeaderV2{
		Header: &block.Header{
			RandSeed: randSeed,
		},
	}

	require.Equal(t, randSeed, h.GetRandSeed())
}
func TestHeaderV2_GetPubKeysBitmapNilPointerReceiverOrInnerHeader(t *testing.T) {
	t.Parallel()

	var h *block.HeaderV2
	require.Equal(t, []byte(nil), h.GetPubKeysBitmap())

	h = &block.HeaderV2{
		Header: nil,
	}
	require.Equal(t, []byte(nil), h.GetPubKeysBitmap())
}

func TestHeaderV2_GetPubKeysBitmap(t *testing.T) {
	t.Parallel()

	pubKeysBitmap := []byte{10, 11, 12, 13}
	h := &block.HeaderV2{
		Header: &block.Header{
			PubKeysBitmap: pubKeysBitmap,
		},
	}

	require.Equal(t, pubKeysBitmap, h.GetPubKeysBitmap())
}

func TestHeaderV2_GetRootHashNilPointerReceiverOrInnerHeader(t *testing.T) {
	t.Parallel()

	var h *block.HeaderV2
	require.Equal(t, []byte(nil), h.GetRootHash())

	h = &block.HeaderV2{
		Header: nil,
	}
	require.Equal(t, []byte(nil), h.GetRootHash())
}

func TestHeaderV2_GetRootHash(t *testing.T) {
	t.Parallel()

	rootHash := []byte("root hash")
	h := &block.HeaderV2{
		Header: &block.Header{
			RootHash: rootHash,
		},
	}

	require.Equal(t, rootHash, h.GetRootHash())
}

func TestHeaderV2_GetRoundNilPointerReceiverOrInnerHeader(t *testing.T) {
	t.Parallel()

	var h *block.HeaderV2
	require.Equal(t, uint64(0), h.GetRound())

	h = &block.HeaderV2{
		Header: nil,
	}
	require.Equal(t, uint64(0), h.GetRound())
}

func TestHeaderV2_GetRound(t *testing.T) {
	t.Parallel()

	round := uint64(1234)
	h := &block.HeaderV2{
		Header: &block.Header{
			Round: round,
		},
	}

	require.Equal(t, round, h.GetRound())
}

func TestHeaderV2_GetSignatureNilPointerReceiverOrInnerHeader(t *testing.T) {
	t.Parallel()

	var h *block.HeaderV2
	require.Equal(t, []byte(nil), h.GetSignature())

	h = &block.HeaderV2{
		Header: nil,
	}
	require.Equal(t, []byte(nil), h.GetSignature())
}

func TestHeaderV2_GetSignature(t *testing.T) {
	t.Parallel()

	signature := []byte("signature")
	h := &block.HeaderV2{
		Header: &block.Header{
			Signature: signature,
		},
	}

	require.Equal(t, signature, h.GetSignature())
}

func TestHeaderV2_GetLeaderSignature(t *testing.T) {
	t.Parallel()

	signature := []byte("signature")
	h := &block.HeaderV2{
		Header: &block.Header{
			LeaderSignature: signature,
		},
	}

	require.Equal(t, signature, h.GetLeaderSignature())
}

func TestHeaderV2_GetChainID(t *testing.T) {
	t.Parallel()

	chainId := []byte("chainId")
	h := &block.HeaderV2{
		Header: &block.Header{
			ChainID: chainId,
		},
	}

	require.Equal(t, chainId, h.GetChainID())
}

func TestHeaderV2_GetSoftwareVersion(t *testing.T) {
	t.Parallel()

	softwareVersion := []byte("softwareVersion")
	h := &block.HeaderV2{
		Header: &block.Header{
			SoftwareVersion: softwareVersion,
		},
	}

	require.Equal(t, softwareVersion, h.GetSoftwareVersion())
}

func TestHeaderV2_GetReceiptHash(t *testing.T) {
	t.Parallel()

	receiptHash := []byte("receiptHash")
	h := &block.HeaderV2{
		Header: &block.Header{
			ReceiptsHash: receiptHash,
		},
	}

	require.Equal(t, receiptHash, h.GetReceiptsHash())
}

func TestHeaderV2_GetAccumulatedFees(t *testing.T) {
	t.Parallel()

	accumulatedFees := big.NewInt(10)
	h := &block.HeaderV2{
		Header: &block.Header{
			AccumulatedFees: accumulatedFees,
		},
	}

	require.Equal(t, accumulatedFees, h.GetAccumulatedFees())
}

func TestHeaderV2_GetDeveloperFees(t *testing.T) {
	t.Parallel()

	developerFees := big.NewInt(10)
	h := &block.HeaderV2{
		Header: &block.Header{
			DeveloperFees: developerFees,
		},
	}

	require.Equal(t, developerFees, h.GetDeveloperFees())
}

func TestHeaderV2_GetReserved(t *testing.T) {
	t.Parallel()

	reserved := []byte("reserved")
	h := &block.HeaderV2{
		Header: &block.Header{
			Reserved: reserved,
		},
	}

	require.Equal(t, reserved, h.GetReserved())
}

func TestHeaderV2_GetMetaBlockHashes(t *testing.T) {
	t.Parallel()

	metablockHashes := [][]byte{
		[]byte("hash1"),
		[]byte("hash2"),
	}
	h := &block.HeaderV2{
		Header: &block.Header{
			MetaBlockHashes: metablockHashes,
		},
	}

	require.True(t, reflect.DeepEqual(metablockHashes, h.GetMetaBlockHashes()))
}

func TestHeaderV2_GetEpochStartMetaHash(t *testing.T) {
	t.Parallel()

	epochStartMetaHash := []byte("hash1")

	h := &block.HeaderV2{
		Header: &block.Header{
			EpochStartMetaHash: epochStartMetaHash,
		},
	}

	require.Equal(t, epochStartMetaHash, h.GetEpochStartMetaHash())
}

func TestHeaderV2_SetLeaderSignature(t *testing.T) {
	t.Parallel()

	leaderSignature := []byte("leaderSig")

	h := &block.HeaderV2{
		Header: &block.Header{},
	}
	err := h.SetLeaderSignature(leaderSignature)

	require.Nil(t, err)
	require.Equal(t, leaderSignature, h.GetLeaderSignature())
}

func TestHeaderV2_SetChainID(t *testing.T) {
	t.Parallel()

	chainId := []byte("chainId")

	h := &block.HeaderV2{
		Header: &block.Header{},
	}
	err := h.SetChainID(chainId)

	require.Nil(t, err)
	require.Equal(t, chainId, h.GetChainID())
}

func TestHeaderV2_SetSoftwareVersion(t *testing.T) {
	t.Parallel()

	softwareVersion := []byte("softwareVersion")

	h := &block.HeaderV2{
		Header: &block.Header{},
	}
	err := h.SetSoftwareVersion(softwareVersion)

	require.Nil(t, err)
	require.Equal(t, softwareVersion, h.GetSoftwareVersion())
}

func TestHeaderV2_SetAccumulatedFees(t *testing.T) {
	t.Parallel()

	accumulatedFees := big.NewInt(10)

	h := &block.HeaderV2{
		Header: &block.Header{},
	}
	err := h.SetAccumulatedFees(accumulatedFees)

	require.Nil(t, err)
	require.Equal(t, accumulatedFees, h.GetAccumulatedFees())
}

func TestHeaderV2_SetDeveloperFees(t *testing.T) {
	t.Parallel()

	developerFees := big.NewInt(10)

	h := &block.HeaderV2{
		Header: &block.Header{},
	}
	err := h.SetDeveloperFees(developerFees)

	require.Nil(t, err)
	require.Equal(t, developerFees, h.GetDeveloperFees())
}

func TestHeaderV2_SetShardID(t *testing.T) {
	t.Parallel()

	shardId := uint32(1)

	h := &block.HeaderV2{
		Header: &block.Header{},
	}
	err := h.SetShardID(shardId)

	require.Nil(t, err)
	require.Equal(t, shardId, h.GetShardID())
}

func TestHeaderV2_GetTxCountNilPointerReceiverOrInnerHeader(t *testing.T) {
	t.Parallel()

	var h *block.HeaderV2
	require.Equal(t, uint32(0), h.GetTxCount())

	h = &block.HeaderV2{
		Header: nil,
	}
	require.Equal(t, uint32(0), h.GetTxCount())
}

func TestHeaderV2_GetTxCount(t *testing.T) {
	t.Parallel()

	txCount := uint32(10)
	h := &block.HeaderV2{
		Header: &block.Header{
			TxCount: txCount,
		},
	}

	require.Equal(t, txCount, h.GetTxCount())
}

func TestHeaderV2_SetEpochNilPointerReceiverOrInnerHeader(t *testing.T) {
	t.Parallel()

	var h *block.HeaderV2
	require.Equal(t, data.ErrNilPointerReceiver, h.SetEpoch(1))

	h = &block.HeaderV2{
		Header: nil,
	}
	require.Equal(t, data.ErrNilPointerReceiver, h.SetEpoch(1))
}

func TestHeaderV2_SetEpoch(t *testing.T) {
	t.Parallel()

	epoch := uint32(10)
	h := &block.HeaderV2{
		Header: &block.Header{},
	}

	err := h.SetEpoch(epoch)
	require.Nil(t, err)
	require.Equal(t, epoch, h.GetEpoch())
}

func TestHeaderV2_SetNonceNilPointerReceiverOrInnerHeader(t *testing.T) {
	t.Parallel()

	var h *block.HeaderV2
	require.Equal(t, data.ErrNilPointerReceiver, h.SetNonce(1))

	h = &block.HeaderV2{
		Header: nil,
	}
	require.Equal(t, data.ErrNilPointerReceiver, h.SetNonce(1))
}

func TestHeaderV2_SetNonce(t *testing.T) {
	t.Parallel()

	nonce := uint64(11)
	h := &block.HeaderV2{
		Header: &block.Header{},
	}

	err := h.SetNonce(nonce)
	require.Nil(t, err)
	require.Equal(t, nonce, h.GetNonce())
}

func TestHeaderV2_SetPrevHashNilPointerReceiverOrInnerHeader(t *testing.T) {
	t.Parallel()

	var h *block.HeaderV2
	require.Equal(t, data.ErrNilPointerReceiver, h.SetPrevHash([]byte("prev hash")))

	h = &block.HeaderV2{
		Header: nil,
	}
	require.Equal(t, data.ErrNilPointerReceiver, h.SetPrevHash([]byte("prev hash")))
}

func TestHeaderV2_SetPrevHash(t *testing.T) {
	t.Parallel()

	prevHash := []byte("prev hash")
	h := &block.HeaderV2{
		Header: &block.Header{},
	}

	err := h.SetPrevHash(prevHash)
	require.Nil(t, err)
	require.Equal(t, prevHash, h.GetPrevHash())
}

func TestHeaderV2_SetPrevRandSeedNilPointerReceiverOrInnerHeader(t *testing.T) {
	t.Parallel()

	var h *block.HeaderV2
	require.Equal(t, data.ErrNilPointerReceiver, h.SetPrevRandSeed([]byte("prev rand seed")))

	h = &block.HeaderV2{
		Header: nil,
	}
	require.Equal(t, data.ErrNilPointerReceiver, h.SetPrevRandSeed([]byte("prev rand seed")))
}

func TestHeaderV2_SetPrevRandSeed(t *testing.T) {
	t.Parallel()

	prevRandSeed := []byte("prev random seed")
	h := &block.HeaderV2{
		Header: &block.Header{},
	}

	err := h.SetPrevRandSeed(prevRandSeed)
	require.Nil(t, err)
	require.Equal(t, prevRandSeed, h.GetPrevRandSeed())
}

func TestHeaderV2_SetRandSeedNilPointerReceiverOrInnerHeader(t *testing.T) {
	t.Parallel()

	var h *block.HeaderV2
	require.Equal(t, data.ErrNilPointerReceiver, h.SetRandSeed([]byte("rand seed")))

	h = &block.HeaderV2{
		Header: nil,
	}
	require.Equal(t, data.ErrNilPointerReceiver, h.SetRandSeed([]byte("rand seed")))
}

func TestHeaderV2_SetRandSeed(t *testing.T) {
	t.Parallel()

	randSeed := []byte("random seed")
	h := &block.HeaderV2{
		Header: &block.Header{},
	}

	err := h.SetRandSeed(randSeed)
	require.Nil(t, err)
	require.Equal(t, randSeed, h.GetRandSeed())
}

func TestHeaderV2_SetPubKeysBitmapNilPointerReceiverOrInnerHeader(t *testing.T) {
	t.Parallel()

	var h *block.HeaderV2
	require.Equal(t, data.ErrNilPointerReceiver, h.SetPubKeysBitmap([]byte("pub key bitmap")))

	h = &block.HeaderV2{
		Header: nil,
	}
	require.Equal(t, data.ErrNilPointerReceiver, h.SetPubKeysBitmap([]byte("pub key bitmap")))
}

func TestHeaderV2_SetPubKeysBitmap(t *testing.T) {
	t.Parallel()

	pubKeysBitmap := []byte{12, 13, 14, 15}
	h := &block.HeaderV2{
		Header: &block.Header{},
	}

	err := h.SetPubKeysBitmap(pubKeysBitmap)
	require.Nil(t, err)
	require.Equal(t, pubKeysBitmap, h.GetPubKeysBitmap())
}

func TestHeaderV2_SetRootHashNilPointerReceiverOrInnerHeader(t *testing.T) {
	t.Parallel()

	var h *block.HeaderV2
	require.Equal(t, data.ErrNilPointerReceiver, h.SetRootHash([]byte("root hash")))

	h = &block.HeaderV2{
		Header: nil,
	}
	require.Equal(t, data.ErrNilPointerReceiver, h.SetRootHash([]byte("root hash")))
}

func TestHeaderV2_SetRootHash(t *testing.T) {
	t.Parallel()

	rootHash := []byte("root hash")
	h := &block.HeaderV2{
		Header: &block.Header{},
	}

	err := h.SetRootHash(rootHash)
	require.Nil(t, err)
	require.Equal(t, rootHash, h.GetRootHash())
}

func TestHeaderV2_SetRoundNilPointerReceiverOrInnerHeader(t *testing.T) {
	t.Parallel()

	var h *block.HeaderV2
	require.Equal(t, data.ErrNilPointerReceiver, h.SetRound(1))

	h = &block.HeaderV2{
		Header: nil,
	}
	require.Equal(t, data.ErrNilPointerReceiver, h.SetRound(1))
}

func TestHeaderV2_SetRound(t *testing.T) {
	t.Parallel()

	round := uint64(10)
	h := &block.HeaderV2{
		Header: &block.Header{},
	}

	err := h.SetRound(round)
	require.Nil(t, err)
	require.Equal(t, round, h.GetRound())
}

func TestHeaderV2_SetSignatureNilPointerReceiverOrInnerHeader(t *testing.T) {
	t.Parallel()

	var h *block.HeaderV2
	require.Equal(t, data.ErrNilPointerReceiver, h.SetSignature([]byte("signature")))

	h = &block.HeaderV2{
		Header: nil,
	}
	require.Equal(t, data.ErrNilPointerReceiver, h.SetSignature([]byte("signature")))
}

func TestHeaderV2_SetSignature(t *testing.T) {
	t.Parallel()

	signature := []byte("signature")
	h := &block.HeaderV2{
		Header: &block.Header{},
	}

	err := h.SetSignature(signature)
	require.Nil(t, err)
	require.Equal(t, signature, h.GetSignature())
}

func TestHeaderV2_SetTimeStampNilPointerReceiverOrInnerHeader(t *testing.T) {
	t.Parallel()

	var h *block.HeaderV2
	require.Equal(t, data.ErrNilPointerReceiver, h.SetTimeStamp(100000))

	h = &block.HeaderV2{
		Header: nil,
	}
	require.Equal(t, data.ErrNilPointerReceiver, h.SetTimeStamp(100000))
}

func TestHeaderV2_SetTimeStamp(t *testing.T) {
	t.Parallel()

	timeStamp := uint64(100000)
	h := &block.HeaderV2{
		Header: &block.Header{},
	}

	err := h.SetTimeStamp(timeStamp)
	require.Nil(t, err)
	require.Equal(t, timeStamp, h.GetTimeStamp())
}

func TestHeaderV2_SetTxCountNilPointerReceiverOrInnerHeader(t *testing.T) {
	t.Parallel()

	var h *block.HeaderV2
	require.Equal(t, data.ErrNilPointerReceiver, h.SetTxCount(10000))

	h = &block.HeaderV2{
		Header: nil,
	}
	require.Equal(t, data.ErrNilPointerReceiver, h.SetTxCount(10000))
}

func TestHeaderV2_SetTxCount(t *testing.T) {
	t.Parallel()

	txCount := uint32(10)
	h := &block.HeaderV2{
		Header: &block.Header{},
	}

	err := h.SetTxCount(txCount)
	require.Nil(t, err)
	require.Equal(t, txCount, h.GetTxCount())
}

func TestHeaderV2_GetMiniBlockHeadersWithDstShouldWork(t *testing.T) {
	t.Parallel()

	hashS0R0 := []byte("hash_0_0")
	hashS0R1 := []byte("hash_0_1")
	hash1S0R2 := []byte("hash_0_2")
	hash2S0R2 := []byte("hash2_0_2")

	h := &block.HeaderV2{
		Header: &block.Header{
			MiniBlockHeaders: []block.MiniBlockHeader{
				{
					SenderShardID:   0,
					ReceiverShardID: 0,
					Hash:            hashS0R0,
				},
				{
					SenderShardID:   0,
					ReceiverShardID: 1,
					Hash:            hashS0R1,
				},
				{
					SenderShardID:   0,
					ReceiverShardID: 2,
					Hash:            hash1S0R2,
				},
				{
					SenderShardID:   0,
					ReceiverShardID: 2,
					Hash:            hash2S0R2,
				},
			},
		},
	}

	hashesWithDest2 := h.GetMiniBlockHeadersWithDst(2)

	require.Equal(t, uint32(0), hashesWithDest2[string(hash1S0R2)])
	require.Equal(t, uint32(0), hashesWithDest2[string(hash2S0R2)])
}

func TestHeaderV2_GetOrderedCrossMiniblocksWithDstShouldWork(t *testing.T) {
	t.Parallel()

	hashSh0ToSh0 := []byte("hash_0_0")
	hashSh0ToSh1 := []byte("hash_0_1")
	hash1Sh0ToSh2 := []byte("hash1_0_2")
	hash2Sh0ToSh2 := []byte("hash2_0_2")

	hdr := &block.Header{
		Round: 10,
		MiniBlockHeaders: []block.MiniBlockHeader{
			{
				SenderShardID:   0,
				ReceiverShardID: 0,
				Hash:            hashSh0ToSh0,
			},
			{
				SenderShardID:   0,
				ReceiverShardID: 1,
				Hash:            hashSh0ToSh1,
			},
			{
				SenderShardID:   0,
				ReceiverShardID: 2,
				Hash:            hash1Sh0ToSh2,
			},
			{
				SenderShardID:   0,
				ReceiverShardID: 2,
				Hash:            hash2Sh0ToSh2,
			},
		},
	}

	h := &block.HeaderV2{
		Header: hdr,
	}

	miniBlocksInfo := h.GetOrderedCrossMiniblocksWithDst(2)

	require.Equal(t, 2, len(miniBlocksInfo))
	require.Equal(t, hash1Sh0ToSh2, miniBlocksInfo[0].Hash)
	require.Equal(t, hdr.Round, miniBlocksInfo[0].Round)
	require.Equal(t, hash2Sh0ToSh2, miniBlocksInfo[1].Hash)
	require.Equal(t, hdr.Round, miniBlocksInfo[1].Round)
}

func TestHeaderV2_SetScheduledRootHash(t *testing.T) {
	t.Parallel()

	hv2 := block.HeaderV2{
		Header: &block.Header{},
	}
	require.Nil(t, hv2.ScheduledRootHash)

	rootHash := []byte("root hash")
	err := hv2.SetScheduledRootHash(rootHash)
	require.Nil(t, err)
	require.Equal(t, rootHash, hv2.ScheduledRootHash)
}

func TestHeaderV2_ValidateHeaderVersion(t *testing.T) {
	t.Parallel()

	hv2 := block.HeaderV2{
		Header: &block.Header{},
	}

	err := hv2.ValidateHeaderVersion()
	require.Equal(t, data.ErrNilScheduledRootHash, err)

	hv2.ScheduledRootHash = make([]byte, 0)
	err = hv2.ValidateHeaderVersion()
	require.Equal(t, data.ErrNilScheduledRootHash, err)

	hv2.ScheduledRootHash = make([]byte, 32)
	err = hv2.ValidateHeaderVersion()
	require.Nil(t, err)
}

func TestHeaderV2_GetMiniBlockHeadersHashes(t *testing.T) {
	t.Parallel()

	hash1 := []byte("hash1")
	hash2 := []byte("hash2")

	hv2 := block.HeaderV2{
		Header: &block.Header{
			MiniBlockHeaders: []block.MiniBlockHeader{
				{Hash: hash1},
				{Hash: hash2},
			},
		},
	}

	mbhh := hv2.GetMiniBlockHeadersHashes()
	require.NotNil(t, mbhh)
	require.Equal(t, 2, len(mbhh))
	require.Equal(t, hash1, mbhh[0])
	require.Equal(t, hash2, mbhh[1])
}

func TestHeaderV2_MapMiniBlockHashesToShards(t *testing.T) {
	t.Parallel()

	hash1 := []byte("hash1")
	hash2 := []byte("hash2")

	hash1Shard := core.MetachainShardId
	hash2Shard := uint32(1)

	hv2 := block.HeaderV2{
		Header: &block.Header{
			MiniBlockHeaders: []block.MiniBlockHeader{
				{Hash: hash1, SenderShardID: hash1Shard},
				{Hash: hash2, SenderShardID: hash2Shard},
			},
		},
	}

	mbhh := hv2.MapMiniBlockHashesToShards()
	require.NotNil(t, mbhh)
	require.Equal(t, 2, len(mbhh))
	require.Equal(t, hash1Shard, mbhh[string(hash1)])
	require.Equal(t, hash2Shard, mbhh[string(hash2)])
}

func TestHeaderV2_IsStartOfEpochBlock(t *testing.T) {
	t.Parallel()

	hv2 := block.HeaderV2{
		Header: &block.Header{
			EpochStartMetaHash: nil,
		},
	}

	isStartOfEpoch := hv2.IsStartOfEpochBlock()
	require.False(t, isStartOfEpoch)

	hv2 = block.HeaderV2{
		Header: &block.Header{
			EpochStartMetaHash: []byte("epochStartMetaHash"),
		},
	}

	isStartOfEpoch = hv2.IsStartOfEpochBlock()
	require.True(t, isStartOfEpoch)
}

func TestHeaderV2_GetBlockBodyTypeInt32(t *testing.T) {
	t.Parallel()

	bodyType := block.SmartContractResultBlock
	hv2 := block.HeaderV2{
		Header: &block.Header{
			BlockBodyType: bodyType,
		},
	}

	bodyTypeInt32 := hv2.GetBlockBodyTypeInt32()
	require.Equal(t, int32(bodyType), bodyTypeInt32)
}

func TestHeaderV2_GetMiniBlockHeaderHandlers(t *testing.T) {
	t.Parallel()

	hash1 := []byte("hash1")
	hash2 := []byte("hash2")

	hash1Shard := core.MetachainShardId
	hash2Shard := uint32(1)

	mbh := []block.MiniBlockHeader{
		{Hash: hash1, SenderShardID: hash1Shard},
		{Hash: hash2, SenderShardID: hash2Shard},
	}

	hv2 := block.HeaderV2{
		Header: &block.Header{
			MiniBlockHeaders: mbh,
		},
	}

	mbhh := hv2.GetMiniBlockHeaderHandlers()
	require.NotNil(t, mbhh)
	require.Equal(t, 2, len(mbhh))
	require.True(t, reflect.DeepEqual(&mbh[0], mbhh[0]))
	require.True(t, reflect.DeepEqual(&mbh[1], mbhh[1]))
}

func TestHeaderV2_SetMiniBlockHeaderHandlers(t *testing.T) {
	t.Parallel()

	hash1 := []byte("hash1")
	hash2 := []byte("hash2")

	hash1Shard := core.MetachainShardId
	hash2Shard := uint32(1)

	mbh := []data.MiniBlockHeaderHandler{
		&block.MiniBlockHeader{Hash: hash1, SenderShardID: hash1Shard},
		&block.MiniBlockHeader{Hash: hash2, SenderShardID: hash2Shard},
	}

	hv2 := block.HeaderV2{
		Header: &block.Header{},
	}

	err := hv2.SetMiniBlockHeaderHandlers(mbh)
	require.Nil(t, err)

	mbhh := hv2.GetMiniBlockHeaderHandlers()
	require.NotNil(t, mbhh)
	require.Equal(t, 2, len(mbhh))
	require.True(t, reflect.DeepEqual(mbh, mbhh))
}

func TestHeaderV2_SetReceiptsHash(t *testing.T) {
	t.Parallel()

	receiptHash := []byte("receiptHash")
	hv2 := block.HeaderV2{
		Header: &block.Header{},
	}
	err := hv2.SetReceiptsHash(receiptHash)
	require.Nil(t, err)
	require.Equal(t, receiptHash, hv2.GetReceiptsHash())
}

func TestHeaderV2_SetMetaBlockHashes(t *testing.T) {
	t.Parallel()

	mbHash1 := []byte("mbHash1")
	mbHash2 := []byte("mbHash2")
	metaBlockHashes := [][]byte{mbHash1, mbHash2}

	hv2 := block.HeaderV2{
		Header: &block.Header{},
	}
	err := hv2.SetMetaBlockHashes(metaBlockHashes)
	require.Nil(t, err)
	require.True(t, reflect.DeepEqual(metaBlockHashes, hv2.GetMetaBlockHashes()))
}

func TestHeaderV2_SetEpochStartMetaHash(t *testing.T) {
	t.Parallel()

	epochStartMetaHash := []byte("epochStartMetaHash")
	hv2 := block.HeaderV2{
		Header: &block.Header{},
	}
	err := hv2.SetEpochStartMetaHash(epochStartMetaHash)
	require.Nil(t, err)
	require.Equal(t, epochStartMetaHash, hv2.GetEpochStartMetaHash())
}

func TestHeaderV2_HasScheduledSupport(t *testing.T) {
	t.Parallel()

	hv2 := block.HeaderV2{}
	require.True(t, hv2.HasScheduledSupport())
}

func TestHeaderV2_SetScheduledGasProvided(t *testing.T) {
	t.Parallel()

	scheduledGasProvided := uint64(10)
	hv2 := block.HeaderV2{}
	err := hv2.SetScheduledGasProvided(scheduledGasProvided)
	require.Nil(t, err)
	require.Equal(t, scheduledGasProvided, hv2.GetScheduledGasProvided())
}

func TestHeaderV2_SetScheduledGasPenalized(t *testing.T) {
	t.Parallel()

	scheduledGasPenalized := uint64(10)
	hv2 := block.HeaderV2{}
	err := hv2.SetScheduledGasPenalized(scheduledGasPenalized)
	require.Nil(t, err)
	require.Equal(t, scheduledGasPenalized, hv2.GetScheduledGasPenalized())
}

func TestHeaderV2_SetScheduledGasRefunded(t *testing.T) {
	t.Parallel()

	scheduledGasRefunded := uint64(10)
	hv2 := block.HeaderV2{}
	err := hv2.SetScheduledGasRefunded(scheduledGasRefunded)
	require.Nil(t, err)
	require.Equal(t, scheduledGasRefunded, hv2.GetScheduledGasRefunded())
}

func TestHeaderV2_GetAdditionalData(t *testing.T) {
	t.Parallel()

	scheduledRootHash := []byte("scheduledRootHash")
	scheduledAccumulatedFees := big.NewInt(10)
	scheduledDeveloperFees := big.NewInt(11)
	scheduledGasProvided := uint64(1)
	scheduledGasPenalized := uint64(2)
	scheduledGasRefunded := uint64(3)

	hv2 := block.HeaderV2{
		ScheduledRootHash:        scheduledRootHash,
		ScheduledAccumulatedFees: scheduledAccumulatedFees,
		ScheduledDeveloperFees:   scheduledDeveloperFees,
		ScheduledGasProvided:     scheduledGasProvided,
		ScheduledGasPenalized:    scheduledGasPenalized,
		ScheduledGasRefunded:     scheduledGasRefunded,
	}
	additionalData := hv2.GetAdditionalData()
	require.NotNil(t, additionalData)
	require.Equal(t, scheduledRootHash, additionalData.GetScheduledRootHash())
	require.Equal(t, scheduledAccumulatedFees, additionalData.GetScheduledAccumulatedFees())
	require.Equal(t, scheduledDeveloperFees, additionalData.GetScheduledDeveloperFees())
	require.Equal(t, scheduledGasProvided, additionalData.GetScheduledGasProvided())
	require.Equal(t, scheduledGasPenalized, additionalData.GetScheduledGasPenalized())
	require.Equal(t, scheduledGasRefunded, additionalData.GetScheduledGasRefunded())
}

func TestHeaderV2_SetAdditionalDataNilAdditionalDataShouldErr(t *testing.T) {
	t.Parallel()

	shardBlock := &block.HeaderV2{
		Header:            &block.Header{},
		ScheduledRootHash: nil,
	}

	err := shardBlock.SetAdditionalData(nil)

	require.NotNil(t, err)
	require.Equal(t, data.ErrNilPointerDereference, err)
}

func TestHeaderV2_SetAdditionalDataEmptyFeesShouldWork(t *testing.T) {
	t.Parallel()

	shardBlock := &block.HeaderV2{
		Header:            &block.Header{},
		ScheduledRootHash: nil,
	}

	scRootHash := []byte("scheduledRootHash")
	err := shardBlock.SetAdditionalData(&headerVersionData.AdditionalData{
		ScheduledRootHash: scRootHash,
	})

	require.Nil(t, err)
	require.Equal(t, scRootHash, shardBlock.ScheduledRootHash)
	require.Equal(t, big.NewInt(0), shardBlock.ScheduledAccumulatedFees)
	require.Equal(t, big.NewInt(0), shardBlock.ScheduledDeveloperFees)
	require.Equal(t, uint64(0), shardBlock.GetScheduledGasPenalized())
	require.Equal(t, uint64(0), shardBlock.GetScheduledGasRefunded())
	require.Equal(t, uint64(0), shardBlock.GetScheduledGasProvided())
}

func TestHeaderV2_SetAdditionalDataShouldWork(t *testing.T) {
	t.Parallel()

	shardBlock := &block.HeaderV2{
		Header:            &block.Header{},
		ScheduledRootHash: nil,
	}

	scRootHash := []byte("scheduledRootHash")
	accFees := big.NewInt(100)
	devFees := big.NewInt(10)
	gasProvided := uint64(60)
	gasRefunded := uint64(10)
	gasPenalized := uint64(20)
	err := shardBlock.SetAdditionalData(&headerVersionData.AdditionalData{
		ScheduledRootHash:        scRootHash,
		ScheduledAccumulatedFees: accFees,
		ScheduledDeveloperFees:   devFees,
		ScheduledGasProvided:     gasProvided,
		ScheduledGasPenalized:    gasPenalized,
		ScheduledGasRefunded:     gasRefunded,
	})

	require.Nil(t, err)
	require.Equal(t, scRootHash, shardBlock.ScheduledRootHash)
	require.Equal(t, accFees, shardBlock.ScheduledAccumulatedFees)
	require.Equal(t, devFees, shardBlock.ScheduledDeveloperFees)
	require.Equal(t, gasPenalized, shardBlock.GetScheduledGasPenalized())
	require.Equal(t, gasRefunded, shardBlock.GetScheduledGasRefunded())
	require.Equal(t, gasProvided, shardBlock.GetScheduledGasProvided())
}

func TestHeaderV2_HasScheduledMiniBlocks(t *testing.T) {
	t.Parallel()

	mbh := &block.MiniBlockHeader{}
	_ = mbh.SetProcessingType(int32(block.Scheduled))

	shardBlock := &block.HeaderV2{
		Header: &block.Header{
			MiniBlockHeaders: []block.MiniBlockHeader{*mbh},
		},
	}
	require.True(t, shardBlock.HasScheduledMiniBlocks())

	_ = mbh.SetProcessingType(int32(block.Normal))
	shardBlock = &block.HeaderV2{
		Header: &block.Header{
			MiniBlockHeaders: []block.MiniBlockHeader{*mbh},
		},
	}

	require.False(t, shardBlock.HasScheduledMiniBlocks())
}
