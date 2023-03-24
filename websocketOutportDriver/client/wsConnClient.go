package client

import (
	"github.com/gorilla/websocket"
	"github.com/multiversx/mx-chain-core-go/core/check"
)

type wsConnClient struct {
	conn *websocket.Conn
}

// NewWSConnClient creates a new wrapper over a websocket connection
func NewWSConnClient() *wsConnClient {
	return &wsConnClient{
		conn: &websocket.Conn{},
	}
}

// OpenConnection will open a new client with a background context
func (wsc *wsConnClient) OpenConnection(url string) error {
	var err error
	wsc.conn, _, err = websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return err
	}

	return nil
}

// Close will try to cleanly close the connection, if possible
func (wsc *wsConnClient) Close() error {
	log.Debug("closing ws connection...")
	if check.IfNilReflect(wsc.conn) {
		return nil
	}

	//Cleanly close the connection by sending a close message and then
	//waiting (with timeout) for the server to close the connection.
	err := wsc.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		log.Error("cannot send close message", "error", err)
	}

	return wsc.conn.Close()
}

// ReadMessage calls the underlying reading message ws connection func
func (wsc *wsConnClient) ReadMessage() (messageType int, p []byte, err error) {
	return wsc.conn.ReadMessage()
}

// WriteMessage calls the underlying write message ws connection func
func (wsc *wsConnClient) WriteMessage(messageType int, data []byte) error {
	return wsc.conn.WriteMessage(messageType, data)
}
