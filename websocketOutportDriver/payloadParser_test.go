package websocketOutportDriver

import (
	"bytes"
	"errors"
	"strings"
	"testing"

	"github.com/multiversx/mx-chain-core-go/core/check"
	"github.com/multiversx/mx-chain-core-go/data/typeConverters/uint64ByteSlice"
	"github.com/multiversx/mx-chain-core-go/testscommon"
	"github.com/multiversx/mx-chain-core-go/websocketOutportDriver/data"
	"github.com/stretchr/testify/require"
)

var uint64ByteSliceConv = uint64ByteSlice.NewBigEndianConverter()

func TestNewWebSocketPayloadParser(t *testing.T) {
	t.Parallel()

	t.Run("nil uint64 byte slice converter", func(t *testing.T) {
		wpp, err := NewWebSocketPayloadParser(nil)
		require.Equal(t, data.ErrNilUint64ByteSliceConverter, err)
		require.Nil(t, wpp)
	})

	t.Run("constructor should work", func(t *testing.T) {
		wpp, err := NewWebSocketPayloadParser(uint64ByteSliceConv)
		require.NoError(t, err)
		require.False(t, check.IfNil(wpp))
	})
}

func TestWebsocketPayloadParser_ExtractPayloadData(t *testing.T) {
	t.Run("invalid payload length", testExtractPayloadDataInvalidLength)
	t.Run("invalid counter byte slice", testExtractPayloadDataInvalidCounterByteSlice)
	t.Run("invalid operation type byte slice", testExtractPayloadDataInvalidOperationTypeByteSlice)
	t.Run("invalid message counter byte slice", testExtractPayloadDataInvalidMessageCounterByteSlice)
	t.Run("invalid payload - message counter vs actual payload size", testExtractPayloadDataMessageCounterDoesNotMatchActualPayloadSize)
	t.Run("should work", testExtractPayloadDataShouldWork)
}

func testExtractPayloadDataInvalidLength(t *testing.T) {
	parser, _ := NewWebSocketPayloadParser(uint64ByteSliceConv)
	res, err := parser.ExtractPayloadData([]byte("invalid"))
	require.Nil(t, res)
	require.Error(t, err)
	require.True(t, strings.Contains(err.Error(), "invalid payload"))
}

func testExtractPayloadDataInvalidCounterByteSlice(t *testing.T) {
	localErr := errors.New("local error")
	uint64ConvStub := &testscommon.Uint64ByteSliceConverterStub{
		ToUint64Called: func(_ []byte) (uint64, error) {
			return 0, localErr
		},
	}
	parser, _ := NewWebSocketPayloadParser(uint64ConvStub)
	res, err := parser.ExtractPayloadData(bytes.Repeat([]byte{0}, minBytesForCorrectPayload))
	require.Nil(t, res)
	require.Error(t, err)
	require.True(t, errors.Is(err, localErr))
}

func testExtractPayloadDataInvalidOperationTypeByteSlice(t *testing.T) {
	localErr := errors.New("local error")
	numCalled := 0
	uint64ConvStub := &testscommon.Uint64ByteSliceConverterStub{
		ToUint64Called: func(_ []byte) (uint64, error) {
			numCalled++
			if numCalled == 2 {
				return 0, localErr
			}

			return 0, nil
		},
	}
	parser, _ := NewWebSocketPayloadParser(uint64ConvStub)
	res, err := parser.ExtractPayloadData(bytes.Repeat([]byte{0}, minBytesForCorrectPayload))
	require.Nil(t, res)
	require.Error(t, err)
	require.True(t, errors.Is(err, localErr))
}

func testExtractPayloadDataInvalidMessageCounterByteSlice(t *testing.T) {
	localErr := errors.New("local error")
	numCalled := 0
	uint64ConvStub := &testscommon.Uint64ByteSliceConverterStub{
		ToUint64Called: func(_ []byte) (uint64, error) {
			numCalled++
			if numCalled == 3 {
				return 0, localErr
			}

			return 0, nil
		},
	}
	parser, _ := NewWebSocketPayloadParser(uint64ConvStub)
	res, err := parser.ExtractPayloadData(bytes.Repeat([]byte{0}, minBytesForCorrectPayload))
	require.Nil(t, res)
	require.Error(t, err)
	require.True(t, errors.Is(err, localErr))
}

func testExtractPayloadDataMessageCounterDoesNotMatchActualPayloadSize(t *testing.T) {
	uint64ConvStub := &testscommon.Uint64ByteSliceConverterStub{
		ToUint64Called: func(_ []byte) (uint64, error) {
			return 0, nil
		},
	}
	parser, _ := NewWebSocketPayloadParser(uint64ConvStub)
	res, err := parser.ExtractPayloadData(bytes.Repeat([]byte{0}, minBytesForCorrectPayload+2))
	require.Nil(t, res)
	require.Error(t, err)
	require.True(t, strings.Contains(err.Error(), "message counter is not equal"))
}

func testExtractPayloadDataShouldWork(t *testing.T) {
	parser, _ := NewWebSocketPayloadParser(uint64ByteSliceConv)

	expectedCounter := uint64(9)
	expectedOperation := data.OperationSaveAccounts
	expectedPayload := []byte("actual payload data")

	payload := make([]byte, 1)
	payload[0] = byte(1) // with ack

	counterBytes := bytes.Repeat([]byte{0}, uint64NumBytes)
	counterBytes[uint64NumBytes-1] = byte(expectedCounter)
	payload = append(payload, counterBytes...)

	operationBytes := bytes.Repeat([]byte{0}, uint32NumBytes)
	operationBytes[uint32NumBytes-1] = byte(expectedOperation.Uint32())
	payload = append(payload, operationBytes...)

	messageLenBytes := bytes.Repeat([]byte{0}, uint32NumBytes)
	messageLenBytes[uint32NumBytes-1] = byte(len(expectedPayload))
	payload = append(payload, messageLenBytes...)

	payload = append(payload, expectedPayload...)

	res, err := parser.ExtractPayloadData(payload)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.True(t, res.WithAcknowledge)
	require.Equal(t, expectedCounter, res.Counter)
	require.Equal(t, expectedOperation, res.OperationType)
	require.Equal(t, expectedPayload, res.Payload)
}
