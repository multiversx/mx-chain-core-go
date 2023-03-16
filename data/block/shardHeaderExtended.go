//go:generate protoc -I=. -I=$GOPATH/src -I=$GOPATH/src/github.com/multiversx/protobuf/protobuf  --gogoslick_out=. shardHeaderExtended.proto
package block

import (
	"fmt"
	"math/big"

	"github.com/multiversx/mx-chain-core-go/core/check"
	"github.com/multiversx/mx-chain-core-go/data"
	"github.com/multiversx/mx-chain-core-go/data/headerVersionData"
)

// GetShardID returns the header shardID
func (she *ShardHeaderExtended) GetShardID() uint32 {
	if she == nil {
		return 0
	}
	if she.Header == nil {
		return 0
	}

	return she.Header.GetShardID()
}

// GetNonce returns the header nonce
func (she *ShardHeaderExtended) GetNonce() uint64 {
	if she == nil {
		return 0
	}
	if she.Header == nil {
		return 0
	}

	return she.Header.GetNonce()
}

// GetEpoch returns the header epoch
func (she *ShardHeaderExtended) GetEpoch() uint32 {
	if she == nil {
		return 0
	}
	if she.Header == nil {
		return 0
	}

	return she.Header.GetEpoch()
}

// GetRound returns the header round
func (she *ShardHeaderExtended) GetRound() uint64 {
	if she == nil {
		return 0
	}
	if she.Header == nil {
		return 0
	}

	return she.Header.GetRound()
}

// GetRootHash returns the header root hash
func (she *ShardHeaderExtended) GetRootHash() []byte {
	if she == nil {
		return nil
	}
	if she.Header == nil {
		return nil
	}

	return she.Header.GetRootHash()
}

// GetPrevHash returns the header previous header hash
func (she *ShardHeaderExtended) GetPrevHash() []byte {
	if she == nil {
		return nil
	}
	if she.Header == nil {
		return nil
	}

	return she.Header.GetPrevHash()
}

// GetPrevRandSeed returns the header previous random seed
func (she *ShardHeaderExtended) GetPrevRandSeed() []byte {
	if she == nil {
		return nil
	}
	if she.Header == nil {
		return nil
	}

	return she.Header.GetPrevRandSeed()
}

// GetRandSeed returns the header random seed
func (she *ShardHeaderExtended) GetRandSeed() []byte {
	if she == nil {
		return nil
	}
	if she.Header == nil {
		return nil
	}

	return she.Header.GetRandSeed()
}

// GetPubKeysBitmap returns the header public key bitmap for the aggregated signatures
func (she *ShardHeaderExtended) GetPubKeysBitmap() []byte {
	if she == nil {
		return nil
	}
	if she.Header == nil {
		return nil
	}

	return she.Header.GetPubKeysBitmap()
}

// GetSignature returns the header aggregated signature
func (she *ShardHeaderExtended) GetSignature() []byte {
	if she == nil {
		return nil
	}
	if she.Header == nil {
		return nil
	}

	return she.Header.GetSignature()
}

// GetLeaderSignature returns the leader signature on top of the finalized (signed) header
func (she *ShardHeaderExtended) GetLeaderSignature() []byte {
	if she == nil {
		return nil
	}
	if she.Header == nil {
		return nil
	}

	return she.Header.GetLeaderSignature()
}

// GetChainID returns the chain ID
func (she *ShardHeaderExtended) GetChainID() []byte {
	if she == nil {
		return nil
	}
	if she.Header == nil {
		return nil
	}

	return she.Header.GetChainID()
}

// GetSoftwareVersion returns the header software version
func (she *ShardHeaderExtended) GetSoftwareVersion() []byte {
	if she == nil {
		return nil
	}
	if she.Header == nil {
		return nil
	}

	return she.Header.GetSoftwareVersion()
}

// GetTimeStamp returns the header timestamp
func (she *ShardHeaderExtended) GetTimeStamp() uint64 {
	if she == nil {
		return 0
	}
	if she.Header == nil {
		return 0
	}

	return she.Header.GetTimeStamp()
}

