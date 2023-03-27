package client

import (
	"github.com/multiversx/mx-chain-core-go/websocketOutportDriver"
	"github.com/multiversx/mx-chain-core-go/websocketOutportDriver/data"
)

// HandlerFunc defines the func responsible for handling received payload data from node
type HandlerFunc func(data []byte) error

// PayloadProcessor defines what a websocket payload processor should do
type PayloadProcessor interface {
	ProcessPayload(payload *websocketOutportDriver.PayloadData) error
	Close() error
}

// PayloadParser defines what a websocket payload parser should do
type PayloadParser interface {
	ExtractPayloadData(payload []byte) (*websocketOutportDriver.PayloadData, error)
}

// WSConnClient extends the existing data.WSConn with an option to OpenConnection on demand
type WSConnClient interface {
	data.WSConn
	OpenConnection(url string) error
}
