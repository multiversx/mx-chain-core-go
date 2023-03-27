package client

import (
	"github.com/multiversx/mx-chain-core-go/websocketOutportDriver"
	"github.com/multiversx/mx-chain-core-go/websocketOutportDriver/data"
)

// HandlerFunc defines the func responsible for handling received payload data from node
type HandlerFunc func(data []byte) error

type PayloadProcessor interface {
	ProcessPayload(payload *websocketOutportDriver.PayloadData) error
	Close() error
}

type PayloadParser interface {
	ExtractPayloadData(payload []byte) (*websocketOutportDriver.PayloadData, error)
}

type WSConnClient interface {
	data.WSConn
	OpenConnection(url string) error
}
