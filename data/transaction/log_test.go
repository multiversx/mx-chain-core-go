package transaction_test

import (
	"testing"

	"github.com/multiversx/mx-chain-core-go/core/check"
	"github.com/multiversx/mx-chain-core-go/data/transaction"
	"github.com/stretchr/testify/require"
)

func TestLog_SettersAndGetters(t *testing.T) {
	t.Parallel()
	address := []byte("address")

	log := &transaction.Log{
		Address: address,
	}

	require.False(t, check.IfNil(log))
	require.Equal(t, address, log.GetAddress())
}

func TestLog_GetEvents(t *testing.T) {
	t.Parallel()
	address := []byte("address")
	evIdentifier := []byte("identifier")

	events := make([]*transaction.Event, 1)
	events[0] = &transaction.Event{
		Identifier: evIdentifier,
	}

	log := &transaction.Log{
		Address: address,
		Events:  events,
	}

	logEvents := log.GetLogEvents()
	require.Equal(t, len(events), len(logEvents))
	require.Equal(t, evIdentifier, logEvents[0].GetIdentifier())
}
