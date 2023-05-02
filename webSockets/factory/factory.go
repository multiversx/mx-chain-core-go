package factory

import (
	"github.com/multiversx/mx-chain-core-go/core"
	"github.com/multiversx/mx-chain-core-go/marshal"
	"github.com/multiversx/mx-chain-core-go/webSockets"
	"github.com/multiversx/mx-chain-core-go/webSockets/client"
	outportData "github.com/multiversx/mx-chain-core-go/webSockets/data"
	"github.com/multiversx/mx-chain-core-go/webSockets/server"
)

// ArgsWebSocketsDriverFactory holds the arguments needed for creating a webSocketsDriverFactory
type ArgsWebSocketsDriverFactory struct {
	WebSocketConfig          outportData.WebSocketConfig
	Marshaller               marshal.Marshalizer
	Uint64ByteSliceConverter webSockets.Uint64ByteSliceConverter
	Log                      core.Logger
	WithAcknowledge          bool
}

// NewWebSocketsDriver will handle the creation of all the components needed to create an outport driver that sends data over
func NewWebSocketsDriver(args ArgsWebSocketsDriverFactory) (webSockets.Driver, error) {
	var host webSockets.HostWebSockets
	var err error
	if args.WebSocketConfig.IsServer {
		host, err = createWebSocketsServer(args)
	} else {
		host, err = createWebSocketsClient(args)
	}

	if err != nil {
		return nil, err
	}

	host.Start()

	return webSockets.NewWebsocketsDriver(
		webSockets.ArgsWebSocketsDriver{
			Marshaller:               args.Marshaller,
			WebsocketSender:          host,
			Uint64ByteSliceConverter: args.Uint64ByteSliceConverter,
			Log:                      args.Log,
		},
	)
}

func createWebSocketsClient(args ArgsWebSocketsDriverFactory) (webSockets.HostWebSockets, error) {
	return client.NewWebSocketsClient(client.ArgsWebSocketsClient{
		RetryDurationInSeconds:   args.WebSocketConfig.RetryDurationInSec,
		WithAcknowledge:          args.WithAcknowledge,
		URL:                      args.WebSocketConfig.URL,
		Uint64ByteSliceConverter: args.Uint64ByteSliceConverter,
		Log:                      args.Log,
		BlockingAckOnError:       false,
	})
}

func createWebSocketsServer(args ArgsWebSocketsDriverFactory) (webSockets.HostWebSockets, error) {
	host, err := server.NewWebSocketsServer(server.ArgsWebSocketsServer{
		RetryDurationInSeconds:   args.WebSocketConfig.RetryDurationInSec,
		WithAcknowledge:          args.WithAcknowledge,
		URL:                      args.WebSocketConfig.URL,
		Uint64ByteSliceConverter: args.Uint64ByteSliceConverter,
		Log:                      args.Log,
		BlockingAckOnError:       false,
	})
	if err != nil {
		return nil, err
	}

	go func() {
		host.Listen()
	}()

	return host, nil
}
