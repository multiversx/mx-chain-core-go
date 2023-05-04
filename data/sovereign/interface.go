package sovereign

import "github.com/multiversx/mx-chain-core-go/data"

type IncomingHeaderHandler interface {
	GetIncomingLogHandlers() []data.LogHandler
	GetHeaderHandler() data.HeaderHandler
}
