package core

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPeerID(t *testing.T) {
	t.Parallel()
	t.Run("invalid string should error", func(t *testing.T) {
		t.Parallel()

		providedStr := "provided str"
		_, err := NewPeerID(providedStr)
		assert.NotNil(t, err)
	})
	t.Run("should work", func(t *testing.T) {
		t.Parallel()

		providedStr := "provided str"
		providedPeerID := PeerID(providedStr)
		pid, err := NewPeerID(providedPeerID.Pretty())
		assert.Nil(t, err)
		assert.True(t, bytes.Equal([]byte(providedPeerID), pid.Bytes()))
	})
}
