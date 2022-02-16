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

func TestMiniBlockHeader_IsFinal(t *testing.T) {
	mbh := &block.MiniBlockHeader{}
	require.True(t, mbh.IsFinal())

	_ = mbh.SetConstructionState(int32(block.Proposed))
	require.False(t, mbh.IsFinal())

	_ = mbh.SetConstructionState(int32(block.Final))
	require.True(t, mbh.IsFinal())
}

func TestMiniBlockHeader_SetProcessingType(t *testing.T) {
	mbh := &block.MiniBlockHeader{}

	_ = mbh.SetProcessingType(int32(block.Scheduled))
	require.Equal(t, int32(block.Scheduled), mbh.GetProcessingType())

	_ = mbh.SetProcessingType(int32(block.Normal))
	require.Equal(t, int32(block.Normal), mbh.GetProcessingType())
}

func TestMiniBlockHeader_SetConstructionState(t *testing.T) {
	mbh := &block.MiniBlockHeader{}

	_ = mbh.SetConstructionState(int32(block.Proposed))
	require.Equal(t, int32(block.Proposed), mbh.GetConstructionState())

	_ = mbh.SetConstructionState(int32(block.Final))
	require.Equal(t, int32(block.Final), mbh.GetConstructionState())
}

func TestMiniBlockHeader_SetProcessingTypeDoesNotChangeConstructionState(t *testing.T) {
	mbh := &block.MiniBlockHeader{}
	state := int32(block.Proposed)
	_ = mbh.SetConstructionState(state)

	_ = mbh.SetProcessingType(int32(block.Scheduled))
	require.Equal(t, state, mbh.GetConstructionState())

	_ = mbh.SetProcessingType(int32(block.Normal))
	require.Equal(t, state, mbh.GetConstructionState())

	state = int32(block.Final)
	_ = mbh.SetConstructionState(state)

	_ = mbh.SetProcessingType(int32(block.Scheduled))
	require.Equal(t, state, mbh.GetConstructionState())

	_ = mbh.SetProcessingType(int32(block.Normal))
	require.Equal(t, state, mbh.GetConstructionState())
}

func TestMiniBlockHeader_SetConstructionStateDoesNotChangeProcessingType(t *testing.T) {
	mbh := &block.MiniBlockHeader{}
	processingType := int32(block.Scheduled)
	_ = mbh.SetProcessingType(processingType)

	_ = mbh.SetConstructionState(int32(block.Proposed))
	require.Equal(t, processingType, mbh.GetConstructionState())

	_ = mbh.SetConstructionState(int32(block.Final))
	require.Equal(t, processingType, mbh.GetProcessingType())

	processingType = int32(block.Normal)
	_ = mbh.SetProcessingType(processingType)

	_ = mbh.SetConstructionState(int32(block.Proposed))
	require.Equal(t, processingType, mbh.GetProcessingType())

	_ = mbh.SetConstructionState(int32(block.Final))
	require.Equal(t, processingType, mbh.GetProcessingType())
}
