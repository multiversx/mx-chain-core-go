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
			args:        &slash.SlashingResult{SlashingLevel: slash.Medium, Headers: nil},
			expectedErr: data.ErrNilHeaderInfoList,
		},
		{
			args:        &slash.SlashingResult{SlashingLevel: slash.Medium, Headers: []data.HeaderInfoHandler{nil}},
			expectedErr: data.ErrNilHeaderInfo,
		},
		{
			args:        &slash.SlashingResult{SlashingLevel: slash.Medium, Headers: []data.HeaderInfoHandler{}},
			expectedErr: nil,
		},
	}

	for _, currTest := range tests {
		_, err := slash.NewMultipleProposalProof(currTest.args)
		require.Equal(t, currTest.expectedErr, err)
	}
}

func TestMultipleProposalProof_GetHeaders_GetLevel_GetType(t *testing.T) {
	h1 := &block.HeaderV2{Header: &block.Header{TimeStamp: 1}}
	h2 := &block.HeaderV2{Header: &block.Header{TimeStamp: 2}}

	hInfo1 := &dataMock.HeaderInfoStub{Header: h1, Hash: []byte("h1")}
	hInfo2 := &dataMock.HeaderInfoStub{Header: h2, Hash: []byte("h2")}

	slashRes := &slash.SlashingResult{
		SlashingLevel: slash.Medium,
		Headers:       []data.HeaderInfoHandler{hInfo1, hInfo2},
	}

	proof, err := slash.NewMultipleProposalProof(slashRes)
	require.Nil(t, err)

	require.Equal(t, slash.MultipleProposal, proof.GetType())
	require.Equal(t, slash.Medium, proof.GetLevel())

	require.Len(t, proof.GetHeaders(), 2)
	require.Equal(t, proof.GetHeaders()[0], h1)
	require.Equal(t, proof.GetHeaders()[1], h2)
}

func TestMultipleHeaderProposalProof_Marshal_Unmarshal(t *testing.T) {
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
