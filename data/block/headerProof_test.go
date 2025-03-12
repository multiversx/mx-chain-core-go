package block

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHeaderProof_IsInterfaceNil(t *testing.T) {
	t.Parallel()

	var proof *HeaderProof
	require.True(t, proof.IsInterfaceNil())

	proof = &HeaderProof{}
	require.False(t, proof.IsInterfaceNil())
}
