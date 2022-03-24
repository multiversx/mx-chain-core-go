package block_test

import (
	"testing"

	"github.com/ElrondNetwork/elrond-go-core/data/block"
	"github.com/stretchr/testify/require"
)

func TestMiniBlockHeader_GetProcessingType(t *testing.T) {
	mbh := &block.MiniBlockHeader{}

	processingType := mbh.GetProcessingType()
	require.Equal(t, int32(block.Normal), processingType)

	mbhReserved := block.MiniBlockHeaderReserved{ExecutionType: block.Normal}
	mbh.Reserved, _ = mbhReserved.Marshal()
	processingType = mbh.GetProcessingType()
	require.Equal(t, int32(block.Normal), processingType)

	mbhReserved.ExecutionType = block.Scheduled
	mbh.Reserved, _ = mbhReserved.Marshal()
	require.Equal(t, int32(block.Scheduled), mbh.GetProcessingType())
}

func TestMiniBlockHeader_GetConstructionState(t *testing.T) {
	mbh := &block.MiniBlockHeader{}

	state := mbh.GetConstructionState()
	require.Equal(t, int32(block.Final), state)

	mbhReserved := block.MiniBlockHeaderReserved{State: block.Proposed}
	reserved, err := mbhReserved.Marshal()
	mbh.Reserved = reserved
	require.Nil(t, err)
	state = mbh.GetConstructionState()
	require.Equal(t, int32(block.Proposed), state)

	mbhReserved.State = block.Final
	mbh.Reserved, _ = mbhReserved.Marshal()
	require.Equal(t, int32(block.Final), mbh.GetConstructionState())
}

func TestMiniBlockHeader_IsFinalFromConstructionState(t *testing.T) {
	mbh := &block.MiniBlockHeader{}
	require.True(t, mbh.IsFinal())

	_ = mbh.SetConstructionState(int32(block.Proposed))
	require.False(t, mbh.IsFinal())

	_ = mbh.SetConstructionState(int32(block.Final))
	require.True(t, mbh.IsFinal())
}

func TestMiniBlockHeader_GetIndexOfLastTxProcessed(t *testing.T) {
	mbh := &block.MiniBlockHeader{
		TxCount: 70,
	}

	index := mbh.GetIndexOfLastTxProcessed()
	require.Equal(t, int32(69), index)

	mbhReserved := block.MiniBlockHeaderReserved{IndexOfLastTxProcessed: 78}
	reserved, err := mbhReserved.Marshal()
	mbh.Reserved = reserved
	require.Nil(t, err)
	index = mbh.GetIndexOfLastTxProcessed()
	require.Equal(t, int32(78), index)

	mbhReserved.IndexOfLastTxProcessed = 21
	mbh.Reserved, _ = mbhReserved.Marshal()
	require.Equal(t, int32(21), mbh.GetIndexOfLastTxProcessed())
}

func TestMiniBlockHeader_SetProcessingType(t *testing.T) {
	mbh := &block.MiniBlockHeader{}

	_ = mbh.SetProcessingType(int32(block.Scheduled))
	require.Equal(t, int32(block.Scheduled), mbh.GetProcessingType())

	_ = mbh.SetProcessingType(int32(block.Processed))
	require.Equal(t, int32(block.Processed), mbh.GetProcessingType())
}

func TestMiniBlockHeader_SetConstructionState(t *testing.T) {
	mbh := &block.MiniBlockHeader{}

	_ = mbh.SetConstructionState(int32(block.Proposed))
	require.Equal(t, int32(block.Proposed), mbh.GetConstructionState())

	_ = mbh.SetConstructionState(int32(block.PartialExecuted))
	require.Equal(t, int32(block.PartialExecuted), mbh.GetConstructionState())
}

func TestMiniBlockHeader_SetIndexOfLastTxProcessed(t *testing.T) {
	mbh := &block.MiniBlockHeader{}

	_ = mbh.SetIndexOfLastTxProcessed(3)
	require.Equal(t, int32(3), mbh.GetIndexOfLastTxProcessed())

	_ = mbh.SetIndexOfLastTxProcessed(5)
	require.Equal(t, int32(5), mbh.GetIndexOfLastTxProcessed())
}

func TestMiniBlockHeader_SetIndexOfFirstTxProcessed(t *testing.T) {
	mbh := &block.MiniBlockHeader{}

	_ = mbh.SetIndexOfFirstTxProcessed(3)
	require.Equal(t, int32(3), mbh.GetIndexOfFirstTxProcessed())

	_ = mbh.SetIndexOfFirstTxProcessed(5)
	require.Equal(t, int32(5), mbh.GetIndexOfFirstTxProcessed())
}

