package client

import (
	"sync"

	"github.com/gorilla/websocket"
)

type wsConnClient struct {
	mut  sync.RWMutex
	conn *websocket.Conn
}

// NewWSConnClient creates a new wrapper over a websocket connection
func NewWSConnClient() *wsConnClient {
	return &wsConnClient{}
}

// OpenConnection will open a new client with a background context
func (wsc *wsConnClient) OpenConnection(url string) error {
	wsc.mut.Lock()
	defer wsc.mut.Unlock()

	if wsc.conn != nil {
		return errConnectionAlreadyOpened
	}

	var err error
	wsc.conn, _, err = websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return err
	}

	return nil
}

// ReadMessage calls the underlying reading message ws connection func
func (wsc *wsConnClient) ReadMessage() (messageType int, p []byte, err error) {
	wsc.mut.RLock()
	defer wsc.mut.RUnlock()

	if wsc.conn == nil {
		return 0, nil, errConnectionNotOpened
	}

	return wsc.conn.ReadMessage()
}

// WriteMessage calls the underlying write message ws connection func
func (wsc *wsConnClient) WriteMessage(messageType int, data []byte) error {
	wsc.mut.RLock()
	defer wsc.mut.RUnlock()

	if wsc.conn == nil {
		return errConnectionNotOpened
	}

	return wsc.conn.WriteMessage(messageType, data)
}

// Close will try to cleanly close the connection, if possible
func (wsc *wsConnClient) Close() error {
	// critical section
	wsc.mut.Lock()
	defer wsc.mut.Unlock()

	if wsc.conn == nil {
		return errConnectionNotOpened
	}

	log.Debug("closing ws connection...")

	//Cleanly close the connection by sending a close message and then
	//waiting (with timeout) for the server to close the connection.
	err := wsc.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		log.Error("cannot send close message", "error", err)
	}

	err = wsc.conn.Close()
	if err != nil {
		return err
	}

	wsc.conn = nil
	return nil
}

// IsInterfaceNil -
func (wsc *wsConnClient) IsInterfaceNil() bool {
	return wsc == nil
}
