package slash

import (
	"sort"

	"github.com/ElrondNetwork/elrond-go-core/core/check"
	"github.com/ElrondNetwork/elrond-go-core/data"
)

// SlashingResult contains the slashable data as well as the severity(slashing level)
// for a possible malicious validator
type SlashingResult struct {
	SlashingLevel ThreatLevel
	Headers       []data.HeaderInfoHandler
}

// ProofTxData represents necessary data to be used in a slashing commitment proof tx by a slashing notifier.
// Each field is required to be added in a transaction.data field
type ProofTxData struct {
	Round   uint64
	ShardID uint32
	ProofID ProofID
}

type ProofID byte

// Used by slashing notifier to create a slashing transaction
// from a proof. Each transaction identifies a different
// slashing event based on this ID
const (
	// MultipleProposalProofID = MultipleProposal's ID
	MultipleProposalProofID ProofID = 0x1
	// MultipleSigningProofID = MultipleSigning's ID
	MultipleSigningProofID ProofID = 0x2
)

func getSortedHeadersV2(headersInfo []data.HeaderInfoHandler) (HeadersV2, error) {
	if headersInfo == nil {
		return HeadersV2{}, data.ErrNilHeaderInfoList
	}

	sortHeadersByHash(headersInfo)
	headers := make([]data.HeaderHandler, 0, len(headersInfo))
	for _, headerInfo := range headersInfo {
		if headerInfo == nil {
			return HeadersV2{}, data.ErrNilHeaderInfo
		}

		headerHandler := headerInfo.GetHeaderHandler()
		hash := headerInfo.GetHash()
		if check.IfNil(headerHandler) {
			return HeadersV2{}, data.ErrNilHeaderHandler
		}
		if hash == nil {
			return HeadersV2{}, data.ErrNilHash
		}

		headers = append(headers, headerHandler)
	}

	headersV2 := HeadersV2{}
	return headersV2, headersV2.SetHeaders(headers)
}

func sortHeadersByHash(headersInfo []data.HeaderInfoHandler) {
	sortFunc := func(i, j int) bool {
		if headersInfo[i] == nil || headersInfo[j] == nil {
			return false
		}
		hash1 := string(headersInfo[i].GetHash())
		hash2 := string(headersInfo[j].GetHash())

		return hash1 < hash2
	}

	sort.Slice(headersInfo, sortFunc)
}
