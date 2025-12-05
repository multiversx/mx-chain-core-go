package transaction

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLogData(t *testing.T) {
	t.Parallel()

	providedHash := "hash"
	providedLog := &Log{
		Address: []byte("address"),
	}
	logData := &LogData{
		Log:    providedLog,
		TxHash: providedHash,
	}
	require.Equal(t, providedHash, logData.GetTxHash())
	require.Equal(t, providedLog, logData.GetLogHandler())
	require.Equal(t, providedLog.GetAddress(), logData.GetAddress())
}
