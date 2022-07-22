package api

import (
	"math/big"
	"time"

	"github.com/ElrondNetwork/elrond-go-core/data/transaction"
)

// Block represents the structure for block that is returned by api routes
type Block struct {
	Nonce                  uint64                 `json:"nonce"`
	Round                  uint64                 `json:"round"`
	Epoch                  uint32                 `json:"epoch"`
	Shard                  uint32                 `json:"shard"`
	NumTxs                 uint32                 `json:"numTxs"`
	Hash                   string                 `json:"hash"`
	PrevBlockHash          string                 `json:"prevBlockHash"`
	StateRootHash          string                 `json:"stateRootHash"`
	AccumulatedFees        string                 `json:"accumulatedFees,omitempty"`
	DeveloperFees          string                 `json:"developerFees,omitempty"`
	AccumulatedFeesInEpoch string                 `json:"accumulatedFeesInEpoch,omitempty"`
	DeveloperFeesInEpoch   string                 `json:"developerFeesInEpoch,omitempty"`
	Status                 string                 `json:"status,omitempty"`
	Timestamp              time.Duration          `json:"timestamp,omitempty"`
	NotarizedBlocks        []*NotarizedBlock      `json:"notarizedBlocks,omitempty"`
	MiniBlocks             []*MiniBlock           `json:"miniBlocks,omitempty"`
	EpochStartInfo         *EpochStartInfo        `json:"epochStartInfo,omitempty"`
	EpochStartShardsData   []*EpochStartShardData `json:"epochStartShardsData,omitempty"`
	ScheduledData          *ScheduledData         `json:"scheduledData,omitempty"`
}

// ScheduledData is a structure that hold information about scheduled events
type ScheduledData struct {
	ScheduledRootHash        string `json:"rootHash,omitempty"`
	ScheduledAccumulatedFees string `json:"accumulatedFees,omitempty"`
	ScheduledDeveloperFees   string `json:"developerFees,omitempty"`
	ScheduledGasProvided     uint64 `json:"gasProvided,omitempty"`
	ScheduledGasPenalized    uint64 `json:"penalized,omitempty"`
	ScheduledGasRefunded     uint64 `json:"gasRefunded,omitempty"`
}

// EpochStartInfo is a structure that holds information about epoch start meta block
type EpochStartInfo struct {
	TotalSupply                      string `json:"totalSupply"`
	TotalToDistribute                string `json:"totalToDistribute"`
	TotalNewlyMinted                 string `json:"totalNewlyMinted"`
	RewardsPerBlock                  string `json:"rewardsPerBlock"`
	RewardsForProtocolSustainability string `json:"rewardsForProtocolSustainability"`
	NodePrice                        string `json:"nodePrice"`
	PrevEpochStartRound              uint64 `json:"prevEpochStartRound"`
	PrevEpochStartHash               string `json:"prevEpochStartHash"`
}

// NotarizedBlock represents a notarized block
type NotarizedBlock struct {
	Hash  string `json:"hash"`
	Nonce uint64 `json:"nonce"`
	Round uint64 `json:"round"`
	Shard uint32 `json:"shard"`
}

// EpochStartShardData is a structure that holds data about the epoch start shard data
type EpochStartShardData struct {
	ShardID                 uint32       `json:"shard"`
	Epoch                   uint32       `json:"epoch"`
	Round                   uint64       `json:"round,omitempty"`
	Nonce                   uint64       `json:"nonce,omitempty"`
	HeaderHash              string       `json:"headerHash,omitempty"`
	RootHash                string       `json:"rootHash,omitempty"`
	ScheduledRootHash       string       `json:"scheduledRootHash,omitempty"`
	FirstPendingMetaBlock   string       `json:"firstPendingMetaBlock,omitempty"`
	LastFinishedMetaBlock   string       `json:"lastFinishedMetaBlock,omitempty"`
	PendingMiniBlockHeaders []*MiniBlock `json:"pendingMiniBlockHeaders,omitempty"`
}

// MiniBlock represents the structure for a miniblock
type MiniBlock struct {
	Hash                    string                              `json:"hash"`
	Type                    string                              `json:"type"`
	ProcessingType          string                              `json:"processingType,omitempty"`
	ConstructionState       string                              `json:"constructionState,omitempty"`
	IsFromReceiptsStorage   bool                                `json:"isFromReceiptsStorage,omitempty"`
	SourceShard             uint32                              `json:"sourceShard"`
	DestinationShard        uint32                              `json:"destinationShard"`
	Transactions            []*transaction.ApiTransactionResult `json:"transactions,omitempty"`
	Receipts                []*transaction.ApiReceipt           `json:"receipts,omitempty"`
	IndexOfFirstTxProcessed int32                               `json:"indexOfFirstTxProcessed"`
	IndexOfLastTxProcessed  int32                               `json:"indexOfLastTxProcessed"`
}

// StakeValues is the structure that contains the total staked value and the total top up value
type StakeValues struct {
	BaseStaked *big.Int
	TopUp      *big.Int
}

// DirectStakedValue holds the total staked value for an address
type DirectStakedValue struct {
	Address    string `json:"address"`
	BaseStaked string `json:"baseStaked"`
	TopUp      string `json:"topUp"`
	Total      string `json:"total"`
}

// DelegatedValue holds the value and the delegation system SC address
type DelegatedValue struct {
	DelegationScAddress string `json:"delegationScAddress"`
	Value               string `json:"value"`
}

// Delegator holds the delegator address and the slice of delegated values
type Delegator struct {
	DelegatorAddress string            `json:"delegatorAddress"`
	DelegatedTo      []*DelegatedValue `json:"delegatedTo"`
	Total            string            `json:"total"`
	TotalAsBigInt    *big.Int          `json:"-"`
}
