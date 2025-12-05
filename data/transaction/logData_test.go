package transaction

import (
	"testing"

	"github.com/multiversx/mx-chain-core-go/marshal"
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
	require.Equal(t, providedLog.GetAddress(), logData.GetLogHandler().GetAddress())
}

func TestLogDataMarshal(t *testing.T) {
	t.Parallel()

	marshaller := &marshal.GogoProtoMarshalizer{}
	providedHash := "hash"
	providedLog := &Log{
		Address: []byte("address"),
	}
	logData := &LogData{
		Log:    providedLog,
		TxHash: providedHash,
	}

	marshalledLogData, err := marshaller.Marshal(logData)
	require.NoError(t, err)

	var logData2 LogData
	err = marshaller.Unmarshal(&logData2, marshalledLogData)
	require.NoError(t, err)
	require.Equal(t, logData.GetTxHash(), logData2.GetTxHash())
}
