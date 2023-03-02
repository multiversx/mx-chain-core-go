package outport

import (
	"time"

	"github.com/multiversx/mx-chain-core-go/data"
	"github.com/multiversx/mx-chain-core-go/data/alteredAccount"
)

// ArgsSaveBlockData will contain all information that are needed to save block data
type ArgsSaveBlockData struct {
	HeaderHash             []byte
	Body                   data.BodyHandler
	Header                 data.HeaderHandler
	SignersIndexes         []uint64
	NotarizedHeadersHashes []string
	HeaderGasConsumption   HeaderGasConsumption
	TransactionsPool       *TransactionPool
	AlteredAccounts        map[string]*alteredAccount.AlteredAccount
	NumberOfShards         uint32
	IsImportDB             bool
	HighestFinalBlockNonce uint64
	HighestFinalBlockHash  []byte
}

// ValidatorRatingInfo is a structure containing validator rating information
type ValidatorRatingInfo struct {
	PublicKey string
	Rating    float32
}

// RoundInfo is a structure containing block signers and shard id
type RoundInfo struct {
	Index            uint64
	SignersIndexes   []uint64
	BlockWasProposed bool
	ShardId          uint32
	Epoch            uint32
	Timestamp        time.Duration
}
