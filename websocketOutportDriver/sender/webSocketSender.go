package sender

import (
	"context"
	"fmt"
	"strings"
	"sync/atomic"
	"time"

	"github.com/ElrondNetwork/elrond-go-core/core"
	"github.com/ElrondNetwork/elrond-go-core/core/check"
	outportData "github.com/ElrondNetwork/elrond-go-core/websocketOutportDriver/data"
	"github.com/gorilla/websocket"
)

var (
	prefixWithAck    = []byte{1}
	prefixWithoutAck = []byte{0}
)

type webSocketClient struct {
	conn       outportData.WSConn
	remoteAddr string
}

type webSocketSender struct {
	log core.Logger
	// TODO: use an interface for http server (or simply provide the URL only) in order to make this component easy testable
	server                   HttpServerHandler
	counter                  uint64
	uint64ByteSliceConverter Uint64ByteSliceConverter
	// TODO: use interfaces instead of direct instances + analyze returning pointers vs values on exported functions
	clientsHolder   *websocketClientsHolder
	acknowledges    *acknowledgesHolder
	withAcknowledge bool
}

// WebSocketSenderArgs holds the arguments needed for creating a new instance of webSocketSender
type WebSocketSenderArgs struct {
	Server                   HttpServerHandler
	Uint64ByteSliceConverter Uint64ByteSliceConverter
	WithAcknowledge          bool
	Log                      core.Logger
}

// NewWebSocketSender returns a new instance of webSocketSender
func NewWebSocketSender(args WebSocketSenderArgs) (*webSocketSender, error) {
	if args.Server == nil {
		return nil, outportData.ErrNilHttpServer
	}
	if check.IfNil(args.Uint64ByteSliceConverter) {
		return nil, outportData.ErrNilUint64ByteSliceConverter
	}
	if check.IfNil(args.Log) {
		return nil, outportData.ErrNilLogger
	}

	ws := &webSocketSender{
		log:                      args.Log,
		server:                   args.Server,
		counter:                  0,
		uint64ByteSliceConverter: args.Uint64ByteSliceConverter,
		clientsHolder:            NewWebsocketClientsHolder(),
		acknowledges:             NewAcknowledgesHolder(),
		withAcknowledge:          args.WithAcknowledge,
	}

	go ws.start()

	return ws, nil
}

// AddClient will add the client to internal maps and will also start
func (w *webSocketSender) AddClient(wss outportData.WSConn, remoteAddr string) {
	if wss == nil {
		w.log.Warn("nil ws connection provider", "remote addr", remoteAddr)
		return
	}

	client := &webSocketClient{
		conn:       wss,
		remoteAddr: remoteAddr,
	}

	err := w.clientsHolder.AddClient(client)
	if err != nil {
		w.log.Warn("cannot AddClient", "error", err)
		return
	}

	// TODO: maybe multiple clients types could be supported: some require ack, while some don't require ack
	if !w.withAcknowledge {
		return
	}

	w.acknowledges.AddEntry(remoteAddr)

	go w.handleReceiveAck(client)
}

func (w *webSocketSender) handleReceiveAck(client *webSocketClient) {
	for {
		mType, message, err := client.conn.ReadMessage()
		if err != nil {
			w.log.Error("cannot read message", "remote addr", client.remoteAddr, "error", err)

			err = w.clientsHolder.CloseAndRemove(client.remoteAddr)
			w.log.LogIfError(err)

			w.acknowledges.RemoveEntryForAddress(client.remoteAddr)

			break
		}

		if mType != websocket.BinaryMessage {
			w.log.Warn("received message is not binary message", "remote addr", client.remoteAddr, "message type", mType)
			continue
		}

		w.log.Trace("received ack", "remote addr", client.remoteAddr, "message", message)
		counter, err := w.uint64ByteSliceConverter.ToUint64(message)
		if err != nil {
			w.log.Warn("cannot decode counter: bytes to uint64",
				"remote addr", client.remoteAddr,
				"counter bytes", message,
				"error", err,
			)
			continue
		}

		w.acknowledges.AddReceivedAcknowledge(client.remoteAddr, counter)
	}
}

func (w *webSocketSender) start() {
	err := w.server.ListenAndServe()
	if err != nil && !strings.Contains(err.Error(), outportData.ErrServerIsClosed.Error()) {
		w.log.Error("could not initialize webserver", "error", err)
	}
}

func (w *webSocketSender) sendDataToClients(
	data []byte,
	counter uint64,
) error {
	numSent := 0
	var err error

	clients := w.clientsHolder.GetAll()
	if len(clients) == 0 {
		return outportData.ErrNoClientToSendTo
	}

	for _, client := range w.clientsHolder.GetAll() {
		err = w.sendData(data, *client, counter)
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

func (w *webSocketSender) sendData(
	data []byte,
	client webSocketClient,
	counter uint64,
) error {
	if len(data) == 0 {
		return outportData.ErrEmptyDataToSend
	}

	errSend := client.conn.WriteMessage(websocket.BinaryMessage, data)
	if errSend != nil {
		// TODO: test if this is a situation when the client connection should be dropped
		w.log.Warn("could not send data to client", "remote addr", client.remoteAddr, "error", errSend)
		return fmt.Errorf("%w while writing message to client %s", errSend, client.remoteAddr)
	}

	if !w.withAcknowledge {
		return nil
	}

	// TODO: might refactor this (send to each clients, then wait for all VS send to one client, wait for it, move to next)
	w.waitForAck(client.remoteAddr, counter)

	return nil
}

func (w *webSocketSender) waitForAck(remoteAddr string, counter uint64) {
	for {
		acksForAddress, ok := w.acknowledges.GetAcknowledgesOfAddress(remoteAddr)
		if !ok {
			w.log.Warn("waiting acknowledge for an address that isn't present anymore in clients map", "remote addr", remoteAddr)
			return
		}

		ok = acksForAddress.ProcessAcknowledged(counter)
		if ok {
			return
		}

		time.Sleep(time.Millisecond)
	}
}

// Send will make the request accordingly to the received arguments
func (w *webSocketSender) Send(args outportData.WsSendArgs) error {
	assignedCounter := atomic.AddUint64(&w.counter, 1)
	ackData := prefixWithoutAck
	if w.withAcknowledge {
		ackData = prefixWithAck
	}

	newPayload := append(ackData, w.uint64ByteSliceConverter.ToByteSlice(assignedCounter)...)
	newPayload = append(newPayload, args.Payload...)

	return w.sendDataToClients(newPayload, assignedCounter)
}

// Close will close the server and the connections with the clients
func (w *webSocketSender) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	err := w.server.Shutdown(ctx)
	if err != nil {
		w.log.Error("cannot close the server", "error", err)
	}

	for _, client := range w.clientsHolder.GetAll() {
		err = client.conn.Close()
		w.log.LogIfError(err)
	}

	return nil
}

// IsInterfaceNil returns true if there is no value under the interface
func (w *webSocketSender) IsInterfaceNil() bool {
	return w == nil
}
