package client

import "github.com/gorilla/websocket"

type wsConnClient struct {
	conn *websocket.Conn
}

func NewWSConnClient() *wsConnClient {
	return &wsConnClient{
		conn: &websocket.Conn{},
	}
}

func (wsc *wsConnClient) OpenConnection(url string) error {
	var err error
	wsc.conn, _, err = websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return err
	}

	return nil
}

func (wsc *wsConnClient) Close() error {
	return wsc.conn.Close()
}

func (wsc *wsConnClient) ReadMessage() (messageType int, p []byte, err error) {
	return wsc.conn.ReadMessage()
}
func (wsc *wsConnClient) WriteMessage(messageType int, data []byte) error {
	return wsc.conn.WriteMessage(messageType, data)
}
