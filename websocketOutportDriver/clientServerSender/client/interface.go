package client

import "github.com/multiversx/mx-chain-core-go/websocketOutportDriver/data"

// WSConnClient extends the existing data.WSConn with an option to OpenConnection on demand
type WSConnClient interface {
	data.WSConn
	OpenConnection(url string) error
	IsInterfaceNil() bool
}

// Uint64ByteSliceConverter converts byte slice to/from uint64
type Uint64ByteSliceConverter interface {
	ToByteSlice(uint64) []byte
	ToUint64([]byte) (uint64, error)
	IsInterfaceNil() bool
}
