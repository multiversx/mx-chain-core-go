//go:generate protoc -I=. -I=$GOPATH/src/github.com/ElrondNetwork/elrond-go-core/data/block -I=$GOPATH/src -I=$GOPATH/src/github.com/ElrondNetwork/protobuf/protobuf --gogoslick_out=. multipleHeaderProposalProof.proto
package slash

import (
	"github.com/ElrondNetwork/elrond-go-core/data"
	"github.com/ElrondNetwork/elrond-go-core/data/block"
)

func (m *MultipleHeaderProposalProof) GetType() SlashingType {
	if m == nil {
		return None
	}
	return MultipleProposal
}

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

func NewMultipleProposalProof(slashResult *SlashingResult) (MultipleProposalProofHandler, error) {
	if slashResult == nil {
		return nil, nil //process.ErrNilSlashResult
	}
	if slashResult.Headers == nil {
		return nil, nil //process.ErrNilHeaderHandler
	}

	headersInfo := block.HeaderInfoList{}
	err := headersInfo.SetHeadersInfo(slashResult.Headers)
	if err != nil {
		return nil, err
	}
	return &MultipleHeaderProposalProof{
		Level:       slashResult.SlashingLevel,
		HeadersInfo: headersInfo,
	}, nil
}
