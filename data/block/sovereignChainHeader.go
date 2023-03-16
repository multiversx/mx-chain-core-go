//go:generate protoc -I=. -I=$GOPATH/src -I=$GOPATH/src/github.com/multiversx/protobuf/protobuf  --gogoslick_out=. sovereignChainHeader.proto
package block

import (
	"fmt"
	"math/big"

	"github.com/multiversx/mx-chain-core-go/data"
	"github.com/multiversx/mx-chain-core-go/data/headerVersionData"
)

// GetAdditionalData returns nil for the sovereign chain header
func (sch *SovereignChainHeader) GetAdditionalData() headerVersionData.HeaderAdditionalData {
	return nil
}

// HasScheduledMiniBlocks returns false for the sovereign chain header
func (sch *SovereignChainHeader) HasScheduledMiniBlocks() bool {
	return false
}

// SetScheduledRootHash does nothing and returns nil for the sovereign chain header
func (sch *SovereignChainHeader) SetScheduledRootHash(_ []byte) error {
	return nil
}

// SetAdditionalData does nothing and returns nil for the sovereign chain header
func (sch *SovereignChainHeader) SetAdditionalData(_ headerVersionData.HeaderAdditionalData) error {
	return nil
}

// ShallowClone returns a clone of the object
func (sch *SovereignChainHeader) ShallowClone() data.HeaderHandler {
	if sch == nil || sch.Header == nil {
		return nil
	}

	internalHeaderCopy := *sch.Header
	headerCopy := *sch
	headerCopy.Header = &internalHeaderCopy

	return &headerCopy
}

// GetShardID returns 0 as the shardID for the sovereign chain header
func (sch *SovereignChainHeader) GetShardID() uint32 {
	return 0
}

// GetNonce returns the header nonce
func (sch *SovereignChainHeader) GetNonce() uint64 {
	if sch == nil {
		return 0
	}

	return sch.Header.GetNonce()
}

// GetEpoch returns the header epoch
func (sch *SovereignChainHeader) GetEpoch() uint32 {
	if sch == nil {
		return 0
	}

	return sch.Header.GetEpoch()
}

// GetRound returns the header round
func (sch *SovereignChainHeader) GetRound() uint64 {
	if sch == nil {
		return 0
	}

	return sch.Header.GetRound()
}

// GetRootHash returns the header root hash
func (sch *SovereignChainHeader) GetRootHash() []byte {
	if sch == nil {
		return nil
	}

	return sch.Header.GetRootHash()
}

// GetPrevHash returns the header previous header hash
func (sch *SovereignChainHeader) GetPrevHash() []byte {
	if sch == nil {
		return nil
	}

	return sch.Header.GetPrevHash()
}

// GetPrevRandSeed returns the header previous random seed
func (sch *SovereignChainHeader) GetPrevRandSeed() []byte {
	if sch == nil {
		return nil
	}

	return sch.Header.GetPrevRandSeed()
}

// GetRandSeed returns the header random seed
func (sch *SovereignChainHeader) GetRandSeed() []byte {
	if sch == nil {
		return nil
	}

	return sch.Header.GetRandSeed()
}

// GetPubKeysBitmap returns the header public key bitmap for the aggregated signatures
func (sch *SovereignChainHeader) GetPubKeysBitmap() []byte {
	if sch == nil {
		return nil
	}

	return sch.Header.GetPubKeysBitmap()
}

// GetSignature returns the header aggregated signature
func (sch *SovereignChainHeader) GetSignature() []byte {
	if sch == nil {
		return nil
	}

	return sch.Header.GetSignature()
}

// GetLeaderSignature returns the leader signature on top of the finalized (signed) header
func (sch *SovereignChainHeader) GetLeaderSignature() []byte {
	if sch == nil {
		return nil
	}

	return sch.Header.GetLeaderSignature()
}

// GetChainID returns the chain ID
func (sch *SovereignChainHeader) GetChainID() []byte {
	if sch == nil {
		return nil
	}

	return sch.Header.GetChainID()
}

// GetSoftwareVersion returns the header software version
func (sch *SovereignChainHeader) GetSoftwareVersion() []byte {
	if sch == nil {
		return nil
	}

	return sch.Header.GetSoftwareVersion()
}

// GetTimeStamp returns the header timestamp
func (sch *SovereignChainHeader) GetTimeStamp() uint64 {
	if sch == nil {
		return 0
	}

	return sch.Header.GetTimeStamp()
}

// GetTxCount returns the number of txs included in the block
func (sch *SovereignChainHeader) GetTxCount() uint32 {
	if sch == nil {
		return 0
	}

	return sch.Header.GetTxCount()
}

// GetReceiptsHash returns the header receipt hash
func (sch *SovereignChainHeader) GetReceiptsHash() []byte {
	if sch == nil {
		return nil
	}

	return sch.Header.GetReceiptsHash()
}

// GetAccumulatedFees returns the block accumulated fees
func (sch *SovereignChainHeader) GetAccumulatedFees() *big.Int {
	if sch == nil {
		return nil
	}

	return sch.Header.GetAccumulatedFees()
}

