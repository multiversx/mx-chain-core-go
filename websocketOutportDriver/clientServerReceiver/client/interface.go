package client

import (
	"github.com/multiversx/mx-chain-core-go/websocketOutportDriver/data"
)

// WSConnClient extends the existing data.WSConn with an option to OpenConnection on demand
type WSConnClient interface {
	data.WSConn
	OpenConnection(url string) error
	GetID() string
	IsInterfaceNil() bool
}

// WSClient defines what a websocket client handler should do
type WSClient interface {
	Start()
	Close()
}
