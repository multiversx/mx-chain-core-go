package data

// LogData holds the data needed for indexing logs and events
type LogData struct {
	LogHandler
	TxHash string
}

// LogDataHandler is an interface implemented by LogData structure
type LogDataHandler interface {
	LogHandler
	GetTxHash() string
}

// KeyValuePair is a tuple of (key, value)
type KeyValuePair struct {
	Key   []byte
	Value []byte
}
