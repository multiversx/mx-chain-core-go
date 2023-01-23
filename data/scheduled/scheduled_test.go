package scheduled_test

import (
	"testing"

	"github.com/multiversx/mx-chain-core-go/data"
	"github.com/multiversx/mx-chain-core-go/data/block"
	"github.com/multiversx/mx-chain-core-go/data/scheduled"
	"github.com/multiversx/mx-chain-core-go/data/smartContractResult"
	"github.com/multiversx/mx-chain-core-go/data/transaction"
	"github.com/stretchr/testify/require"
)

func TestScheduledSCRs_GetTransactionHandlersMapNilSCRsList(t *testing.T) {
	nb := 20
	scheduledSCRs := &scheduled.ScheduledSCRs{
		RootHash:            []byte("root hash"),
		Scrs:                nil,
		InvalidTransactions: createInitializedInvalidTxsPointerArray(nb),
	}

	expectedTxHandlersMap := createInitializedTransactionHandlerMap(nb)
	expectedTxHandlersMap[block.SmartContractResultBlock] = nil
	txHandlersMap := scheduledSCRs.GetTransactionHandlersMap()
	require.NotNil(t, txHandlersMap)
	require.Equal(t, expectedTxHandlersMap, txHandlersMap)
}

func TestScheduledSCRs_GetTransactionHandlersMapNilSCRsMap(t *testing.T) {
	scheduledSCRs := &scheduled.ScheduledSCRs{
		RootHash:            []byte("root hash"),
		Scrs:                nil,
		InvalidTransactions: nil,
	}

	txHandlersMap := scheduledSCRs.GetTransactionHandlersMap()
	require.Nil(t, txHandlersMap)
}

func TestScheduledSCRs_GetTransactionHandlersMapNilInvalidTransactions(t *testing.T) {
	nb := 20
	scheduledSCRs := &scheduled.ScheduledSCRs{
		RootHash:            []byte("root hash"),
		Scrs:                createInitializedSCRPointerArray(nb),
		InvalidTransactions: nil,
	}

	expectedTxHandlersMap := createInitializedTransactionHandlerMap(nb)
	expectedTxHandlersMap[block.InvalidBlock] = nil
	txHandlersMap := scheduledSCRs.GetTransactionHandlersMap()
	require.NotNil(t, txHandlersMap)
	require.Equal(t, expectedTxHandlersMap, txHandlersMap)
}

func TestScheduledSCRs_GetTransactionHandlersMapOK(t *testing.T) {
	nb := 20
	scheduledSCRs := &scheduled.ScheduledSCRs{
		RootHash:            []byte("root hash"),
		Scrs:                createInitializedSCRPointerArray(nb),
		InvalidTransactions: createInitializedInvalidTxsPointerArray(nb),
	}

	expectedTxHandlersMap := createInitializedTransactionHandlerMap(nb)
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
	nb := 20
	scheduledSCRs := &scheduled.ScheduledSCRs{
		RootHash:            []byte("root hash"),
		Scrs:                createInitializedSCRPointerArray(nb),
		InvalidTransactions: createInitializedInvalidTxsPointerArray(nb),
	}
	expectedScheduledSCRs := &scheduled.ScheduledSCRs{
		RootHash:            []byte("root hash"),
		InvalidTransactions: createInitializedInvalidTxsPointerArray(nb),
	}

	txHandlersMap := createInitializedTransactionHandlerMap(nb)
	txHandlersMap[block.SmartContractResultBlock] = nil

	err := scheduledSCRs.SetTransactionHandlersMap(txHandlersMap)
	require.Nil(t, err)
	require.Equal(t, scheduledSCRs, expectedScheduledSCRs)
}

