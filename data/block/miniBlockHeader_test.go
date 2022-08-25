package block_test

import (
	"math"
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

	require.Equal(t, mbh.TxCount-1, uint32(mbh.GetIndexOfLastTxProcessed()))

	mbhReserved := block.MiniBlockHeaderReserved{}

	mbhReserved.IndexOfLastTxProcessed = 78
	reserved, err := mbhReserved.Marshal()
	require.Nil(t, err)

	mbh.Reserved = reserved
	require.Equal(t, mbhReserved.IndexOfLastTxProcessed, mbh.GetIndexOfLastTxProcessed())

	mbhReserved.IndexOfLastTxProcessed = -1
	mbh.Reserved, _ = mbhReserved.Marshal()
	require.Equal(t, mbhReserved.IndexOfLastTxProcessed, mbh.GetIndexOfLastTxProcessed())

	mbhReserved.IndexOfLastTxProcessed = math.MaxInt32
	mbh.Reserved, _ = mbhReserved.Marshal()
	require.Equal(t, mbhReserved.IndexOfLastTxProcessed, mbh.GetIndexOfLastTxProcessed())
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

	providedIndexOfLastTxProcessed := int32(3)
	_ = mbh.SetIndexOfLastTxProcessed(providedIndexOfLastTxProcessed)
	require.Equal(t, providedIndexOfLastTxProcessed, mbh.GetIndexOfLastTxProcessed())

	providedIndexOfLastTxProcessed = int32(-1)
	_ = mbh.SetIndexOfLastTxProcessed(providedIndexOfLastTxProcessed)
	require.Equal(t, providedIndexOfLastTxProcessed, mbh.GetIndexOfLastTxProcessed())

	providedIndexOfLastTxProcessed = math.MaxInt32
	_ = mbh.SetIndexOfLastTxProcessed(providedIndexOfLastTxProcessed)
	require.Equal(t, providedIndexOfLastTxProcessed, mbh.GetIndexOfLastTxProcessed())
}

func TestMiniBlockHeader_SetIndexOfFirstTxProcessed(t *testing.T) {
	mbh := &block.MiniBlockHeader{}

	providedIndexOfLastTxProcessed := int32(3)
	_ = mbh.SetIndexOfFirstTxProcessed(providedIndexOfLastTxProcessed)
	require.Equal(t, providedIndexOfLastTxProcessed, mbh.GetIndexOfFirstTxProcessed())

	providedIndexOfLastTxProcessed = int32(-1)
	_ = mbh.SetIndexOfFirstTxProcessed(providedIndexOfLastTxProcessed)
	require.Equal(t, providedIndexOfLastTxProcessed, mbh.GetIndexOfFirstTxProcessed())

	providedIndexOfLastTxProcessed = math.MaxInt32
	_ = mbh.SetIndexOfFirstTxProcessed(providedIndexOfLastTxProcessed)
	require.Equal(t, providedIndexOfLastTxProcessed, mbh.GetIndexOfFirstTxProcessed())
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

	index = int32(0)
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

	index = math.MaxInt32
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

	index = int32(0)
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

	index = math.MaxInt32
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

	index = int32(0)
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

	index = math.MaxInt32
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

	index = int32(0)
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

	index = math.MaxInt32
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

	_ = mbh.SetIndexOfFirstTxProcessed(-1)
	_ = mbh.SetIndexOfLastTxProcessed(-1)
	require.Equal(t, processingType, mbh.GetProcessingType())
	require.Equal(t, state, mbh.GetConstructionState())

	_ = mbh.SetIndexOfFirstTxProcessed(math.MaxInt32)
	_ = mbh.SetIndexOfLastTxProcessed(math.MaxInt32)
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

	_ = mbh.SetIndexOfFirstTxProcessed(-1)
	_ = mbh.SetIndexOfLastTxProcessed(-1)
	require.Equal(t, processingType, mbh.GetProcessingType())
	require.Equal(t, state, mbh.GetConstructionState())

	_ = mbh.SetIndexOfFirstTxProcessed(math.MaxInt32)
	_ = mbh.SetIndexOfLastTxProcessed(math.MaxInt32)
	require.Equal(t, processingType, mbh.GetProcessingType())
	require.Equal(t, state, mbh.GetConstructionState())
}

func TestMiniBlockHeader_GetIndexOfLastTxProcessedShouldWorkWhenPreviousVersionDoesNotSet(t *testing.T) {
	mbh := &block.MiniBlockHeader{
		TxCount: 70,
	}

	mbhReserved := block.MiniBlockHeaderReserved{}

	mbhReserved.IndexOfLastTxProcessed = 0
	mbhReserved.State = block.PartialExecuted
	reserved, _ := mbhReserved.Marshal()
	mbh.Reserved = reserved
	require.Equal(t, int32(0), mbh.GetIndexOfLastTxProcessed())

	mbhReserved.IndexOfLastTxProcessed = 0
	mbhReserved.State = block.Proposed
	reserved, _ = mbhReserved.Marshal()
	mbh.Reserved = reserved
	require.Equal(t, int32(mbh.TxCount-1), mbh.GetIndexOfLastTxProcessed())

	mbhReserved.IndexOfLastTxProcessed = -1
	mbhReserved.State = block.Proposed
	reserved, _ = mbhReserved.Marshal()
	mbh.Reserved = reserved
	require.Equal(t, int32(-1), mbh.GetIndexOfLastTxProcessed())
}
