package client

import (
	"fmt"
	"testing"

	"github.com/multiversx/mx-chain-core-go/testscommon"
	"github.com/stretchr/testify/require"
)

func TestCreateWsClient(t *testing.T) {
	t.Parallel()

	args := ArgsCreateWsClient{
		Url:                "url",
		RetryDurationInSec: 5,
		BlockingAckOnError: true,
		PayloadProcessor:   &testscommon.PayloadProcessorStub{},
	}

	wsClient, err := CreateWsClient(args)
	require.Nil(t, err)
	require.Equal(t, "*client.client", fmt.Sprintf("%T", wsClient))
}
