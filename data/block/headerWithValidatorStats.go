//go:generate protoc -I=. -I=$GOPATH/src -I=$GOPATH/src/github.com/ElrondNetwork/protobuf/protobuf  --gogoslick_out=. headerWithValidatorStats.proto
package block

import (
	"math/big"

	"github.com/multiversx/mx-chain-core-go/data"
	"github.com/multiversx/mx-chain-core-go/data/headerVersionData"
)

// GetAdditionalData returns nil
func (hv *HeaderWithValidatorStats) GetAdditionalData() headerVersionData.HeaderAdditionalData {
	return nil
}

// HasScheduledMiniBlocks returns false
func (hv *HeaderWithValidatorStats) HasScheduledMiniBlocks() bool {
	return false
}

// SetScheduledRootHash returns nil
func (hv *HeaderWithValidatorStats) SetScheduledRootHash(_ []byte) error {
	return nil
}

// SetAdditionalData will not do anything
func (hv *HeaderWithValidatorStats) SetAdditionalData(_ headerVersionData.HeaderAdditionalData) error {
	return nil
}

// ShallowClone  will return a clone of the object
func (hv *HeaderWithValidatorStats) ShallowClone() data.HeaderHandler {
	if hv == nil || hv.Header == nil {
		return nil
	}

	internalHeaderCopy := *hv.Header
	headerCopy := *hv
	headerCopy.Header = &internalHeaderCopy

	return &headerCopy
}

// GetShardID returns the header shardID
func (hv *HeaderWithValidatorStats) GetShardID() uint32 {
	return 0
}

// GetNonce returns the header nonce
func (hv *HeaderWithValidatorStats) GetNonce() uint64 {
	if hv == nil {
		return 0
	}

	return hv.Header.GetNonce()
}

// GetEpoch returns the header epoch
func (hv *HeaderWithValidatorStats) GetEpoch() uint32 {
	if hv == nil {
		return 0
	}

	return hv.Header.GetEpoch()
}

// GetRound returns the header round
func (hv *HeaderWithValidatorStats) GetRound() uint64 {
	if hv == nil {
		return 0
	}

	return hv.Header.GetRound()
}

// GetRootHash returns the header root hash
func (hv *HeaderWithValidatorStats) GetRootHash() []byte {
	if hv == nil {
		return nil
	}

	return hv.Header.GetRootHash()
}

// GetPrevHash returns the header previous header hash
func (hv *HeaderWithValidatorStats) GetPrevHash() []byte {
	if hv == nil {
		return nil
	}

	return hv.Header.GetPrevHash()
}

// GetPrevRandSeed returns the header previous random seed
func (hv *HeaderWithValidatorStats) GetPrevRandSeed() []byte {
	if hv == nil {
		return nil
	}

	return hv.Header.GetPrevRandSeed()
}

// GetRandSeed returns the header random seed
func (hv *HeaderWithValidatorStats) GetRandSeed() []byte {
	if hv == nil {
		return nil
	}

	return hv.Header.GetRandSeed()
}

// GetPubKeysBitmap returns the header public key bitmap for the aggregated signatures
func (hv *HeaderWithValidatorStats) GetPubKeysBitmap() []byte {
	if hv == nil {
		return nil
	}

	return hv.Header.GetPubKeysBitmap()
}

// GetSignature returns the header aggregated signature
func (hv *HeaderWithValidatorStats) GetSignature() []byte {
	if hv == nil {
		return nil
	}

	return hv.Header.GetSignature()
}

// GetLeaderSignature returns the leader signature on top of the finalized (signed) header
func (hv *HeaderWithValidatorStats) GetLeaderSignature() []byte {
	if hv == nil {
		return nil
	}

	return hv.Header.GetLeaderSignature()
}

// GetChainID returns the chain ID
func (hv *HeaderWithValidatorStats) GetChainID() []byte {
	if hv == nil {
		return nil
	}

	return hv.Header.GetChainID()
}

// GetSoftwareVersion returns the header software version
func (hv *HeaderWithValidatorStats) GetSoftwareVersion() []byte {
	if hv == nil {
		return nil
	}

	return hv.Header.GetSoftwareVersion()
}

// GetTimeStamp returns the header timestamp
func (hv *HeaderWithValidatorStats) GetTimeStamp() uint64 {
	if hv == nil {
		return 0
	}

	return hv.Header.GetTimeStamp()
}

// GetTxCount returns the number of txs included in the block
func (hv *HeaderWithValidatorStats) GetTxCount() uint32 {
	if hv == nil {
		return 0
	}

	return hv.Header.GetTxCount()
}

// GetReceiptsHash returns the header receipt hash
func (hv *HeaderWithValidatorStats) GetReceiptsHash() []byte {
	if hv == nil {
		return nil
	}

	return hv.Header.GetReceiptsHash()
}

