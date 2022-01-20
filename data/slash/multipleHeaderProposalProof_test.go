package slash_test

import (
	"strings"
	"testing"

	"github.com/ElrondNetwork/elrond-go-core/data"
	"github.com/ElrondNetwork/elrond-go-core/data/block"
	dataMock "github.com/ElrondNetwork/elrond-go-core/data/mock"
	"github.com/ElrondNetwork/elrond-go-core/data/slash"
	"github.com/ElrondNetwork/elrond-go-core/marshal"
	"github.com/stretchr/testify/require"
)

func TestNewMultipleProposalProof(t *testing.T) {
	tests := []struct {
		args        *slash.SlashingResult
		expectedErr error
	}{
		{
			args:        nil,
			expectedErr: data.ErrNilSlashResult,
		},
		{
			args: &slash.SlashingResult{
				SlashingLevel: slash.Medium,
				Headers:       nil,
			},
			expectedErr: data.ErrEmptyHeaderInfoList,
		},
		{
			args: &slash.SlashingResult{
				SlashingLevel: slash.Medium,
				Headers:       []data.HeaderInfoHandler{},
			},
			expectedErr: data.ErrEmptyHeaderInfoList,
		},
		{
			args: &slash.SlashingResult{
				SlashingLevel: slash.Medium,
				Headers:       []data.HeaderInfoHandler{nil, &dataMock.HeaderInfoStub{}},
			},
			expectedErr: data.ErrNilHeaderInfo,
		},
		{
			args: &slash.SlashingResult{
				SlashingLevel: slash.Medium,
				Headers: []data.HeaderInfoHandler{
					&dataMock.HeaderInfoStub{Header: &block.HeaderV2{}, Hash: []byte("h")},
					&dataMock.HeaderInfoStub{Header: &block.HeaderV2{}, Hash: []byte("h")},
				},
			},
			expectedErr: data.ErrHeadersSameHash,
		},
	}

	for _, currTest := range tests {
		_, err := slash.NewMultipleProposalProof(currTest.args)
		require.Error(t, err)
		require.True(t, strings.Contains(err.Error(), currTest.expectedErr.Error()))
	}
}

func TestMultipleProposalProof_GetHeadersGetLevel(t *testing.T) {
	h1 := &block.HeaderV2{Header: &block.Header{TimeStamp: 1}}
	h2 := &block.HeaderV2{Header: &block.Header{TimeStamp: 2}}
	h3 := &block.HeaderV2{Header: &block.Header{TimeStamp: 3}}

	hInfo1 := &dataMock.HeaderInfoStub{Header: h1, Hash: []byte("h1")}
	hInfo2 := &dataMock.HeaderInfoStub{Header: h2, Hash: []byte("h2")}
	hInfo3 := &dataMock.HeaderInfoStub{Header: h3, Hash: []byte("h3")}

	slashRes := &slash.SlashingResult{
		SlashingLevel: slash.Medium,
		Headers:       []data.HeaderInfoHandler{hInfo2, hInfo1, hInfo3},
	}

	proof, err := slash.NewMultipleProposalProof(slashRes)
	require.Nil(t, err)
	require.Equal(t, slash.Medium, proof.GetLevel())
	require.Equal(t, []data.HeaderHandler{h1, h2, h3}, proof.GetHeaders())
}

func TestMultipleHeaderProposalProof_GetProofTxDataNotEnoughHeadersExpectError(t *testing.T) {
	proof := slash.MultipleHeaderProposalProof{}

	proofTxData, err := proof.GetProofTxData()
	require.Nil(t, proofTxData)
	require.Equal(t, data.ErrNotEnoughHeadersProvided, err)
}

func TestMultipleHeaderProposalProof_GetProofTxDataNilHeaderExpectError(t *testing.T) {
	proof := slash.MultipleHeaderProposalProof{
		HeadersV2: slash.HeadersV2{
			Headers: []*block.HeaderV2{nil},
		},
	}

	proofTxData, err := proof.GetProofTxData()
	require.Nil(t, proofTxData)
	require.Equal(t, data.ErrNilHeaderHandler, err)
}

func TestMultipleHeaderProposalProof_GetProofTxData(t *testing.T) {
	round := uint64(1)
	shardID := uint32(2)

	header := &block.HeaderV2{
		Header: &block.Header{
			Round:   round,
			ShardID: shardID,
		},
	}
	proof := slash.MultipleHeaderProposalProof{
		HeadersV2: slash.HeadersV2{
			Headers: []*block.HeaderV2{header},
		},
	}
	expectedProofTxData := &slash.ProofTxData{
		Round:   round,
		ShardID: shardID,
		ProofID: slash.MultipleProposalProofID,
	}

	proofTxData, err := proof.GetProofTxData()
	require.Equal(t, expectedProofTxData, proofTxData)
	require.Nil(t, err)
}

func TestMultipleHeaderProposalProof_MarshalUnmarshal(t *testing.T) {
	h1 := &block.HeaderV2{Header: &block.Header{TimeStamp: 1, LeaderSignature: []byte("sig1")}}
	h2 := &block.HeaderV2{Header: &block.Header{TimeStamp: 2, LeaderSignature: []byte("sig2")}}

	hInfo1 := &dataMock.HeaderInfoStub{Header: h1, Hash: []byte("h1")}
	hInfo2 := &dataMock.HeaderInfoStub{Header: h2, Hash: []byte("h2")}

	slashRes1 := &slash.SlashingResult{
		SlashingLevel: slash.Medium,
		Headers:       []data.HeaderInfoHandler{hInfo1, hInfo2},
	}
	slashRes2 := &slash.SlashingResult{
		SlashingLevel: slash.Medium,
		Headers:       []data.HeaderInfoHandler{hInfo2, hInfo1},
	}

	proof1, err1 := slash.NewMultipleProposalProof(slashRes1)
	proof2, err2 := slash.NewMultipleProposalProof(slashRes2)
	require.Nil(t, err1)
	require.Nil(t, err2)
	require.Equal(t, proof1, proof2)

	marshaller := marshal.GogoProtoMarshalizer{}
	proof1Bytes, err1 := marshaller.Marshal(proof1)
	proof2Bytes, err2 := marshaller.Marshal(proof2)
	require.Nil(t, err1)
	require.Nil(t, err2)
	require.Equal(t, proof1Bytes, proof2Bytes)

	proof1Unmarshalled := &slash.MultipleHeaderProposalProof{}
	proof2Unmarshalled := &slash.MultipleHeaderProposalProof{}
	err1 = marshaller.Unmarshal(proof1Unmarshalled, proof1Bytes)
	err2 = marshaller.Unmarshal(proof2Unmarshalled, proof2Bytes)
	require.Nil(t, err1)
	require.Nil(t, err2)
	require.Equal(t, proof1Unmarshalled, proof1)
	require.Equal(t, proof2Unmarshalled, proof2)
}
