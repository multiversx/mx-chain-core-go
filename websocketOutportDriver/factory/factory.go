package factory

import (
	"net/http"

	"github.com/ElrondNetwork/elrond-go-core/core"
	"github.com/ElrondNetwork/elrond-go-core/core/check"
	"github.com/ElrondNetwork/elrond-go-core/marshal"
	outportdriverwebsocketsender "github.com/ElrondNetwork/elrond-go-core/websocketOutportDriver"
	data2 "github.com/ElrondNetwork/elrond-go-core/websocketOutportDriver/data"
	"github.com/ElrondNetwork/elrond-go-core/websocketOutportDriver/sender"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

const operationsPath = "/operations"

// OutportDriverWebSocketSenderFactoryArgs holds the arguments needed for creating a outportDriverWebSocketSenderFactory
type OutportDriverWebSocketSenderFactoryArgs struct {
	WebSocketConfig          data2.WebSocketConfig
	Marshaller               marshal.Marshalizer
	Actions                  map[string]struct{}
	Uint64ByteSliceConverter Uint64ByteSliceConverter
	Log                      core.Logger
	WithAcknowledge          bool
}

type outportDriverWebSocketSenderFactory struct {
	webSocketConfig          data2.WebSocketConfig
	marshaller               marshal.Marshalizer
	uint64ByteSliceConverter Uint64ByteSliceConverter
	withAcknowledge          bool
	log                      core.Logger
}

// NewOutportDriverWebSocketSenderFactory will return a new instance of outportDriverWebSocketSenderFactory
func NewOutportDriverWebSocketSenderFactory(args OutportDriverWebSocketSenderFactoryArgs) (*outportDriverWebSocketSenderFactory, error) {
	if check.IfNil(args.Marshaller) {
		return nil, ErrNilMarshaller
	}
	if check.IfNil(args.Uint64ByteSliceConverter) {
		return nil, ErrNilUint64ByteSliceConverter
	}
	if check.IfNil(args.Log) {
		return nil, ErrNilLogger
	}
	return &outportDriverWebSocketSenderFactory{
		webSocketConfig:          args.WebSocketConfig,
		marshaller:               args.Marshaller,
		uint64ByteSliceConverter: args.Uint64ByteSliceConverter,
		withAcknowledge:          args.WithAcknowledge,
		log:                      args.Log,
	}, nil
}

// Create will handle the creation of all the components needed to create an outport driver that sends data over
// web socket and return it afterwards
func (o *outportDriverWebSocketSenderFactory) Create() (Driver, error) {
	webSocketSender, err := o.createWebSocketSender()
	if err != nil {
		return nil, err
	}

	return outportdriverwebsocketsender.NewWebsocketOutportDriverNodePart(
		outportdriverwebsocketsender.WebsocketOutportDriverNodePartArgs{
			Enabled:                  false,
			Marshalizer:              o.marshaller,
			WebsocketSender:          webSocketSender,
			WebSocketConfig:          data2.WebSocketConfig{},
			Uint64ByteSliceConverter: o.uint64ByteSliceConverter,
			Log:                      o.log,
		},
	)
}

func (o *outportDriverWebSocketSenderFactory) createWebSocketSender() (WebSocketSenderHandler, error) {
	router := mux.NewRouter()
	server := &http.Server{
		Addr:    o.webSocketConfig.URL,
		Handler: router,
	}

	webSocketSenderArgs := sender.WebSocketSenderArgs{
		Server:                   server,
		Uint64ByteSliceConverter: o.uint64ByteSliceConverter,
		WithAcknowledge:          o.withAcknowledge,
		Log:                      o.log,
	}
	webSocketSender, err := sender.NewWebSocketSender(webSocketSenderArgs)
	if err != nil {
		return nil, err
	}

	err = o.registerRoute(router, webSocketSender, operationsPath)
	if err != nil {
		return nil, err
	}

	return webSocketSender, nil
}

func (o *outportDriverWebSocketSenderFactory) registerRoute(router *mux.Router, webSocketHandler WebSocketSenderHandler, path string) error {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	routeSendData := router.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		o.log.Info("new connection", "route", path, "remote address", r.RemoteAddr)

		upgrader.CheckOrigin = func(r *http.Request) bool { return true }

		ws, errUpgrade := upgrader.Upgrade(w, r, nil)
		if errUpgrade != nil {
			o.log.Warn("could not upgrade http connection to sender", "error", errUpgrade)
			return
		}

		webSocketHandler.AddClient(ws, ws.RemoteAddr().String())
	})

	if routeSendData.GetError() != nil {
		o.log.Error("sender router failed to handle send data",
			"route", routeSendData.GetName(),
			"error", routeSendData.GetError())
	}

	return nil
}
