//go:generate protoc -I=. -I=$GOPATH/src -I=$GOPATH/src/github.com/multiversx/protobuf/protobuf  --gogoslick_out=. sovereignChainHeader.proto
package block

import (
	"fmt"
	"math/big"

	"github.com/multiversx/mx-chain-core-go/core"
	"github.com/multiversx/mx-chain-core-go/core/check"
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

	if !check.IfNil(sch.OutGoingMiniBlockHeader) {
		internalOutGoingMbHeader := *sch.OutGoingMiniBlockHeader
		headerCopy.OutGoingMiniBlockHeader = &internalOutGoingMbHeader
	}

	return &headerCopy
}

// GetShardID returns internal header shard id
func (sch *SovereignChainHeader) GetShardID() uint32 {
	if sch == nil {
		return 0
	}

	return sch.Header.ShardID
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

// SetShardID sets the shard id
func (sch *SovereignChainHeader) SetShardID(shardID uint32) error {
	if sch == nil {
		return data.ErrNilPointerReceiver
	}

	return sch.Header.SetShardID(shardID)
}

// SetValidatorStatsRootHash sets the root hash for the validator statistics trie
func (sch *SovereignChainHeader) SetValidatorStatsRootHash(rootHash []byte) error {
	if sch == nil {
		return data.ErrNilPointerReceiver
	}

	sch.ValidatorStatsRootHash = rootHash

	return nil
}

// SetExtendedShardHeaderHashes sets the extended shard header hashes
func (sch *SovereignChainHeader) SetExtendedShardHeaderHashes(hdrHashes [][]byte) error {
	if sch == nil {
		return data.ErrNilPointerReceiver
	}

	sch.ExtendedShardHeaderHashes = hdrHashes

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
	if sch == nil {
		return false
	}

	return sch.IsStartOfEpoch
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

// GetOutGoingMiniBlockHeaderHandler returns the outgoing mini block header
func (sch *SovereignChainHeader) GetOutGoingMiniBlockHeaderHandler() data.OutGoingMiniBlockHeaderHandler {
	if sch == nil {
		return nil
	}

	return sch.GetOutGoingMiniBlockHeader()
}

// SetOutGoingMiniBlockHeaderHandler returns the outgoing mini block header
func (sch *SovereignChainHeader) SetOutGoingMiniBlockHeaderHandler(mbHeader data.OutGoingMiniBlockHeaderHandler) error {
	if sch == nil {
		return data.ErrNilPointerReceiver
	}

	if check.IfNil(mbHeader) {
		sch.OutGoingMiniBlockHeader = nil
		return nil
	}

	sch.OutGoingMiniBlockHeader = &OutGoingMiniBlockHeader{
		Hash:                                  mbHeader.GetHash(),
		OutGoingOperationsHash:                mbHeader.GetOutGoingOperationsHash(),
		AggregatedSignatureOutGoingOperations: mbHeader.GetAggregatedSignatureOutGoingOperations(),
		LeaderSignatureOutGoingOperations:     mbHeader.GetLeaderSignatureOutGoingOperations(),
	}

	return nil
}

// SetBlockBodyTypeInt32 sets the blockBodyType in the header
func (sch *SovereignChainHeader) SetBlockBodyTypeInt32(blockBodyType int32) error {
	if sch == nil {
		return data.ErrNilPointerReceiver
	}

	return sch.Header.SetBlockBodyTypeInt32(blockBodyType)
}

// SetStartOfEpochHeader sets the bool flag for epoch start header
func (sch *SovereignChainHeader) SetStartOfEpochHeader() error {
	if sch == nil {
		return data.ErrNilPointerReceiver
	}

	sch.IsStartOfEpoch = true
	return nil
}

// SetDevFeesInEpoch sets the developer fees in the header
func (sch *SovereignChainHeader) SetDevFeesInEpoch(value *big.Int) error {
	if sch == nil {
		return data.ErrNilPointerReceiver
	}
	if value == nil {
		return data.ErrInvalidValue
	}
	if sch.DevFeesInEpoch == nil {
		sch.DevFeesInEpoch = big.NewInt(0)
	}

	sch.DevFeesInEpoch.Set(value)

	return nil
}

// SetAccumulatedFeesInEpoch sets the epoch accumulated fees in the header
func (sch *SovereignChainHeader) SetAccumulatedFeesInEpoch(value *big.Int) error {
	if sch == nil {
		return data.ErrNilPointerReceiver
	}
	if value == nil {
		return data.ErrInvalidValue
	}
	if sch.AccumulatedFeesInEpoch == nil {
		sch.AccumulatedFeesInEpoch = big.NewInt(0)
	}

	sch.AccumulatedFeesInEpoch.Set(value)

	return nil
}

// GetEpochStartHandler returns epoch start header handler as for metachain, but with last finalized headers from main chain, if found.
func (sch *SovereignChainHeader) GetEpochStartHandler() data.EpochStartHandler {
	if sch == nil {
		return nil
	}

	return &sch.EpochStart
}

// GetLastFinalizedCrossChainHeaderHandler returns the last finalized cross chain header data
func (sch *SovereignChainHeader) GetLastFinalizedCrossChainHeaderHandler() data.EpochStartChainDataHandler {
	if sch == nil {
		return nil
	}

	return &sch.EpochStart.LastFinalizedCrossChainHeader
}

// SetLastFinalizedCrossChainHeaderHandler sets the last finalized cross chain header handler
func (sch *SovereignChainHeader) SetLastFinalizedCrossChainHeaderHandler(crossChainData data.EpochStartChainDataHandler) error {
	if sch == nil {
		return data.ErrNilPointerReceiver
	}

	sch.EpochStart.LastFinalizedCrossChainHeader = EpochStartCrossChainData{
		ShardID:    crossChainData.GetShardID(),
		Epoch:      crossChainData.GetEpoch(),
		Round:      crossChainData.GetRound(),
		Nonce:      crossChainData.GetNonce(),
		HeaderHash: crossChainData.GetHeaderHash(),
	}

	return nil
}

// GetShardInfoHandlers returns empty slice
func (sch *SovereignChainHeader) GetShardInfoHandlers() []data.ShardDataHandler {
	if sch == nil {
		return nil
	}

	return make([]data.ShardDataHandler, 0)
}

// SetShardInfoHandlers does nothing
func (sch *SovereignChainHeader) SetShardInfoHandlers(_ []data.ShardDataHandler) error {
	if sch == nil {
		return data.ErrNilPointerReceiver
	}

	return nil
}

// SetHash returns the hash
func (omb *OutGoingMiniBlockHeader) SetHash(hash []byte) error {
	if omb == nil {
		return data.ErrNilPointerReceiver
	}

	omb.Hash = hash
	return nil
}

// SetOutGoingOperationsHash returns the outgoing operations hash
func (omb *OutGoingMiniBlockHeader) SetOutGoingOperationsHash(hash []byte) error {
	if omb == nil {
		return data.ErrNilPointerReceiver
	}

	omb.OutGoingOperationsHash = hash
	return nil
}

// SetLeaderSignatureOutGoingOperations returns the leader signature
func (omb *OutGoingMiniBlockHeader) SetLeaderSignatureOutGoingOperations(sig []byte) error {
	if omb == nil {
		return data.ErrNilPointerReceiver
	}

	omb.LeaderSignatureOutGoingOperations = sig
	return nil
}

// SetAggregatedSignatureOutGoingOperations returns the aggregated signature
func (omb *OutGoingMiniBlockHeader) SetAggregatedSignatureOutGoingOperations(sig []byte) error {
	if omb == nil {
		return data.ErrNilPointerReceiver
	}

	omb.AggregatedSignatureOutGoingOperations = sig
	return nil
}

// IsInterfaceNil checks if the underlying interface is nil
func (omb *OutGoingMiniBlockHeader) IsInterfaceNil() bool {
	return omb == nil
}

// SetShardID sets the epoch start shardID
func (essd *EpochStartCrossChainData) SetShardID(shardID uint32) error {
	if essd == nil {
		return data.ErrNilPointerReceiver
	}

	essd.ShardID = shardID

	return nil
}

// SetEpoch sets the epoch start epoch
func (essd *EpochStartCrossChainData) SetEpoch(epoch uint32) error {
	if essd == nil {
		return data.ErrNilPointerReceiver
	}

	essd.Epoch = epoch

	return nil
}

// SetRound sets the epoch start round
func (essd *EpochStartCrossChainData) SetRound(round uint64) error {
	if essd == nil {
		return data.ErrNilPointerReceiver
	}

	essd.Round = round

	return nil
}

// SetNonce sets the epoch start nonce
func (essd *EpochStartCrossChainData) SetNonce(nonce uint64) error {
	if essd == nil {
		return data.ErrNilPointerReceiver
	}

	essd.Nonce = nonce

	return nil
}

// SetHeaderHash sets the epoch start header hash
func (essd *EpochStartCrossChainData) SetHeaderHash(hash []byte) error {
	if essd == nil {
		return data.ErrNilPointerReceiver
	}

	essd.HeaderHash = hash

	return nil
}

// GetRootHash returns nothing
func (essd *EpochStartCrossChainData) GetRootHash() []byte {
	return nil
}

// GetFirstPendingMetaBlock returns nothing
func (essd *EpochStartCrossChainData) GetFirstPendingMetaBlock() []byte {
	return nil
}

// GetLastFinishedMetaBlock returns nothing
func (essd *EpochStartCrossChainData) GetLastFinishedMetaBlock() []byte {
	return nil
}

// GetPendingMiniBlockHeaderHandlers returns empty slice
func (essd *EpochStartCrossChainData) GetPendingMiniBlockHeaderHandlers() []data.MiniBlockHeaderHandler {
	return make([]data.MiniBlockHeaderHandler, 0)
}

// SetRootHash does nothing
func (essd *EpochStartCrossChainData) SetRootHash([]byte) error {
	return nil
}

// SetFirstPendingMetaBlock does nothing
func (essd *EpochStartCrossChainData) SetFirstPendingMetaBlock([]byte) error {
	return nil
}

// SetLastFinishedMetaBlock does nothing
func (essd *EpochStartCrossChainData) SetLastFinishedMetaBlock([]byte) error {
	return nil
}

// SetPendingMiniBlockHeaders does nothing
func (essd *EpochStartCrossChainData) SetPendingMiniBlockHeaders(_ []data.MiniBlockHeaderHandler) error {
	return nil
}

// GetLastFinalizedHeaderHandlers returns last cross main chain finalized header in a slice w.r.t to the interface
func (m *EpochStartSovereign) GetLastFinalizedHeaderHandlers() []data.EpochStartShardDataHandler {
	if m == nil {
		return nil
	}

	epochStartShardData := make([]data.EpochStartShardDataHandler, 0)
	if m.LastFinalizedCrossChainHeader.ShardID == core.MainChainShardId {
		epochStartShardData = append(epochStartShardData, &m.LastFinalizedCrossChainHeader)
	}

	return epochStartShardData
}

// GetEconomicsHandler returns the economics
func (m *EpochStartSovereign) GetEconomicsHandler() data.EconomicsHandler {
	if m == nil {
		return nil
	}

	return &m.Economics
}

// SetLastFinalizedHeaders sets epoch start data for main chain chain only
func (m *EpochStartSovereign) SetLastFinalizedHeaders(epochStartShardDataHandlers []data.EpochStartShardDataHandler) error {
	if m == nil {
		return data.ErrNilPointerReceiver
	}

	for _, epochStartShardData := range epochStartShardDataHandlers {
		if epochStartShardData.GetShardID() == core.MainChainShardId {
			m.LastFinalizedCrossChainHeader = EpochStartCrossChainData{
				ShardID:    epochStartShardData.GetShardID(),
				Epoch:      epochStartShardData.GetEpoch(),
				Round:      epochStartShardData.GetRound(),
				Nonce:      epochStartShardData.GetNonce(),
				HeaderHash: epochStartShardData.GetHeaderHash(),
			}
		}
	}

	return nil
}

// SetEconomics sets economics
func (m *EpochStartSovereign) SetEconomics(economicsHandler data.EconomicsHandler) error {
	if m == nil {
		return data.ErrNilPointerReceiver
	}

	ec, ok := economicsHandler.(*Economics)
	if !ok {
		return data.ErrInvalidTypeAssertion
	}
	if ec == nil {
		return data.ErrNilPointerDereference
	}

	m.Economics = *ec

	return nil
}
