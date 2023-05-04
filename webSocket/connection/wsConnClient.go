package connection

import (
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/multiversx/mx-chain-core-go/webSocket/data"
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
		return data.ErrConnectionAlreadyOpen
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
	conn, err := wsc.getConn()
	if err != nil {
		return 0, nil, err
	}

	return conn.ReadMessage()
}

// WriteMessage calls the underlying write message ws connection func
func (wsc *wsConnClient) WriteMessage(messageType int, payload []byte) error {
	conn, err := wsc.getConn()
	if err != nil {
		return err
	}

	return conn.WriteMessage(messageType, payload)
}

func (wsc *wsConnClient) getConn() (*websocket.Conn, error) {
	wsc.mut.RLock()
	defer wsc.mut.RUnlock()

	if wsc.conn == nil {
		return nil, data.ErrConnectionNotOpen
	}

	conn := wsc.conn

	return conn, nil
}

// GetID will return the unique id of the client
func (wsc *wsConnClient) GetID() string {
	return wsc.clientID
}

// Close will try to cleanly close the connection, if possible
func (wsc *wsConnClient) Close() error {
	// critical section
	wsc.mut.Lock()
	defer wsc.mut.Unlock()

	if wsc.conn == nil {
		return data.ErrConnectionNotOpen
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

// IsInterfaceNil -
func (wsc *wsConnClient) IsInterfaceNil() bool {
	return wsc == nil
}
