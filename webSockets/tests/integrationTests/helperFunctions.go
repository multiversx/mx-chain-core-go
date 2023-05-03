package integrationTests

import (
	"github.com/multiversx/mx-chain-core-go/core"
	"github.com/multiversx/mx-chain-core-go/core/mock"
	"github.com/multiversx/mx-chain-core-go/data/typeConverters/uint64ByteSlice"
	"github.com/multiversx/mx-chain-core-go/webSockets"
	"github.com/multiversx/mx-chain-core-go/webSockets/client"
	"github.com/multiversx/mx-chain-core-go/webSockets/server"
)

const retryDurationInSeconds = 1

var (
	uint64Converter     = uint64ByteSlice.NewBigEndianConverter()
	payloadConverter, _ = webSockets.NewWebSocketPayloadConverter(uint64Converter)
)

func createClient(url string) (webSockets.HostWebSockets, error) {

	return client.NewWebSocketsClient(client.ArgsWebSocketsClient{
		RetryDurationInSeconds: retryDurationInSeconds,
		WithAcknowledge:        true,
		URL:                    url,
		PayloadConverter:       payloadConverter,
		Log:                    &mock.LoggerMock{},
	})
}

func createServer(url string, log core.Logger) (webSockets.HostWebSockets, error) {
	return server.NewWebSocketsServer(server.ArgsWebSocketsServer{
		RetryDurationInSeconds: retryDurationInSeconds,
		WithAcknowledge:        true,
		URL:                    url,
		PayloadConverter:       payloadConverter,
		Log:                    log,
	})
}