// GetDeveloperFees returns the block developer fees
func (sch *SovereignChainHeader) GetDeveloperFees() *big.Int {
	if sch == nil {
		return nil
	}

	return sch.Header.GetDeveloperFees()
}

// GetReserved returns the reserved field
func (sch *SovereignChainHeader) GetReserved() []byte {
	if sch == nil {
		return nil
	}

	return sch.Header.GetReserved()
}

// GetMetaBlockHashes returns the metaBlock hashes
func (sch *SovereignChainHeader) GetMetaBlockHashes() [][]byte {
	if sch == nil {
		return nil
	}

	return sch.Header.GetMetaBlockHashes()
}

// GetEpochStartMetaHash returns the epoch start metaBlock hash
func (sch *SovereignChainHeader) GetEpochStartMetaHash() []byte {
	if sch == nil {
		return nil
	}

	return sch.Header.GetEpochStartMetaHash()
}

// SetNonce sets the header nonce
func (sch *SovereignChainHeader) SetNonce(n uint64) error {
	if sch == nil {
		return data.ErrNilPointerReceiver
	}
	return sch.Header.SetNonce(n)
}

// SetEpoch sets the header epoch
func (sch *SovereignChainHeader) SetEpoch(e uint32) error {
	if sch == nil {
		return data.ErrNilPointerReceiver
	}

	return sch.Header.SetEpoch(e)
}

// SetRound sets the header round
func (sch *SovereignChainHeader) SetRound(r uint64) error {
	if sch == nil {
		return data.ErrNilPointerReceiver
	}

	return sch.Header.SetRound(r)
}

// SetRootHash sets the root hash
func (sch *SovereignChainHeader) SetRootHash(rHash []byte) error {
	if sch == nil {
		return data.ErrNilPointerReceiver
	}

	return sch.Header.SetRootHash(rHash)
}

// SetPrevHash sets the previous hash
func (sch *SovereignChainHeader) SetPrevHash(pvHash []byte) error {
	if sch == nil {
		return data.ErrNilPointerReceiver
	}

	return sch.Header.SetPrevHash(pvHash)
}

// SetPrevRandSeed sets the previous random seed
func (sch *SovereignChainHeader) SetPrevRandSeed(pvRandSeed []byte) error {
	if sch == nil {
		return data.ErrNilPointerReceiver
	}

	return sch.Header.SetPrevRandSeed(pvRandSeed)
}

// SetRandSeed sets the random seed
func (sch *SovereignChainHeader) SetRandSeed(randSeed []byte) error {
	if sch == nil {
		return data.ErrNilPointerReceiver
	}

	return sch.Header.SetRandSeed(randSeed)
}

// SetPubKeysBitmap sets the public key bitmap
func (sch *SovereignChainHeader) SetPubKeysBitmap(pkbm []byte) error {
	if sch == nil {
		return data.ErrNilPointerReceiver
	}

	return sch.Header.SetPubKeysBitmap(pkbm)
}

// SetSignature sets the header signature
func (sch *SovereignChainHeader) SetSignature(sg []byte) error {
	if sch == nil {
		return data.ErrNilPointerReceiver
	}

	return sch.Header.SetSignature(sg)
}

// SetLeaderSignature sets the leader's signature
func (sch *SovereignChainHeader) SetLeaderSignature(sg []byte) error {
	if sch == nil {
		return data.ErrNilPointerReceiver
	}

	return sch.Header.SetLeaderSignature(sg)
}

// SetChainID sets the chain ID on which this block is valid on
func (sch *SovereignChainHeader) SetChainID(chainID []byte) error {
	if sch == nil {
		return data.ErrNilPointerReceiver
	}

	return sch.Header.SetChainID(chainID)
}

// SetSoftwareVersion sets the software version of the header
func (sch *SovereignChainHeader) SetSoftwareVersion(version []byte) error {
	if sch == nil {
		return data.ErrNilPointerReceiver
	}

	return sch.Header.SetSoftwareVersion(version)
}

// SetTimeStamp sets the header timestamp
func (sch *SovereignChainHeader) SetTimeStamp(ts uint64) error {
	if sch == nil {
		return data.ErrNilPointerReceiver
	}

	return sch.Header.SetTimeStamp(ts)
}

// SetAccumulatedFees sets the accumulated fees in the header
func (sch *SovereignChainHeader) SetAccumulatedFees(value *big.Int) error {
	if sch == nil {
		return data.ErrNilPointerReceiver
	}

	return sch.Header.SetAccumulatedFees(value)
}

// SetDeveloperFees sets the developer fees in the header
func (sch *SovereignChainHeader) SetDeveloperFees(value *big.Int) error {
	if sch == nil {
		return data.ErrNilPointerReceiver
	}

	return sch.Header.SetDeveloperFees(value)
}

// SetTxCount sets the transaction count of the block associated with this header
func (sch *SovereignChainHeader) SetTxCount(txCount uint32) error {
	if sch == nil {
		return data.ErrNilPointerReceiver
	}

	return sch.Header.SetTxCount(txCount)
}

// SetShardID does nothing and returns nil for the sovereign chain header
func (sch *SovereignChainHeader) SetShardID(_ uint32) error {
	return nil
}

