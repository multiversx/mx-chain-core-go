package api

// AccountQueryOptions holds options for account queries
type AccountQueryOptions struct {
	OnFinalBlock   bool
	OnStartOfEpoch uint32
}

// BlockQueryOptions holds options for block queries
type BlockQueryOptions struct {
	WithTransactions bool
	WithLogs         bool
}
