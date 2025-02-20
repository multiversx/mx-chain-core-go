package sovereign

import "github.com/multiversx/mx-chain-core-go/data"

// IncomingHeaderHandler defines the incoming header to a sovereign chain that is sent by a notifier
type IncomingHeaderHandler interface {
	GetIncomingEventHandlers() []data.EventHandler
	GetHeaderHandler() data.HeaderHandler
	IsInterfaceNil() bool
}

// Proof defines the proof of proposed mini blocks, using provided incoming header and events
type Proof interface {
	GetHeaderHandler() data.HeaderHandler
	GetIncomingEventHandlers() []data.EventHandler
	GetIncomingMiniBlockHandlers() []data.MiniBlockHandler
	IsInterfaceNil() bool
}
