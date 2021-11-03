package slash_test

import (
	"fmt"
	"testing"

	"github.com/ElrondNetwork/elrond-go-core/data"
	"github.com/ElrondNetwork/elrond-go-core/data/block"
	"github.com/ElrondNetwork/elrond-go-core/data/slash"
	"github.com/stretchr/testify/require"
)

func TestNewMultipleProposalProof(t *testing.T) {
	tests := []struct {
		args        func() *slash.SlashingResult
		expectedErr error
	}{
		{
			args: func() *slash.SlashingResult {
				return nil
			},
			expectedErr: data.ErrNilSlashResult,
		},
		{
			args: func() *slash.SlashingResult {
				return &slash.SlashingResult{SlashingLevel: slash.Medium, Headers: nil}
			},
			expectedErr: data.ErrNilHeaderHandler,
		},
		{
			args: func() *slash.SlashingResult {
				return &slash.SlashingResult{SlashingLevel: slash.Medium, Headers: []data.HeaderHandler{}}
			},
			expectedErr: nil,
		},
	}

	for _, currTest := range tests {
		_, err := slash.NewMultipleProposalProof(currTest.args())
		require.Equal(t, currTest.expectedErr, err)
	}
}

func TestMultipleProposalProof_GetHeaders_GetLevel_GetType(t *testing.T) {
	h1 := &block.HeaderV2{Header: &block.Header{TimeStamp: 1}}
	h2 := &block.HeaderV2{Header: &block.Header{TimeStamp: 2}}

	slashRes := &slash.SlashingResult{
		SlashingLevel: slash.Medium,
		Headers:       []data.HeaderHandler{h1, h2}}

	proof, err := slash.NewMultipleProposalProof(slashRes)
	require.Nil(t, err)

	require.Equal(t, slash.MultipleProposal, proof.GetType())
	require.Equal(t, slash.Medium, proof.GetLevel())

	x := proof.GetHeaders()
	fmt.Println(len(x))
	require.Len(t, proof.GetHeaders(), 2)
	require.Contains(t, proof.GetHeaders(), h1)
	require.Contains(t, proof.GetHeaders(), h2)
}
