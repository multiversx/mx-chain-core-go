package factory

import (
	"github.com/multiversx/mx-chain-core-go/core"
	"github.com/multiversx/mx-chain-core-go/core/check"
	"github.com/multiversx/mx-chain-core-go/marshal"
	"github.com/multiversx/mx-chain-core-go/webSockets"
	"github.com/multiversx/mx-chain-core-go/webSockets/clientServerSender"
	"github.com/multiversx/mx-chain-core-go/webSockets/common"
	outportData "github.com/multiversx/mx-chain-core-go/webSockets/data"
)

// ArgsWebSocketsDriverFactory holds the arguments needed for creating a webSocketsDriverFactory
type ArgsWebSocketsDriverFactory struct {
	WebSocketConfig          outportData.WebSocketConfig
	Marshaller               marshal.Marshalizer
	Uint64ByteSliceConverter common.Uint64ByteSliceConverter
	Log                      core.Logger
	WithAcknowledge          bool
}

type webSocketsDriverFactory struct {
	webSocketConfig          outportData.WebSocketConfig
	marshaller               marshal.Marshalizer
	uint64ByteSliceConverter common.Uint64ByteSliceConverter
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
	webSocketSender, err := clientServerSender.NewClientServerSender(clientServerSender.ArgsWSClientServerSender{
		Url:                      o.webSocketConfig.URL,
		IsServer:                 o.webSocketConfig.IsServer,
		Uint64ByteSliceConverter: o.uint64ByteSliceConverter,
		RetryDurationInSec:       o.webSocketConfig.RetryDurationInSec,
		WithAcknowledge:          o.withAcknowledge,
		Log:                      o.log,
	})
	if err != nil {
		return nil, err
	}

	return webSockets.NewWebsocketsDriver(
		webSockets.ArgsWebSocketsDriver{
			Marshaller:               o.marshaller,
			WebsocketSender:          webSocketSender,
			Uint64ByteSliceConverter: o.uint64ByteSliceConverter,
			Log:                      o.log,
		},
	)
}
