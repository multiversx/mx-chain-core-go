package webSocket

import (
	"testing"

	"github.com/multiversx/mx-chain-core-go/core/mock"
	"github.com/multiversx/mx-chain-core-go/webSocket/data"
	"github.com/stretchr/testify/require"
)

func TestNewWebSocketPayloadConverter(t *testing.T) {
	t.Parallel()

	payloadConverter, err := NewWebSocketPayloadConverter(nil)
	require.Nil(t, payloadConverter)
	require.Equal(t, data.ErrNilMarshaller, err)

	payloadConverter, _ = NewWebSocketPayloadConverter(&mock.MarshalizerMock{})
	require.NotNil(t, payloadConverter)
	require.False(t, payloadConverter.IsInterfaceNil())
}

func TestWebSocketsPayloadConverter_ConstructPayload(t *testing.T) {
	t.Parallel()

	payloadConverter, _ := NewWebSocketPayloadConverter(&mock.MarshalizerMock{})

	wsMessage := &data.WsMessage{
		WithAcknowledge: true,
		PayloadData: &data.PayloadData{
			Payload:       []byte("test"),
			OperationType: data.OperationSaveAccounts.Uint32(),
		},
		Counter:     10,
		MessageType: data.PayloadMessage,
	}

	payload, err := payloadConverter.ConstructPayload(wsMessage)
	require.Nil(t, err)

	newWsMessage, err := payloadConverter.ExtractWsMessage(payload)
	require.Nil(t, err)
	require.Equal(t, wsMessage, newWsMessage)
}
