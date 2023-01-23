package data

import (
	"math/big"

	"github.com/multiversx/mx-chain-core-go/data/headerVersionData"
)

// TriggerRegistryHandler defines getters and setters for the trigger registry
type TriggerRegistryHandler interface {
	GetIsEpochStart() bool
	GetNewEpochHeaderReceived() bool
	GetEpoch() uint32
	GetMetaEpoch() uint32
	GetCurrentRoundIndex() int64
	GetEpochStartRound() uint64
	GetEpochFinalityAttestingRound() uint64
	GetEpochMetaBlockHash() []byte
	GetEpochStartHeaderHandler() HeaderHandler

	SetIsEpochStart(isEpochStart bool) error
	SetNewEpochHeaderReceived(newEpochHeaderReceived bool) error
	SetEpoch(epoch uint32) error
	SetMetaEpoch(metaEpoch uint32) error
	SetCurrentRoundIndex(roundIndex int64) error
	SetEpochStartRound(startRound uint64) error
	SetEpochFinalityAttestingRound(finalityAttestingRound uint64) error
	SetEpochMetaBlockHash(epochMetaBlockHash []byte) error
	SetEpochStartHeaderHandler(epochStartHeaderHandler HeaderHandler) error
}

// HeaderHandler defines getters and setters for header data holder
type HeaderHandler interface {
	GetShardID() uint32
	GetNonce() uint64
	GetEpoch() uint32
	GetRound() uint64
	GetRootHash() []byte
	GetPrevHash() []byte
	GetPrevRandSeed() []byte
	GetRandSeed() []byte
	GetPubKeysBitmap() []byte
	GetSignature() []byte
	GetLeaderSignature() []byte
	GetChainID() []byte
	GetSoftwareVersion() []byte
	GetTimeStamp() uint64
	GetTxCount() uint32
	GetReceiptsHash() []byte
	GetAccumulatedFees() *big.Int
	GetDeveloperFees() *big.Int
	GetReserved() []byte
	GetMiniBlockHeadersWithDst(destId uint32) map[string]uint32
	GetOrderedCrossMiniblocksWithDst(destId uint32) []*MiniBlockInfo
	GetMiniBlockHeadersHashes() [][]byte
	GetMiniBlockHeaderHandlers() []MiniBlockHeaderHandler
	HasScheduledSupport() bool
	GetAdditionalData() headerVersionData.HeaderAdditionalData
	HasScheduledMiniBlocks() bool

	SetAccumulatedFees(value *big.Int) error
	SetDeveloperFees(value *big.Int) error
	SetShardID(shId uint32) error
	SetNonce(n uint64) error
	SetEpoch(e uint32) error
	SetRound(r uint64) error
	SetTimeStamp(ts uint64) error
	SetRootHash(rHash []byte) error
	SetPrevHash(pvHash []byte) error
	SetPrevRandSeed(pvRandSeed []byte) error
	SetRandSeed(randSeed []byte) error
	SetPubKeysBitmap(pkbm []byte) error
	SetSignature(sg []byte) error
	SetLeaderSignature(sg []byte) error
	SetChainID(chainID []byte) error
	SetSoftwareVersion(version []byte) error
	SetTxCount(txCount uint32) error
	SetMiniBlockHeaderHandlers(mbHeaderHandlers []MiniBlockHeaderHandler) error
	SetReceiptsHash(hash []byte) error
	SetScheduledRootHash(rootHash []byte) error
	ValidateHeaderVersion() error
	SetAdditionalData(headerVersionData headerVersionData.HeaderAdditionalData) error
	IsStartOfEpochBlock() bool
	ShallowClone() HeaderHandler
	IsInterfaceNil() bool
}

// ShardHeaderHandler defines getters and setters for the shard block header
type ShardHeaderHandler interface {
	HeaderHandler
	GetMetaBlockHashes() [][]byte
	GetEpochStartMetaHash() []byte
	SetEpochStartMetaHash(hash []byte) error
	GetBlockBodyTypeInt32() int32
	SetMetaBlockHashes(hashes [][]byte) error
	MapMiniBlockHashesToShards() map[string]uint32
}

// MetaHeaderHandler defines getters and setters for the meta block header
type MetaHeaderHandler interface {
	HeaderHandler
	GetValidatorStatsRootHash() []byte
	GetEpochStartHandler() EpochStartHandler
	GetDevFeesInEpoch() *big.Int
	GetShardInfoHandlers() []ShardDataHandler
	SetValidatorStatsRootHash(rHash []byte) error
	SetDevFeesInEpoch(value *big.Int) error
	SetShardInfoHandlers(shardInfo []ShardDataHandler) error
	SetAccumulatedFeesInEpoch(value *big.Int) error
}

