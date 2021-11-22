package slash

import (
	"testing"

	"github.com/ElrondNetwork/elrond-go-core/data"
	"github.com/ElrondNetwork/elrond-go-core/data/block"
	"github.com/ElrondNetwork/elrond-go-core/data/mock"
	"github.com/stretchr/testify/require"
)

func TestGetSortedHeadersV2_NilHeaderInfoListExpectError(t *testing.T) {
	sortedHeaders, err := getSortedHeaders(nil)
	require.Nil(t, sortedHeaders)
	require.Equal(t, data.ErrNilHeaderInfoList, err)
}

func TestGetSortedHeadersV2_NilHeaderInfoExpectError(t *testing.T) {
	headerInfoList := []data.HeaderInfoHandler{nil}
	sortedHeaders, err := getSortedHeaders(headerInfoList)
	require.Nil(t, sortedHeaders)
	require.Equal(t, data.ErrNilHeaderInfo, err)
}

func TestGetSortedHeadersV2_EmptyHeaderInfoListExpectEmptyResult(t *testing.T) {
	sortedHeaders, err := getSortedHeaders([]data.HeaderInfoHandler{})
	require.Len(t, sortedHeaders, 0)
	require.Nil(t, err)
}

func TestGetSortedHeadersV2_NilHeaderHandlerExpectError(t *testing.T) {
	headerInfo := &mock.HeaderInfoStub{
		Header: nil,
		Hash:   []byte("hash1"),
	}
	headerInfoList := []data.HeaderInfoHandler{headerInfo}

	sortedHeaders, err := getSortedHeaders(headerInfoList)
	require.Nil(t, sortedHeaders)
	require.Equal(t, data.ErrNilHeaderHandler, err)
}

func TestGetSortedHeadersV2_HeadersSameHash_ExpectError(t *testing.T) {
	header1 := &block.HeaderV2{Header: &block.Header{TimeStamp: 1}}
	header2 := &block.HeaderV2{Header: &block.Header{TimeStamp: 2}}
	header3 := &block.HeaderV2{Header: &block.Header{TimeStamp: 3}}
	headerInfo1 := &mock.HeaderInfoStub{
		Header: header1,
		Hash:   []byte("hash1"),
	}
	headerInfo2 := &mock.HeaderInfoStub{
		Header: header2,
		Hash:   []byte("hash2"),
	}
	headerInfo3 := &mock.HeaderInfoStub{
		Header: header3,
		Hash:   []byte("hash2"),
	}
	headerInfoList := []data.HeaderInfoHandler{headerInfo1, headerInfo2, headerInfo3}

	sortedHeaders, err := getSortedHeaders(headerInfoList)
	require.Nil(t, sortedHeaders)
	require.Equal(t, data.ErrHeadersSameHash, err)
}

func TestGetSortedHeadersV2_NilHashExpectError(t *testing.T) {
	header1 := &block.HeaderV2{Header: &block.Header{TimeStamp: 1}}
	header2 := &block.HeaderV2{Header: &block.Header{TimeStamp: 2}}
	headerInfo1 := &mock.HeaderInfoStub{
		Header: header1,
		Hash:   []byte("hash"),
	}
	headerInfo2 := &mock.HeaderInfoStub{
		Header: header2,
		Hash:   nil,
	}
	headerInfoList := []data.HeaderInfoHandler{headerInfo1, headerInfo2}

	sortedHeaders, err := getSortedHeaders(headerInfoList)
	require.Nil(t, sortedHeaders)
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
	sortedHeaders, err := getSortedHeaders(headersInfo)
	require.Nil(t, err)

	require.Equal(t, h1, sortedHeaders[0])
	require.Equal(t, h2, sortedHeaders[1])
	require.Equal(t, h3, sortedHeaders[2])
}
