//go:generate protoc -I=. -I=$GOPATH/src/github.com/multiversx/mx-chain-core-go/data/block -I=$GOPATH/src -I=$GOPATH/src/github.com/multiversx/protobuf/protobuf  --gogoslick_out=. incomingHeader.proto

package sovereign

import "github.com/multiversx/mx-chain-core-go/data"

// GetIncomingLogHandlers returns the incoming logs as an array of log handlers
func (ih *IncomingHeader) GetIncomingLogHandlers() []data.LogHandler {
	if ih == nil {
		return nil
	}

	logs := ih.GetIncomingLogs()
	logHandlers := make([]data.LogHandler, len(logs))

	for i := range logs {
		logHandlers[i] = logs[i]
	}

	return logHandlers
}

// GetHeaderHandler returns the incoming headerV2 as a header handler
func (ih *IncomingHeader) GetHeaderHandler() data.HeaderHandler {
	if ih == nil {
		return nil
	}

	return ih.GetHeader()
}
