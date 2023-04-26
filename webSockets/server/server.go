package server

import (
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/multiversx/mx-chain-core-go/core"
	"github.com/multiversx/mx-chain-core-go/webSockets"
	"github.com/multiversx/mx-chain-core-go/webSockets/connection"
	"github.com/multiversx/mx-chain-core-go/webSockets/data"
	"github.com/multiversx/mx-chain-core-go/webSockets/receiver"
	"github.com/multiversx/mx-chain-core-go/webSockets/sender"
)

type ArgsWebSocketsServer struct {
	RetryDurationInSeconds   int
	BlockingAckOnError       bool
	WithAcknowledge          bool
	URL                      string
	Uint64ByteSliceConverter connection.Uint64ByteSliceConverter
	Log                      core.Logger
}

type server struct {
	blockingAckOnError       bool
	url                      string
	uint64ByteSliceConverter connection.Uint64ByteSliceConverter
	retryDuration            time.Duration
	safeCloser               core.SafeCloser
	log                      core.Logger
	sender                   Sender
	server                   connection.HttpServerHandler
	receivers                ReceiversHolder
	connectionHandler        func(connection connection.WSConClient)
}

func NewWebSocketsServer(args ArgsWebSocketsServer) (*server, error) {
	sender, err := sender.NewSender(sender.ArgsSender{
		WithAcknowledge:          args.WithAcknowledge,
		RetryDurationInSeconds:   args.RetryDurationInSeconds,
		Uint64ByteSliceConverter: args.Uint64ByteSliceConverter,
		Log:                      args.Log,
	})
	if err != nil {
		return nil, err
	}

	s := &server{
		sender:             sender,
		receivers:          NewReceiversHolder(),
		blockingAckOnError: args.BlockingAckOnError,
		log:                args.Log,
		retryDuration:      time.Duration(args.RetryDurationInSeconds) * time.Second,
	}
	s.connectionHandler = s.defaultConnectionHandler

	s.initializeServer(args.URL, data.WSRoute)

	return &server{}, nil
}

func (s *server) defaultConnectionHandler(conn connection.WSConClient) {
	_ = s.sender.AddConnection(conn)
}

func (s *server) initializeServer(wsURL string, wsPath string) {
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
		client := connection.NewWSConnClientWithConn(ws, clientID)
		s.connectionHandler(client)
	}

	routeSendData := router.HandleFunc(wsPath, addClientFunc)
	if routeSendData.GetError() != nil {
		s.log.Error("sender router failed to handle send data",
			"route", routeSendData.GetName(),
			"error", routeSendData.GetError())
	}

	s.server = httpServer
}

func (s *server) Send(args data.WsSendArgs) error {
	return s.sender.Send(args.Payload)
}

func (s *server) RegisterPayloadHandler(handler webSockets.PayloadHandler) {
	s.connectionHandler = func(connection connection.WSConClient) {
		receiver, err := receiver.NewReceiver(receiver.ArgsReceiver{
			Uint64ByteSliceConverter: s.uint64ByteSliceConverter,
			Log:                      s.log,
			RetryDurationInSec:       int(s.retryDuration.Seconds()),
			BlockingAckOnError:       s.blockingAckOnError,
		})
		if err != nil {
			s.log.Warn("s.connectionHandler cannot create receiver", "error", err)
		}
		receiver.SetPayloadHandler(handler)

		go func() {
			s.receivers.AddReceiver(connection.GetID(), receiver)
			// this method is blocking
			_ = receiver.Listen(connection)
			// if method listen will end, the client was disconnected should remove the listener from the list
			s.receivers.RemoveReceiver(connection.GetID())
		}()
	}
}

func (s *server) Listen() {
	err := s.server.ListenAndServe()
	if err != nil && !strings.Contains(err.Error(), data.ErrServerIsClosed.Error()) {
		s.log.Error("could not initialize webserver", "error", err)
	}

	s.log.Info("server was closed")
}

func (s *server) Close() error {
	defer s.safeCloser.Close()

	err := s.sender.Close()
	if err != nil {
		s.log.Warn("server.Close() cannot close the sender", "error", err)
	}

	for _, receiver := range s.receivers.GetAll() {
		err = receiver.Close()
		if err != nil {
			s.log.Warn("server.Close() cannot close receiver", "error", err)
		}
	}

	return nil
}

func (s *server) IsInterfaceNil() bool {
	return s == nil
}