// MiniBlockHeaderHandler defines setters and getters for miniBlock headers
type MiniBlockHeaderHandler interface {
	GetHash() []byte
	GetSenderShardID() uint32
	GetReceiverShardID() uint32
	GetTxCount() uint32
	GetTypeInt32() int32
	GetReserved() []byte
	GetProcessingType() int32
	GetConstructionState() int32
	IsFinal() bool
	GetIndexOfFirstTxProcessed() int32
	GetIndexOfLastTxProcessed() int32

	SetHash(hash []byte) error
	SetSenderShardID(shardID uint32) error
	SetReceiverShardID(shardID uint32) error
	SetTxCount(count uint32) error
	SetTypeInt32(t int32) error
	SetReserved(reserved []byte) error
	SetProcessingType(procType int32) error
	SetConstructionState(state int32) error
	SetIndexOfLastTxProcessed(indexOfLastTxProcessed int32) error
	SetIndexOfFirstTxProcessed(indexOfFirstTxProcessed int32) error
	ShallowClone() MiniBlockHeaderHandler
}

// PeerChangeHandler defines setters and getters for PeerChange
type PeerChangeHandler interface {
	GetPubKey() []byte
	GetShardIdDest() uint32

	SetPubKey(pubKey []byte) error
	SetShardIdDest(shardID uint32) error
}

// ShardDataHandler defines setters and getters for ShardDataHandler
type ShardDataHandler interface {
	GetHeaderHash() []byte
	GetShardMiniBlockHeaderHandlers() []MiniBlockHeaderHandler
	GetPrevRandSeed() []byte
	GetPubKeysBitmap() []byte
	GetSignature() []byte
	GetRound() uint64
	GetPrevHash() []byte
	GetNonce() uint64
	GetAccumulatedFees() *big.Int
	GetDeveloperFees() *big.Int
	GetNumPendingMiniBlocks() uint32
	GetLastIncludedMetaNonce() uint64
	GetShardID() uint32
	GetTxCount() uint32

	SetHeaderHash(hash []byte) error
	SetShardMiniBlockHeaderHandlers(mbHeaderHandlers []MiniBlockHeaderHandler) error
	SetPrevRandSeed(prevRandSeed []byte) error
	SetPubKeysBitmap(pubKeysBitmap []byte) error
	SetSignature(signature []byte) error
	SetRound(round uint64) error
	SetPrevHash(prevHash []byte) error
	SetNonce(nonce uint64) error
	SetAccumulatedFees(fees *big.Int) error
	SetDeveloperFees(fees *big.Int) error
	SetNumPendingMiniBlocks(num uint32) error
	SetLastIncludedMetaNonce(nonce uint64) error
	SetShardID(shardID uint32) error
	SetTxCount(txCount uint32) error

	ShallowClone() ShardDataHandler
}

// EpochStartShardDataHandler defines setters and getters for EpochStartShardData
type EpochStartShardDataHandler interface {
	GetShardID() uint32
	GetEpoch() uint32
	GetRound() uint64
	GetNonce() uint64
	GetHeaderHash() []byte
	GetRootHash() []byte
	GetFirstPendingMetaBlock() []byte
	GetLastFinishedMetaBlock() []byte
	GetPendingMiniBlockHeaderHandlers() []MiniBlockHeaderHandler

	SetShardID(uint32) error
	SetEpoch(uint32) error
	SetRound(uint64) error
	SetNonce(uint64) error
	SetHeaderHash([]byte) error
	SetRootHash([]byte) error
	SetFirstPendingMetaBlock([]byte) error
	SetLastFinishedMetaBlock([]byte) error
	SetPendingMiniBlockHeaders([]MiniBlockHeaderHandler) error
}

// EconomicsHandler defines setters and getters for Economics
type EconomicsHandler interface {
	GetTotalSupply() *big.Int
	GetTotalToDistribute() *big.Int
	GetTotalNewlyMinted() *big.Int
	GetRewardsPerBlock() *big.Int
	GetRewardsForProtocolSustainability() *big.Int
	GetNodePrice() *big.Int
	GetPrevEpochStartRound() uint64
	GetPrevEpochStartHash() []byte

	SetTotalSupply(totalSupply *big.Int) error
	SetTotalToDistribute(totalToDistribute *big.Int) error
	SetTotalNewlyMinted(totalNewlyMinted *big.Int) error
	SetRewardsPerBlock(rewardsPerBlock *big.Int) error
	SetRewardsForProtocolSustainability(rewardsForProtocolSustainability *big.Int) error
	SetNodePrice(nodePrice *big.Int) error
	SetPrevEpochStartRound(prevEpochStartRound uint64) error
	SetPrevEpochStartHash(prevEpochStartHash []byte) error
}

// EpochStartHandler defines setters and getters for EpochStart
type EpochStartHandler interface {
	GetLastFinalizedHeaderHandlers() []EpochStartShardDataHandler
	GetEconomicsHandler() EconomicsHandler

	SetLastFinalizedHeaders(epochStartShardDataHandlers []EpochStartShardDataHandler) error
	SetEconomics(economicsHandler EconomicsHandler) error
}

// BodyHandler interface for a block body
type BodyHandler interface {
	Clone() BodyHandler
	// IntegrityAndValidity checks the integrity and validity of the block
	IntegrityAndValidity() error
	// IsInterfaceNil returns true if there is no value under the interface
	IsInterfaceNil() bool
}

