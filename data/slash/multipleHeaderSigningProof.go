//go:generate protoc -I=. -I=$GOPATH/src/github.com/ElrondNetwork/elrond-go-core/data/block -I=$GOPATH/src -I=$GOPATH/src/github.com/ElrondNetwork/protobuf/protobuf --gogoslick_out=. multipleHeaderSigningProof.proto
package slash

import (
	"github.com/ElrondNetwork/elrond-go-core/core/check"
	"github.com/ElrondNetwork/elrond-go-core/core/sliceUtil"
	"github.com/ElrondNetwork/elrond-go-core/data"
)

const byteSize = 8

// GetPubKeys - returns all validator's public keys which have signed multiple headers
func (m *MultipleHeaderSigningProof) GetPubKeys() [][]byte {
	if m == nil {
		return nil
	}

	ret := make([][]byte, 0, len(m.SignersSlashData))
	for pubKey := range m.SignersSlashData {
		ret = append(ret, []byte(pubKey))
	}

	return ret
}

// GetLevel returns the ThreatLevel of a possible malicious validator
func (m *MultipleHeaderSigningProof) GetLevel(pubKey []byte) ThreatLevel {
	if m == nil {
		return Zero
	}

	slashData, exists := m.SignersSlashData[string(pubKey)]
	if !exists {
		return Zero
	}

	return slashData.ThreatLevel
}

// GetHeaders returns all headers that have been signed by a possible malicious validator
func (m *MultipleHeaderSigningProof) GetHeaders(pubKey []byte) []data.HeaderHandler {
	if m == nil {
		return nil
	}

	slashData, exists := m.SignersSlashData[string(pubKey)]
	if !exists {
		return nil
	}

	idx := uint32(0)
	bitmap := slashData.GetSignedHeadersBitMap()
	headers := m.HeadersV2.GetHeaderHandlers()

	ret := make([]data.HeaderHandler, 0)
	for _, header := range headers {
		if sliceUtil.IsIndexSetInBitmap(idx, bitmap) {
			ret = append(ret, header)
		}
		idx++
	}

	return ret
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

	headersInfo, err := getAllUniqueHeaders(slashResult)
	if err != nil {
		return nil, err
	}
	sortHeadersByHash(headersInfo)

	idx := uint32(0)
	hashIndexMap := make(map[string]uint32)
	sortedHeaders := make([]data.HeaderHandler, 0, len(headersInfo))
	for _, headerInfo := range headersInfo {
		hashIndexMap[string(headerInfo.GetHash())] = idx
		sortedHeaders = append(sortedHeaders, headerInfo.GetHeaderHandler())
		idx++
	}

	signersSlashData := make(map[string]SignerSlashingData)
	bitMapLen := len(sortedHeaders)/byteSize + 1
	for pubKey, res := range slashResult {
		bitmap := make([]byte, bitMapLen)
		for _, header := range res.Headers {
			index, exists := hashIndexMap[string(header.GetHash())]
			if exists {
				sliceUtil.SetIndexInBitmap(index, bitmap)
			}
		}
		signersSlashData[pubKey] = SignerSlashingData{
			SignedHeadersBitMap: bitmap,
			ThreatLevel:         res.SlashingLevel,
		}
	}

	headersV2 := HeadersV2{}
	err = headersV2.SetHeaders(sortedHeaders)
	if err != nil {
		return nil, err
	}

	return &MultipleHeaderSigningProof{
		HeadersV2:        headersV2,
		SignersSlashData: signersSlashData,
	}, nil
}

func getAllUniqueHeaders(slashResult map[string]SlashingResult) ([]data.HeaderInfoHandler, error) {
	headersInfo := make([]data.HeaderInfoHandler, 0, len(slashResult))
	hashes := make(map[string]struct{})

	for _, res := range slashResult {
		for _, currHeaderInfo := range res.Headers {
			if currHeaderInfo == nil {
				return nil, data.ErrNilHeaderInfo
			}

			currHash := string(currHeaderInfo.GetHash())
			_, exists := hashes[currHash]
			if exists {
				continue
			}

			hashes[currHash] = struct{}{}
			headersInfo = append(headersInfo, currHeaderInfo)
		}
	}

	return headersInfo, nil
}
