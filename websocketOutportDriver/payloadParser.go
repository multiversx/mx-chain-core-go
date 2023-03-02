package websocketOutportDriver

import (
	"bytes"
	"fmt"

	"github.com/multiversx/mx-chain-core-go/core/check"
	"github.com/multiversx/mx-chain-core-go/websocketOutportDriver/data"
)

const (
	withAcknowledgeNumBytes = 1
	uint64NumBytes          = 8
	uint32NumBytes          = 4
)

var (
	minBytesForCorrectPayload = withAcknowledgeNumBytes + uint64NumBytes + uint32NumBytes + uint32NumBytes
)

// PayloadData holds the arguments that should be parsed from a websocket payload
type PayloadData struct {
	WithAcknowledge bool
	Counter         uint64
	OperationType   data.OperationType
	Payload         []byte
}

type websocketPayloadParser struct {
	uint64ByteSliceConverter Uint64ByteSliceConverter
}

// NewWebSocketPayloadParser returns a new instance of websocketPayloadParser
func NewWebSocketPayloadParser(uint64ByteSliceConverter Uint64ByteSliceConverter) (*websocketPayloadParser, error) {
	if check.IfNil(uint64ByteSliceConverter) {
		return nil, data.ErrNilUint64ByteSliceConverter
	}

	return &websocketPayloadParser{
		uint64ByteSliceConverter: uint64ByteSliceConverter,
	}, nil
}

// ExtractPayloadData will extract the data from the received payload
// It should have the following form:
// first byte - with acknowledge or not
// next 8 bytes - counter (uint64 big endian)
// next 4 bytes - operation type (uint32 big endian)
// next 4 bytes - message length (uint32 big endian)
// next X bytes - the actual data to parse
func (wpp *websocketPayloadParser) ExtractPayloadData(payload []byte) (*PayloadData, error) {
	if len(payload) < minBytesForCorrectPayload {
		return nil, fmt.Errorf("invalid payload. minimum required length is %d bytes, but only provided %d",
			minBytesForCorrectPayload,
			len(payload))
	}

	var err error
	payloadData := &PayloadData{
		WithAcknowledge: false,
	}

	if payload[0] == byte(1) {
		payloadData.WithAcknowledge = true
	}
	payload = payload[withAcknowledgeNumBytes:]

	counterBytes := payload[:uint64NumBytes]
	payloadData.Counter, err = wpp.uint64ByteSliceConverter.ToUint64(counterBytes)
	if err != nil {
		return nil, fmt.Errorf("%w while extracting the counter from the payload", err)
	}
	payload = payload[uint64NumBytes:]

	operationTypeBytes := payload[:uint32NumBytes]
	var operationTypeUint64 uint64
	operationTypeUint64, err = wpp.uint64ByteSliceConverter.ToUint64(padUint32ByteSlice(operationTypeBytes))
	if err != nil {
		return nil, fmt.Errorf("%w while extracting the counter from the payload", err)
	}
	payloadData.OperationType = data.OperationTypeFromUint64(operationTypeUint64)
	payload = payload[uint32NumBytes:]

	var messageLen uint64
	messageLen, err = wpp.uint64ByteSliceConverter.ToUint64(padUint32ByteSlice(payload[:uint32NumBytes]))
	if err != nil {
		return nil, fmt.Errorf("%w while extracting the message length", err)
	}
	payload = payload[uint32NumBytes:]

	if messageLen != uint64(len(payload)) {
		return nil, fmt.Errorf("message counter is not equal to the actual payload. provided: %d, actual: %d",
			messageLen, len(payload))
	}

	payloadData.Payload = payload

	return payloadData, nil
}

func padUint32ByteSlice(initial []byte) []byte {
	padding := bytes.Repeat([]byte{0}, 4)
	return append(padding, initial...)
}
