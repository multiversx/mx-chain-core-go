package outport

import (
	"math/big"

	"github.com/ElrondNetwork/elrond-go-core/data"
)

// FeeInfo holds information about the fee and gas used
type FeeInfo struct {
	GasUsed        uint64
	Fee            *big.Int
	InitialPaidFee *big.Int
}

// TransactionHandlerWithGasAndFee holds a data.TransactionHandler and information about fee and gas used
type TransactionHandlerWithGasAndFee struct {
	data.TransactionHandler
	FeeInfo
}

// NewTransactionHandlerWithGasAndFee returns a new instance of transactionHandlerWithGasAndFee which matches the interface
func NewTransactionHandlerWithGasAndFee(txHandler data.TransactionHandler, gasUsed uint64, fee *big.Int) data.TransactionHandlerWithGasUsedAndFee {
	return &TransactionHandlerWithGasAndFee{
		TransactionHandler: txHandler,
		FeeInfo: FeeInfo{
			GasUsed: gasUsed,
			Fee:     fee,
		},
	}
}

// SetInitialPaidFee will set the initial paid fee
func (t *TransactionHandlerWithGasAndFee) SetInitialPaidFee(fee *big.Int) {
	t.InitialPaidFee = fee
}

// GetInitialPaidFee returns the initial paid fee of the transactions
func (t *TransactionHandlerWithGasAndFee) GetInitialPaidFee() *big.Int {
	return t.InitialPaidFee
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

// WrapTxsMap will wrap the provided transactions map in a map fo transactions with fee and gas used
func WrapTxsMap(txs map[string]data.TransactionHandler) map[string]data.TransactionHandlerWithGasUsedAndFee {
	newMap := make(map[string]data.TransactionHandlerWithGasUsedAndFee, len(txs))
	for txHash, tx := range txs {
		newMap[txHash] = NewTransactionHandlerWithGasAndFee(tx, 0, big.NewInt(0))
	}

	return newMap
}