// GetTxCount returns the number of txs included in the block
func (she *ShardHeaderExtended) GetTxCount() uint32 {
	if she == nil {
		return 0
	}
	if she.Header == nil {
		return 0
	}

	return she.Header.GetTxCount()
}

// GetReceiptsHash returns the header receipt hash
func (she *ShardHeaderExtended) GetReceiptsHash() []byte {
	if she == nil {
		return nil
	}
	if she.Header == nil {
		return nil
	}

	return she.Header.GetReceiptsHash()
}

// GetAccumulatedFees returns the block accumulated fees
func (she *ShardHeaderExtended) GetAccumulatedFees() *big.Int {
	if she == nil {
		return nil
	}
	if she.Header == nil {
		return nil
	}

	return she.Header.GetAccumulatedFees()
}

// GetDeveloperFees returns the block developer fees
func (she *ShardHeaderExtended) GetDeveloperFees() *big.Int {
	if she == nil {
		return nil
	}
	if she.Header == nil {
		return nil
	}

	return she.Header.GetDeveloperFees()
}

// GetReserved returns the reserved field
func (she *ShardHeaderExtended) GetReserved() []byte {
	if she == nil {
		return nil
	}
	if she.Header == nil {
		return nil
	}

	return she.Header.GetReserved()
}

// GetMetaBlockHashes returns the metaBlock hashes
func (she *ShardHeaderExtended) GetMetaBlockHashes() [][]byte {
	if she == nil {
		return nil
	}
	if she.Header == nil {
		return nil
	}

	return she.Header.GetMetaBlockHashes()
}

// GetEpochStartMetaHash returns the epoch start metaBlock hash
func (she *ShardHeaderExtended) GetEpochStartMetaHash() []byte {
	if she == nil {
		return nil
	}
	if she.Header == nil {
		return nil
	}

	return she.Header.GetEpochStartMetaHash()
}

// SetNonce sets the header nonce
func (she *ShardHeaderExtended) SetNonce(n uint64) error {
	if she == nil {
		return data.ErrNilPointerReceiver
	}
	if she.Header == nil {
		return data.ErrNilHeader
	}

	return she.Header.SetNonce(n)
}

// SetEpoch sets the header epoch
func (she *ShardHeaderExtended) SetEpoch(e uint32) error {
	if she == nil {
		return data.ErrNilPointerReceiver
	}
	if she.Header == nil {
		return data.ErrNilHeader
	}

	return she.Header.SetEpoch(e)
}

// SetRound sets the header round
func (she *ShardHeaderExtended) SetRound(r uint64) error {
	if she == nil {
		return data.ErrNilPointerReceiver
	}
	if she.Header == nil {
		return data.ErrNilHeader
	}

	return she.Header.SetRound(r)
}

// SetRootHash sets the root hash
func (she *ShardHeaderExtended) SetRootHash(rHash []byte) error {
	if she == nil {
		return data.ErrNilPointerReceiver
	}
	if she.Header == nil {
		return data.ErrNilHeader
	}

	return she.Header.SetRootHash(rHash)
}

// SetValidatorStatsRootHash does nothing and returns nil
func (she *ShardHeaderExtended) SetValidatorStatsRootHash(_ []byte) error {
	return nil
}

// SetPrevHash sets the previous hash
func (she *ShardHeaderExtended) SetPrevHash(pvHash []byte) error {
	if she == nil {
		return data.ErrNilPointerReceiver
	}
	if she.Header == nil {
		return data.ErrNilHeader
	}

	return she.Header.SetPrevHash(pvHash)
}

// SetPrevRandSeed sets the previous random seed
func (she *ShardHeaderExtended) SetPrevRandSeed(pvRandSeed []byte) error {
	if she == nil {
		return data.ErrNilPointerReceiver
	}
	if she.Header == nil {
		return data.ErrNilHeader
	}

	return she.Header.SetPrevRandSeed(pvRandSeed)
}

// SetRandSeed sets the random seed
func (she *ShardHeaderExtended) SetRandSeed(randSeed []byte) error {
	if she == nil {
		return data.ErrNilPointerReceiver
	}
	if she.Header == nil {
		return data.ErrNilHeader
	}

	return she.Header.SetRandSeed(randSeed)
}

