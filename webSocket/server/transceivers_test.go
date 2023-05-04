package server

import (
	"testing"

	"github.com/multiversx/mx-chain-core-go/testscommon"
	"github.com/stretchr/testify/require"
)

func TestTransceiversHolderAddAndRemove(t *testing.T) {
	t.Parallel()

	recsHolder := newTransceiversAndConnHolder()

	recsHolder.addTransceiverAndConn(&testscommon.WebSocketTransceiverStub{}, &testscommon.WebsocketConnectionStub{
		GetIDCalled: func() string {
			return "id1"
		},
	})
	recsHolder.addTransceiverAndConn(&testscommon.WebSocketTransceiverStub{}, &testscommon.WebsocketConnectionStub{
		GetIDCalled: func() string {
			return "id2"
		},
	})

	recsHolder.remove("id1")

	allReceivers := recsHolder.getAll()
	require.Equal(t, 1, len(allReceivers))

	_, found := allReceivers["id2"]
	require.True(t, found)
}

func TestTransceiversHolderAddAndGetAll(t *testing.T) {
	t.Parallel()

	recsHolder := newTransceiversAndConnHolder()

	recsHolder.addTransceiverAndConn(&testscommon.WebSocketTransceiverStub{}, &testscommon.WebsocketConnectionStub{
		GetIDCalled: func() string {
			return "1"
		},
	})
	recsHolder.addTransceiverAndConn(&testscommon.WebSocketTransceiverStub{}, &testscommon.WebsocketConnectionStub{
		GetIDCalled: func() string {
			return "2"
		},
	})
	recsHolder.addTransceiverAndConn(&testscommon.WebSocketTransceiverStub{}, &testscommon.WebsocketConnectionStub{
		GetIDCalled: func() string {
			return "3"
		},
	})

	allReceivers := recsHolder.getAll()
	require.Equal(t, 3, len(allReceivers))

	_, found := allReceivers["1"]
	require.True(t, found)
	_, found = allReceivers["2"]
	require.True(t, found)
	_, found = allReceivers["3"]
	require.True(t, found)
}
