package server

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/multiversx/mx-chain-core-go/core"
	"github.com/multiversx/mx-chain-core-go/core/check"
	"github.com/multiversx/mx-chain-core-go/websocketOutportDriver/common"
	outportData "github.com/multiversx/mx-chain-core-go/websocketOutportDriver/data"
)

// ArgsServerSender holds the arguments needed for creating a new instance of webSocketSender
type ArgsServerSender struct {
	Uint64ByteSliceConverter common.Uint64ByteSliceConverter
	Log                      core.Logger
	URL                      string
	WithAcknowledge          bool
}

type serverSender struct {
	withAcknowledge          bool
	log                      core.Logger
	uint64ByteSliceConverter common.Uint64ByteSliceConverter
	clientsHolder            ClientsHandler
	server                   common.HttpServerHandler
	acknowledges             ClientAcknowledgesHolder
}

// NewServerSender returns a new instance of serverSender
func NewServerSender(args ArgsServerSender) (*serverSender, error) {
	if check.IfNil(args.Uint64ByteSliceConverter) {
		return nil, outportData.ErrNilUint64ByteSliceConverter
	}
	if check.IfNil(args.Log) {
		return nil, outportData.ErrNilLogger
	}
	if len(args.URL) == 0 {
		return nil, outportData.ErrEmptyUrl
	}

	ws := &serverSender{
		log:                      args.Log,
		uint64ByteSliceConverter: args.Uint64ByteSliceConverter,
		clientsHolder:            common.NewWebsocketClientsHolder(),
		acknowledges:             common.NewAcknowledgesHolder(),
		withAcknowledge:          args.WithAcknowledge,
	}

	ws.initializeServer(args.URL, outportData.WSRoute)

	go ws.start()

	return ws, nil
}

func (w *serverSender) initializeServer(wsURL string, wsPath string) {
	router := mux.NewRouter()
	server := &http.Server{
		Addr:    wsURL,
		Handler: router,
	}

	w.log.Info("webSocketSender initializeServer(): initializing web-sockets server", "url", wsURL, "path", wsPath)

	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	addClientFunc := func(writer http.ResponseWriter, r *http.Request) {
		// generate a unique client ID for the new client
		clientID := uuid.New().String()
		w.log.Info("new connection", "route", wsPath, "remote address", r.RemoteAddr, "id", clientID)

		upgrader.CheckOrigin = func(r *http.Request) bool { return true }

		ws, errUpgrade := upgrader.Upgrade(writer, r, nil)
		if errUpgrade != nil {
			w.log.Warn("could not update websocket connection", "remote address", r.RemoteAddr, "error", errUpgrade)
			return
		}
		client := common.NewWSConnClientWithConn(ws, clientID)
		w.addClient(client)
	}

	routeSendData := router.HandleFunc(wsPath, addClientFunc)

	if routeSendData.GetError() != nil {
		w.log.Error("sender router failed to handle send data",
			"route", routeSendData.GetName(),
			"error", routeSendData.GetError())
	}

	w.server = server
}

func (w *serverSender) addClient(client common.WSClient) {
	err := w.clientsHolder.AddClient(client)
	if err != nil {
		w.log.Warn("webSocketSender.handleNewClient cannot add client", "error", err, "id", client.GetID())
	}

	if !w.withAcknowledge {
		return
	}

	w.acknowledges.AddEntry(client.GetID())
	go w.handleReceiveAck(client)
}

func (w *serverSender) start() {
	err := w.server.ListenAndServe()
	if err != nil && !strings.Contains(err.Error(), outportData.ErrServerIsClosed.Error()) {
		w.log.Error("could not initialize webserver", "error", err)
	}

	w.log.Info("server was closed")
}

// Send will make the request accordingly to the received arguments
func (w *serverSender) Send(counter uint64, payload []byte) error {
	return w.sendDataToClients(counter, payload)
}

func (w *serverSender) sendDataToClients(
	counter uint64,
	data []byte,
) error {
	numSent := 0
	var err error

	clients := w.clientsHolder.GetAll()
	if len(clients) == 0 {
		return outportData.ErrNoClientToSendTo
	}

	for _, client := range w.clientsHolder.GetAll() {
		err = w.sendData(counter, data, client)
		if err != nil {
			w.log.Error("couldn't send data to client", "error", err)
			continue
		}

		numSent++
	}

	if numSent == 0 {
		return fmt.Errorf("data wasn't sent to any client. last known error: %w", err)
	}

	return nil
}

func (w *serverSender) sendData(
	counter uint64,
	data []byte,
	client common.WSClient,
) error {
	if len(data) == 0 {
		return outportData.ErrEmptyDataToSend
	}

	errSend := client.WriteMessage(websocket.BinaryMessage, data)
	if errSend != nil {
		// TODO: test if this is a situation when the client connection should be dropped
		w.log.Warn("could not send data to client", "remote addr", client.GetID(), "error", errSend)
		return fmt.Errorf("%w while writing message to client %s", errSend, client.GetID())
	}

	if !w.withAcknowledge {
		return nil
	}

	w.waitForAck(client.GetID(), counter)

	return nil
}

// Close will close the server and the connections with the clients
func (w *serverSender) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	w.log.Info("webSocketSender.Close(): closing the web-sockets server")

	err := w.server.Shutdown(ctx)
	if err != nil {
		w.log.Error("cannot close the server", "error", err)
	}

	for _, client := range w.clientsHolder.GetAll() {
		err = client.Close()
		w.log.LogIfError(err)
	}

	return nil
}

// IsInterfaceNil returns true if there is no value under the interface
func (w *serverSender) IsInterfaceNil() bool {
	return w == nil
}
