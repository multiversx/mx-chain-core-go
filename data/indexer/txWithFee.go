package indexer

import (
	"math/big"

	"github.com/ElrondNetwork/elrond-go-core/data"
)

type transactionHandlerWithGasAndFee struct {
	data.TransactionHandler

	gasUsed uint64
	fee     *big.Int
}

// NewTransactionHandlerWithGasAndFee returns a new instance of transactionHandlerWithGasAndFee which matches the interface
func NewTransactionHandlerWithGasAndFee(txHandler data.TransactionHandler, gasUsed uint64, fee *big.Int) data.TransactionHandlerWithGasUsedAndFee {
	return &transactionHandlerWithGasAndFee{
		TransactionHandler: txHandler,
		gasUsed:            gasUsed,
		fee:                fee,
	}
}

// SetGasUsed sets the used gas internally
func (t *transactionHandlerWithGasAndFee) SetGasUsed(gasUsed uint64) {
	t.gasUsed = gasUsed
}

// GetGasUsed returns the used gas of the transaction
func (t *transactionHandlerWithGasAndFee) GetGasUsed() uint64 {
	return t.gasUsed
}

// SetFee sets the fee internally
func (t *transactionHandlerWithGasAndFee) SetFee(fee *big.Int) {
	t.fee = fee
}

// GetFee returns the fee of the transaction
func (t *transactionHandlerWithGasAndFee) GetFee() *big.Int {
	return t.fee
}
