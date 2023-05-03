package server

import (
	"testing"

	"github.com/multiversx/mx-chain-core-go/testscommon"
	"github.com/stretchr/testify/require"
)

func TestReceiversHolderAddAndRemove(t *testing.T) {
	t.Parallel()

	recsHolder := NewTransceiversAndConnHolder()

	recsHolder.AddTransceiverAndConn(&testscommon.WebSocketReceiverStub{}, &testscommon.WebsocketConnectionStub{
		GetIDCalled: func() string {
			return "id1"
		},
	})
	recsHolder.AddTransceiverAndConn(&testscommon.WebSocketReceiverStub{}, &testscommon.WebsocketConnectionStub{
		GetIDCalled: func() string {
			return "id2"
		},
	})

	recsHolder.Remove("id1")

	allReceivers := recsHolder.GetAll()
	require.Equal(t, 1, len(allReceivers))

	_, found := allReceivers["id2"]
	require.True(t, found)
}
