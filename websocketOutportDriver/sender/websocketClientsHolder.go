package sender

import (
	"sync"

	"github.com/multiversx/mx-chain-core-go/websocketOutportDriver/data"
)

type websocketClientsHolder struct {
	clients map[string]*webSocketClient
	mut     sync.RWMutex
}

// NewWebsocketClientsHolder will return a new instance of websocketClientsHolder
func NewWebsocketClientsHolder() *websocketClientsHolder {
	return &websocketClientsHolder{
		clients: make(map[string]*webSocketClient),
	}
}

// AddClient will add the provided client to the internal members
func (wch *websocketClientsHolder) AddClient(client *webSocketClient) error {
	if client == nil {
		return data.ErrNilWebSocketClient
	}

	wch.mut.Lock()
	wch.clients[client.remoteAddr] = client
	wch.mut.Unlock()

	return nil
}

// GetAll will return all the clients
func (wch *websocketClientsHolder) GetAll() map[string]*webSocketClient {
	wch.mut.RLock()
	defer wch.mut.RUnlock()

	clientsMap := make(map[string]*webSocketClient, len(wch.clients))
	for remoteAddr, client := range wch.clients {
		clientsMap[remoteAddr] = client
	}

	return clientsMap
}

// CloseAndRemove will handle the closing of the connection and the deletion from the internal map
func (wch *websocketClientsHolder) CloseAndRemove(remoteAddr string) error {
	wch.mut.Lock()
	defer wch.mut.Unlock()

	client, ok := wch.clients[remoteAddr]
	if !ok {
		return data.ErrWebSocketClientNotFound
	}

	delete(wch.clients, remoteAddr)

	return client.conn.Close()
}
