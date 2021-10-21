package slash_test

import (
	"testing"

	"github.com/ElrondNetwork/elrond-go-core/data"
	"github.com/ElrondNetwork/elrond-go-core/data/block"
	"github.com/ElrondNetwork/elrond-go-core/data/slash"
	"github.com/stretchr/testify/require"
)

func TestHeaders_SetHeaders_InvalidHeaders_ExpectError(t *testing.T) {
	header := &block.Header{TimeStamp: 1}
	headers := slash.Headers{}

	err := headers.SetHeaders([]data.HeaderHandler{header})
	require.Equal(t, data.ErrInvalidTypeAssertion, err)
}

func TestHeaders_SetHeaders(t *testing.T) {
	header1 := &block.HeaderV2{
		Header: &block.Header{
			TimeStamp: 1,
		},
	}
	header2 := &block.HeaderV2{
		Header: &block.Header{
			TimeStamp: 2,
		},
	}

	in := []data.HeaderHandler{header1, header2}
	headers := slash.Headers{}

	err := headers.SetHeaders(in)
	require.Nil(t, err)

	require.Len(t, headers.GetHeaders(), 2)
	require.Contains(t, headers.GetHeaders(), header1)
	require.Contains(t, headers.GetHeaders(), header2)
}
