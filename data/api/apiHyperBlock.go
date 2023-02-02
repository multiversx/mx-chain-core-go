package api

import (
	"time"

	"github.com/multiversx/mx-chain-core-go/data/transaction"
)

// Hyperblock contains all fully executed (both in source and in destination shards) transactions notarized in a given metablock
type Hyperblock struct {
	Hash                   string                              `json:"hash"`
	PrevBlockHash          string                              `json:"prevBlockHash"`
	StateRootHash          string                              `json:"stateRootHash"`
	Nonce                  uint64                              `json:"nonce"`
	Round                  uint64                              `json:"round"`
	Epoch                  uint32                              `json:"epoch"`
	NumTxs                 uint32                              `json:"numTxs"`
	AccumulatedFees        string                              `json:"accumulatedFees,omitempty"`
	DeveloperFees          string                              `json:"developerFees,omitempty"`
	AccumulatedFeesInEpoch string                              `json:"accumulatedFeesInEpoch,omitempty"`
	DeveloperFeesInEpoch   string                              `json:"developerFeesInEpoch,omitempty"`
	Timestamp              time.Duration                       `json:"timestamp,omitempty"`
	EpochStartInfo         *EpochStartInfo                     `json:"epochStartInfo,omitempty"`
	EpochStartShardsData   []*EpochStartShardData              `json:"epochStartShardsData,omitempty"`
	ShardBlocks            []*NotarizedBlock                   `json:"shardBlocks"`
	Transactions           []*transaction.ApiTransactionResult `json:"transactions"`
	Status                 string                              `json:"status,omitempty"`
}
