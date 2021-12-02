package slash

import (
	"encoding/hex"
	"fmt"
	"sort"

	"github.com/ElrondNetwork/elrond-go-core/core/check"
	"github.com/ElrondNetwork/elrond-go-core/data"
)

// MinSlashableNoOfHeaders represents the min number of headers required for a
// multiple proposal/signing proof to be considered slashable
const MinSlashableNoOfHeaders = 2

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
	ProofID byte
}

// Used by slashing notifier to create a slashing transaction
// from a proof. Each transaction identifies a different
// slashing event based on this ID
const (
	// MultipleProposalProofID = MultipleProposal's ID
	MultipleProposalProofID byte = 0x1
	// MultipleSigningProofID = MultipleSigning's ID
	MultipleSigningProofID byte = 0x2
)

func sortAndGetHeadersV2(headersInfo []data.HeaderInfoHandler) (*HeadersV2, error) {
	sortedHeaders, err := sortHeaders(headersInfo)
	if err != nil {
		return nil, err
	}

	headersV2 := &HeadersV2{}
	err = headersV2.SetHeaders(sortedHeaders)
	if err != nil {
		return nil, err
	}

	return headersV2, nil
}

func sortHeaders(headersInfo []data.HeaderInfoHandler) ([]data.HeaderHandler, error) {
	if len(headersInfo) == 0 {
		return nil, data.ErrEmptyHeaderInfoList
	}

	sortHeadersByHash(headersInfo)
	headers := make([]data.HeaderHandler, 0, len(headersInfo))
	hashes := make(map[string]struct{})
	for idx, headerInfo := range headersInfo {
		if headerInfo == nil {
			return nil, fmt.Errorf("%w in sorted headers at index: %v", data.ErrNilHeaderInfo, idx)
		}

		hash := headerInfo.GetHash()
		if hash == nil {
			return nil, fmt.Errorf("%w in sorted headers at index: %v", data.ErrNilHash, idx)
		}

		hashStr := string(hash)
		_, exists := hashes[hashStr]
		if exists {
			return nil, fmt.Errorf("%w, duplicated hash: %s", data.ErrHeadersSameHash, hex.EncodeToString(hash))
		}

		headerHandler := headerInfo.GetHeaderHandler()
		if check.IfNil(headerHandler) {
			return nil, fmt.Errorf("%w in sorted headers for hash: %s", data.ErrNilHeaderHandler, hex.EncodeToString(hash))
		}

		headers = append(headers, headerHandler)
		hashes[hashStr] = struct{}{}
	}

	return headers, nil
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
