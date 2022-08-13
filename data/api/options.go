package api

import "github.com/ElrondNetwork/elrond-go-core/core"

// AccountQueryOptions holds options for account queries
type AccountQueryOptions struct {
	OnFinalBlock   bool
	OnStartOfEpoch core.OptionalUint32
	BlockNonce     core.OptionalUint64
	BlockHash      []byte
	BlockRootHash  []byte
}

// BlockQueryOptions holds options for block queries
type BlockQueryOptions struct {
	WithTransactions bool
	WithLogs         bool
}
