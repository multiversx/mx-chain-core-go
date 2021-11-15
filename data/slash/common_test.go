package slash

import (
	"testing"

	"github.com/ElrondNetwork/elrond-go-core/data"
	"github.com/ElrondNetwork/elrond-go-core/data/block"
	"github.com/ElrondNetwork/elrond-go-core/data/mock"
	"github.com/stretchr/testify/require"
)

func TestGetSortedHeadersV2_NilHeaderInfoList_ExpectError(t *testing.T) {
	sortedHeaders, err := getSortedHeadersV2(nil)
	require.Equal(t, HeadersV2{}, sortedHeaders)
	require.Equal(t, data.ErrNilHeaderInfoList, err)
}

func TestGetSortedHeadersV2_NilHeaderInfo_ExpectError(t *testing.T) {
	headerInfoList := []data.HeaderInfoHandler{nil}
	sortedHeaders, err := getSortedHeadersV2(headerInfoList)
	require.Equal(t, HeadersV2{}, sortedHeaders)
	require.Equal(t, data.ErrNilHeaderInfo, err)
}

func TestGetSortedHeadersV2_EmptyHeaderInfoList_ExpectEmptyResult(t *testing.T) {
	sortedHeaders, err := getSortedHeadersV2([]data.HeaderInfoHandler{})
	require.Equal(t, HeadersV2{}, sortedHeaders)
	require.Nil(t, err)
}

func TestGetSortedHeadersV2_NilHeaderHandler_ExpectError(t *testing.T) {
	headerInfo := &mock.HeaderInfoStub{
		Header: nil,
		Hash:   []byte("hash1"),
	}
	headerInfoList := []data.HeaderInfoHandler{headerInfo}

	sortedHeaders, err := getSortedHeadersV2(headerInfoList)
	require.Equal(t, HeadersV2{}, sortedHeaders)
	require.Equal(t, data.ErrNilHeaderHandler, err)
}

func TestGetSortedHeadersV2_NilHash_ExpectError(t *testing.T) {
	header := &block.HeaderV2{Header: &block.Header{TimeStamp: 1}}
	headerInfo := &mock.HeaderInfoStub{
		Header: header,
		Hash:   nil,
	}
	headerInfoList := []data.HeaderInfoHandler{headerInfo}

	sortedHeaders, err := getSortedHeadersV2(headerInfoList)
	require.Equal(t, HeadersV2{}, sortedHeaders)
	require.Equal(t, data.ErrNilHash, err)
}

func TestGetSortedHeadersV2(t *testing.T) {
	h1 := &block.HeaderV2{Header: &block.Header{TimeStamp: 1}}
	h2 := &block.HeaderV2{Header: &block.Header{TimeStamp: 2}}
	h3 := &block.HeaderV2{Header: &block.Header{TimeStamp: 3}}

	hInfo1 := &mock.HeaderInfoStub{
		Header: h1,
		Hash:   []byte("h1"),
	}
	hInfo2 := &mock.HeaderInfoStub{
		Header: h2,
		Hash:   []byte("h2"),
	}
	hInfo3 := &mock.HeaderInfoStub{
		Header: h3,
		Hash:   []byte("h3"),
	}

	headersInfo := []data.HeaderInfoHandler{hInfo3, hInfo1, hInfo2}
	sortedHeaders, _ := getSortedHeadersV2(headersInfo)

	require.Equal(t, h1, sortedHeaders.Headers[0])
	require.Equal(t, h2, sortedHeaders.Headers[1])
	require.Equal(t, h3, sortedHeaders.Headers[2])
}
