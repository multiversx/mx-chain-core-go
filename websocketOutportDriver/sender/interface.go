package sender

import "context"

// Uint64ByteSliceConverter converts byte slice to/from uint64
type Uint64ByteSliceConverter interface {
	ToByteSlice(uint64) []byte
	ToUint64([]byte) (uint64, error)
	IsInterfaceNil() bool
}

// HttpServerHandler defines the minimum behaviour of a http server
type HttpServerHandler interface {
	ListenAndServe() error
	Shutdown(ctx context.Context) error
}