// SetPubKeysBitmap sets the public key bitmap
func (she *ShardHeaderExtended) SetPubKeysBitmap(pkbm []byte) error {
	if she == nil {
		return data.ErrNilPointerReceiver
	}
	if she.Header == nil {
		return data.ErrNilHeader
	}

	return she.Header.SetPubKeysBitmap(pkbm)
}

// SetSignature sets the header signature
func (she *ShardHeaderExtended) SetSignature(sg []byte) error {
	if she == nil {
		return data.ErrNilPointerReceiver
	}
	if she.Header == nil {
		return data.ErrNilHeader
	}

	return she.Header.SetSignature(sg)
}

// SetLeaderSignature sets the leader's signature
func (she *ShardHeaderExtended) SetLeaderSignature(sg []byte) error {
	if she == nil {
		return data.ErrNilPointerReceiver
	}
	if she.Header == nil {
		return data.ErrNilHeader
	}

	return she.Header.SetLeaderSignature(sg)
}

// SetChainID sets the chain ID on which this block is valid on
func (she *ShardHeaderExtended) SetChainID(chainID []byte) error {
	if she == nil {
		return data.ErrNilPointerReceiver
	}
	if she.Header == nil {
		return data.ErrNilHeader
	}

	return she.Header.SetChainID(chainID)
}

// SetSoftwareVersion sets the software version of the header
func (she *ShardHeaderExtended) SetSoftwareVersion(version []byte) error {
	if she == nil {
		return data.ErrNilPointerReceiver
	}
	if she.Header == nil {
		return data.ErrNilHeader
	}

	return she.Header.SetSoftwareVersion(version)
}

// SetTimeStamp sets the header timestamp
func (she *ShardHeaderExtended) SetTimeStamp(ts uint64) error {
	if she == nil {
		return data.ErrNilPointerReceiver
	}
	if she.Header == nil {
		return data.ErrNilHeader
	}

	return she.Header.SetTimeStamp(ts)
}

// SetAccumulatedFees sets the accumulated fees in the header
func (she *ShardHeaderExtended) SetAccumulatedFees(value *big.Int) error {
	if she == nil {
		return data.ErrNilPointerReceiver
	}
	if she.Header == nil {
		return data.ErrNilHeader
	}

	return she.Header.SetAccumulatedFees(value)
}

// SetDeveloperFees sets the developer fees in the header
func (she *ShardHeaderExtended) SetDeveloperFees(value *big.Int) error {
	if she == nil {
		return data.ErrNilPointerReceiver
	}
	if she.Header == nil {
		return data.ErrNilHeader
	}

	return she.Header.SetDeveloperFees(value)
}

// SetTxCount sets the transaction count of the block associated with this header
func (she *ShardHeaderExtended) SetTxCount(txCount uint32) error {
	if she == nil {
		return data.ErrNilPointerReceiver
	}
	if she.Header == nil {
		return data.ErrNilHeader
	}

	return she.Header.SetTxCount(txCount)
}

// SetShardID sets the header shard ID
func (she *ShardHeaderExtended) SetShardID(shId uint32) error {
	if she == nil {
		return data.ErrNilPointerReceiver
	}
	if she.Header == nil {
		return data.ErrNilHeader
	}

	return she.Header.SetShardID(shId)
}

// GetMiniBlockHeadersWithDst gets the map of miniBlockHeader hashes and sender IDs
func (she *ShardHeaderExtended) GetMiniBlockHeadersWithDst(destId uint32) map[string]uint32 {
	if she == nil {
		return nil
	}
	if she.Header == nil {
		return nil
	}

	return she.Header.GetMiniBlockHeadersWithDst(destId)
}

// GetOrderedCrossMiniblocksWithDst gets all the cross miniBlocks with the given destination shard ID, ordered in a
// chronological way, taking into consideration the round in which they were created/executed in the sender shard
func (she *ShardHeaderExtended) GetOrderedCrossMiniblocksWithDst(destId uint32) []*data.MiniBlockInfo {
	if she == nil {
		return nil
	}
	if she.Header == nil {
		return nil
	}

	return she.Header.GetOrderedCrossMiniblocksWithDst(destId)
}

