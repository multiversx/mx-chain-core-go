package slash_test

import (
	"testing"

	"github.com/ElrondNetwork/elrond-go-core/data"
	"github.com/ElrondNetwork/elrond-go-core/data/block"
	dataMock "github.com/ElrondNetwork/elrond-go-core/data/mock"
	"github.com/ElrondNetwork/elrond-go-core/data/slash"
	"github.com/ElrondNetwork/elrond-go-core/marshal"
	"github.com/stretchr/testify/require"
)

func TestNewMultipleSigningProof(t *testing.T) {
	tests := []struct {
		args        map[string]slash.SlashingResult
		expectedErr error
	}{
		{
			args:        nil,
			expectedErr: data.ErrNilSlashResult,
		},
		{
			args: map[string]slash.SlashingResult{
				"pubKey": {Headers: nil},
			},
			expectedErr: data.ErrNilHeaderInfoList,
		},
		{
			args: map[string]slash.SlashingResult{
				"pubKey": {Headers: []data.HeaderInfoHandler{nil}},
			},
			expectedErr: data.ErrNilHeaderInfo,
		},
		{
			args:        make(map[string]slash.SlashingResult),
			expectedErr: nil,
		},
	}

	for _, currTest := range tests {
		_, err := slash.NewMultipleSigningProof(currTest.args)
		require.Equal(t, currTest.expectedErr, err)
	}
}

func TestMultipleSigningProof_GetHeaders_GetLevel_GetType(t *testing.T) {
	h1 := &block.HeaderV2{Header: &block.Header{TimeStamp: 1}}
	h2 := &block.HeaderV2{Header: &block.Header{TimeStamp: 2}}
	h3 := &block.HeaderV2{Header: &block.Header{TimeStamp: 3}}

	hInfo1 := &dataMock.HeaderInfoStub{Header: h1, Hash: []byte("h1")}
	hInfo2 := &dataMock.HeaderInfoStub{Header: h2, Hash: []byte("h2")}
	hInfo3 := &dataMock.HeaderInfoStub{Header: h3, Hash: []byte("h3")}

	slashRes1 := slash.SlashingResult{
		SlashingLevel: slash.High,
		Headers:       []data.HeaderInfoHandler{hInfo1, hInfo2},
	}
	slashRes2 := slash.SlashingResult{
		SlashingLevel: slash.Medium,
		Headers:       []data.HeaderInfoHandler{hInfo3},
	}
	slashRes := map[string]slash.SlashingResult{
		"pubKey1": slashRes1,
		"pubKey2": slashRes2,
	}

	proof, err := slash.NewMultipleSigningProof(slashRes)
	require.Nil(t, err)

	require.Equal(t, slash.High, proof.GetLevel([]byte("pubKey1")))
	require.Equal(t, slash.Medium, proof.GetLevel([]byte("pubKey2")))
	require.Equal(t, slash.Zero, proof.GetLevel([]byte("pubKey3")))

	require.Len(t, proof.GetHeaders([]byte("pubKey1")), 2)
	require.Len(t, proof.GetHeaders([]byte("pubKey2")), 1)
	require.Len(t, proof.GetHeaders([]byte("pubKey3")), 0)

	require.Equal(t, proof.GetHeaders([]byte("pubKey1"))[0], h1)
	require.Equal(t, proof.GetHeaders([]byte("pubKey1"))[1], h2)
	require.Equal(t, proof.GetHeaders([]byte("pubKey2"))[0], h3)
	require.Nil(t, proof.GetHeaders([]byte("pubKey3")))
}

func TestMultipleSigningProof_GetProofTxData_EnoughPublicKeysProvided_ExpectError(t *testing.T) {
	proof := slash.MultipleHeaderSigningProof{}

	proofTxData, err := proof.GetProofTxData()
	require.Nil(t, proofTxData)
	require.Equal(t, data.ErrNotEnoughPublicKeysProvided, err)
}

func TestMultipleSigningProof_GetProofTxData_NotEnoughHeadersProvided_ExpectError(t *testing.T) {
	proof := slash.MultipleHeaderSigningProof{PubKeys: [][]byte{[]byte("pub key")}}

	proofTxData, err := proof.GetProofTxData()
	require.Nil(t, proofTxData)
	require.Equal(t, data.ErrNotEnoughHeadersProvided, err)
}

