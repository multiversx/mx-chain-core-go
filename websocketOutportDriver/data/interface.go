package data

import "io"

// WSConn defines what a sender shall do
type WSConn interface {
	io.Closer
	ReadMessage() (messageType int, payload []byte, err error)
	WriteMessage(messageType int, data []byte) error
}
