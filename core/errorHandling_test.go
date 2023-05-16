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

func TestUnwrapGetNodeFromDBErr(t *testing.T) {
	t.Parallel()

	key := []byte("key")
	identifier := "identifier"
	err := fmt.Errorf("key not found")

	getNodeFromDbErr := NewGetNodeFromDBErrWithKey(key, err, identifier)
	wrappedErr1 := fmt.Errorf("wrapped error 1: %w", getNodeFromDbErr)
	wrappedErr2 := fmt.Errorf("wrapped error 2: %w", wrappedErr1)
	wrappedErr3 := fmt.Errorf("wrapped error 3: %w", wrappedErr2)

	assert.Nil(t, UnwrapGetNodeFromDBErr(nil))
	assert.Nil(t, UnwrapGetNodeFromDBErr(err))
	assert.Equal(t, getNodeFromDbErr, UnwrapGetNodeFromDBErr(getNodeFromDbErr))
	assert.Equal(t, getNodeFromDbErr, UnwrapGetNodeFromDBErr(wrappedErr1))
	assert.Equal(t, getNodeFromDbErr, UnwrapGetNodeFromDBErr(wrappedErr2))
	assert.Equal(t, getNodeFromDbErr, UnwrapGetNodeFromDBErr(wrappedErr3))
}