// GetMiniBlockHeadersHashes gets the miniBlock hashes
func (she *ShardHeaderExtended) GetMiniBlockHeadersHashes() [][]byte {
	if she == nil {
		return nil
	}
	if she.Header == nil {
		return nil
	}

	return she.Header.GetMiniBlockHeadersHashes()
}

// MapMiniBlockHashesToShards gets the map of miniBlock hashes and sender IDs
func (she *ShardHeaderExtended) MapMiniBlockHashesToShards() map[string]uint32 {
	if she == nil {
		return nil
	}
	if she.Header == nil {
		return nil
	}

	return she.Header.MapMiniBlockHashesToShards()
}

// ShallowClone returns a clone of the object
func (she *ShardHeaderExtended) ShallowClone() data.HeaderHandler {
	if she == nil {
		return nil
	}
	if she.Header == nil {
		return nil
	}

	internalHeaderCopy := *she.Header
	headerCopy := *she
	headerCopy.Header = &internalHeaderCopy

	return &headerCopy
}

// IsInterfaceNil returns true if there is no value under the interface
func (she *ShardHeaderExtended) IsInterfaceNil() bool {
	return she == nil
}

// IsStartOfEpochBlock verifies if the block is of type start of epoch
func (she *ShardHeaderExtended) IsStartOfEpochBlock() bool {
	if she == nil {
		return false
	}
	if she.Header == nil {
		return false
	}

	return she.Header.IsStartOfEpochBlock()
}

// GetBlockBodyTypeInt32 returns the blockBody type as int32
func (she *ShardHeaderExtended) GetBlockBodyTypeInt32() int32 {
	if she == nil {
		return -1
	}
	if she.Header == nil {
		return -1
	}

	return she.Header.GetBlockBodyTypeInt32()
}

// GetMiniBlockHeaderHandlers returns the miniBlock headers as an array of miniBlock header handlers
func (she *ShardHeaderExtended) GetMiniBlockHeaderHandlers() []data.MiniBlockHeaderHandler {
	if she == nil {
		return nil
	}
	if she.Header == nil {
		return nil
	}

	return she.Header.GetMiniBlockHeaderHandlers()
}

// SetMiniBlockHeaderHandlers sets the miniBlock headers from the given miniBlock header handlers
func (she *ShardHeaderExtended) SetMiniBlockHeaderHandlers(mbHeaderHandlers []data.MiniBlockHeaderHandler) error {
	if she == nil {
		return data.ErrNilPointerReceiver
	}
	if she.Header == nil {
		return data.ErrNilHeader
	}

	return she.Header.SetMiniBlockHeaderHandlers(mbHeaderHandlers)
}

// SetReceiptsHash sets the receipts hash
func (she *ShardHeaderExtended) SetReceiptsHash(hash []byte) error {
	if she == nil {
		return data.ErrNilPointerReceiver
	}
	if she.Header == nil {
		return data.ErrNilHeader
	}

	return she.Header.SetReceiptsHash(hash)
}

// SetMetaBlockHashes sets the metaBlock hashes
func (she *ShardHeaderExtended) SetMetaBlockHashes(hashes [][]byte) error {
	if she == nil {
		return data.ErrNilPointerReceiver
	}
	if she.Header == nil {
		return data.ErrNilHeader
	}

	return she.Header.SetMetaBlockHashes(hashes)
}

// SetEpochStartMetaHash sets the epoch start metaBlock hash
func (she *ShardHeaderExtended) SetEpochStartMetaHash(hash []byte) error {
	if she == nil {
		return data.ErrNilPointerReceiver
	}
	if she.Header == nil {
		return data.ErrNilHeader
	}

	return she.Header.SetEpochStartMetaHash(hash)
}

// HasScheduledSupport returns true as the second block version does support scheduled data
func (she *ShardHeaderExtended) HasScheduledSupport() bool {
	return true
}

// HasScheduledMiniBlocks returns true if the header has scheduled miniBlock headers
func (she *ShardHeaderExtended) HasScheduledMiniBlocks() bool {
	if she == nil {
		return false
	}

	mbHeaderHandlers := she.GetMiniBlockHeaderHandlers()
	for _, mbHeader := range mbHeaderHandlers {
		processingType := ProcessingType(mbHeader.GetProcessingType())
		if processingType == Scheduled {
			return true
		}
	}

	return false
}

