//go:generate protoc -I=. -I=$GOPATH/src/github.com/ElrondNetwork/elrond-go-core/data/block -I=$GOPATH/src -I=$GOPATH/src/github.com/ElrondNetwork/protobuf/protobuf --gogoslick_out=. multipleHeaderSigningProof.proto
package slash

import (
	"github.com/ElrondNetwork/elrond-go-core/data"
)

func (m *MultipleHeaderSigningProof) GetLevel(pubKey []byte) ThreatLevel {
	if m == nil {
		return Low
	}

	level, exists := m.Levels[string(pubKey)]
	if !exists {
		return Low
	}

	return level
}

func (m *MultipleHeaderSigningProof) GetHeaders(pubKey []byte) []data.HeaderInfoHandler {
	if m == nil {
		return nil
	}

	headersInfo, exist := m.HeadersInfo[string(pubKey)]
	if !exist {
		return nil
	}

	ret := make([]data.HeaderInfoHandler, len(headersInfo.Headers))
	for _, headerInfo := range headersInfo.GetHeaders() {
		ret = append(ret, headerInfo)
	}

	return ret
}
