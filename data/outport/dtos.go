package outport

import (
	"time"

	"github.com/ElrondNetwork/elrond-go-core/data"
)

// TokenMetaData is the api metaData struct for tokens
type TokenMetaData struct {
	Nonce      uint64   `json:"nonce"`
	Name       string   `json:"name"`
	Creator    string   `json:"creator"`
	Royalties  uint32   `json:"royalties"`
	Hash       []byte   `json:"hash"`
	URIs       [][]byte `json:"uris"`
	Attributes []byte   `json:"attributes"`
}

// AccountTokenData holds the data needed for indexing a token of an altered account
type AccountTokenData struct {
	Nonce          uint64                      `json:"nonce"`
	Identifier     string                      `json:"identifier"`
	Balance        string                      `json:"balance"`
	Properties     string                      `json:"properties"`
	MetaData       *TokenMetaData              `json:"metadata,omitempty"`
	AdditionalData *AdditionalAccountTokenData `json:"additionalData,omitempty"`
}

// AlteredAccount holds the data needed of an altered account in a block
type AlteredAccount struct {
	Nonce          uint64                 `json:"nonce"`
	Address        string                 `json:"address"`
	Balance        string                 `json:"balance,omitempty"`
	CurrentOwner   string                 `json:"currentOwner,omitempty"`
	Tokens         []*AccountTokenData    `json:"tokens"`
	AdditionalData *AdditionalAccountData `json:"additionalData,omitempty"`
}

// AdditionalAccountData holds the additional data for an altered account
type AdditionalAccountData struct {
	IsSender       bool `json:"isSender,omitempty"`
	BalanceChanged bool `json:"balanceChanged,omitempty"`
}

// AdditionalAccountTokenData holds the additional data for indexing a token of an altered account
type AdditionalAccountTokenData struct {
	IsNFTCreate bool `json:"isNFTCreate,omitempty"`
}

// ArgsSaveBlockData will contain all information that are needed to save block data
type ArgsSaveBlockData struct {
	HeaderHash             []byte
	Body                   data.BodyHandler
	Header                 data.HeaderHandler
	SignersIndexes         []uint64
	NotarizedHeadersHashes []string
	HeaderGasConsumption   HeaderGasConsumption
	TransactionsPool       *Pool
	AlteredAccounts        map[string]*AlteredAccount
	NumberOfShards         uint32
	IsImportDB             bool
}

// HeaderGasConsumption holds the data needed to save the gas consumption of a header
type HeaderGasConsumption struct {
	GasProvided    uint64
	GasRefunded    uint64
	GasPenalized   uint64
	MaxGasPerBlock uint64
}

// Pool will hold all types of transaction
type Pool struct {
	Txs                                        map[string]data.TransactionHandlerWithGasUsedAndFee
	Scrs                                       map[string]data.TransactionHandlerWithGasUsedAndFee
	Rewards                                    map[string]data.TransactionHandlerWithGasUsedAndFee
	Invalid                                    map[string]data.TransactionHandlerWithGasUsedAndFee
	Receipts                                   map[string]data.TransactionHandlerWithGasUsedAndFee
	Logs                                       []*data.LogData
	ScheduledExecutedSCRSHashesPrevBlock       []string
	ScheduledExecutedInvalidTxsHashesPrevBlock []string
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