func TestScheduledSCRs_SetTransactionHandlersMapOK(t *testing.T) {
	nbInitial := 20
	scheduledSCRs := &scheduled.ScheduledSCRs{
		RootHash:            []byte("root hash"),
		Scrs:                createInitializedSCRPointerArray(nbInitial),
		InvalidTransactions: createInitializedInvalidTxsPointerArray(nbInitial),
	}

	nbFinal := 30
	expectedTxHandlersMap := createInitializedTransactionHandlerMap(nbFinal)

	err := scheduledSCRs.SetTransactionHandlersMap(expectedTxHandlersMap)
	require.Nil(t, err)
	actualTxHandlersMap := scheduledSCRs.GetTransactionHandlersMap()

	require.NotNil(t, actualTxHandlersMap)
	require.Equal(t, actualTxHandlersMap, expectedTxHandlersMap)
}

func TestScheduledSCRs_SetTransactionHandlersMapInvalidTypeAssertionInvalidTxsShouldNotSetScrs(t *testing.T) {
	nbInitial := 20
	initialScrs := createInitializedSCRPointerArray(nbInitial)
	initialInvalid := createInitializedInvalidTxsPointerArray(nbInitial)

	scheduledSCRs := &scheduled.ScheduledSCRs{
		RootHash:            []byte("root hash"),
		Scrs:                initialScrs,
		InvalidTransactions: initialInvalid,
	}

	nbFinal := 2
	expectedTxHandlersMap := createInitializedTransactionHandlerMap(nbFinal)
	expectedTxHandlersMap[block.InvalidBlock] = append(expectedTxHandlersMap[block.InvalidBlock], nil)

	err := scheduledSCRs.SetTransactionHandlersMap(expectedTxHandlersMap)
	require.NotNil(t, err)
	require.Equal(t, initialScrs, scheduledSCRs.Scrs)
	require.Equal(t, initialInvalid, scheduledSCRs.InvalidTransactions)
}

func TestScheduledSCRs_SetTransactionHandlersMapInvalidTypeAssertionScrsShouldNotSetInvalid(t *testing.T) {
	nbInitial := 20
	initialScrs := createInitializedSCRPointerArray(nbInitial)
	initialInvalid := createInitializedInvalidTxsPointerArray(nbInitial)

	scheduledSCRs := &scheduled.ScheduledSCRs{
		RootHash:            []byte("root hash"),
		Scrs:                initialScrs,
		InvalidTransactions: initialInvalid,
	}

	nbFinal := 2
	expectedTxHandlersMap := createInitializedTransactionHandlerMap(nbFinal)
	expectedTxHandlersMap[block.SmartContractResultBlock] = append(expectedTxHandlersMap[block.SmartContractResultBlock], nil)

	err := scheduledSCRs.SetTransactionHandlersMap(expectedTxHandlersMap)
	require.NotNil(t, err)
	require.Equal(t, initialScrs, scheduledSCRs.Scrs)
	require.Equal(t, initialInvalid, scheduledSCRs.InvalidTransactions)
}

func createInitializedTransactionHandlerMap(nbTxsPerIndex int) map[block.Type][]data.TransactionHandler {
	result := make(map[block.Type][]data.TransactionHandler)
	result[block.InvalidBlock] = createInitializedInvalidTxsAsTransactionHandlerArray(nbTxsPerIndex)
	result[block.SmartContractResultBlock] = createInitializedSCRsAsTransactionHandlerArray(nbTxsPerIndex)

	return result
}

func createInitializedSCRsAsTransactionHandlerArray(nbTxs int) []data.TransactionHandler {
	result := make([]data.TransactionHandler, nbTxs)
	scrs := createInitializedSCRPointerArray(nbTxs)
	for i := range scrs {
		result[i] = scrs[i]
	}

	return result
}

func createInitializedInvalidTxsAsTransactionHandlerArray(nbTxs int) []data.TransactionHandler {
	result := make([]data.TransactionHandler, nbTxs)
	scrs := createInitializedInvalidTxsPointerArray(nbTxs)
	for i := range scrs {
		result[i] = scrs[i]
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

func createInitializedInvalidTxsPointerArray(nbInvalidTxs int) []*transaction.Transaction {
	result := make([]*transaction.Transaction, nbInvalidTxs)
	for i := 0; i < nbInvalidTxs; i++ {
		result[i] = &transaction.Transaction{
			Nonce: uint64(i),
		}
	}
	return result
}
