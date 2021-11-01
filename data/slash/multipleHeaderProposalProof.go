//go:generate protoc -I=. -I=$GOPATH/src/github.com/ElrondNetwork/elrond-go-core/data/block -I=$GOPATH/src -I=$GOPATH/src/github.com/ElrondNetwork/protobuf/protobuf --gogoslick_out=. multipleHeaderProposalProof.proto
package slash

import "github.com/ElrondNetwork/elrond-go-core/data"

func (m *MultipleHeaderProposalProof) GetHeaders() []data.HeaderInfoHandler {
	if m == nil {
		return nil
	}
	ret := make([]data.HeaderInfoHandler, len(m.HeadersInfo.Headers))

	for _, headerInfo := range m.HeadersInfo.GetHeaders() {
		ret = append(ret, headerInfo)
	}

	return ret
}
