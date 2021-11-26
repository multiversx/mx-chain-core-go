package sender

import "io"

// WSConn defines what a sender shall do
type WSConn interface {
	io.Closer
	ReadMessage() (messageType int, payload []byte, err error)
	WriteMessage(messageType int, data []byte) error
}

// Uint64ByteSliceConverter converts byte slice to/from uint64
type Uint64ByteSliceConverter interface {
	ToByteSlice(uint64) []byte
	ToUint64([]byte) (uint64, error)
	IsInterfaceNil() bool
}
