//go:generate protoc -I=. -I=$GOPATH/src -I=$GOPATH/src/github.com/ElrondNetwork/protobuf/protobuf  --gogoslick_out=. scheduled.proto
package scheduled

import (
	"github.com/ElrondNetwork/elrond-go-core/data"
	"github.com/ElrondNetwork/elrond-go-core/data/block"
	"github.com/ElrondNetwork/elrond-go-core/data/smartContractResult"
)

// GetTransactionHandlersMap returns the smart contract results as a map of transaction handlers
func (sscr *ScheduledSCRs) GetTransactionHandlersMap() map[block.Type][]data.TransactionHandler {
	if sscr == nil {
		return nil
	}
	if len(sscr.Scrs) == 0 {
		return nil
	}

	result := make(map[block.Type][]data.TransactionHandler)
	for i, scrs := range sscr.Scrs {
		if len(scrs.TxHandlers) == 0 {
			result[block.Type(i)] = nil
			continue
		}
		transactionHandlers := make([]data.TransactionHandler, len(scrs.TxHandlers))
		for j := range scrs.TxHandlers {
			transactionHandlers[j] = scrs.TxHandlers[j]
		}
		result[block.Type(i)] = transactionHandlers
	}

	return result
}

// SetTransactionHandlersMap fills the smart contract results map from the given transaction handlers map
func (sscr *ScheduledSCRs) SetTransactionHandlersMap(txHandlersMap map[block.Type][]data.TransactionHandler) error {
	if sscr == nil {
		return data.ErrNilPointerReceiver
	}
	if txHandlersMap == nil {
		sscr.Scrs = nil
		return nil
	}

	sscr.Scrs = make(map[int32]SmartContractResults)
	for i, txHandlers := range txHandlersMap {
		if len(txHandlers) == 0 {
			sscr.Scrs[int32(i)] = SmartContractResults{}
			continue
		}
		scrs := make([]*smartContractResult.SmartContractResult, len(txHandlersMap[i]))
		for j := range txHandlers {
			scr, ok := txHandlers[j].(*smartContractResult.SmartContractResult)
			if !ok {
				return data.ErrInvalidTypeAssertion
			}
			scrs[j] = scr
		}
		sscr.Scrs[int32(i)] = SmartContractResults{
			TxHandlers: scrs,
		}
	}

	return nil
}
