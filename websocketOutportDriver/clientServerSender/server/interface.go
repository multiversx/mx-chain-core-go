package server

import (
	"github.com/multiversx/mx-chain-core-go/websocketOutportDriver/common"
)

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
