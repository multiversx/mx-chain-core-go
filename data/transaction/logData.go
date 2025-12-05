package transaction

import "github.com/multiversx/mx-chain-core-go/data"

// LogData defines an extended log data needed for indexing logs and events
type LogData struct {
	*Log
	TxHash string
}

// GetLogHandler returns the log pointer
func (l *LogData) GetLogHandler() data.LogHandler {
	return l.Log
}

// GetTxHash returns the tx hash
func (l *LogData) GetTxHash() string {
	return l.TxHash
}
