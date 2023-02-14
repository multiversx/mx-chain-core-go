package outport

import (
	"math/big"
	"testing"

	"github.com/multiversx/mx-chain-core-go/data/transaction"
	"github.com/stretchr/testify/require"
)

func TestNewTransactionWithFee(t *testing.T) {
	t.Parallel()

	txWithFee := NewTransactionHandlerWithGasAndFee(&transaction.Transaction{
		Nonce: 1,
	}, 100, big.NewInt(1000))
	txWithFee.SetInitialPaidFee(big.NewInt(2000))

	require.Equal(t, uint64(100), txWithFee.GetGasUsed())
	require.Equal(t, big.NewInt(1000), txWithFee.GetFee())
	require.Equal(t, big.NewInt(2000), txWithFee.GetInitialPaidFee())
}
