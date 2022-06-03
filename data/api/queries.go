package api

// AccountQueryOptions holds options for account queries.
type AccountQueryOptions struct {
	OnFinalBlock bool
}

// BlockQueryOptions holds options for block queries.
type BlockQueryOptions struct {
	WithTransactions bool
	WithLogs         bool
}
