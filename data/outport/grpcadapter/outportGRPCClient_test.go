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

	t.Run("transport credentials are required", func(t *testing.T) {
		client, err := NewOutportGRPCClient("localhost:12345")

		require.Nil(t, client)
		require.Error(t, err)
	})
}
