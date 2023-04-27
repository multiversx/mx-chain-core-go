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
	"github.com/multiversx/mx-chain-core-go/webSockets"
	"github.com/multiversx/mx-chain-core-go/webSockets/connection"
	"github.com/multiversx/mx-chain-core-go/webSockets/data"
	"github.com/multiversx/mx-chain-core-go/webSockets/receiver"
	"github.com/multiversx/mx-chain-core-go/webSockets/sender"
)

// ArgsWebSocketsServer holds all the components needed to create a server
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
	connectionHandler        func(connection connection.WSConClient)
	uint64ByteSliceConverter connection.Uint64ByteSliceConverter
	retryDuration            time.Duration
	log                      core.Logger
	httpServer               connection.HttpServerHandler
	sender                   Sender
	receivers                ReceiversHolder
}

//NewWebSocketsServer will create a new instance of server
func NewWebSocketsServer(args ArgsWebSocketsServer) (*server, error) {
	if err := checkArgs(args); err != nil {
		return nil, err
	}

	webSocketsSender, err := sender.NewSender(sender.ArgsSender{
		WithAcknowledge:          args.WithAcknowledge,
		RetryDurationInSeconds:   args.RetryDurationInSeconds,
		Uint64ByteSliceConverter: args.Uint64ByteSliceConverter,
		Log:                      args.Log,
	})
	if err != nil {
		return nil, err
	}

	wsServer := &server{
		sender:                   webSocketsSender,
		receivers:                NewReceiversHolder(),
		blockingAckOnError:       args.BlockingAckOnError,
		log:                      args.Log,
		retryDuration:            time.Duration(args.RetryDurationInSeconds) * time.Second,
		uint64ByteSliceConverter: args.Uint64ByteSliceConverter,
	}
	wsServer.connectionHandler = wsServer.defaultConnectionHandler

	wsServer.initializeServer(args.URL, data.WSRoute)

	return wsServer, nil
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

	s.httpServer = httpServer
}

// Send will send the provided payload from args
func (s *server) Send(args data.WsSendArgs) error {
	return s.sender.Send(args.Payload)
}

// RegisterPayloadHandler will register the provided payload handler
func (s *server) RegisterPayloadHandler(handler webSockets.PayloadHandler) {
	s.connectionHandler = func(connection connection.WSConClient) {
		webSocketsReceiver, err := receiver.NewReceiver(receiver.ArgsReceiver{
			Uint64ByteSliceConverter: s.uint64ByteSliceConverter,
			Log:                      s.log,
			RetryDurationInSec:       int(s.retryDuration.Seconds()),
			BlockingAckOnError:       s.blockingAckOnError,
		})
		if err != nil {
			s.log.Warn("s.connectionHandler cannot create receiver", "error", err)
		}
		webSocketsReceiver.SetPayloadHandler(handler)

		go func() {
			s.receivers.AddReceiver(connection.GetID(), webSocketsReceiver)
			// this method is blocking
			_ = webSocketsReceiver.Listen(connection)
			// if method listen will end, the client was disconnected should remove the listener from the list
			s.receivers.RemoveReceiver(connection.GetID())
		}()
	}
}

// Listen will start the server
func (s *server) Listen() {
	err := s.httpServer.ListenAndServe()
	if err != nil && !strings.Contains(err.Error(), data.ErrServerIsClosed.Error()) {
		s.log.Error("could not initialize webserver", "error", err)
	}

	s.log.Info("server was closed")
}

// Close will close the server
func (s *server) Close() error {
	err := s.httpServer.Shutdown(context.Background())
	if err != nil {
		s.log.Warn("server.Close() cannot close http server", "error", err)
	}

	err = s.sender.Close()
	if err != nil {
		s.log.Warn("server.Close() cannot close the sender", "error", err)
	}

	for _, webSocketsReceiver := range s.receivers.GetAll() {
		err = webSocketsReceiver.Close()
		if err != nil {
			s.log.Warn("server.Close() cannot close receiver", "error", err)
		}
	}

	return nil
}

func checkArgs(args ArgsWebSocketsServer) error {
	if check.IfNil(args.Log) {
		return core.ErrNilLogger
	}
	if check.IfNil(args.Uint64ByteSliceConverter) {
		return data.ErrNilUint64ByteSliceConverter
	}
	if args.URL == "" {
		return data.ErrEmptyUrl
	}
	if args.RetryDurationInSeconds == 0 {
		return data.ErrZeroValueRetryDuration
	}
	return nil
}

func (s *server) IsInterfaceNil() bool {
	return s == nil
}
