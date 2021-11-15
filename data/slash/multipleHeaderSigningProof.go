//go:generate protoc -I=. -I=$GOPATH/src/github.com/ElrondNetwork/elrond-go-core/data/block -I=$GOPATH/src -I=$GOPATH/src/github.com/ElrondNetwork/protobuf/protobuf --gogoslick_out=. multipleHeaderSigningProof.proto
package slash

import (
	"sort"

	"github.com/ElrondNetwork/elrond-go-core/data"
)

// GetType returns MultipleSigning
func (m *MultipleHeaderSigningProof) GetType() SlashingType {
	if m == nil {
		return None
	}
	return MultipleSigning
}

// GetLevel returns the ThreatLevel of a possible malicious validator
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
