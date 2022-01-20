package slash

import (
	"strings"
	"testing"

	"github.com/ElrondNetwork/elrond-go-core/data"
	"github.com/ElrondNetwork/elrond-go-core/data/block"
	"github.com/ElrondNetwork/elrond-go-core/data/mock"
	"github.com/stretchr/testify/require"
)

func TestSortHeaders_NilHeaderInfoListExpectError(t *testing.T) {
	sortedHeaders, err := sortHeaders(nil)
	require.Nil(t, sortedHeaders)
	require.Equal(t, data.ErrEmptyHeaderInfoList, err)
}

func TestSortHeaders_EmptyHeaderInfoListExpectEmptyResult(t *testing.T) {
	sortedHeaders, err := sortHeaders([]data.HeaderInfoHandler{})
	require.Nil(t, sortedHeaders)
	require.Equal(t, data.ErrEmptyHeaderInfoList, err)
}

func TestSortHeaders_NilHeaderInfoExpectError(t *testing.T) {
	headerInfoList := []data.HeaderInfoHandler{nil}
	sortedHeaders, err := sortHeaders(headerInfoList)
	require.Nil(t, sortedHeaders)
	require.Error(t, err)
	require.True(t, strings.Contains(err.Error(), data.ErrNilHeaderInfo.Error()))
}

func TestSortHeaders_NilHeaderHandlerExpectError(t *testing.T) {
	headerInfo := &mock.HeaderInfoStub{
		Header: nil,
		Hash:   []byte("hash1"),
	}
	headerInfoList := []data.HeaderInfoHandler{headerInfo}

	sortedHeaders, err := sortHeaders(headerInfoList)
	require.Nil(t, sortedHeaders)
	require.Error(t, err)
	require.True(t, strings.Contains(err.Error(), data.ErrNilHeaderHandler.Error()))
}

func TestSortHeaders_NilHashExpectError(t *testing.T) {
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

	sortedHeaders, err := sortHeaders(headerInfoList)
	require.Nil(t, sortedHeaders)
	require.Error(t, err)
	require.True(t, strings.Contains(err.Error(), data.ErrNilHash.Error()))
}

func TestSortHeaders_HeadersSameHashExpectError(t *testing.T) {
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

	sortedHeaders, err := sortHeaders(headerInfoList)
	require.Nil(t, sortedHeaders)
	require.Error(t, err)
	require.True(t, strings.Contains(err.Error(), data.ErrHeadersSameHash.Error()))
}

func TestSortHeaders(t *testing.T) {
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
	sortedHeaders, err := sortHeaders(headersInfo)
	require.Nil(t, err)
	require.Equal(t, h1, sortedHeaders[0])
	require.Equal(t, h2, sortedHeaders[1])
	require.Equal(t, h3, sortedHeaders[2])

	require.Equal(t, hInfo1, headersInfo[0])
	require.Equal(t, hInfo2, headersInfo[1])
	require.Equal(t, hInfo3, headersInfo[2])
}
