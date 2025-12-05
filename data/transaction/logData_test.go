package transaction

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLogData_GetTxHash(t *testing.T) {
	t.Parallel()

	providedHash := "hash"
	logData := &LogData{
		TxHash: providedHash,
	}
	require.Equal(t, providedHash, logData.GetTxHash())
}
