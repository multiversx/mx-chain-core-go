package slash_test

import (
	"testing"

	"github.com/ElrondNetwork/elrond-go-core/data"
	"github.com/ElrondNetwork/elrond-go-core/data/block"
	"github.com/ElrondNetwork/elrond-go-core/data/slash"
	"github.com/stretchr/testify/require"
)

func TestHeaders_SetHeadersInvalidHeadersExpectError(t *testing.T) {
	header := &block.Header{TimeStamp: 1}
	headers := slash.HeadersV2{}

	err := headers.SetHeaders([]data.HeaderHandler{header})
	require.Equal(t, data.ErrInvalidTypeAssertion, err)
}

func TestHeaders_SetHeadersGetHeadersGetHeaderHandlers(t *testing.T) {
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

	headers := []data.HeaderHandler{header1, header2}
	headersV2 := slash.HeadersV2{}

	err := headersV2.SetHeaders(headers)
	require.Nil(t, err)

	require.Len(t, headersV2.GetHeaders(), 2)
	require.Contains(t, headersV2.GetHeaders(), header1)
	require.Contains(t, headersV2.GetHeaders(), header2)

	require.Len(t, headersV2.GetHeaderHandlers(), 2)
	require.Contains(t, headersV2.GetHeaderHandlers(), header1)
	require.Contains(t, headersV2.GetHeaderHandlers(), header2)
}
