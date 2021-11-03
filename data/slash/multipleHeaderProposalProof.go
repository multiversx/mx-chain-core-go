//go:generate protoc -I=. -I=$GOPATH/src/github.com/ElrondNetwork/elrond-go-core/data/block -I=$GOPATH/src -I=$GOPATH/src/github.com/ElrondNetwork/protobuf/protobuf --gogoslick_out=. multipleHeaderProposalProof.proto
package slash

import (
	"github.com/ElrondNetwork/elrond-go-core/data"
)

func (m *MultipleHeaderProposalProof) GetType() SlashingType {
	if m == nil {
		return None
	}
	return MultipleProposal
}

func (m *MultipleHeaderProposalProof) GetHeaders() []data.HeaderHandler {
	if m == nil {
		return nil
	}

	return m.HeadersV2.GetHeaderHandlers()
}

func NewMultipleProposalProof(slashResult *SlashingResult) (MultipleProposalProofHandler, error) {
	if slashResult == nil {
		return nil, data.ErrNilSlashResult
	}
	if slashResult.Headers == nil {
		return nil, data.ErrNilHeaderHandler
	}

	headersV2 := HeadersV2{}
	err := headersV2.SetHeaders(slashResult.Headers)
	if err != nil {
		return nil, err
	}

	return &MultipleHeaderProposalProof{
		Level:     slashResult.SlashingLevel,
		HeadersV2: headersV2,
	}, nil
}
