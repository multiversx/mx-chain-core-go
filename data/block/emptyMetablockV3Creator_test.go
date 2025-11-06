package block

import (
	"fmt"
	"testing"

	"github.com/multiversx/mx-chain-core-go/core/check"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewEmptyMetaBlockV3Creator(t *testing.T) {
	t.Parallel()

	creator := NewEmptyMetaBlockCreator()
	require.False(t, check.IfNil(creator))
}

func TestEmptyMetaBlockV3Creator_CreateNewHeader(t *testing.T) {
	t.Parallel()

	creator := NewEmptyMetaBlockV3Creator()
	header := creator.CreateNewHeader()
	require.False(t, check.IfNil(header))
	require.False(t, check.IfNil(header.(*MetaBlockV3)))
	require.True(t, header.IsHeaderV3())
	assert.Equal(t, "*block.MetaBlockV3", fmt.Sprintf("%T", header))
}
