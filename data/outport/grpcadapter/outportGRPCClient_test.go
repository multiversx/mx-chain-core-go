package grpcadapter

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewOutportGRPCClient(t *testing.T) {
	t.Run("empty target should error", func(t *testing.T) {
		client, err := NewOutportGRPCClient("")

		require.Nil(t, client)
		require.True(t, errors.Is(err, ErrEmptyOutportGRPCAddress))
	})
}
