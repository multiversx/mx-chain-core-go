package connection

import (
	"sync"

	"github.com/multiversx/mx-chain-core-go/webSocket"
	"github.com/multiversx/mx-chain-core-go/webSocket/data"
)

type websocketClientsHolder struct {
	clients map[string]webSocket.WSConClient
	mut     sync.RWMutex
}

// NewWebsocketClientsHolder will return a new instance of websocketClientsHolder
func NewWebsocketClientsHolder() *websocketClientsHolder {
	return &websocketClientsHolder{
		clients: make(map[string]webSocket.WSConClient),
	}
}

// AddClient will add the provided client to the internal members
func (wch *websocketClientsHolder) AddClient(client webSocket.WSConClient) error {
	if client == nil {
		return data.ErrNilWebSocketClient
	}

	wch.mut.Lock()
	wch.clients[client.GetID()] = client
	wch.mut.Unlock()

	return nil
}

// GetAll will return all the clients
func (wch *websocketClientsHolder) GetAll() map[string]webSocket.WSConClient {
	wch.mut.RLock()
	defer wch.mut.RUnlock()

	clientsMap := make(map[string]webSocket.WSConClient, len(wch.clients))
	for remoteAddr, client := range wch.clients {
		clientsMap[remoteAddr] = client
	}

	return clientsMap
}

// CloseAndRemove will handle the closing of the connection and the deletion from the internal map
func (wch *websocketClientsHolder) CloseAndRemove(remoteAddr string) error {
	wch.mut.Lock()
	client, ok := wch.clients[remoteAddr]
	if !ok {
		wch.mut.Unlock()
		return data.ErrWebSocketClientNotFound
	}

	delete(wch.clients, remoteAddr)
	wch.mut.Unlock()

	return client.Close()
}