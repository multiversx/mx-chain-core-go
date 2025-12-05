package transaction

// LogData defines an extended log data needed for indexing logs and events
type LogData struct {
	*Log
	TxHash string
}

// GetTxHash returns the tx hash
func (l *LogData) GetTxHash() string {
	return l.TxHash
}
