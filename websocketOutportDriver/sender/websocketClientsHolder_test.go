package sender

import (
	"testing"

	"github.com/ElrondNetwork/elrond-go-core/websocketOutportDriver/data"
	"github.com/ElrondNetwork/elrond-go-core/websocketOutportDriver/mock"
	"github.com/stretchr/testify/require"
)

func TestNewWebsocketClientsHolder(t *testing.T) {
	t.Parallel()

	wch := NewWebsocketClientsHolder()
	require.NotNil(t, wch)
}

func TestWebsocketClientsHolder_AddClient(t *testing.T) {
	t.Parallel()

	t.Run("nil web socket client", func(t *testing.T) {
		t.Parallel()

		wch := NewWebsocketClientsHolder()
		err := wch.AddClient(nil)
		require.Equal(t, data.ErrNilWebSocketClient, err)
	})

	t.Run("should work", func(t *testing.T) {
		t.Parallel()

		cl := &webSocketClient{}
		wch := NewWebsocketClientsHolder()
		err := wch.AddClient(cl)
		require.NoError(t, err)
	})
}

func TestWebsocketClientsHolder_GetAll(t *testing.T) {
	t.Parallel()

	cl0 := &webSocketClient{remoteAddr: "cl0"}
	cl1 := &webSocketClient{remoteAddr: "cl1"}

	wch := NewWebsocketClientsHolder()

	_ = wch.AddClient(cl0)
	_ = wch.AddClient(cl1)

	clients := wch.GetAll()
	require.Equal(t, cl0, clients["cl0"])
	require.Equal(t, cl1, clients["cl1"])
	require.Equal(t, 2, len(clients))
}

func TestWebsocketClientsHolder_CloseAndRemove(t *testing.T) {
	t.Parallel()

	t.Run("CloseAndRemove should error because the client is not found", func(t *testing.T) {
		t.Parallel()

		wch := NewWebsocketClientsHolder()

		err := wch.CloseAndRemove("new address")
		require.Equal(t, data.ErrWebSocketClientNotFound, err)
	})
	t.Run("CloseAndRemove should work", func(t *testing.T) {
		t.Parallel()

		wch := NewWebsocketClientsHolder()
		closeWasCalled := false
		_ = wch.AddClient(&webSocketClient{remoteAddr: "cl", conn: &mock.WebsocketConnectionStub{
			CloseCalled: func() error {
				closeWasCalled = true
				return nil
			},
		}})

		err := wch.CloseAndRemove("cl")
		require.NoError(t, err)
		require.True(t, closeWasCalled)
	})

}
