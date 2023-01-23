package versioning

import (
	"testing"

	"github.com/multiversx/mx-chain-core-go/core"
	"github.com/multiversx/mx-chain-core-go/data/transaction"
	"github.com/stretchr/testify/require"
)

func TestTxVersionChecker_IsSignedWithHashOptionsZeroShouldReturnFalse(t *testing.T) {
	t.Parallel()

	minTxVersion := uint32(1)
	tx := &transaction.Transaction{
		Options: 0,
		Version: minTxVersion,
	}
	tvc := NewTxVersionChecker(minTxVersion)

	res := tvc.IsSignedWithHash(tx)
	require.False(t, res)
}

func TestTxVersionChecker_IsSignedWithHash(t *testing.T) {
	t.Parallel()

	minTxVersion := uint32(1)
	tx := &transaction.Transaction{
		Options: MaskSignedWithHash,
		Version: minTxVersion + 1,
	}
	tvc := NewTxVersionChecker(minTxVersion)

	res := tvc.IsSignedWithHash(tx)
	require.True(t, res)
}

func TestTxVersionChecker_CheckTxVersionShouldReturnErrorOptionsNotZero(t *testing.T) {
	minTxVersion := uint32(1)
	tx := &transaction.Transaction{
		Options: MaskSignedWithHash,
		Version: minTxVersion,
	}

	tvc := NewTxVersionChecker(minTxVersion)
	err := tvc.CheckTxVersion(tx)
	require.Equal(t, core.ErrInvalidTransactionVersion, err)
}

func TestTxVersionChecker_CheckTxVersionShould(t *testing.T) {
	minTxVersion := uint32(1)
	tx := &transaction.Transaction{
		Options: 0,
		Version: minTxVersion,
	}

	tvc := NewTxVersionChecker(minTxVersion)
	err := tvc.CheckTxVersion(tx)
	require.Nil(t, err)
}
