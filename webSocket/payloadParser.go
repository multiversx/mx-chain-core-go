package webSocket

import (
	"bytes"
	"fmt"

	"github.com/multiversx/mx-chain-core-go/core/check"
	"github.com/multiversx/mx-chain-core-go/webSocket/data"
)

const (
	prefixAckMessageLen     = 5
	withAcknowledgeNumBytes = 1
	uint64NumBytes          = 8
	uint32NumBytes          = 4
)

var (
	minBytesForCorrectPayload = withAcknowledgeNumBytes + uint64NumBytes + uint32NumBytes + uint32NumBytes

	prefixPayloadWithoutAck = []byte{0}
	prefixPayloadWithAck    = []byte{1}
	prefixAckMessage        = []byte("#ack_")
)

type webSocketsPayloadConverter struct {
	uint64ByteSliceConverter Uint64ByteSliceConverter
}

// NewWebSocketPayloadConverter returns a new instance of websocketPayloadParser
func NewWebSocketPayloadConverter(uint64ByteSliceConverter Uint64ByteSliceConverter) (*webSocketsPayloadConverter, error) {
	if check.IfNil(uint64ByteSliceConverter) {
		return nil, data.ErrNilUint64ByteSliceConverter
	}

	return &webSocketsPayloadConverter{
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
func (wpc *webSocketsPayloadConverter) ExtractPayloadData(payload []byte) (*data.PayloadData, error) {
	if len(payload) < minBytesForCorrectPayload {
		return nil, fmt.Errorf("invalid payload. minimum required length is %d bytes, but only provided %d",
			minBytesForCorrectPayload,
			len(payload))
	}

	var err error
	payloadData := &data.PayloadData{
		WithAcknowledge: false,
	}

	if payload[0] == byte(1) {
		payloadData.WithAcknowledge = true
	}
	payload = payload[withAcknowledgeNumBytes:]

	counterBytes := payload[:uint64NumBytes]
	payloadData.Counter, err = wpc.uint64ByteSliceConverter.ToUint64(counterBytes)
	if err != nil {
		return nil, fmt.Errorf("%w while extracting the counter from the payload", err)
	}
	payload = payload[uint64NumBytes:]

	operationTypeBytes := payload[:uint32NumBytes]
	var operationTypeUint64 uint64
	operationTypeUint64, err = wpc.uint64ByteSliceConverter.ToUint64(padUint32ByteSlice(operationTypeBytes))
	if err != nil {
		return nil, fmt.Errorf("%w while extracting the counter from the payload", err)
	}
	payloadData.OperationType = data.OperationTypeFromUint64(operationTypeUint64)
	payload = payload[uint32NumBytes:]

	var messageLen uint64
	messageLen, err = wpc.uint64ByteSliceConverter.ToUint64(padUint32ByteSlice(payload[:uint32NumBytes]))
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

// ConstructPayloadData will construct the payload data
func (wpc *webSocketsPayloadConverter) ConstructPayloadData(args data.WsSendArgs, counter uint64, withAcknowledge bool) []byte {
	opBytes := wpc.encodeUint64(uint64(args.OpType.Uint32()))
	opBytes = opBytes[uint32NumBytes:]

	messageLength := uint64(len(args.Payload))
	messageLengthBytes := wpc.uint64ByteSliceConverter.ToByteSlice(messageLength)
	messageLengthBytes = messageLengthBytes[uint32NumBytes:]

	newPayload := append(opBytes, messageLengthBytes...)
	newPayload = append(newPayload, args.Payload...)

	ackData := prefixPayloadWithoutAck
	if withAcknowledge {
		ackData = prefixPayloadWithAck
	}

	payloadWithCounter := append(ackData, wpc.encodeUint64(counter)...)
	payloadWithCounter = append(payloadWithCounter, newPayload...)
	return payloadWithCounter
}

// EncodeUint64 will encode the provided counter in a byte array
func (wpc *webSocketsPayloadConverter) encodeUint64(counter uint64) []byte {
	return wpc.uint64ByteSliceConverter.ToByteSlice(counter)
}

// PrepareUint64Ack will prepare the provided uint64 value in a ack message
func (wpc *webSocketsPayloadConverter) PrepareUint64Ack(counter uint64) []byte {
	counterBytes := wpc.encodeUint64(counter)

	return append(prefixAckMessage, counterBytes...)
}

// IsAckPayload will return true if the provided payload contains an ack message
func (wpc *webSocketsPayloadConverter) IsAckPayload(payload []byte) bool {
	if len(payload) < prefixAckMessageLen+1 {
		return false
	}

	prefixAckMessageFromPayload := payload[:prefixAckMessageLen]

	return bytes.Equal(prefixAckMessageFromPayload, prefixAckMessage)
}

// ExtractUint64FromAckMessage will decode the provided payload in an uint64 value
func (wpc *webSocketsPayloadConverter) ExtractUint64FromAckMessage(payload []byte) (uint64, error) {
	if len(payload) < prefixAckMessageLen+1 {
		return 0, data.ErrInvalidPayloadForAckMessage
	}

	counterBytes := payload[prefixAckMessageLen:]

	return wpc.uint64ByteSliceConverter.ToUint64(counterBytes)
}

func padUint32ByteSlice(initial []byte) []byte {
	padding := bytes.Repeat([]byte{0}, 4)
	return append(padding, initial...)
}

// IsInterfaceNil -
func (wpc *webSocketsPayloadConverter) IsInterfaceNil() bool {
	return wpc == nil
}
