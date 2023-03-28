package client

import (
	"fmt"
	"testing"
	"time"

	"github.com/multiversx/mx-chain-core-go/core/atomic"
	"github.com/multiversx/mx-chain-core-go/core/closing"
	"github.com/multiversx/mx-chain-core-go/testscommon"
	"github.com/stretchr/testify/require"
)

func TestClient_Start(t *testing.T) {
	t.Parallel()

	openConnectionCalledCt := &atomic.Counter{}
	args := ArgsWsClient{
		Url:                      "url",
		RetryDurationInSec:       1,
		BlockingAckOnError:       false,
		PayloadProcessor:         &testscommon.PayloadProcessorStub{},
		PayloadParser:            &testscommon.PayloadParserStub{},
		Uint64ByteSliceConverter: &testscommon.Uint64ByteSliceConverterStub{},
		WSConnClient: &testscommon.WebsocketConnectionStub{
			OpenConnectionCalled: func(url string) error {
				openConnectionCalledCt.Increment()
				return fmt.Errorf("dsa")
			},
		},
		SafeCloser: closing.NewSafeChanCloser(),
	}
	wsClient, _ := NewWsClientHandler(args)

	go wsClient.Start()

	time.Sleep(2 * time.Second)
	args.SafeCloser.Close()

	time.Sleep(1 * time.Second)
	require.Equal(t, int64(3), openConnectionCalledCt.Get())
}