// SetValidatorStatsRootHash sets the root hash for the validator statistics trie
func (sch *SovereignChainHeader) SetValidatorStatsRootHash(rootHash []byte) error {
	if sch == nil {
		return data.ErrNilPointerReceiver
	}

	sch.ValidatorStatsRootHash = rootHash

	return nil
}

// SetMainChainShardHeaderHashes sets the main chain shard header hashes
func (sch *SovereignChainHeader) SetMainChainShardHeaderHashes(hdrHashes [][]byte) error {
	if sch == nil {
		return data.ErrNilPointerReceiver
	}

	sch.MainChainShardHeaderHashes = hdrHashes

	return nil
}

// GetMiniBlockHeadersWithDst returns the miniblocks headers hashes for the destination shard
func (sch *SovereignChainHeader) GetMiniBlockHeadersWithDst(destId uint32) map[string]uint32 {
	if sch == nil {
		return nil
	}

	return sch.Header.GetMiniBlockHeadersWithDst(destId)
}

// GetOrderedCrossMiniblocksWithDst gets all the cross miniblocks with the given destination shard ID, ordered in a
// chronological way, taking into consideration the round in which they were created/executed in the sender shard
func (sch *SovereignChainHeader) GetOrderedCrossMiniblocksWithDst(destId uint32) []*data.MiniBlockInfo {
	if sch == nil {
		return nil
	}

	return sch.Header.GetOrderedCrossMiniblocksWithDst(destId)
}

// GetMiniBlockHeadersHashes gets the miniBlock hashes
func (sch *SovereignChainHeader) GetMiniBlockHeadersHashes() [][]byte {
	if sch == nil {
		return nil
	}

	return sch.Header.GetMiniBlockHeadersHashes()
}

// MapMiniBlockHashesToShards gets the map of miniBlock hashes and sender IDs
func (sch *SovereignChainHeader) MapMiniBlockHashesToShards() map[string]uint32 {
	if sch == nil {
		return nil
	}

	return sch.Header.MapMiniBlockHashesToShards()
}

// IsInterfaceNil returns true if there is no value under the interface
func (sch *SovereignChainHeader) IsInterfaceNil() bool {
	return sch == nil
}

// IsStartOfEpochBlock returns false for the sovereign chain header
func (sch *SovereignChainHeader) IsStartOfEpochBlock() bool {
	return false
}

// GetBlockBodyTypeInt32 returns the blockBody type as int32
func (sch *SovereignChainHeader) GetBlockBodyTypeInt32() int32 {
	if sch == nil {
		return -1
	}

	return sch.Header.GetBlockBodyTypeInt32()
}

// GetMiniBlockHeaderHandlers returns the miniBlock headers as an array of miniBlock header handlers
func (sch *SovereignChainHeader) GetMiniBlockHeaderHandlers() []data.MiniBlockHeaderHandler {
	if sch == nil {
		return nil
	}

	return sch.Header.GetMiniBlockHeaderHandlers()
}

// SetMiniBlockHeaderHandlers sets the miniBlock headers from the given miniBlock header handlers
func (sch *SovereignChainHeader) SetMiniBlockHeaderHandlers(mbHeaderHandlers []data.MiniBlockHeaderHandler) error {
	if sch == nil {
		return data.ErrNilPointerReceiver
	}

	return sch.Header.SetMiniBlockHeaderHandlers(mbHeaderHandlers)
}

// SetReceiptsHash sets the receipts hash
func (sch *SovereignChainHeader) SetReceiptsHash(hash []byte) error {
	if sch == nil {
		return data.ErrNilPointerReceiver
	}

	return sch.Header.SetReceiptsHash(hash)
}

// SetMetaBlockHashes sets the metaBlock hashes
func (sch *SovereignChainHeader) SetMetaBlockHashes(hashes [][]byte) error {
	if sch == nil {
		return data.ErrNilPointerReceiver
	}

	return sch.Header.SetMetaBlockHashes(hashes)
}

// SetEpochStartMetaHash sets the epoch start metaBlock hash
func (sch *SovereignChainHeader) SetEpochStartMetaHash(hash []byte) error {
	if sch == nil {
		return data.ErrNilPointerReceiver
	}

	return sch.Header.SetEpochStartMetaHash(hash)
}

// HasScheduledSupport returns false for the sovereign chain header
func (sch *SovereignChainHeader) HasScheduledSupport() bool {
	return false
}

// ValidateHeaderVersion does extra validation for the header version
func (sch *SovereignChainHeader) ValidateHeaderVersion() error {
	if sch == nil {
		return data.ErrNilPointerReceiver
	}

	return nil
}

// CheckFieldsForNil checks a predefined set of fields for nil values
func (sch *SovereignChainHeader) CheckFieldsForNil() error {
	if sch == nil {
		return data.ErrNilPointerReceiver
	}
	err := sch.Header.CheckFieldsForNil()
	if err != nil {
		return err
	}

	if sch.ValidatorStatsRootHash == nil {
		return fmt.Errorf("%w in sch.ValidatorStatsRootHash", data.ErrNilValue)
	}

	return nil
}