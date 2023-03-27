package client

import (
	"fmt"
	"testing"
	"time"

	"github.com/multiversx/mx-chain-core-go/core/atomic"
	"github.com/multiversx/mx-chain-core-go/core/closing"
	"github.com/multiversx/mx-chain-core-go/data/mock"
	mock2 "github.com/multiversx/mx-chain-core-go/websocketOutportDriver/mock"
	"github.com/stretchr/testify/require"
)

func TestClient_Start(t *testing.T) {
	t.Parallel()

	openConnectionCalledCt := &atomic.Counter{}
	args := ArgsWsClient{
		Url:                      "url",
		RetryDurationInSec:       1,
		BlockingAckOnError:       false,
		PayloadProcessor:         &mock.PayloadProcessorStub{},
		PayloadParser:            &mock.PayloadParserStub{},
		Uint64ByteSliceConverter: &mock2.Uint64ByteSliceConverterStub{},
		WSConnClient: &mock2.WebsocketConnectionStub{
			OpenConnectionCalled: func(url string) error {
				openConnectionCalledCt.Increment()
				return fmt.Errorf("dsa")
			},
		},
		SafeCloser: closing.NewSafeChanCloser(),
	}
	wsClient, _ := NewWsClientHandler(args)

	go wsClient.Start()

	time.Sleep(3 * time.Second)
	args.SafeCloser.Close()

	time.Sleep(1 * time.Second)
	require.Equal(t, int64(4), openConnectionCalledCt.Get())
}
