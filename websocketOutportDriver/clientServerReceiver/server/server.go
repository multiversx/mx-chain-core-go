package server

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/multiversx/mx-chain-core-go/core"
	"github.com/multiversx/mx-chain-core-go/core/check"
	"github.com/multiversx/mx-chain-core-go/websocketOutportDriver/clientServerReceiver/payload"
	"github.com/multiversx/mx-chain-core-go/websocketOutportDriver/common"
	"github.com/multiversx/mx-chain-core-go/websocketOutportDriver/data"
)

type ArgsWsServer struct {
	URL                      string
	RetryDurationInSec       uint32
	BlockingAckOnError       bool
	Log                      core.Logger
	PayloadParser            common.PayloadParser
	PayloadProcessor         common.PayloadProcessor
	Uint64ByteSliceConverter common.Uint64ByteSliceConverter
}

type wsServer struct {
	log                      core.Logger
	payloadParser            common.PayloadParser
	payloadProcessor         common.PayloadProcessor
	uint64ByteSliceConverter common.Uint64ByteSliceConverter
	safeCloser               core.SafeCloser
	server                   common.HttpServerHandler
	listeners                common.ListenersHolder
	retryDurationInSec       uint32
	blockingAckOnError       bool
}

func NewWsServer(args ArgsWsServer) (*wsServer, error) {
	err := checkArgs(args)
	if err != nil {
		return nil, err
	}

	s := &wsServer{
		log:                      args.Log,
		payloadParser:            args.PayloadParser,
		payloadProcessor:         args.PayloadProcessor,
		uint64ByteSliceConverter: args.Uint64ByteSliceConverter,
		retryDurationInSec:       args.RetryDurationInSec,
		blockingAckOnError:       args.BlockingAckOnError,
		listeners:                common.NewListenersHolder(),
	}
	s.initializeServer(args.URL, data.WSRoute)

	return s, nil
}

// Start will start the web-sockets server
func (s *wsServer) Start() {
	err := s.server.ListenAndServe()
	if err != nil && !strings.Contains(err.Error(), data.ErrServerIsClosed.Error()) {
		s.log.Error("could not initialize webserver", "error", err)
	}

	s.log.Info("server was closed")
}

func (s *wsServer) initializeServer(wsURL string, wsPath string) {
	router := mux.NewRouter()
	httpServer := &http.Server{
		Addr:    wsURL,
		Handler: router,
	}

	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	s.log.Info("wsServer.initializeServer(): initializing web-sockets server", "url", wsURL, "path", wsPath)

	addClientFunc := func(writer http.ResponseWriter, r *http.Request) {
		// generate a unique client ID for the new client
		clientID := uuid.New().String()
		s.log.Info("new connection", "route", wsPath, "remote address", r.RemoteAddr, "id", clientID)

		upgrader.CheckOrigin = func(r *http.Request) bool { return true }

		ws, errUpgrade := upgrader.Upgrade(writer, r, nil)
		if errUpgrade != nil {
			s.log.Warn("could not update websocket connection", "remote address", r.RemoteAddr, "error", errUpgrade)
			return
		}
		client := common.NewWSConnClientWithConn(ws, clientID)
		go s.handleMessages(client)

	}

	routeSendData := router.HandleFunc(wsPath, addClientFunc)

	if routeSendData.GetError() != nil {
		s.log.Error("sender router failed to handle send data",
			"route", routeSendData.GetName(),
			"error", routeSendData.GetError())
	}

	s.server = httpServer
}

func (s *wsServer) handleMessages(client common.WSConClient) {
	listener, err := payload.NewMessagesListener(payload.ArgsMessagesProcessor{
		Log:                      s.log,
		PayloadParser:            s.payloadParser,
		PayloadProcessor:         s.payloadProcessor,
		WsClient:                 client,
		Uint64ByteSliceConverter: s.uint64ByteSliceConverter,
		RetryDurationInSec:       s.retryDurationInSec,
		BlockingAckOnError:       s.blockingAckOnError,
	})
	if err != nil {
		s.log.Error("wsServer.handleMessages: cannot create messages listener", "error", err, "clientID", client.GetID())
		return
	}

	s.listeners.Add(client.GetID(), listener)
	// this method is blocking
	_ = listener.Listen()
	// if method listen will end, the client was disconnected should remove the listener from the list
	s.listeners.Remove(client.GetID())

}

// Close will close the web-sockets server
func (s *wsServer) Close() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	err := s.server.Shutdown(ctx)
	if err != nil {
		s.log.Error("cannot close the server", "error", err)
	}

	for _, listener := range s.listeners.GetAll() {
		listener.Close()
	}
}

func checkArgs(args ArgsWsServer) error {
	if check.IfNil(args.PayloadProcessor) {
		return data.ErrNilPayloadProcessor
	}
	if check.IfNil(args.PayloadParser) {
		return data.ErrNilPayloadParser
	}
	if check.IfNil(args.Uint64ByteSliceConverter) {
		return data.ErrNilUint64ByteSliceConverter
	}
	if len(args.URL) == 0 {
		return data.ErrEmptyUrl
	}
	if args.RetryDurationInSec == 0 {
		return data.ErrZeroValueRetryDuration
	}
	if check.IfNil(args.Log) {
		return data.ErrNilLogger
	}

	return nil
}
