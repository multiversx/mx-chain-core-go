//go:generate protoc -I=. -I=$GOPATH/src/github.com/ElrondNetwork/elrond-go-core/data/block -I=$GOPATH/src -I=$GOPATH/src/github.com/ElrondNetwork/protobuf/protobuf --gogoslick_out=. multipleHeaderSigningProof.proto
package slash

import (
	"github.com/ElrondNetwork/elrond-go-core/data"
	"github.com/ElrondNetwork/elrond-go-core/data/block"
)

func (m *MultipleHeaderSigningProof) GetType() SlashingType {
	if m == nil {
		return None
	}
	return MultipleSigning
}

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

func NewMultipleSigningProof(slashResult map[string]SlashingResult) (MultipleSigningProofHandler, error) {
	if slashResult == nil {
		return nil, nil //process.ErrNilSlashResult
	}

	pubKeys := make([][]byte, 0, len(slashResult))
	levels := make(map[string]ThreatLevel, len(slashResult))
	headers := make(map[string]block.HeaderInfoList, len(slashResult))

	for pubKey, res := range slashResult {
		pubKeys = append(pubKeys, []byte(pubKey))
		levels[pubKey] = res.SlashingLevel

		tmp := block.HeaderInfoList{}
		err := tmp.SetHeadersInfo(res.Headers)
		if err != nil {
			return nil, err
		}
		headers[pubKey] = tmp
	}

	return &MultipleHeaderSigningProof{
		PubKeys:     pubKeys,
		Levels:      levels,
		HeadersInfo: headers,
	}, nil
}
