package sovereign

import "github.com/multiversx/mx-chain-core-go/data"

// IncomingHeaderHandler defines the incoming header to a sovereign chain that is sent by a notifier
type IncomingHeaderHandler interface {
	GetIncomingLogHandlers() []data.LogHandler
	GetHeaderHandler() data.HeaderHandler
}
