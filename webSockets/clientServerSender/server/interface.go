package server

import (
	"github.com/multiversx/mx-chain-core-go/webSockets/common"
)

// ClientsHandler defines what a clients handler should be able to do
type ClientsHandler interface {
	AddClient(client common.WSConClient) error
	GetAll() map[string]common.WSConClient
	CloseAndRemove(remoteAddr string) error
}

// ClientAcknowledgesHolder defines what a client acknowledges holder should be able to do
type ClientAcknowledgesHolder interface {
	AddEntry(id string)
	Exists(id string) bool
	AddReceivedAcknowledge(id string, counter uint64) bool
	RemoveEntryForAddress(id string)
	GetAcknowledgesOfAddress(id string) (common.AcknowledgesHandler, bool)
}
