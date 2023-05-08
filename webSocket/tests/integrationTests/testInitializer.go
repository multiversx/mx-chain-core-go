package integrationTests

import (
	"github.com/multiversx/mx-chain-core-go/core"
	"github.com/multiversx/mx-chain-core-go/marshal/factory"
	"github.com/multiversx/mx-chain-core-go/webSocket"
	"github.com/multiversx/mx-chain-core-go/webSocket/client"
	"github.com/multiversx/mx-chain-core-go/webSocket/server"
)

const retryDurationInSeconds = 1

var (
	marshaller, _       = factory.NewMarshalizer("gogo protobuf")
	payloadConverter, _ = webSocket.NewWebSocketPayloadConverter(marshaller)
)

func createClient(url string, log core.Logger) (webSocket.HostWebSocket, error) {
	return client.NewWebSocketClient(client.ArgsWebSocketClient{
		RetryDurationInSeconds: retryDurationInSeconds,
		WithAcknowledge:        true,
		URL:                    url,
		PayloadConverter:       payloadConverter,
		Log:                    log,
	})
}

func createServer(url string, log core.Logger) (webSocket.HostWebSocket, error) {
	return server.NewWebSocketServer(server.ArgsWebSocketServer{
		RetryDurationInSeconds: retryDurationInSeconds,
		WithAcknowledge:        true,
		URL:                    url,
		PayloadConverter:       payloadConverter,
		Log:                    log,
	})
}
