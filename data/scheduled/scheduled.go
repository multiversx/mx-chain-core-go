//go:generate protoc -I=. -I=$GOPATH/src -I=$GOPATH/src/github.com/ElrondNetwork/protobuf/protobuf  --gogoslick_out=. scheduled.proto
package scheduled

import (
	"github.com/ElrondNetwork/elrond-go-core/data"
	"github.com/ElrondNetwork/elrond-go-core/data/block"
	"github.com/ElrondNetwork/elrond-go-core/data/smartContractResult"
	"github.com/ElrondNetwork/elrond-go-core/data/transaction"
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
	sscr.Scrs = smartContractResults

	var invalidTxs []*transaction.Transaction
	txHandlers = txHandlersMap[block.InvalidBlock]
	for j := range txHandlers {
		invalidTx, ok := txHandlers[j].(*transaction.Transaction)
		if !ok {
			return data.ErrInvalidTypeAssertion
		}
		invalidTxs = append(invalidTxs, invalidTx)
	}
	sscr.InvalidTransactions = invalidTxs

	return nil
}
