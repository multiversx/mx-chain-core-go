package connection

import (
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/multiversx/mx-chain-core-go/webSockets/data"
	logger "github.com/multiversx/mx-chain-logger-go"
)

var log = logger.GetOrCreate("connection")

type wsConnClient struct {
	mut      sync.RWMutex
	conn     *websocket.Conn
	clientID string
}

// NewWSConnClient creates a new wrapper over a websocket connection
func NewWSConnClient() *wsConnClient {
	return &wsConnClient{}
}

// NewWSConnClientWithConn creates a new wrapper over a provided websocket connection
func NewWSConnClientWithConn(conn *websocket.Conn) *wsConnClient {
	wsc := &wsConnClient{
		conn: conn,
	}
	wsc.clientID = fmt.Sprintf("%p", wsc)

	return wsc
}

// OpenConnection will open a new client with a background context
func (wsc *wsConnClient) OpenConnection(url string) error {
	wsc.mut.Lock()
	defer wsc.mut.Unlock()

	if wsc.conn != nil {
		return data.ErrConnectionAlreadyOpened
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
		return 0, nil, data.ErrConnectionNotOpened
	}

	return wsc.conn.ReadMessage()
}

// WriteMessage calls the underlying write message ws connection func
func (wsc *wsConnClient) WriteMessage(messageType int, payload []byte) error {
	wsc.mut.RLock()
	defer wsc.mut.RUnlock()

	if wsc.conn == nil {
		return data.ErrConnectionNotOpened
	}

	return wsc.conn.WriteMessage(messageType, payload)
}

// GetID will return the uniq id of the client
func (wsc *wsConnClient) GetID() string {
	return wsc.clientID
}

// Close will try to cleanly close the connection, if possible
func (wsc *wsConnClient) Close() error {
	// critical section
	wsc.mut.Lock()
	defer wsc.mut.Unlock()

	if wsc.conn == nil {
		return data.ErrConnectionNotOpened
	}

	log.Debug("closing ws connection...")

	//Cleanly close the connection by sending a close message and then
	//waiting (with timeout) for the server to close the connection.
	err := wsc.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		log.Error("cannot send close message", "error", err)
	}
	wsc.conn.CloseHandler()

	err = wsc.conn.Close()
	if err != nil {
		return err
	}

	wsc.conn = nil
	return nil
}

// SetCloseHandler will set the close handler
func (wsc *wsConnClient) SetCloseHandler(closeHandler func(code int, text string) error) {
	// critical section
	wsc.mut.Lock()
	defer wsc.mut.Unlock()

	wsc.conn.SetCloseHandler(closeHandler)
}

// IsInterfaceNil -
func (wsc *wsConnClient) IsInterfaceNil() bool {
	return wsc == nil
}
