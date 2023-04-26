package factory

import (
	"github.com/multiversx/mx-chain-core-go/core"
	"github.com/multiversx/mx-chain-core-go/core/check"
	"github.com/multiversx/mx-chain-core-go/marshal"
	"github.com/multiversx/mx-chain-core-go/webSockets"
	"github.com/multiversx/mx-chain-core-go/webSockets/client"
	"github.com/multiversx/mx-chain-core-go/webSockets/connection"
	outportData "github.com/multiversx/mx-chain-core-go/webSockets/data"
	"github.com/multiversx/mx-chain-core-go/webSockets/server"
)

// ArgsWebSocketsDriverFactory holds the arguments needed for creating a webSocketsDriverFactory
type ArgsWebSocketsDriverFactory struct {
	WebSocketConfig          outportData.WebSocketConfig
	Marshaller               marshal.Marshalizer
	Uint64ByteSliceConverter connection.Uint64ByteSliceConverter
	Log                      core.Logger
	WithAcknowledge          bool
}

type webSocketsDriverFactory struct {
	webSocketConfig          outportData.WebSocketConfig
	marshaller               marshal.Marshalizer
	uint64ByteSliceConverter connection.Uint64ByteSliceConverter
	log                      core.Logger
	withAcknowledge          bool
}

// NewWebSocketsDriverFactory will return a new instance of outportDriverWebSocketSenderFactory
func NewWebSocketsDriverFactory(args ArgsWebSocketsDriverFactory) (*webSocketsDriverFactory, error) {
	if check.IfNil(args.Marshaller) {
		return nil, outportData.ErrNilMarshaller
	}
	if check.IfNil(args.Uint64ByteSliceConverter) {
		return nil, outportData.ErrNilUint64ByteSliceConverter
	}
	if check.IfNil(args.Log) {
		return nil, outportData.ErrNilLogger
	}
	return &webSocketsDriverFactory{
		webSocketConfig:          args.WebSocketConfig,
		marshaller:               args.Marshaller,
		uint64ByteSliceConverter: args.Uint64ByteSliceConverter,
		withAcknowledge:          args.WithAcknowledge,
		log:                      args.Log,
	}, nil
}

// Create will handle the creation of all the components needed to create an outport driver that sends data over
// web socket and return it afterwards
func (o *webSocketsDriverFactory) Create() (webSockets.Driver, error) {
	var host webSockets.HostWebSockets
	var err error
	if o.webSocketConfig.IsServer {
		host, err = o.createWebSocketsServer()
	} else {
		host, err = client.NewWebSocketsClient(client.ArgsWebSocketsClient{
			RetryDurationInSeconds:   o.webSocketConfig.RetryDurationInSec,
			WithAcknowledge:          o.withAcknowledge,
			URL:                      o.webSocketConfig.URL,
			Uint64ByteSliceConverter: o.uint64ByteSliceConverter,
			Log:                      o.log,
			BlockingAckOnError:       false,
		})
	}

	if err != nil {
		return nil, err
	}

	return webSockets.NewWebsocketsDriver(
		webSockets.ArgsWebSocketsDriver{
			Marshaller:               o.marshaller,
			WebsocketSender:          host,
			Uint64ByteSliceConverter: o.uint64ByteSliceConverter,
			Log:                      o.log,
		},
	)
}

func (o *webSocketsDriverFactory) createWebSocketsServer() (webSockets.HostWebSockets, error) {
	host, err := server.NewWebSocketsServer(server.ArgsWebSocketsServer{
		RetryDurationInSeconds:   o.webSocketConfig.RetryDurationInSec,
		WithAcknowledge:          o.withAcknowledge,
		URL:                      o.webSocketConfig.URL,
		Uint64ByteSliceConverter: o.uint64ByteSliceConverter,
		Log:                      o.log,
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
