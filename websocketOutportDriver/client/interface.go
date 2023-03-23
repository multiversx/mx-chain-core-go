package client

import (
	"github.com/multiversx/mx-chain-core-go/websocketOutportDriver"
	"github.com/multiversx/mx-chain-core-go/websocketOutportDriver/data"
)

// HandlerFunc defines the func responsible for handling received payload data from node
type HandlerFunc func(data []byte) error

// OperationHandler defines a HandlerFunc for each indexer operation type from node
type OperationHandler interface {
	GetOperationHandler(operation data.OperationType) (HandlerFunc, bool)
	Close() error
}

type PayloadParser interface {
	ExtractPayloadData(payload []byte) (*websocketOutportDriver.PayloadData, error)
}
