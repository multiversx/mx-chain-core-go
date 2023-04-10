package server

import (
	"context"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/multiversx/mx-chain-core-go/core"
	"github.com/multiversx/mx-chain-core-go/data/typeConverters/uint64ByteSlice"
	"github.com/multiversx/mx-chain-core-go/websocketOutportDriver"
	"github.com/multiversx/mx-chain-core-go/websocketOutportDriver/clientServerReceiver/payload"
	"github.com/multiversx/mx-chain-core-go/websocketOutportDriver/clientServerSender/server"
	"github.com/multiversx/mx-chain-core-go/websocketOutportDriver/common"
	outportData "github.com/multiversx/mx-chain-core-go/websocketOutportDriver/data"
)

type ArgsWsServer struct {
	URL                string
	RetryDurationInSec uint32
	BlockingAckOnError bool
	Log                core.Logger
	PayloadProcessor   common.PayloadProcessor
}

type wsServer struct {
	mutex                    sync.Mutex
	log                      core.Logger
	payloadParser            common.PayloadParser
	payloadProcessor         common.PayloadProcessor
	uint64ByteSliceConverter server.Uint64ByteSliceConverter
	safeCloser               core.SafeCloser
	server                   common.HttpServerHandler
	listeners                map[string]common.MessagesListener
	retryDurationInSec       uint32
	blockingAckOnError       bool
}

func NewWsServer(args ArgsWsServer) (*wsServer, error) {
	uint64ByteSliceConverter := uint64ByteSlice.NewBigEndianConverter()
	payloadParser, err := websocketOutportDriver.NewWebSocketPayloadParser(uint64ByteSliceConverter)
	if err != nil {
		return nil, err
	}

	return &wsServer{
		log:                      args.Log,
		payloadParser:            payloadParser,
		payloadProcessor:         args.PayloadProcessor,
		uint64ByteSliceConverter: uint64ByteSliceConverter,
		retryDurationInSec:       args.RetryDurationInSec,
		blockingAckOnError:       args.BlockingAckOnError,
	}, nil
}

// Start will start the web-sockets server
func (s *wsServer) Start() {
	err := s.server.ListenAndServe()
	if err != nil && !strings.Contains(err.Error(), outportData.ErrServerIsClosed.Error()) {
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

func (s *wsServer) handleMessages(client common.WSClient) {
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
	}

	s.mutex.Lock()
	s.listeners[client.GetID()] = listener
	s.mutex.Unlock()

	// this method is blocking
	_ = listener.Listen()
	// if method listen will end the client was disconnected should remove the listener from the list
	s.mutex.Lock()
	delete(s.listeners, client.GetID())
	s.mutex.Unlock()

}

// Close will close the web-sockets server
func (s *wsServer) Close() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	err := s.server.Shutdown(ctx)
	if err != nil {
		s.log.Error("cannot close the server", "error", err)
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()
	for _, listener := range s.listeners {
		listener.Close()
	}
}