// GetAccumulatedFees returns the block accumulated fees
func (hv *HeaderWithValidatorStats) GetAccumulatedFees() *big.Int {
	if hv == nil {
		return nil
	}

	return hv.Header.GetAccumulatedFees()
}

// GetDeveloperFees returns the block developer fees
func (hv *HeaderWithValidatorStats) GetDeveloperFees() *big.Int {
	if hv == nil {
		return nil
	}

	return hv.Header.GetDeveloperFees()
}

// GetReserved returns the reserved field
func (hv *HeaderWithValidatorStats) GetReserved() []byte {
	if hv == nil {
		return nil
	}

	return hv.Header.GetReserved()
}

// GetMetaBlockHashes returns the metaBlock hashes
func (hv *HeaderWithValidatorStats) GetMetaBlockHashes() [][]byte {
	if hv == nil {
		return nil
	}

	return hv.Header.GetMetaBlockHashes()
}

// GetEpochStartMetaHash returns the epoch start metaBlock hash
func (hv *HeaderWithValidatorStats) GetEpochStartMetaHash() []byte {
	if hv == nil {
		return nil
	}

	return hv.Header.GetEpochStartMetaHash()
}

// SetNonce sets header nonce
func (hv *HeaderWithValidatorStats) SetNonce(n uint64) error {
	if hv == nil {
		return data.ErrNilPointerReceiver
	}
	return hv.Header.SetNonce(n)
}

// SetEpoch sets header epoch
func (hv *HeaderWithValidatorStats) SetEpoch(e uint32) error {
	if hv == nil {
		return data.ErrNilPointerReceiver
	}

	return hv.Header.SetEpoch(e)
}

// SetRound sets header round
func (hv *HeaderWithValidatorStats) SetRound(r uint64) error {
	if hv == nil {
		return data.ErrNilPointerReceiver
	}

	return hv.Header.SetRound(r)
}

// SetRootHash sets root hash
func (hv *HeaderWithValidatorStats) SetRootHash(rHash []byte) error {
	if hv == nil {
		return data.ErrNilPointerReceiver
	}

	return hv.Header.SetRootHash(rHash)
}

// SetPrevHash sets prev hash
func (hv *HeaderWithValidatorStats) SetPrevHash(pvHash []byte) error {
	if hv == nil {
		return data.ErrNilPointerReceiver
	}

	return hv.Header.SetPrevHash(pvHash)
}

// SetPrevRandSeed sets previous random seed
func (hv *HeaderWithValidatorStats) SetPrevRandSeed(pvRandSeed []byte) error {
	if hv == nil {
		return data.ErrNilPointerReceiver
	}

	return hv.Header.SetPrevRandSeed(pvRandSeed)
}

// SetRandSeed sets previous random seed
func (hv *HeaderWithValidatorStats) SetRandSeed(randSeed []byte) error {
	if hv == nil {
		return data.ErrNilPointerReceiver
	}

	return hv.Header.SetRandSeed(randSeed)
}

// SetPubKeysBitmap sets public key bitmap
func (hv *HeaderWithValidatorStats) SetPubKeysBitmap(pkbm []byte) error {
	if hv == nil {
		return data.ErrNilPointerReceiver
	}

	return hv.Header.SetPubKeysBitmap(pkbm)
}

// SetSignature sets header signature
func (hv *HeaderWithValidatorStats) SetSignature(sg []byte) error {
	if hv == nil {
		return data.ErrNilPointerReceiver
	}

	return hv.Header.SetSignature(sg)
}

// SetLeaderSignature will set the leader's signature
func (hv *HeaderWithValidatorStats) SetLeaderSignature(sg []byte) error {
	if hv == nil {
		return data.ErrNilPointerReceiver
	}

	return hv.Header.SetLeaderSignature(sg)
}

// SetChainID sets the chain ID on which this block is valid on
func (hv *HeaderWithValidatorStats) SetChainID(chainID []byte) error {
	if hv == nil {
		return data.ErrNilPointerReceiver
	}

	return hv.Header.SetChainID(chainID)
}

// SetSoftwareVersion sets the software version of the header
func (hv *HeaderWithValidatorStats) SetSoftwareVersion(version []byte) error {
	if hv == nil {
		return data.ErrNilPointerReceiver
	}

	return hv.Header.SetSoftwareVersion(version)
}

// SetTimeStamp sets header timestamp
func (hv *HeaderWithValidatorStats) SetTimeStamp(ts uint64) error {
	if hv == nil {
		return data.ErrNilPointerReceiver
	}

	return hv.Header.SetTimeStamp(ts)
}

