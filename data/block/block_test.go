package block_test

import (
	"math/big"
	"reflect"
	"testing"

	"github.com/ElrondNetwork/elrond-go-core/data"
	"github.com/ElrondNetwork/elrond-go-core/data/block"
	"github.com/ElrondNetwork/elrond-go-core/data/headerVersionData"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHeader_GetEpoch(t *testing.T) {
	t.Parallel()

	epoch := uint32(1)
	h := block.Header{
		Epoch: epoch,
	}

	assert.Equal(t, epoch, h.GetEpoch())
}

func TestHeader_GetShard(t *testing.T) {
	t.Parallel()

	shardId := uint32(2)
	h := block.Header{
		ShardID: shardId,
	}

	assert.Equal(t, shardId, h.GetShardID())
}

func TestHeader_GetNonce(t *testing.T) {
	t.Parallel()

	nonce := uint64(2)
	h := block.Header{
		Nonce: nonce,
	}

	assert.Equal(t, nonce, h.GetNonce())
}

func TestHeader_GetPrevHash(t *testing.T) {
	t.Parallel()

	prevHash := []byte("prev hash")
	h := block.Header{
		PrevHash: prevHash,
	}

	assert.Equal(t, prevHash, h.GetPrevHash())
}

func TestHeader_GetPrevRandSeed(t *testing.T) {
	t.Parallel()

	prevRandSeed := []byte("prev random seed")
	h := block.Header{
		PrevRandSeed: prevRandSeed,
	}

	assert.Equal(t, prevRandSeed, h.GetPrevRandSeed())
}

func TestHeader_GetRandSeed(t *testing.T) {
	t.Parallel()

	randSeed := []byte("random seed")
	h := block.Header{
		RandSeed: randSeed,
	}

	assert.Equal(t, randSeed, h.GetRandSeed())
}

func TestHeader_GetPubKeysBitmap(t *testing.T) {
	t.Parallel()

	pubKeysBitmap := []byte{10, 11, 12, 13}
	h := block.Header{
		PubKeysBitmap: pubKeysBitmap,
	}

	assert.Equal(t, pubKeysBitmap, h.GetPubKeysBitmap())
}

func TestHeader_GetRootHash(t *testing.T) {
	t.Parallel()

	rootHash := []byte("root hash")
	h := block.Header{
		RootHash: rootHash,
	}

	assert.Equal(t, rootHash, h.GetRootHash())
}

func TestHeader_GetRound(t *testing.T) {
	t.Parallel()

	round := uint64(1234)
	h := block.Header{
		Round: round,
	}

	assert.Equal(t, round, h.GetRound())
}

func TestHeader_GetSignature(t *testing.T) {
	t.Parallel()

	signature := []byte("signature")
	h := block.Header{
		Signature: signature,
	}

	assert.Equal(t, signature, h.GetSignature())
}

func TestHeader_GetTxCount(t *testing.T) {
	t.Parallel()

	txCount := uint32(10)
	h := block.Header{
		TxCount: txCount,
	}

	assert.Equal(t, txCount, h.GetTxCount())
}

func TestHeader_SetEpoch(t *testing.T) {
	t.Parallel()

	epoch := uint32(10)
	h := block.Header{}
	err := h.SetEpoch(epoch)

	assert.Nil(t, err)
	assert.Equal(t, epoch, h.GetEpoch())
}

func TestHeader_SetNonce(t *testing.T) {
	t.Parallel()

	nonce := uint64(11)
	h := block.Header{}
	err := h.SetNonce(nonce)

	assert.Nil(t, err)
	assert.Equal(t, nonce, h.GetNonce())
}

func TestHeader_SetPrevHash(t *testing.T) {
	t.Parallel()

	prevHash := []byte("prev hash")
	h := block.Header{}
	err := h.SetPrevHash(prevHash)

	assert.Nil(t, err)
	assert.Equal(t, prevHash, h.GetPrevHash())
}

func TestHeader_SetPrevRandSeed(t *testing.T) {
	t.Parallel()

	prevRandSeed := []byte("prev random seed")
	h := block.Header{}
	err := h.SetPrevRandSeed(prevRandSeed)

	assert.Nil(t, err)
	assert.Equal(t, prevRandSeed, h.GetPrevRandSeed())
}

func TestHeader_SetRandSeed(t *testing.T) {
	t.Parallel()

	randSeed := []byte("random seed")
	h := block.Header{}
	err := h.SetRandSeed(randSeed)

	assert.Nil(t, err)
	assert.Equal(t, randSeed, h.GetRandSeed())
}

func TestHeader_SetPubKeysBitmap(t *testing.T) {
	t.Parallel()

	pubKeysBitmap := []byte{12, 13, 14, 15}
	h := block.Header{}
	err := h.SetPubKeysBitmap(pubKeysBitmap)

	assert.Nil(t, err)
	assert.Equal(t, pubKeysBitmap, h.GetPubKeysBitmap())
}

func TestHeader_SetRootHash(t *testing.T) {
	t.Parallel()

	rootHash := []byte("root hash")
	h := block.Header{}
	err := h.SetRootHash(rootHash)

	assert.Nil(t, err)
	assert.Equal(t, rootHash, h.GetRootHash())
}

func TestHeader_SetRound(t *testing.T) {
	t.Parallel()

	round := uint64(10)
	h := block.Header{}
	err := h.SetRound(round)

	assert.Nil(t, err)
	assert.Equal(t, round, h.GetRound())
}

func TestHeader_SetSignature(t *testing.T) {
	t.Parallel()

	signature := []byte("signature")
	h := block.Header{}
	err := h.SetSignature(signature)

	assert.Nil(t, err)
	assert.Equal(t, signature, h.GetSignature())
}

func TestHeader_SetLeaderSignature(t *testing.T) {
	t.Parallel()

	leaderSig := []byte("leaderSig")
	h := block.Header{}
	err := h.SetLeaderSignature(leaderSig)

	assert.Nil(t, err)
	assert.Equal(t, leaderSig, h.GetLeaderSignature())
}

func TestHeader_SetChainID(t *testing.T) {
	t.Parallel()

	chainId := []byte("chainId")
	h := block.Header{}
	err := h.SetChainID(chainId)

	assert.Nil(t, err)
	assert.Equal(t, chainId, h.GetChainID())
}

func TestHeader_SetSoftwareVersion(t *testing.T) {
	t.Parallel()

	version := []byte("version")
	h := block.Header{}
	err := h.SetSoftwareVersion(version)

	assert.Nil(t, err)
	assert.Equal(t, version, h.GetSoftwareVersion())
}

func TestHeader_SetTimeStamp(t *testing.T) {
	t.Parallel()

	timeStamp := uint64(100000)
	h := block.Header{}
	err := h.SetTimeStamp(timeStamp)

	assert.Nil(t, err)
	assert.Equal(t, timeStamp, h.GetTimeStamp())
}

func TestHeader_SetAccumulatedFees(t *testing.T) {
	t.Parallel()

	accumulatedFees := big.NewInt(10)
	h := block.Header{}
	err := h.SetAccumulatedFees(accumulatedFees)

	assert.Nil(t, err)
	assert.Equal(t, accumulatedFees, h.GetAccumulatedFees())
}

func TestHeader_SetDeveloperFees(t *testing.T) {
	t.Parallel()

	developerFees := big.NewInt(10)
	h := block.Header{}
	err := h.SetDeveloperFees(developerFees)

	assert.Nil(t, err)
	assert.Equal(t, developerFees, h.GetDeveloperFees())
}

func TestHeader_SetTxCount(t *testing.T) {
	t.Parallel()

	txCount := uint32(10)
	h := block.Header{}
	err := h.SetTxCount(txCount)

	assert.Nil(t, err)
	assert.Equal(t, txCount, h.GetTxCount())
}

func TestHeader_SetShardID(t *testing.T) {
	t.Parallel()

	shardId := uint32(2)
	h := block.Header{}
	err := h.SetShardID(shardId)

	assert.Nil(t, err)
	assert.Equal(t, shardId, shardId, h.GetShardID())
}

func TestBody_IntegrityAndValidityNil(t *testing.T) {
	t.Parallel()

	var body *block.Body = nil
	assert.Equal(t, data.ErrNilPointerReceiver, body.IntegrityAndValidity())
}

func TestBody_IntegrityAndValidityEmptyMiniblockShouldThrowException(t *testing.T) {
	t.Parallel()

	txHash0 := []byte("txHash0")
	mb0 := block.MiniBlock{
		ReceiverShardID: 0,
		SenderShardID:   0,
		TxHashes:        [][]byte{txHash0},
	}

	mb1 := block.MiniBlock{}

	body := &block.Body{}
	body.MiniBlocks = append(body.MiniBlocks, &mb0)
	body.MiniBlocks = append(body.MiniBlocks, &mb1)

	assert.Equal(t, data.ErrMiniBlockEmpty, body.IntegrityAndValidity())
}

func TestBody_IntegrityAndValidityOK(t *testing.T) {
	t.Parallel()

	txHash0 := []byte("txHash0")
	mb0 := block.MiniBlock{
		ReceiverShardID: 0,
		SenderShardID:   0,
		TxHashes:        [][]byte{txHash0},
	}

	body := &block.Body{}
	body.MiniBlocks = append(body.MiniBlocks, &mb0)

	assert.Equal(t, nil, body.IntegrityAndValidity())
}

func TestBody_Clone(t *testing.T) {
	t.Parallel()

	txHash0 := []byte("txHash0")
	mb0 := block.MiniBlock{
		ReceiverShardID: 0,
		SenderShardID:   0,
		TxHashes:        [][]byte{txHash0},
	}

	body := &block.Body{}
	body.MiniBlocks = append(body.MiniBlocks, &mb0)

	clonedBody := body.Clone()

	assert.True(t, reflect.DeepEqual(body, clonedBody))
}

func TestHeader_SetMiniBlockHeaderHandlersNilMiniBlockHeaderHandlers(t *testing.T) {
	t.Parallel()

	hdr := block.Header{}
	err := hdr.SetMiniBlockHeaderHandlers(make([]data.MiniBlockHeaderHandler, 0))
	assert.Nil(t, err)
	assert.Nil(t, hdr.MiniBlockHeaders)
}

func TestHeader_SetMiniBlockHeaderHandlersWithError(t *testing.T) {
	t.Parallel()

	t.Run("invalid type assertion", func(t *testing.T) {
		t.Parallel()

		mb0 := &block.MiniBlockHeader{}

		mbHeaderHandlers := []data.MiniBlockHeaderHandler{
			mb0,
			nil,
		}

		hdr := block.Header{}
		err := hdr.SetMiniBlockHeaderHandlers(mbHeaderHandlers)
		assert.Equal(t, data.ErrInvalidTypeAssertion, err)
		assert.Equal(t, 0, len(hdr.MiniBlockHeaders))
	})
	t.Run("nil pointer dereference", func(t *testing.T) {
		t.Parallel()

		mb0 := &block.MiniBlockHeader{}
		var mb1 data.MiniBlockHeaderHandler = (*block.MiniBlockHeader)(nil)

		mbHeaderHandlers := []data.MiniBlockHeaderHandler{
			mb0,
			mb1,
		}

		hdr := block.Header{}
		err := hdr.SetMiniBlockHeaderHandlers(mbHeaderHandlers)
		assert.Equal(t, data.ErrNilPointerDereference, err)
		assert.Equal(t, 0, len(hdr.MiniBlockHeaders))
	})
	t.Run("set error overwrites the previous value", func(t *testing.T) {
		t.Parallel()

		hdr := block.Header{}

		mb0 := &block.MiniBlockHeader{}
		mbHeaderHandlers := []data.MiniBlockHeaderHandler{
			mb0,
		}

		err := hdr.SetMiniBlockHeaderHandlers(mbHeaderHandlers)
		assert.Nil(t, err)
		expectedLen := len(mbHeaderHandlers)
		assert.Equal(t, expectedLen, len(hdr.GetMiniBlockHeaderHandlers()))

		mbHeaderHandlers = []data.MiniBlockHeaderHandler{
			nil,
			nil,
		}

		err = hdr.SetMiniBlockHeaderHandlers(mbHeaderHandlers)
		assert.NotNil(t, err)
		assert.Equal(t, expectedLen, len(hdr.GetMiniBlockHeaderHandlers()))
	})
}

func TestHeader_SetMiniBlockHeaderHandlers(t *testing.T) {
	t.Parallel()

	hashS0R0 := []byte("hash_0_0")
	hashS0R1 := []byte("hash_0_1")

	mb0 := &block.MiniBlockHeader{
		SenderShardID:   0,
		ReceiverShardID: 0,
		Hash:            hashS0R0,
	}
	mb1 := &block.MiniBlockHeader{
		SenderShardID:   0,
		ReceiverShardID: 1,
		Hash:            hashS0R1,
	}
	mbHeaderHandlers := []data.MiniBlockHeaderHandler{
		mb0,
		mb1,
	}

	hdr := block.Header{}
	err := hdr.SetMiniBlockHeaderHandlers(mbHeaderHandlers)

	assert.Nil(t, err)

	mbhh := hdr.GetMiniBlockHeaderHandlers()
	assert.True(t, reflect.DeepEqual(mb0, mbhh[0]))
	assert.True(t, reflect.DeepEqual(mb1, mbhh[1]))
}

func TestHeader_SetReceiptsHash(t *testing.T) {
	t.Parallel()

	receiptHash := []byte("hash")
	hdr := &block.Header{}
	err := hdr.SetReceiptsHash(receiptHash)

	assert.Nil(t, err)
	assert.Equal(t, receiptHash, hdr.GetReceiptsHash())
}

func TestHeader_SetMetaBlockHashes(t *testing.T) {
	t.Parallel()

	hash1 := []byte("hash1")
	hash2 := []byte("hash2")
	hdr := &block.Header{}
	err := hdr.SetMetaBlockHashes([][]byte{hash1, hash2})

	assert.Nil(t, err)

	mbHashes := hdr.GetMetaBlockHashes()
	assert.Equal(t, 2, len(mbHashes))
	assert.Equal(t, hash1, mbHashes[0])
	assert.Equal(t, hash2, mbHashes[1])
}

func TestHeader_SetEpochStartMetaHash(t *testing.T) {
	t.Parallel()

	epochStartMetaHash := []byte("hash")
	hdr := &block.Header{}
	err := hdr.SetEpochStartMetaHash(epochStartMetaHash)

	assert.Nil(t, err)
	assert.Equal(t, epochStartMetaHash, hdr.GetEpochStartMetaHash())
}

func TestHeader_GetMiniBlockHeadersWithDstShouldWork(t *testing.T) {
	t.Parallel()

	hashS0R0 := []byte("hash_0_0")
	hashS0R1 := []byte("hash_0_1")
	hash1S0R2 := []byte("hash_0_2")
	hash2S0R2 := []byte("hash2_0_2")

	hdr := &block.Header{
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
	}

	hashesWithDest2 := hdr.GetMiniBlockHeadersWithDst(2)

	assert.Equal(t, uint32(0), hashesWithDest2[string(hash1S0R2)])
	assert.Equal(t, uint32(0), hashesWithDest2[string(hash2S0R2)])
}

func TestHeader_GetOrderedCrossMiniblocksWithDstShouldWork(t *testing.T) {
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

	miniBlocksInfo := hdr.GetOrderedCrossMiniblocksWithDst(2)

	require.Equal(t, 2, len(miniBlocksInfo))
	assert.Equal(t, hash1Sh0ToSh2, miniBlocksInfo[0].Hash)
	assert.Equal(t, hdr.Round, miniBlocksInfo[0].Round)
	assert.Equal(t, hash2Sh0ToSh2, miniBlocksInfo[1].Hash)
	assert.Equal(t, hdr.Round, miniBlocksInfo[1].Round)
}

func TestMiniBlock_Clone(t *testing.T) {
	t.Parallel()

	miniBlock := &block.MiniBlock{
		TxHashes:        [][]byte{[]byte("something"), []byte("something2")},
		ReceiverShardID: 1,
		SenderShardID:   2,
		Type:            0,
		Reserved:        []byte("something"),
	}

	clonedMB := miniBlock.Clone()

	assert.True(t, reflect.DeepEqual(miniBlock, clonedMB))
}

func TestMiniBlock_SetMiniBlockReservedToNil(t *testing.T) {
	t.Parallel()

	miniBlock := &block.MiniBlock{
		Reserved: []byte("something"),
	}

	err := miniBlock.SetMiniBlockReserved(nil)
	assert.Nil(t, err)

	mbReserved, err := miniBlock.GetMiniBlockReserved()
	assert.Nil(t, err)
	assert.Nil(t, mbReserved)
}

func TestMiniBlock_SetMiniBlockReservedShouldWork(t *testing.T) {
	t.Parallel()

	miniBlock := &block.MiniBlock{
		Reserved: nil,
	}

	executionType := block.ProcessingType(7)
	transactionType := []byte("txType")
	mbr := &block.MiniBlockReserved{
		ExecutionType:    executionType,
		TransactionsType: transactionType,
	}
	err := miniBlock.SetMiniBlockReserved(mbr)
	assert.Nil(t, err)

	mbReserved, err := miniBlock.GetMiniBlockReserved()
	assert.Nil(t, err)
	assert.NotNil(t, mbReserved)
	assert.Equal(t, executionType, mbReserved.ExecutionType)
	assert.Equal(t, transactionType, mbReserved.TransactionsType)
}

func TestMiniBlock_IsScheduledMiniBlock(t *testing.T) {
	t.Parallel()

	miniBlock := &block.MiniBlock{
		Reserved: nil,
	}

	isScheduledMB := miniBlock.IsScheduledMiniBlock()
	assert.False(t, isScheduledMB)

	miniBlock.Reserved = []byte("not marshal")
	isScheduledMB = miniBlock.IsScheduledMiniBlock()
	assert.False(t, isScheduledMB)

	_ = miniBlock.SetMiniBlockReserved(
		&block.MiniBlockReserved{
			ExecutionType: block.Scheduled,
		})
	isScheduledMB = miniBlock.IsScheduledMiniBlock()
	assert.True(t, isScheduledMB)
}

func TestMiniBlock_GetProcessingType(t *testing.T) {
	t.Parallel()

	miniBlock := &block.MiniBlock{}

	_ = miniBlock.SetMiniBlockReserved(&block.MiniBlockReserved{ExecutionType: block.Scheduled})
	assert.Equal(t, int32(block.Scheduled), miniBlock.GetProcessingType())

	_ = miniBlock.SetMiniBlockReserved(&block.MiniBlockReserved{ExecutionType: block.Processed})
	assert.Equal(t, int32(block.Processed), miniBlock.GetProcessingType())
}

func TestMiniBlock_GetTxsTypeFromMiniBlock(t *testing.T) {
	t.Parallel()

	miniBlock := &block.MiniBlock{
		Reserved: nil,
	}
	txTypes, err := miniBlock.GetTxsTypeFromMiniBlock()

	assert.Nil(t, err)
	assert.NotNil(t, txTypes)
	assert.Equal(t, 0, len(txTypes))

	miniBlock = &block.MiniBlock{
		Reserved: []byte("not marshall"),
	}
	txTypes, err = miniBlock.GetTxsTypeFromMiniBlock()

	assert.NotNil(t, err)
	assert.Nil(t, txTypes)

	numTxs := 2
	// first txBlockType, second smartcontractResultType
	txType := byte(2)
	miniBlock = &block.MiniBlock{
		TxHashes: make([][]byte, numTxs),
	}
	_ = miniBlock.SetMiniBlockReserved(&block.MiniBlockReserved{
		TransactionsType: []byte{txType},
	})
	txTypes, err = miniBlock.GetTxsTypeFromMiniBlock()

	assert.Nil(t, err)
	assert.NotNil(t, txTypes)
	assert.Equal(t, numTxs, len(txTypes))
	assert.Equal(t, block.TxBlock, txTypes[0])
	assert.Equal(t, block.SmartContractResultBlock, txTypes[1])
}

func TestHeader_SetScheduledRootHash(t *testing.T) {
	t.Parallel()

	header := &block.Header{}
	err := header.SetScheduledRootHash([]byte("root hash"))
	require.Equal(t, data.ErrScheduledRootHashNotSupported, err)
}

func TestHeader_ValidateHeaderVersion(t *testing.T) {
	t.Parallel()

	header := &block.Header{}
	err := header.ValidateHeaderVersion()
	require.Nil(t, err)
}

func TestHeader_SetAdditionalDataShouldDoNothing(t *testing.T) {
	t.Parallel()

	var shardBlock *block.Header

	//goland:noinspection ALL
	err := shardBlock.SetAdditionalData(&headerVersionData.AdditionalData{})
	assert.Nil(t, err)
}

func TestMiniBlockHeader_SetMiniBlockHeaderReserved(t *testing.T) {
	t.Parallel()

	mbh := &block.MiniBlockHeader{}
	err := mbh.SetMiniBlockHeaderReserved(nil)
	assert.Nil(t, err)
	assert.Nil(t, mbh.Reserved)

	mbh = &block.MiniBlockHeader{}
	mbhr := &block.MiniBlockHeaderReserved{
		ExecutionType: block.Scheduled,
		State:         block.Final,
	}
	err = mbh.SetMiniBlockHeaderReserved(mbhr)
	assert.Nil(t, err)

	newMbhr, err := mbh.GetMiniBlockHeaderReserved()
	assert.Nil(t, err)
	assert.True(t, reflect.DeepEqual(mbhr, newMbhr))
}

func TestMiniBlockHeader_IsFinal(t *testing.T) {
	t.Parallel()

	mbh := &block.MiniBlockHeader{}
	mbh.Reserved = []byte("non marshall")
	isFinal := mbh.IsFinal()
	assert.True(t, isFinal)

	mbh = &block.MiniBlockHeader{}
	mbh.Reserved = nil
	isFinal = mbh.IsFinal()
	assert.True(t, isFinal)

	mbhr := &block.MiniBlockHeaderReserved{
		ExecutionType: block.Scheduled,
		State:         block.Final,
	}
	_ = mbh.SetMiniBlockHeaderReserved(mbhr)
	isFinal = mbh.IsFinal()
	assert.True(t, isFinal)

	mbhr = &block.MiniBlockHeaderReserved{
		ExecutionType: block.Scheduled,
		State:         block.Proposed,
	}
	_ = mbh.SetMiniBlockHeaderReserved(mbhr)
	isFinal = mbh.IsFinal()
	assert.False(t, isFinal)
}

func TestHeader_HasScheduledSupport(t *testing.T) {
	t.Parallel()

	h := &block.Header{}
	hasScheduleSupport := h.HasScheduledSupport()
	assert.False(t, hasScheduleSupport)
}

func TestHeader_GetAdditionalData(t *testing.T) {
	t.Parallel()

	h := &block.Header{}
	additionalData := h.GetAdditionalData()
	assert.Nil(t, additionalData)
}

func TestHeader_HasScheduledMiniBlocks(t *testing.T) {
	t.Parallel()

	h := &block.Header{}
	require.False(t, h.HasScheduledMiniBlocks())

	mbHeader := &block.MiniBlockHeader{}
	_ = mbHeader.SetProcessingType(int32(block.Normal))
	h.MiniBlockHeaders = []block.MiniBlockHeader{*mbHeader}
	require.False(t, h.HasScheduledMiniBlocks())

	// scheduled miniBlocks not supported for v1 header, so it should return false
	_ = mbHeader.SetProcessingType(int32(block.Scheduled))
	h.MiniBlockHeaders = []block.MiniBlockHeader{*mbHeader}
	require.False(t, h.HasScheduledMiniBlocks())
}