// SetScheduledRootHash sets the scheduled root hash
func (she *ShardHeaderExtended) SetScheduledRootHash(rootHash []byte) error {
	if she == nil {
		return data.ErrNilPointerReceiver
	}
	if she.Header == nil {
		return data.ErrNilHeader
	}

	she.Header.ScheduledRootHash = rootHash

	return nil
}

// SetScheduledAccumulatedFees sets the scheduled accumulated fees
func (she *ShardHeaderExtended) SetScheduledAccumulatedFees(value *big.Int) error {
	if she == nil {
		return data.ErrNilPointerReceiver
	}
	if she.Header == nil {
		return data.ErrNilHeader
	}

	if she.Header.ScheduledAccumulatedFees == nil {
		she.Header.ScheduledAccumulatedFees = big.NewInt(0)
	}
	if value == nil {
		value = big.NewInt(0)
	}

	she.Header.ScheduledAccumulatedFees.Set(value)

	return nil
}

// SetScheduledDeveloperFees sets the scheduled developer fees
func (she *ShardHeaderExtended) SetScheduledDeveloperFees(value *big.Int) error {
	if she == nil {
		return data.ErrNilPointerReceiver
	}
	if she.Header == nil {
		return data.ErrNilHeader
	}

	if she.Header.ScheduledDeveloperFees == nil {
		she.Header.ScheduledDeveloperFees = big.NewInt(0)
	}
	if value == nil {
		value = big.NewInt(0)
	}

	she.Header.ScheduledDeveloperFees.Set(value)

	return nil
}

// SetScheduledGasProvided sets the scheduled SC calls provided gas
func (she *ShardHeaderExtended) SetScheduledGasProvided(gasProvided uint64) error {
	if she == nil {
		return data.ErrNilPointerReceiver
	}
	if she.Header == nil {
		return data.ErrNilHeader
	}

	she.Header.ScheduledGasProvided = gasProvided

	return nil
}

// SetScheduledGasPenalized sets the scheduled SC calls penalized gas
func (she *ShardHeaderExtended) SetScheduledGasPenalized(gasPenalized uint64) error {
	if she == nil {
		return data.ErrNilPointerReceiver
	}
	if she.Header == nil {
		return data.ErrNilHeader
	}

	she.Header.ScheduledGasPenalized = gasPenalized

	return nil
}

// SetScheduledGasRefunded sets the scheduled SC calls refunded gas
func (she *ShardHeaderExtended) SetScheduledGasRefunded(gasRefunded uint64) error {
	if she == nil {
		return data.ErrNilPointerReceiver
	}
	if she.Header == nil {
		return data.ErrNilHeader
	}

	she.Header.ScheduledGasRefunded = gasRefunded

	return nil
}

// ValidateHeaderVersion does extra validation for header version
func (she *ShardHeaderExtended) ValidateHeaderVersion() error {
	if she == nil {
		return data.ErrNilPointerReceiver
	}
	if she.Header == nil {
		return data.ErrNilHeader
	}

	// the header needs to have a not nil & not empty scheduled root hash
	if len(she.Header.ScheduledRootHash) == 0 {
		return data.ErrNilScheduledRootHash
	}

	return she.Header.ValidateHeaderVersion()
}

// SetAdditionalData sets the additional version related data for the header
func (she *ShardHeaderExtended) SetAdditionalData(headerVersionData headerVersionData.HeaderAdditionalData) error {
	if she == nil {
		return data.ErrNilPointerReceiver
	}
	if she.Header == nil {
		return data.ErrNilHeader
	}
	if check.IfNil(headerVersionData) {
		return data.ErrNilPointerDereference
	}

	err := she.SetScheduledRootHash(headerVersionData.GetScheduledRootHash())
	if err != nil {
		return err
	}

	she.Header.ScheduledGasProvided = headerVersionData.GetScheduledGasProvided()
	she.Header.ScheduledGasPenalized = headerVersionData.GetScheduledGasPenalized()
	she.Header.ScheduledGasRefunded = headerVersionData.GetScheduledGasRefunded()

	err = she.SetScheduledAccumulatedFees(headerVersionData.GetScheduledAccumulatedFees())
	if err != nil {
		return err
	}

	return she.SetScheduledDeveloperFees(headerVersionData.GetScheduledDeveloperFees())
}

