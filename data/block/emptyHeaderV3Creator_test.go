package block

import (
	"fmt"
	"testing"

	"github.com/multiversx/mx-chain-core-go/core/check"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewEmptyHeaderV3Creator(t *testing.T) {
	t.Parallel()

	creator := NewEmptyHeaderV3Creator()
	require.False(t, check.IfNil(creator))
}

func TestEmptyHeaderV3Creator_CreateNewHeader(t *testing.T) {
	t.Parallel()

	creator := NewEmptyHeaderV3Creator()
	header := creator.CreateNewHeader()
	require.False(t, check.IfNil(header))
	require.False(t, check.IfNil(header.(*HeaderV3)))
	require.True(t, header.IsHeaderV3())
	assert.Equal(t, "*block.HeaderV3", fmt.Sprintf("%T", header))
}
