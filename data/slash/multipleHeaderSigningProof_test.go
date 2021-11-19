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
			expectedErr: nil,
		},
		{
			args: map[string]slash.SlashingResult{
				"pubKey": {Headers: []data.HeaderInfoHandler{nil}},
			},
			expectedErr: data.ErrNilHeaderInfo,
		},
		{
			args: map[string]slash.SlashingResult{
				"pubKey": {Headers: []data.HeaderInfoHandler{
					&dataMock.HeaderInfoStub{Header: &block.HeaderV2{}, Hash: []byte("h")},
					&dataMock.HeaderInfoStub{Header: &block.HeaderV2{}, Hash: []byte("h")}},
				},
			},
			expectedErr: data.ErrHeadersSameHash,
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

func TestMultipleSigningProof_GetHeaders_GetLevel(t *testing.T) {
	h1 := &block.HeaderV2{Header: &block.Header{TimeStamp: 1}}
	h2 := &block.HeaderV2{Header: &block.Header{TimeStamp: 2}}
	h3 := &block.HeaderV2{Header: &block.Header{TimeStamp: 3}}
	h4 := &block.HeaderV2{Header: &block.Header{TimeStamp: 4}}
	h5 := &block.HeaderV2{Header: &block.Header{TimeStamp: 5}}

	hInfo1 := &dataMock.HeaderInfoStub{Header: h1, Hash: []byte("h1")}
	hInfo2 := &dataMock.HeaderInfoStub{Header: h2, Hash: []byte("h2")}
	hInfo3 := &dataMock.HeaderInfoStub{Header: h3, Hash: []byte("h3")}
	hInfo4 := &dataMock.HeaderInfoStub{Header: h4, Hash: []byte("h4")}
	hInfo5 := &dataMock.HeaderInfoStub{Header: h5, Hash: []byte("h5")}

	slashRes1 := slash.SlashingResult{
		SlashingLevel: slash.High,
		Headers:       []data.HeaderInfoHandler{hInfo4, hInfo2, hInfo1, hInfo3},
	}
	slashRes2 := slash.SlashingResult{
		SlashingLevel: slash.Medium,
		Headers:       []data.HeaderInfoHandler{hInfo3, hInfo5, hInfo4},
	}
	slashRes3 := slash.SlashingResult{
		SlashingLevel: slash.Zero,
		Headers:       []data.HeaderInfoHandler{},
	}
	slashRes4 := slash.SlashingResult{
		SlashingLevel: slash.Zero,
		Headers:       nil,
	}
	slashRes := map[string]slash.SlashingResult{
		"pubKey1": slashRes1,
		"pubKey2": slashRes2,
		"pubKey3": slashRes3,
		"pubKey4": slashRes4,
	}

	proof, err := slash.NewMultipleSigningProof(slashRes)
	require.Nil(t, err)

	require.Equal(t, slash.High, proof.GetLevel([]byte("pubKey1")))
	require.Equal(t, slash.Medium, proof.GetLevel([]byte("pubKey2")))
	require.Equal(t, slash.Zero, proof.GetLevel([]byte("pubKey3")))
	require.Equal(t, slash.Zero, proof.GetLevel([]byte("pubKey4")))
	require.Equal(t, slash.Zero, proof.GetLevel([]byte("pubKey5")))

	require.Len(t, proof.GetHeaders([]byte("pubKey1")), 4)
	require.Len(t, proof.GetHeaders([]byte("pubKey2")), 3)
	require.Len(t, proof.GetHeaders([]byte("pubKey3")), 0)
	require.Len(t, proof.GetHeaders([]byte("pubKey4")), 0)
	require.Len(t, proof.GetHeaders([]byte("pubKey5")), 0)

	require.Equal(t, proof.GetHeaders([]byte("pubKey1"))[0], h1)
	require.Equal(t, proof.GetHeaders([]byte("pubKey1"))[1], h2)
	require.Equal(t, proof.GetHeaders([]byte("pubKey1"))[2], h3)
	require.Equal(t, proof.GetHeaders([]byte("pubKey1"))[3], h4)
	require.Equal(t, proof.GetHeaders([]byte("pubKey2"))[0], h3)
	require.Equal(t, proof.GetHeaders([]byte("pubKey2"))[1], h4)
	require.Equal(t, proof.GetHeaders([]byte("pubKey2"))[2], h5)
}

func TestMultipleSigningProof_GetProofTxData_NotEnoughPublicKeysProvided_ExpectError(t *testing.T) {
	proof := slash.MultipleHeaderSigningProof{}

	proofTxData, err := proof.GetProofTxData()
	require.Nil(t, proofTxData)
	require.Equal(t, data.ErrNotEnoughPublicKeysProvided, err)
}

func TestMultipleSigningProof_GetProofTxData_NotEnoughHeadersProvided_ExpectError(t *testing.T) {
	slashResPubKey1 := slash.SlashingResult{
		SlashingLevel: slash.High,
		Headers:       []data.HeaderInfoHandler{},
	}
	slashRes := map[string]slash.SlashingResult{
		"pubKey1": slashResPubKey1,
	}
	proof, _ := slash.NewMultipleSigningProof(slashRes)

	proofTxData, err := proof.GetProofTxData()
	require.Nil(t, proofTxData)
	require.Equal(t, data.ErrNotEnoughHeadersProvided, err)
}

func TestMultipleSigningProof_GetProofTxData_NilHeaderHandler_ExpectError(t *testing.T) {
	proof := &slash.MultipleHeaderSigningProof{
		HeadersV2: slash.HeadersV2{Headers: []*block.HeaderV2{nil}},
		SignersSlashData: map[string]slash.SignerSlashingData{
			"pubKey1": {SignedHeadersBitMap: []byte{0x1}},
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
	headerInfo := &dataMock.HeaderInfoStub{
		Header: header,
		Hash:   []byte("hash"),
	}

	slashResPubKey1 := slash.SlashingResult{
		SlashingLevel: slash.High,
		Headers:       []data.HeaderInfoHandler{headerInfo},
	}
	slashRes := map[string]slash.SlashingResult{
		"pubKey1": slashResPubKey1,
	}
	proof, _ := slash.NewMultipleSigningProof(slashRes)

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
	h5 := &block.HeaderV2{Header: &block.Header{TimeStamp: 5}}

	hInfo1 := &dataMock.HeaderInfoStub{Header: h1, Hash: []byte("h1")}
	hInfo2 := &dataMock.HeaderInfoStub{Header: h2, Hash: []byte("h2")}
	hInfo3 := &dataMock.HeaderInfoStub{Header: h3, Hash: []byte("h3")}
	hInfo4 := &dataMock.HeaderInfoStub{Header: h4, Hash: []byte("h4")}
	hInfo5 := &dataMock.HeaderInfoStub{Header: h5, Hash: []byte("h5")}

	slashResPubKey1 := slash.SlashingResult{
		SlashingLevel: slash.Medium,
		Headers:       []data.HeaderInfoHandler{hInfo1, hInfo2, hInfo3},
	}
	slashResPubKey2 := slash.SlashingResult{
		SlashingLevel: slash.High,
		Headers:       []data.HeaderInfoHandler{hInfo3, hInfo4, hInfo5},
	}
	slashResProof1 := map[string]slash.SlashingResult{
		"pubKey1": slashResPubKey1,
		"pubKey2": slashResPubKey2,
	}

	// Same slash result for pubKey1, but change headers order
	slashResPubKey1 = slash.SlashingResult{
		SlashingLevel: slash.Medium,
		Headers:       []data.HeaderInfoHandler{hInfo3, hInfo1, hInfo2},
	}
	// Same slash result for pubKey2, but change headers order
	slashResPubKey2 = slash.SlashingResult{
		SlashingLevel: slash.High,
		Headers:       []data.HeaderInfoHandler{hInfo4, hInfo3, hInfo5},
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
