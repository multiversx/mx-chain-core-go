package server

import (
	"github.com/multiversx/mx-chain-core-go/websocketOutportDriver/common"
)

// Uint64ByteSliceConverter converts byte slice to/from uint64
type Uint64ByteSliceConverter interface {
	ToByteSlice(uint64) []byte
	ToUint64([]byte) (uint64, error)
	IsInterfaceNil() bool
}

type ClientsHandler interface {
	AddClient(client common.WSClient) error
	GetAll() map[string]common.WSClient
	CloseAndRemove(remoteAddr string) error
}

type ClientAcknowledgesHolder interface {
	AddEntry(remoteAddr string)
	Exists(id string) bool
	AddReceivedAcknowledge(remoteAddr string, counter uint64) bool
	RemoveEntryForAddress(remoteAddr string)
	GetAcknowledgesOfAddress(remoteAddr string) (common.AcknowledgesHandler, bool)
}