func TestMiniBlockHeader_SetProcessingTypeDoesNotChangeConstructionStateOrIndexOfTxProcessed(t *testing.T) {
	mbh := &block.MiniBlockHeader{}
	state := int32(block.Proposed)
	_ = mbh.SetConstructionState(state)
	index := int32(21)
	_ = mbh.SetIndexOfFirstTxProcessed(index - 1)
	_ = mbh.SetIndexOfLastTxProcessed(index)

	_ = mbh.SetProcessingType(int32(block.Scheduled))
	require.Equal(t, state, mbh.GetConstructionState())
	require.Equal(t, index-1, mbh.GetIndexOfFirstTxProcessed())
	require.Equal(t, index, mbh.GetIndexOfLastTxProcessed())

	_ = mbh.SetProcessingType(int32(block.Processed))
	require.Equal(t, state, mbh.GetConstructionState())
	require.Equal(t, index-1, mbh.GetIndexOfFirstTxProcessed())
	require.Equal(t, index, mbh.GetIndexOfLastTxProcessed())

	state = int32(block.PartialExecuted)
	_ = mbh.SetConstructionState(state)
	index = int32(78)
	_ = mbh.SetIndexOfFirstTxProcessed(index - 1)
	_ = mbh.SetIndexOfLastTxProcessed(index)

	_ = mbh.SetProcessingType(int32(block.Scheduled))
	require.Equal(t, state, mbh.GetConstructionState())
	require.Equal(t, index-1, mbh.GetIndexOfFirstTxProcessed())
	require.Equal(t, index, mbh.GetIndexOfLastTxProcessed())

	_ = mbh.SetProcessingType(int32(block.Processed))
	require.Equal(t, state, mbh.GetConstructionState())
	require.Equal(t, index-1, mbh.GetIndexOfFirstTxProcessed())
	require.Equal(t, index, mbh.GetIndexOfLastTxProcessed())
}

func TestMiniBlockHeader_SetConstructionStateDoesNotChangeProcessingTypeOrIndexOfTxProcessed(t *testing.T) {
	mbh := &block.MiniBlockHeader{}
	processingType := int32(block.Scheduled)
	_ = mbh.SetProcessingType(processingType)
	index := int32(21)
	_ = mbh.SetIndexOfFirstTxProcessed(index - 1)
	_ = mbh.SetIndexOfLastTxProcessed(index)

	_ = mbh.SetConstructionState(int32(block.Proposed))
	require.Equal(t, processingType, mbh.GetProcessingType())
	require.Equal(t, index-1, mbh.GetIndexOfFirstTxProcessed())
	require.Equal(t, index, mbh.GetIndexOfLastTxProcessed())

	_ = mbh.SetConstructionState(int32(block.PartialExecuted))
	require.Equal(t, processingType, mbh.GetProcessingType())
	require.Equal(t, index-1, mbh.GetIndexOfFirstTxProcessed())
	require.Equal(t, index, mbh.GetIndexOfLastTxProcessed())

	processingType = int32(block.Processed)
	_ = mbh.SetProcessingType(processingType)
	index = int32(78)
	_ = mbh.SetIndexOfFirstTxProcessed(index - 1)
	_ = mbh.SetIndexOfLastTxProcessed(index)

	_ = mbh.SetConstructionState(int32(block.Proposed))
	require.Equal(t, processingType, mbh.GetProcessingType())
	require.Equal(t, index-1, mbh.GetIndexOfFirstTxProcessed())
	require.Equal(t, index, mbh.GetIndexOfLastTxProcessed())

	_ = mbh.SetConstructionState(int32(block.PartialExecuted))
	require.Equal(t, processingType, mbh.GetProcessingType())
	require.Equal(t, index-1, mbh.GetIndexOfFirstTxProcessed())
	require.Equal(t, index, mbh.GetIndexOfLastTxProcessed())
}

func TestMiniBlockHeader_SetIndexOfTxProcessedDoesNotChangeConstructionStateOrProcessingType(t *testing.T) {
	mbh := &block.MiniBlockHeader{}
	processingType := int32(block.Scheduled)
	_ = mbh.SetProcessingType(processingType)
	state := int32(block.Proposed)
	_ = mbh.SetConstructionState(state)

	_ = mbh.SetIndexOfFirstTxProcessed(20)
	_ = mbh.SetIndexOfLastTxProcessed(21)
	require.Equal(t, processingType, mbh.GetProcessingType())
	require.Equal(t, state, mbh.GetConstructionState())

	_ = mbh.SetIndexOfFirstTxProcessed(77)
	_ = mbh.SetIndexOfLastTxProcessed(78)
	require.Equal(t, processingType, mbh.GetProcessingType())
	require.Equal(t, state, mbh.GetConstructionState())

	processingType = int32(block.Processed)
	_ = mbh.SetProcessingType(processingType)
	state = int32(block.PartialExecuted)
	_ = mbh.SetConstructionState(state)

	_ = mbh.SetIndexOfFirstTxProcessed(20)
	_ = mbh.SetIndexOfLastTxProcessed(21)
	require.Equal(t, processingType, mbh.GetProcessingType())
	require.Equal(t, state, mbh.GetConstructionState())

	_ = mbh.SetIndexOfFirstTxProcessed(77)
	_ = mbh.SetIndexOfLastTxProcessed(78)
	require.Equal(t, processingType, mbh.GetProcessingType())
	require.Equal(t, state, mbh.GetConstructionState())
}
