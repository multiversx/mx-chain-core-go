package server

import (
	"testing"

	"github.com/multiversx/mx-chain-core-go/testscommon"
	"github.com/stretchr/testify/require"
)

func TestReceiversHolderAddAndRemove(t *testing.T) {
	t.Parallel()

	recsHolder := NewReceiversHolder()

	recsHolder.AddReceiver("id1", &testscommon.WebSocketReceiverStub{})
	recsHolder.AddReceiver("id2", &testscommon.WebSocketReceiverStub{})

	recsHolder.RemoveReceiver("id1")

	allReceivers := recsHolder.GetAll()
	require.Equal(t, 1, len(allReceivers))

	_, found := allReceivers["id2"]
	require.True(t, found)
}