// ChainHandler is the interface defining the functionality a blockchain should implement
type ChainHandler interface {
	GetGenesisHeader() HeaderHandler
	SetGenesisHeader(gb HeaderHandler) error
	GetGenesisHeaderHash() []byte
	SetGenesisHeaderHash(hash []byte)
	GetCurrentBlockHeader() HeaderHandler
	SetCurrentBlockHeaderAndRootHash(bh HeaderHandler, rootHash []byte) error
	GetCurrentBlockHeaderHash() []byte
	SetCurrentBlockHeaderHash(hash []byte)
	GetCurrentBlockRootHash() []byte
	SetFinalBlockInfo(nonce uint64, blockHash []byte, rootHash []byte)
	GetFinalBlockInfo() (nonce uint64, blockHash []byte, rootHash []byte)
	IsInterfaceNil() bool
}

// TransactionHandler defines the type of executable transaction
type TransactionHandler interface {
	IsInterfaceNil() bool

	GetValue() *big.Int
	GetNonce() uint64
	GetData() []byte
	GetRcvAddr() []byte
	GetRcvUserName() []byte
	GetSndAddr() []byte
	GetGasLimit() uint64
	GetGasPrice() uint64

	SetValue(*big.Int)
	SetData([]byte)
	SetRcvAddr([]byte)
	SetSndAddr([]byte)
	Size() int

	CheckIntegrity() error
}

// TransactionHandlerWithGasUsedAndFee extends TransactionHandler by also including used gas and fee
type TransactionHandlerWithGasUsedAndFee interface {
	TransactionHandler

	SetInitialPaidFee(fee *big.Int)
	SetGasUsed(gasUsed uint64)
	SetFee(fee *big.Int)
	GetInitialPaidFee() *big.Int
	GetGasUsed() uint64
	GetFee() *big.Int
	GetTxHandler() TransactionHandler
	SetExecutionOrder(order int)
	GetExecutionOrder() int
}

// LogHandler defines the type for a log resulted from executing a transaction or smart contract call
type LogHandler interface {
	// GetAddress returns the address of the sc that was originally called by the user
	GetAddress() []byte
	// GetLogEvents returns the events from a transaction log entry
	GetLogEvents() []EventHandler

	IsInterfaceNil() bool
}

// EventHandler defines the type for an event resulted from a smart contract call contained in a log
type EventHandler interface {
	// GetAddress returns the address of the contract that generated this event
	//  - in sc calling another sc situation this will differ from the
	//    LogHandler's GetAddress, whereas in the single sc situation
	//    they will be the same
	GetAddress() []byte
	// GetIdentifier returns identifier of the event, that together with the ABI can
	//   be used to understand the type of the event by other applications
	GetIdentifier() []byte
	// GetTopics returns the data that can be indexed so that it would be searchable
	//  by other applications
	GetTopics() [][]byte
	// GetData returns the rest of the event data, which will not be indexed, so storing
	//  information here should be cheaper
	GetData() []byte

	IsInterfaceNil() bool
}

// ValidatorInfoHandler is used to store multiple validatorInfo properties
type ValidatorInfoHandler interface {
	GetPublicKey() []byte
	GetShardId() uint32
	GetList() string
	GetIndex() uint32
	GetTempRating() uint32
	GetRating() uint32
	String() string
	IsInterfaceNil() bool
}

// ShardValidatorInfoHandler is used to store multiple validatorInfo properties required in shards
type ShardValidatorInfoHandler interface {
	GetPublicKey() []byte
	GetTempRating() uint32
	String() string
	IsInterfaceNil() bool
}

// GoRoutineThrottler can monitor the number of the currently running go routines
type GoRoutineThrottler interface {
	CanProcess() bool
	StartProcessing()
	EndProcessing()
	IsInterfaceNil() bool
}

// MiniBlockInfo holds information about a cross miniblock referenced in a received block
type MiniBlockInfo struct {
	Hash          []byte
	SenderShardID uint32
	Round         uint64
}

// SyncStatisticsHandler defines the methods for a component able to store the sync statistics for a trie
type SyncStatisticsHandler interface {
	Reset()
	AddNumProcessed(value int)
	AddNumLarge(value int)
	SetNumMissing(rootHash []byte, value int)
	NumProcessed() int
	NumLarge() int
	NumMissing() int
	IsInterfaceNil() bool
}

// TransactionWithFeeHandler represents a transaction structure that has economics variables defined
type TransactionWithFeeHandler interface {
	GetGasLimit() uint64
	GetGasPrice() uint64
	GetData() []byte
	GetRcvAddr() []byte
	GetValue() *big.Int
}

// UserAccountHandler models a user account
type UserAccountHandler interface {
	RetrieveValue(key []byte) ([]byte, uint32, error)
	GetBalance() *big.Int
	GetNonce() uint64
	AddressBytes() []byte
	IsInterfaceNil() bool
}
