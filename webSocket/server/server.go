package server

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/multiversx/mx-chain-core-go/core"
	"github.com/multiversx/mx-chain-core-go/core/check"
	"github.com/multiversx/mx-chain-core-go/webSocket"
	"github.com/multiversx/mx-chain-core-go/webSocket/connection"
	"github.com/multiversx/mx-chain-core-go/webSocket/data"
	"github.com/multiversx/mx-chain-core-go/webSocket/transceiver"
)

// ArgsWebSocketServer holds all the components needed to create a server
type ArgsWebSocketServer struct {
	RetryDurationInSeconds int
	BlockingAckOnError     bool
	WithAcknowledge        bool
	URL                    string
	PayloadConverter       webSocket.PayloadConverter
	Log                    core.Logger
}

type server struct {
	blockingAckOnError  bool
	withAcknowledge     bool
	payloadConverter    webSocket.PayloadConverter
	retryDuration       time.Duration
	log                 core.Logger
	httpServer          webSocket.HttpServerHandler
	transceiversAndConn transceiversAndConnHandler
	payloadHandler      webSocket.PayloadHandler
}

//NewWebSocketServer will create a new instance of server
func NewWebSocketServer(args ArgsWebSocketServer) (*server, error) {
	if err := checkArgs(args); err != nil {
		return nil, err
	}

	wsServer := &server{
		transceiversAndConn: newTransceiversAndConnHolder(),
		blockingAckOnError:  args.BlockingAckOnError,
		log:                 args.Log,
		retryDuration:       time.Duration(args.RetryDurationInSeconds) * time.Second,
		payloadConverter:    args.PayloadConverter,
		payloadHandler:      webSocket.NewNilPayloadHandler(),
		withAcknowledge:     args.WithAcknowledge,
	}

	wsServer.initializeServer(args.URL, data.WSRoute)

	return wsServer, nil
}

func checkArgs(args ArgsWebSocketServer) error {
	if check.IfNil(args.Log) {
		return core.ErrNilLogger
	}
	if check.IfNil(args.PayloadConverter) {
		return data.ErrNilPayloadConverter
	}
	if args.URL == "" {
		return data.ErrEmptyUrl
	}
	if args.RetryDurationInSeconds == 0 {
		return data.ErrZeroValueRetryDuration
	}
	return nil
}

func (s *server) connectionHandler(connection webSocket.WSConClient) {
	webSocketTransceiver, err := transceiver.NewTransceiver(transceiver.ArgsTransceiver{
		PayloadConverter:   s.payloadConverter,
		Log:                s.log,
		RetryDurationInSec: int(s.retryDuration.Seconds()),
		BlockingAckOnError: s.blockingAckOnError,
		WithAcknowledge:    s.withAcknowledge,
	})
	if err != nil {
		s.log.Warn("s.connectionHandler cannot create transceiver", "error", err)
		return
	}
	err = webSocketTransceiver.SetPayloadHandler(s.payloadHandler)
	if err != nil {
		s.log.Warn("s.SetPayloadHandler cannot set payload handler", "error", err)
	}

	go func() {
		s.transceiversAndConn.addTransceiverAndConn(webSocketTransceiver, connection)
		// this method is blocking
		_ = webSocketTransceiver.Listen(connection)
		s.log.Info("connection closed", "client id", connection.GetID())
		// if method listen will end, the client was disconnected, and we should remove the listener from the list
		s.transceiversAndConn.remove(connection.GetID())
	}()
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

	s.log.Info("wsServer.initializeServer(): initializing WebSocket server", "url", wsURL, "path", wsPath)

	addClientFunc := func(writer http.ResponseWriter, r *http.Request) {
		s.log.Info("new connection", "route", wsPath, "remote address", r.RemoteAddr)

		upgrader.CheckOrigin = func(r *http.Request) bool { return true }

		ws, errUpgrade := upgrader.Upgrade(writer, r, nil)
		if errUpgrade != nil {
			s.log.Warn("could not update websocket connection", "remote address", r.RemoteAddr, "error", errUpgrade)
			return
		}
		client := connection.NewWSConnClientWithConn(ws)
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
func (s *server) Send(args data.PayloadData) error {
	for _, tuple := range s.transceiversAndConn.getAll() {
		err := tuple.transceiver.Send(args, tuple.conn)
		if err != nil {
			s.log.Debug("s.Send() cannot send message", "id", tuple.conn.GetID(), "error", err.Error())
		}
	}

	return nil
}

// Start will start the websockets server
func (s *server) Start() {
	go func() {
		err := s.httpServer.ListenAndServe()
		shouldLogError := err != nil && !strings.Contains(err.Error(), data.ErrServerIsClosed.Error())
		if shouldLogError {
			s.log.Error("could not initialize webserver", "error", err)
			return
		}

		s.log.Info("server was closed")
	}()
}

// SetPayloadHandler will set the provided payload handler
func (s *server) SetPayloadHandler(handler webSocket.PayloadHandler) error {
	s.payloadHandler = handler
	return nil
}

// Close will close the server
func (s *server) Close() error {
	var lastError error

	err := s.httpServer.Shutdown(context.Background())
	if err != nil {
		s.log.Debug("server.Close() cannot close http server", "error", err)
		lastError = err
	}

	for _, tuple := range s.transceiversAndConn.getAll() {
		err = tuple.transceiver.Close()
		if err != nil {
			s.log.Debug("server.Close() cannot close transceiver", "error", err)
			lastError = err
		}
		err = tuple.conn.Close()
		if err != nil {
			s.log.Debug("server.Close() cannot close connection", "id", tuple.conn.GetID(), "error", err.Error())
			lastError = err
		}
	}

	return lastError
}

// IsInterfaceNil returns true if there is no value under the interface
func (s *server) IsInterfaceNil() bool {
	return s == nil
}