// SetAccumulatedFees sets the accumulated fees in the header
func (hv *HeaderWithValidatorStats) SetAccumulatedFees(value *big.Int) error {
	if hv == nil {
		return data.ErrNilPointerReceiver
	}

	return hv.Header.SetAccumulatedFees(value)
}

// SetDeveloperFees sets the developer fees in the header
func (hv *HeaderWithValidatorStats) SetDeveloperFees(value *big.Int) error {
	if hv == nil {
		return data.ErrNilPointerReceiver
	}

	return hv.Header.SetDeveloperFees(value)
}

// SetTxCount sets the transaction count of the block associated with this header
func (hv *HeaderWithValidatorStats) SetTxCount(txCount uint32) error {
	if hv == nil {
		return data.ErrNilPointerReceiver
	}

	return hv.Header.SetTxCount(txCount)
}

// SetShardID sets header shard ID
func (hv *HeaderWithValidatorStats) SetShardID(_ uint32) error {
	return nil
}

// SetValidatorStatsRootHash sets the root hash for the validator statistics trie
func (hv *HeaderWithValidatorStats) SetValidatorStatsRootHash(rHash []byte) error {
	if hv == nil {
		return data.ErrNilPointerReceiver
	}

	hv.ValidatorStatsRootHash = rHash

	return nil
}

// GetMiniBlockHeadersWithDst as a map of hashes and sender IDs
func (hv *HeaderWithValidatorStats) GetMiniBlockHeadersWithDst(destId uint32) map[string]uint32 {
	if hv == nil {
		return nil
	}

	return hv.Header.GetMiniBlockHeadersWithDst(destId)
}

// GetOrderedCrossMiniblocksWithDst gets all cross miniblocks with the given destination shard ID, ordered in a
// chronological way, taking into consideration the round in which they were created/executed in the sender shard
func (hv *HeaderWithValidatorStats) GetOrderedCrossMiniblocksWithDst(destId uint32) []*data.MiniBlockInfo {
	if hv == nil {
		return nil
	}

	return hv.Header.GetOrderedCrossMiniblocksWithDst(destId)
}

// GetMiniBlockHeadersHashes gets the miniblock hashes
func (hv *HeaderWithValidatorStats) GetMiniBlockHeadersHashes() [][]byte {
	if hv == nil {
		return nil
	}

	return hv.Header.GetMiniBlockHeadersHashes()
}

// MapMiniBlockHashesToShards is a map of mini block hashes and sender IDs
func (hv *HeaderWithValidatorStats) MapMiniBlockHashesToShards() map[string]uint32 {
	if hv == nil {
		return nil
	}

	return hv.Header.MapMiniBlockHashesToShards()
}

// IsInterfaceNil returns true if there is no value under the interface
func (hv *HeaderWithValidatorStats) IsInterfaceNil() bool {
	return hv == nil
}

// IsStartOfEpochBlock verifies if the block is of type start of epoch
func (hv *HeaderWithValidatorStats) IsStartOfEpochBlock() bool {
	return false
}

// GetBlockBodyTypeInt32 returns the block body type as int32
func (hv *HeaderWithValidatorStats) GetBlockBodyTypeInt32() int32 {
	if hv == nil {
		return -1
	}

	return hv.Header.GetBlockBodyTypeInt32()
}

// GetMiniBlockHeaderHandlers returns the miniBlock headers as an array of miniBlock header handlers
func (hv *HeaderWithValidatorStats) GetMiniBlockHeaderHandlers() []data.MiniBlockHeaderHandler {
	if hv == nil {
		return nil
	}

	return hv.Header.GetMiniBlockHeaderHandlers()
}

// SetMiniBlockHeaderHandlers sets the miniBlock headers from the given miniBlock header handlers
func (hv *HeaderWithValidatorStats) SetMiniBlockHeaderHandlers(mbHeaderHandlers []data.MiniBlockHeaderHandler) error {
	if hv == nil {
		return data.ErrNilPointerReceiver
	}

	return hv.Header.SetMiniBlockHeaderHandlers(mbHeaderHandlers)
}

// SetReceiptsHash sets the receipts hash
func (hv *HeaderWithValidatorStats) SetReceiptsHash(hash []byte) error {
	if hv == nil {
		return data.ErrNilPointerReceiver
	}

	return hv.Header.SetReceiptsHash(hash)
}

// HasScheduledSupport returns true as the second block version does support scheduled data
func (hv *HeaderWithValidatorStats) HasScheduledSupport() bool {
	return false
}

// ValidateHeaderVersion does extra validation for header version
func (hv *HeaderWithValidatorStats) ValidateHeaderVersion() error {
	if hv == nil {
		return data.ErrNilPointerReceiver
	}

	return nil
}
