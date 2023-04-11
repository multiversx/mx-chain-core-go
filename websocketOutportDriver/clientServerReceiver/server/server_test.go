package server

import (
	"testing"

	"github.com/multiversx/mx-chain-core-go/core/mock"
	"github.com/multiversx/mx-chain-core-go/testscommon"
	"github.com/multiversx/mx-chain-core-go/websocketOutportDriver/data"
	"github.com/stretchr/testify/require"
)

func createArgs() ArgsWsServer {
	return ArgsWsServer{
		URL:                      "url",
		RetryDurationInSec:       1,
		BlockingAckOnError:       false,
		PayloadProcessor:         &testscommon.PayloadProcessorStub{},
		PayloadParser:            &testscommon.PayloadParserStub{},
		Uint64ByteSliceConverter: &testscommon.Uint64ByteSliceConverterStub{},
		Log:                      &mock.LoggerMock{},
	}
}

func TestNewWsServer(t *testing.T) {
	t.Parallel()

	t.Run("should work", func(t *testing.T) {
		args := createArgs()
		ws, err := NewWsServer(args)
		require.NotNil(t, ws)
		require.Nil(t, err)
	})

	t.Run("nil payload parser, should return error", func(t *testing.T) {
		args := createArgs()
		args.PayloadParser = nil
		ws, err := NewWsServer(args)
		require.Nil(t, ws)
		require.Equal(t, err, data.ErrNilPayloadParser)
	})

	t.Run("nil payload processor, should return error", func(t *testing.T) {
		args := createArgs()
		args.PayloadProcessor = nil
		ws, err := NewWsServer(args)
		require.Nil(t, ws)
		require.Equal(t, err, data.ErrNilPayloadProcessor)
	})

	t.Run("nil uint64 byte slice converter, should return error", func(t *testing.T) {
		args := createArgs()
		args.Uint64ByteSliceConverter = nil
		ws, err := NewWsServer(args)
		require.Nil(t, ws)
		require.Equal(t, err, data.ErrNilUint64ByteSliceConverter)
	})
	t.Run("nil logger, should return error", func(t *testing.T) {
		args := createArgs()
		args.Log = nil
		ws, err := NewWsServer(args)
		require.Nil(t, ws)
		require.Equal(t, err, data.ErrNilLogger)
	})

	t.Run("empty url, should return error", func(t *testing.T) {
		args := createArgs()
		args.URL = ""
		ws, err := NewWsServer(args)
		require.Nil(t, ws)
		require.Equal(t, err, data.ErrEmptyUrl)
	})

	t.Run("zero value retry duration, should return error", func(t *testing.T) {
		args := createArgs()
		args.RetryDurationInSec = 0
		ws, err := NewWsServer(args)
		require.Nil(t, ws)
		require.Equal(t, err, data.ErrZeroValueRetryDuration)
	})
}
