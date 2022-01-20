package slash

import (
	"encoding/hex"
	"fmt"
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

func getSortedHeadersV2(headersInfo []data.HeaderInfoHandler) (*HeadersV2, error) {
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

func getHeaderHashIfUnique(headerInfo data.HeaderInfoHandler, hashes map[string]struct{}) (string, error) {
	if headerInfo == nil {
		return "", data.ErrNilHeaderInfo
	}

	hash := headerInfo.GetHash()
	if hash == nil {
		return "", data.ErrNilHash
	}

	headerHandler := headerInfo.GetHeaderHandler()
	if check.IfNil(headerHandler) {
		return "", fmt.Errorf("%w in header info for hash: %s", data.ErrNilHeaderHandler, hex.EncodeToString(hash))
	}

	hashStr := string(hash)
	_, exists := hashes[hashStr]
	if exists {
		return "", fmt.Errorf("%w, duplicated hash: %s", data.ErrHeadersSameHash, hex.EncodeToString(hash))
	}

	return hashStr, nil
}

func sortHeaders(headersInfo []data.HeaderInfoHandler) ([]data.HeaderHandler, error) {
	if len(headersInfo) == 0 {
		return nil, data.ErrEmptyHeaderInfoList
	}

	sortHeadersByHash(headersInfo)
	headers := make([]data.HeaderHandler, 0, len(headersInfo))
	hashes := make(map[string]struct{})
	for idx, headerInfo := range headersInfo {
		hash, err := getHeaderHashIfUnique(headerInfo, hashes)
		if err != nil {
			return nil, fmt.Errorf("%w in sorted header list at index: %d", err, idx)
		}

		headers = append(headers, headerInfo.GetHeaderHandler())
		hashes[hash] = struct{}{}
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
