package outport

import (
	"math/big"

	"github.com/ElrondNetwork/elrond-go-core/data"
)

// TransactionHandlerWithGasAndFee hold a data.TransactionHandler and information about fee and gas used0
type TransactionHandlerWithGasAndFee struct {
	data.TransactionHandler

	GasUsed uint64
	Fee     *big.Int
}

// NewTransactionHandlerWithGasAndFee returns a new instance of transactionHandlerWithGasAndFee which matches the interface
func NewTransactionHandlerWithGasAndFee(txHandler data.TransactionHandler, gasUsed uint64, fee *big.Int) data.TransactionHandlerWithGasUsedAndFee {
	return &TransactionHandlerWithGasAndFee{
		TransactionHandler: txHandler,
		GasUsed:            gasUsed,
		Fee:                fee,
	}
}

// SetGasUsed sets the used gas internally
func (t *TransactionHandlerWithGasAndFee) SetGasUsed(gasUsed uint64) {
	t.GasUsed = gasUsed
}

// GetGasUsed returns the used gas of the transaction
func (t *TransactionHandlerWithGasAndFee) GetGasUsed() uint64 {
	return t.GasUsed
}

// SetFee sets the fee internally
func (t *TransactionHandlerWithGasAndFee) SetFee(fee *big.Int) {
	t.Fee = fee
}

// GetFee returns the fee of the transaction
func (t *TransactionHandlerWithGasAndFee) GetFee() *big.Int {
	return t.Fee
}

// GetTxHandler will return the TransactionHandler
func (t *TransactionHandlerWithGasAndFee) GetTxHandler() data.TransactionHandler {
	return t.TransactionHandler
}
