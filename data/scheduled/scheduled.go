//go:generate protoc -I=. -I=$GOPATH/src -I=$GOPATH/src/github.com/multiversx/protobuf/protobuf  --gogoslick_out=. scheduled.proto
package scheduled

import (
	"github.com/multiversx/mx-chain-core-go/data"
	"github.com/multiversx/mx-chain-core-go/data/block"
	"github.com/multiversx/mx-chain-core-go/data/smartContractResult"
	"github.com/multiversx/mx-chain-core-go/data/transaction"
)

// GetTransactionHandlersMap returns the smart contract results as a map of transaction handlers
func (sscr *ScheduledSCRs) GetTransactionHandlersMap() map[block.Type][]data.TransactionHandler {
	if sscr == nil {
		return nil
	}
	if len(sscr.Scrs) == 0 && len(sscr.InvalidTransactions) == 0 {
		return nil
	}

	result := make(map[block.Type][]data.TransactionHandler)
	var smartContractResults []data.TransactionHandler
	for i := range sscr.Scrs {
		smartContractResults = append(smartContractResults, sscr.Scrs[i])
	}
	result[block.SmartContractResultBlock] = smartContractResults

	var invalidTxs []data.TransactionHandler
	for i := range sscr.InvalidTransactions {
		invalidTxs = append(invalidTxs, sscr.InvalidTransactions[i])
	}
	result[block.InvalidBlock] = invalidTxs

	return result
}

// SetTransactionHandlersMap fills the smart contract results map from the given transaction handlers map
func (sscr *ScheduledSCRs) SetTransactionHandlersMap(txHandlersMap map[block.Type][]data.TransactionHandler) error {
	if sscr == nil {
		return data.ErrNilPointerReceiver
	}
	if txHandlersMap == nil {
		sscr.Scrs = nil
		sscr.InvalidTransactions = nil
		return nil
	}

	var smartContractResults []*smartContractResult.SmartContractResult
	txHandlers := txHandlersMap[block.SmartContractResultBlock]
	for j := range txHandlers {
		scr, ok := txHandlers[j].(*smartContractResult.SmartContractResult)
		if !ok {
			return data.ErrInvalidTypeAssertion
		}
		smartContractResults = append(smartContractResults, scr)
	}

	var invalidTxs []*transaction.Transaction
	txHandlers = txHandlersMap[block.InvalidBlock]
	for j := range txHandlers {
		invalidTx, ok := txHandlers[j].(*transaction.Transaction)
		if !ok {
			return data.ErrInvalidTypeAssertion
		}
		invalidTxs = append(invalidTxs, invalidTx)
	}

	sscr.Scrs = smartContractResults
	sscr.InvalidTransactions = invalidTxs

	return nil
}