func TestMultipleSigningProof_GetProofTxData_NilHeaderHandler_ExpectError(t *testing.T) {
	proof := slash.MultipleHeaderSigningProof{
		PubKeys: [][]byte{[]byte("pub key")},
		HeadersV2: map[string]slash.HeadersV2{
			"pub key": {Headers: []*block.HeaderV2{nil}},
		},
	}

	proofTxData, err := proof.GetProofTxData()
	require.Nil(t, proofTxData)
	require.Equal(t, data.ErrNilHeaderHandler, err)
}

func TestMultipleSigningProof_GetProofTxData(t *testing.T) {
	round := uint64(1)
	shardID := uint32(2)

	header := &block.HeaderV2{
		Header: &block.Header{
			Round:   round,
			ShardID: shardID,
		},
	}
	proof := slash.MultipleHeaderSigningProof{
		PubKeys: [][]byte{[]byte("pub key")},
		HeadersV2: map[string]slash.HeadersV2{
			"pub key": {Headers: []*block.HeaderV2{header}},
		},
	}
	expectedProofTxData := &slash.ProofTxData{
		Round:   round,
		ShardID: shardID,
		ProofID: slash.MultipleSigningProofID,
	}

	proofTxData, err := proof.GetProofTxData()
	require.Equal(t, expectedProofTxData, proofTxData)
	require.Nil(t, err)
}

func TestMultipleSigningProof_Marshal_Unmarshal(t *testing.T) {
	h1 := &block.HeaderV2{Header: &block.Header{TimeStamp: 1}}
	h2 := &block.HeaderV2{Header: &block.Header{TimeStamp: 2}}
	h3 := &block.HeaderV2{Header: &block.Header{TimeStamp: 3}}
	h4 := &block.HeaderV2{Header: &block.Header{TimeStamp: 4}}

	hInfo1 := &dataMock.HeaderInfoStub{Header: h1, Hash: []byte("h1")}
	hInfo2 := &dataMock.HeaderInfoStub{Header: h2, Hash: []byte("h2")}
	hInfo3 := &dataMock.HeaderInfoStub{Header: h3, Hash: []byte("h3")}
	hInfo4 := &dataMock.HeaderInfoStub{Header: h4, Hash: []byte("h4")}

	slashResPubKey1 := slash.SlashingResult{
		SlashingLevel: slash.Medium,
		Headers:       []data.HeaderInfoHandler{hInfo1, hInfo2},
	}
	slashResPubKey2 := slash.SlashingResult{
		SlashingLevel: slash.High,
		Headers:       []data.HeaderInfoHandler{hInfo3, hInfo4},
	}
	slashResProof1 := map[string]slash.SlashingResult{
		"pubKey1": slashResPubKey1,
		"pubKey2": slashResPubKey2,
	}

	// Same slash result for pubKey1, but change headers order
	slashResPubKey1 = slash.SlashingResult{
		SlashingLevel: slash.Medium,
		Headers:       []data.HeaderInfoHandler{hInfo2, hInfo1},
	}
	// Same slash result for pubKey2, but change headers order
	slashResPubKey2 = slash.SlashingResult{
		SlashingLevel: slash.High,
		Headers:       []data.HeaderInfoHandler{hInfo4, hInfo3},
	}

	// Change pub keys order in map
	slashResProof2 := map[string]slash.SlashingResult{
		"pubKey2": slashResPubKey2,
		"pubKey1": slashResPubKey1,
	}

	proof1, err1 := slash.NewMultipleSigningProof(slashResProof1)
	proof2, err2 := slash.NewMultipleSigningProof(slashResProof2)
	require.Nil(t, err1)
	require.Nil(t, err2)
	require.Equal(t, proof1, proof2)

	marshaller := marshal.GogoProtoMarshalizer{}
	proof1Bytes, err1 := marshaller.Marshal(proof1)
	proof2Bytes, err2 := marshaller.Marshal(proof2)
	require.Nil(t, err1)
	require.Nil(t, err2)
	require.Equal(t, proof1Bytes, proof2Bytes)

	proof1Unmarshalled := &slash.MultipleHeaderSigningProof{}
	proof2Unmarshalled := &slash.MultipleHeaderSigningProof{}
	err1 = marshaller.Unmarshal(proof1Unmarshalled, proof1Bytes)
	err2 = marshaller.Unmarshal(proof2Unmarshalled, proof2Bytes)
	require.Nil(t, err1)
	require.Nil(t, err2)
	require.Equal(t, proof1Unmarshalled, proof1)
	require.Equal(t, proof2Unmarshalled, proof2)
}
