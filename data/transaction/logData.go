package transaction

import "github.com/multiversx/mx-chain-core-go/data"

// GetLogHandler returns the log pointer
func (l *LogData) GetLogHandler() data.LogHandler {
	return l.Log
}
