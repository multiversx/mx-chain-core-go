package common

import (
	"context"
	"io"

	"github.com/multiversx/mx-chain-core-go/websocketOutportDriver/data"
)

type AcknowledgesHandler interface {
	ProcessAcknowledged(counter uint64) bool
}

type WSClient interface {
	io.Closer
	WriteMessage(messageType int, data []byte) error
	ReadMessage() (int, []byte, error)
	GetID() string
}

// PayloadProcessor defines what a websocket payload processor should do
type PayloadProcessor interface {
	ProcessPayload(payload *data.PayloadData) error
	IsInterfaceNil() bool
	Close() error
}

// PayloadParser defines what a websocket payload parser should do
type PayloadParser interface {
	ExtractPayloadData(payload []byte) (*data.PayloadData, error)
	IsInterfaceNil() bool
}

// HttpServerHandler defines the minimum behaviour of a http server
type HttpServerHandler interface {
	ListenAndServe() error
	Shutdown(ctx context.Context) error
}

type MessagesListener interface {
	Listen() bool
	Close()
}

// Uint64ByteSliceConverter converts byte slice to/from uint64
type Uint64ByteSliceConverter interface {
	ToByteSlice(uint64) []byte
	ToUint64([]byte) (uint64, error)
	IsInterfaceNil() bool
}

// ListenersHolder will hold a map with all the MessagesListener
type ListenersHolder interface {
	Add(id string, listener MessagesListener)
	Remove(id string)
	GetAll() map[string]MessagesListener
}
