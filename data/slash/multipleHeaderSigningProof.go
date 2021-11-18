//go:generate protoc -I=. -I=$GOPATH/src/github.com/ElrondNetwork/elrond-go-core/data/block -I=$GOPATH/src -I=$GOPATH/src/github.com/ElrondNetwork/protobuf/protobuf --gogoslick_out=. multipleHeaderSigningProof.proto
package slash

import (
	"sort"

	"github.com/ElrondNetwork/elrond-go-core/core/check"
	"github.com/ElrondNetwork/elrond-go-core/data"
)

// GetLevel returns the ThreatLevel of a possible malicious validator
func (m *MultipleHeaderSigningProof) GetLevel(pubKey []byte) ThreatLevel {
	if m == nil {
		return Zero
	}

	level, exists := m.Levels[string(pubKey)]
	if !exists {
		return Zero
	}

	return level
}

// GetHeaders returns all headers that have been signed by a possible malicious validator
func (m *MultipleHeaderSigningProof) GetHeaders(pubKey []byte) []data.HeaderHandler {
	if m == nil {
		return nil
	}

	headersV2, exist := m.HeadersV2[string(pubKey)]
	if !exist {
		return nil
	}

	return headersV2.GetHeaderHandlers()
}

func (m *MultipleHeaderSigningProof) GetProofTxData() (*ProofTxData, error) {
	if m == nil {
		return nil, data.ErrNilPointerReceiver
	}

	pubKeys := m.GetPubKeys()
	if len(pubKeys) == 0 {
		return nil, data.ErrNotEnoughPublicKeysProvided
	}
	pubKey := pubKeys[0]
	headers := m.GetHeaders(pubKey)
	if len(headers) == 0 {
		return nil, data.ErrNotEnoughHeadersProvided
	}
	if check.IfNil(headers[0]) {
		return nil, data.ErrNilHeaderHandler
	}

	return &ProofTxData{
		Round:   headers[0].GetRound(),
		ShardID: headers[0].GetShardID(),
		ProofID: MultipleSigningProofID,
	}, nil
}

// NewMultipleSigningProof returns a MultipleSigningProofHandler from a slashing result
func NewMultipleSigningProof(slashResult map[string]SlashingResult) (MultipleSigningProofHandler, error) {
	if slashResult == nil {
		return nil, data.ErrNilSlashResult
	}

	pubKeys := make([][]byte, 0, len(slashResult))
	levels := make(map[string]ThreatLevel, len(slashResult))
	headers := make(map[string]HeadersV2, len(slashResult))

	sortedPubKeys := getSortedPubKeys(slashResult)
	for _, pubKey := range sortedPubKeys {
		pubKeys = append(pubKeys, []byte(pubKey))
		levels[pubKey] = slashResult[pubKey].SlashingLevel

		sortedHeaders, err := getSortedHeadersV2(slashResult[pubKey].Headers)
		if err != nil {
			return nil, err
		}
		headers[pubKey] = sortedHeaders
	}

	return &MultipleHeaderSigningProof{
		PubKeys:   pubKeys,
		Levels:    levels,
		HeadersV2: headers,
	}, nil
}

func getSortedPubKeys(slashResult map[string]SlashingResult) []string {
	sortedPubKeys := make([]string, 0, len(slashResult))

	for pubKey := range slashResult {
		sortedPubKeys = append(sortedPubKeys, pubKey)
	}
	sort.Strings(sortedPubKeys)

	return sortedPubKeys
}
