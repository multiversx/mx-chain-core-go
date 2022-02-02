package scheduled_test

import (
	"testing"

	"github.com/ElrondNetwork/elrond-go-core/data"
	"github.com/ElrondNetwork/elrond-go-core/data/block"
	"github.com/ElrondNetwork/elrond-go-core/data/scheduled"
	"github.com/ElrondNetwork/elrond-go-core/data/smartContractResult"
	"github.com/stretchr/testify/require"
)

func TestScheduledSCRs_GetTransactionHandlersMapNilSCRsList(t *testing.T) {
	scheduledSCRs := &scheduled.ScheduledSCRs{
		RootHash: []byte("root hash"),
		Scrs:     createInitializedSCRsMap(20),
	}
	scheduledSCRs.Scrs[int32(block.TxBlock)] = scheduled.SmartContractResults{
		TxHandlers: nil,
	}

	expectedTxHandlersMap := createInitializedTransactionHandlerMap(20)
	expectedTxHandlersMap[block.TxBlock] = nil
	txHandlersMap := scheduledSCRs.GetTransactionHandlersMap()
	require.NotNil(t, txHandlersMap)
	require.Equal(t, expectedTxHandlersMap, txHandlersMap)
}

func TestScheduledSCRs_GetTransactionHandlersMapNilSCRsMap(t *testing.T) {
	scheduledSCRs := &scheduled.ScheduledSCRs{
		RootHash: []byte("root hash"),
		Scrs:     nil,
	}

	txHandlersMap := scheduledSCRs.GetTransactionHandlersMap()
	require.Nil(t, txHandlersMap)
}

func TestScheduledSCRs_GetTransactionHandlersMapOK(t *testing.T) {
	scheduledSCRs := &scheduled.ScheduledSCRs{
		RootHash: []byte("root hash"),
		Scrs:     createInitializedSCRsMap(20),
	}

	expectedTxHandlersMap := createInitializedTransactionHandlerMap(20)
	txHandlersMap := scheduledSCRs.GetTransactionHandlersMap()
	require.NotNil(t, txHandlersMap)
	require.Equal(t, expectedTxHandlersMap, txHandlersMap)
}

func TestScheduledSCRs_SetTransactionHandlersMapNilSCRsMap(t *testing.T) {
	scheduledSCRs := &scheduled.ScheduledSCRs{
		RootHash: []byte("root hash"),
	}
	expectedScheduledSCRs := &scheduled.ScheduledSCRs{
		RootHash: []byte("root hash"),
	}
	err := scheduledSCRs.SetTransactionHandlersMap(nil)
	require.Nil(t, err)
	require.Equal(t, scheduledSCRs, expectedScheduledSCRs)
}

func TestScheduledSCRs_SetTransactionHandlersMapNilSCRsList(t *testing.T) {
	scheduledSCRs := &scheduled.ScheduledSCRs{
		RootHash: []byte("root hash"),
	}
	expectedScheduledSCRs := &scheduled.ScheduledSCRs{
		RootHash: []byte("root hash"),
	}
	expectedScheduledSCRs.Scrs = createInitializedSCRsMap(20)
	expectedScheduledSCRs.Scrs[int32(block.TxBlock)] = scheduled.SmartContractResults{TxHandlers: nil}

	txHandlersMap := createInitializedTransactionHandlerMap(20)
	txHandlersMap[block.TxBlock] = nil

	err := scheduledSCRs.SetTransactionHandlersMap(txHandlersMap)
	require.Nil(t, err)
	require.Equal(t, scheduledSCRs, expectedScheduledSCRs)
}

func TestScheduledSCRs_SetTransactionHandlersMapOK(t *testing.T) {
	scheduledSCRs := &scheduled.ScheduledSCRs{
		RootHash: []byte("root hash"),
		Scrs:     createInitializedSCRsMap(20),
	}

	expectedTxHandlersMap := createInitializedTransactionHandlerMap(20)
	actualTxHandlersMap := scheduledSCRs.GetTransactionHandlersMap()
	require.NotNil(t, actualTxHandlersMap)
	require.Equal(t, actualTxHandlersMap, expectedTxHandlersMap)
}

func createInitializedTransactionHandlerMap(nbTxsPerIndex int) map[block.Type][]data.TransactionHandler {
	result := make(map[block.Type][]data.TransactionHandler)
	result[block.TxBlock] = createInitializedTransactionHandlerArray(nbTxsPerIndex)
	result[block.RewardsBlock] = createInitializedTransactionHandlerArray(nbTxsPerIndex)

	return result
}

func createInitializedSCRsMap(nbSCRsPerIndex int) map[int32]scheduled.SmartContractResults {
	result := make(map[int32]scheduled.SmartContractResults)
	result[int32(block.TxBlock)] = scheduled.SmartContractResults{
		TxHandlers: createInitializedSCRPointerArray(nbSCRsPerIndex),
	}
	result[int32(block.RewardsBlock)] = scheduled.SmartContractResults{
		TxHandlers: createInitializedSCRPointerArray(nbSCRsPerIndex),
	}

	return result
}

func createInitializedTransactionHandlerArray(nbTxs int) []data.TransactionHandler {
	result := make([]data.TransactionHandler, nbTxs)
	for i := 0; i < nbTxs; i++ {
		result[i] = &smartContractResult.SmartContractResult{
			Nonce: uint64(i),
		}
	}
	return result
}

func createInitializedSCRPointerArray(nbSCRs int) []*smartContractResult.SmartContractResult {
	result := make([]*smartContractResult.SmartContractResult, nbSCRs)
	for i := 0; i < nbSCRs; i++ {
		result[i] = &smartContractResult.SmartContractResult{
			Nonce: uint64(i),
		}
	}
	return result
}
