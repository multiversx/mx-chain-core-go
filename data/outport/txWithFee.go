package outport

import (
	"math/big"

	"github.com/multiversx/mx-chain-core-go/data"
)

// TransactionHandlerWithGasAndFee holds a data.TransactionHandler and information about fee and gas used
type TransactionHandlerWithGasAndFee struct {
	data.TransactionHandler
	data.FeeInfoHandler
	ExecutionOrder int
}

// NewTransactionHandlerWithGasAndFee returns a new instance of transactionHandlerWithGasAndFee which matches the interface
func NewTransactionHandlerWithGasAndFee(txHandler data.TransactionHandler, gasUsed uint64, fee *big.Int) data.TransactionHandlerWithGasUsedAndFee {
	return &TransactionHandlerWithGasAndFee{
		TransactionHandler: txHandler,
		FeeInfoHandler: &FeeInfo{
			GasUsed: gasUsed,
			Fee:     fee,
		},
	}
}

// GetTxHandler will return the TransactionHandler
func (t *TransactionHandlerWithGasAndFee) GetTxHandler() data.TransactionHandler {
	return t.TransactionHandler
}

// SetExecutionOrder will set the execution order of the TransactionHandler
func (t *TransactionHandlerWithGasAndFee) SetExecutionOrder(order int) {
	t.ExecutionOrder = order
}

// GetExecutionOrder will return the execution order of the TransactionHandler
func (t *TransactionHandlerWithGasAndFee) GetExecutionOrder() int {
	return t.ExecutionOrder
}

// WrapTxsMap will wrap the provided transactions map in a map fo transactions with fee and gas used
func WrapTxsMap(txs map[string]data.TransactionHandler) map[string]data.TransactionHandlerWithGasUsedAndFee {
	newMap := make(map[string]data.TransactionHandlerWithGasUsedAndFee, len(txs))
	for txHash, tx := range txs {
		newMap[txHash] = NewTransactionHandlerWithGasAndFee(tx, 0, big.NewInt(0))
	}

	return newMap
}
