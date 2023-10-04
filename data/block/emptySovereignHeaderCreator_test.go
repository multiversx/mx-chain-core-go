package block

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_emptySovereignHeaderCreator_CreateNewHeader(t *testing.T) {
	creator := NewEmptySovereignHeaderCreator()
	require.False(t, creator.IsInterfaceNil())

	hdr := creator.CreateNewHeader()
	require.IsType(t, &SovereignChainHeader{}, hdr)
}
