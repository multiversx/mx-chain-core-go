package core

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsGetNodeFromDBError(t *testing.T) {
	t.Parallel()

	t.Run("nil error", func(t *testing.T) {
		t.Parallel()

		assert.False(t, IsGetNodeFromDBError(nil))
	})

	t.Run("closing error", func(t *testing.T) {
		t.Parallel()

		assert.False(t, IsGetNodeFromDBError(ErrContextClosing))
		assert.False(t, IsGetNodeFromDBError(ErrDBIsClosed))
	})

	t.Run("get node from db error", func(t *testing.T) {
		t.Parallel()

		assert.True(t, IsGetNodeFromDBError(fmt.Errorf("trie error: %s", GetNodeFromDBErrorString)))
	})

	t.Run("other error", func(t *testing.T) {
		t.Parallel()

		assert.False(t, IsGetNodeFromDBError(fmt.Errorf("trie error: %s", "other error")))
	})
}

func TestIsClosingError(t *testing.T) {
	t.Parallel()

	t.Run("nil error", func(t *testing.T) {
		t.Parallel()

		assert.False(t, IsClosingError(nil))
	})

	t.Run("closing error", func(t *testing.T) {
		t.Parallel()

		assert.True(t, IsClosingError(ErrContextClosing))
		assert.True(t, IsClosingError(ErrDBIsClosed))
	})

	t.Run("other error", func(t *testing.T) {
		t.Parallel()

		assert.False(t, IsClosingError(fmt.Errorf("trie error: %s", "other error")))
	})
}
