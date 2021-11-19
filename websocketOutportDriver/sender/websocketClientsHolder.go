package sender

import "sync"

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
func (wch *websocketClientsHolder) AddClient(client *webSocketClient) {
	wch.mut.Lock()
	wch.clients[client.remoteAddr] = client
	wch.mut.Unlock()
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

// Remove will remove the client from the map
func (wch *websocketClientsHolder) Remove(remoteAddr string) {
	wch.mut.Lock()
	delete(wch.clients, remoteAddr)
	wch.mut.Unlock()
}