// GetAdditionalData gets the additional version related data for the header
func (she *ShardHeaderExtended) GetAdditionalData() headerVersionData.HeaderAdditionalData {
	if she == nil {
		return nil
	}
	if she.Header == nil {
		return nil
	}

	accFees := big.NewInt(0)
	if she.Header.GetScheduledAccumulatedFees() != nil {
		accFees = big.NewInt(0).Set(she.Header.GetScheduledAccumulatedFees())
	}
	devFees := big.NewInt(0)
	if she.Header.GetScheduledDeveloperFees() != nil {
		devFees = big.NewInt(0).Set(she.Header.GetScheduledDeveloperFees())
	}

	additionalVersionData := &headerVersionData.AdditionalData{
		ScheduledRootHash:        she.Header.GetScheduledRootHash(),
		ScheduledAccumulatedFees: accFees,
		ScheduledDeveloperFees:   devFees,
		ScheduledGasProvided:     she.Header.GetScheduledGasProvided(),
		ScheduledGasPenalized:    she.Header.GetScheduledGasPenalized(),
		ScheduledGasRefunded:     she.Header.GetScheduledGasRefunded(),
	}

	return additionalVersionData
}

// GetValidatorStatsRootHash returns an empty byte slice
func (she *ShardHeaderExtended) GetValidatorStatsRootHash() []byte {
	return make([]byte, 0)
}

// CheckFieldsForNil checks a predefined set of fields for nil values
func (she *ShardHeaderExtended) CheckFieldsForNil() error {
	if she == nil {
		return data.ErrNilPointerReceiver
	}
	if she.Header == nil {
		return data.ErrNilHeader
	}
	err := she.Header.CheckFieldsForNil()
	if err != nil {
		return err
	}
	if she.Header.ScheduledAccumulatedFees == nil {
		return fmt.Errorf("%w in ShardHeaderExtended.ScheduledAccumulatedFees", data.ErrNilValue)
	}
	if she.Header.ScheduledDeveloperFees == nil {
		return fmt.Errorf("%w in ShardHeaderExtended.ScheduledDeveloperFees", data.ErrNilValue)
	}

	return nil
}

// GetIncomingMiniBlockHeaderHandlers gets the incoming mini blocks headers as an array of mini blocks headers handlers
func (she *ShardHeaderExtended) GetIncomingMiniBlockHeaderHandlers() []data.MiniBlockHeaderHandler {
	if she == nil {
		return nil
	}

	mbHeaders := she.GetIncomingMiniBlockHeaders()
	mbHeaderHandlers := make([]data.MiniBlockHeaderHandler, len(mbHeaders))

	for i := range mbHeaders {
		mbHeaderHandlers[i] = &mbHeaders[i]
	}

	return mbHeaderHandlers
}

// SetIncomingMiniBlockHeaderHandlers sets the incoming mini blocks headers from the given array of mini blocks headers handlers
func (she *ShardHeaderExtended) SetIncomingMiniBlockHeaderHandlers(mbHeaderHandlers []data.MiniBlockHeaderHandler) error {
	if she == nil {
		return data.ErrNilPointerReceiver
	}
	if len(mbHeaderHandlers) == 0 {
		she.IncomingMiniBlockHeaders = nil
		return nil
	}

	incomingMiniBlockHeaders := make([]MiniBlockHeader, len(mbHeaderHandlers))
	for i, mbHeaderHandler := range mbHeaderHandlers {
		mbHeader, ok := mbHeaderHandler.(*MiniBlockHeader)
		if !ok {
			return data.ErrInvalidTypeAssertion
		}

		if mbHeader == nil {
			return data.ErrNilPointerDereference
		}

		incomingMiniBlockHeaders[i] = *mbHeader
	}

	she.IncomingMiniBlockHeaders = incomingMiniBlockHeaders

	return nil
}
