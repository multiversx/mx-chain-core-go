//go:generate protoc -I=. -I=$GOPATH/src -I=$GOPATH/src/github.com/ElrondNetwork/protobuf/protobuf  --gogoslick_out=. headerInfo.proto
package block

import "github.com/ElrondNetwork/elrond-go-core/data"

func (m *HeaderInfo) GetHeaderHandler() data.HeaderHandler {
	if m == nil {
		return nil
	}

	return m.Header
}
