package api

import "github.com/multiversx/mx-chain-core-go/core"

// AccountQueryOptions holds options for account queries
type AccountQueryOptions struct {
	OnFinalBlock   bool
	OnStartOfEpoch core.OptionalUint32
	BlockNonce     core.OptionalUint64
	BlockHash      []byte
	BlockRootHash  []byte
	HintEpoch      core.OptionalUint32
}

// BlockQueryOptions holds options for block queries
type BlockQueryOptions struct {
	WithTransactions bool
	WithLogs         bool
}
