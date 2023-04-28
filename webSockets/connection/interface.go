package connection

import (
	"context"
	"io"

	"github.com/multiversx/mx-chain-core-go/webSockets/data"
)

// WSConClient defines what a web-sockets connection client should be able to do
type WSConClient interface {
	io.Closer
	SetCloseHandler(closeHandler func(code int, text string) error)
	OpenConnection(url string) error
	WriteMessage(messageType int, data []byte) error
	ReadMessage() (int, []byte, error)
	GetID() string
	IsInterfaceNil() bool
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

// Uint64ByteSliceConverter converts byte slice to/from uint64
type Uint64ByteSliceConverter interface {
	ToByteSlice(uint64) []byte
	ToUint64([]byte) (uint64, error)
	IsInterfaceNil() bool
}
