package factory

import (
	"github.com/multiversx/mx-chain-core-go/core"
	"github.com/multiversx/mx-chain-core-go/marshal"
	"github.com/multiversx/mx-chain-core-go/webSocket"
	"github.com/multiversx/mx-chain-core-go/webSocket/client"
	outportData "github.com/multiversx/mx-chain-core-go/webSocket/data"
	"github.com/multiversx/mx-chain-core-go/webSocket/server"
)

// ArgsWebSocketDriverFactory holds the arguments needed for creating a webSocketsDriverFactory
type ArgsWebSocketDriverFactory struct {
	WebSocketConfig          outportData.WebSocketConfig
	Marshaller               marshal.Marshalizer
	Uint64ByteSliceConverter webSocket.Uint64ByteSliceConverter
	Log                      core.Logger
	WithAcknowledge          bool
}

// NewWebSocketDriver will handle the creation of all the components needed to create an outport driver that sends data over WebSocket
func NewWebSocketDriver(args ArgsWebSocketDriverFactory) (webSocket.Driver, error) {
	var host webSocket.HostWebSocket
	var err error
	if args.WebSocketConfig.IsServer {
		host, err = createWebSocketServer(args)
	} else {
		host, err = createWebSocketClient(args)
	}

	if err != nil {
		return nil, err
	}

	host.Start()

	return webSocket.NewWebsocketDriver(
		webSocket.ArgsWebSocketDriver{
			Marshaller:      args.Marshaller,
			WebsocketSender: host,
			Log:             args.Log,
		},
	)
}

// TODO merge the ArgsWebSocketClient and ArgsWebSocketServer as they look the same and remove the duplicated arguments build
func createWebSocketClient(args ArgsWebSocketDriverFactory) (webSocket.HostWebSocket, error) {
	payloadConverter, err := webSocket.NewWebSocketPayloadConverter(args.Uint64ByteSliceConverter)
	if err != nil {
		return nil, err
	}

	return client.NewWebSocketClient(client.ArgsWebSocketClient{
		RetryDurationInSeconds: args.WebSocketConfig.RetryDurationInSec,
		WithAcknowledge:        args.WithAcknowledge,
		URL:                    args.WebSocketConfig.URL,
		PayloadConverter:       payloadConverter,
		Log:                    args.Log,
		BlockingAckOnError:     false,
	})
}

func createWebSocketServer(args ArgsWebSocketDriverFactory) (webSocket.HostWebSocket, error) {
	payloadConverter, err := webSocket.NewWebSocketPayloadConverter(args.Uint64ByteSliceConverter)
	if err != nil {
		return nil, err
	}

	host, err := server.NewWebSocketServer(server.ArgsWebSocketServer{
		RetryDurationInSeconds: args.WebSocketConfig.RetryDurationInSec,
		WithAcknowledge:        args.WithAcknowledge,
		URL:                    args.WebSocketConfig.URL,
		PayloadConverter:       payloadConverter,
		Log:                    args.Log,
		BlockingAckOnError:     false,
	})
	if err != nil {
		return nil, err
	}

	go func() {
		host.Listen()
	}()

	return host, nil
}
