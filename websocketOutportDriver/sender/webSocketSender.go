package sender

import (
	"net/http"
	"sync"
	"time"

	"github.com/ElrondNetwork/elrond-go-core/core"
	"github.com/ElrondNetwork/elrond-go-core/core/atomic"
	"github.com/ElrondNetwork/elrond-go-core/core/check"
	"github.com/ElrondNetwork/elrond-go-core/websocketOutportDriver/data"
	"github.com/gorilla/websocket"
)

const retrialTimeoutMS = 50

var (
	prefixWithAck    = []byte{1}
	prefixWithoutAck = []byte{0}
)

type webSocketClient struct {
	conn       WSConn
	remoteAddr string
}

type webSocketSender struct {
	log                      core.Logger
	server                   *http.Server
	counter                  atomic.Uint64
	uint64ByteSliceConverter Uint64ByteSliceConverter
	clients                  map[string]*webSocketClient
	mutClients               sync.RWMutex
	acknowledges             map[string]map[uint64]struct{}
	mutAcknowledges          sync.RWMutex
	withAcknowledge          bool
}

// WebSocketSenderArgs holds the arguments needed for creating a new instance of webSockerSender
type WebSocketSenderArgs struct {
	Server                   *http.Server
	Uint64ByteSliceConverter Uint64ByteSliceConverter
	WithAcknowledge          bool
	Log                      core.Logger
}

// NewWebSocketSender returns a new instance of webSocketSender
func NewWebSocketSender(args WebSocketSenderArgs) (*webSocketSender, error) {
	if args.Server == nil {
		return nil, ErrNilHttpServer
	}
	if check.IfNil(args.Uint64ByteSliceConverter) {
		return nil, ErrNilUint64ByteSliceConverter
	}
	if check.IfNil(args.Log) {
		return nil, ErrNilLogger
	}

	atomicCounter := atomic.Uint64{}
	atomicCounter.Set(0)

	ws := &webSocketSender{
		log:                      args.Log,
		server:                   args.Server,
		counter:                  atomicCounter,
		uint64ByteSliceConverter: args.Uint64ByteSliceConverter,
		clients:                  make(map[string]*webSocketClient),
		acknowledges:             make(map[string]map[uint64]struct{}),
		withAcknowledge:          args.WithAcknowledge,
	}

	go ws.start()

	return ws, nil
}

func (w *webSocketSender) AddClient(wss WSConn, remoteAddr string) {
	client := &webSocketClient{
		conn:       wss,
		remoteAddr: remoteAddr,
	}
	w.mutClients.Lock()
	w.clients[remoteAddr] = client
	w.mutClients.Unlock()

	w.mutAcknowledges.Lock()
	w.acknowledges[remoteAddr] = make(map[uint64]struct{})
	w.mutAcknowledges.Unlock()

	go w.handleReceiveAck(client)
}

func (w *webSocketSender) getClients() map[string]*webSocketClient {
	w.mutClients.RLock()
	defer w.mutClients.RUnlock()

	return w.clients
}

func (w *webSocketSender) handleReceiveAck(client *webSocketClient) {
	for {
		mType, message, err := client.conn.ReadMessage()
		if err != nil {
			w.log.Error("cannot read message", "remote addr", client.remoteAddr, "error", err)
			w.mutClients.Lock()
			delete(w.clients, client.remoteAddr)
			w.mutClients.Unlock()

			break
		}

		if mType != websocket.BinaryMessage {
			w.log.Warn("received message is not binary message", "remote addr", client.remoteAddr, "message type", mType)
			continue
		}

		w.log.Info("received ack", "remote addr", client.remoteAddr, "message", message)
		counter, err := w.uint64ByteSliceConverter.ToUint64(message)
		if err != nil {
			w.log.Warn("cannot decode counter: bytes to uint64",
				"remote addr", client.remoteAddr,
				"counter bytes", message,
				"error", err,
			)
			continue
		}

		w.mutAcknowledges.Lock()
		w.acknowledges[client.remoteAddr][counter] = struct{}{}
		w.mutAcknowledges.Unlock()
	}
}

func (w *webSocketSender) start() {
	err := w.server.ListenAndServe()
	if err != nil {
		w.log.Error("could not initialize webserver", "error", err)
	}
}

func (w *webSocketSender) sendWithRetrial(data []byte, ackData []byte, counter uint64) {
	clients := w.getClients()
	if len(clients) == 0 {
		w.log.Warn("no client to send to")
	}

	ticker := time.NewTicker(time.Millisecond * retrialTimeoutMS)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			clients = w.getClients()
			if len(clients) == 0 {
				continue
			}
			dataSent := w.sendDataToClients(data, ackData, clients, counter)
			if dataSent {
				return
			}
		}
	}
}

func (w *webSocketSender) sendDataToClients(
	data []byte,
	ackData []byte,
	clients map[string]*webSocketClient,
	counter uint64,
) bool {

	for _, client := range clients {
		result := w.sendData(data, ackData, *client, counter)
		if !result {
			return false
		}
	}

	return true
}

func (w *webSocketSender) sendData(
	data []byte,
	ackData []byte,
	client webSocketClient,
	counter uint64,
) bool {
	errSend := client.conn.WriteMessage(websocket.BinaryMessage, data)
	if errSend != nil {
		w.log.Warn("could not send data to client", "remote addr", client.remoteAddr, "error", errSend)
		return false
	}

	if len(data) == 0 || len(ackData) == 0 {
		return false
	}

	if ackData[0] == prefixWithoutAck[0] {
		return true
	}

	// TODO: might refactor this (send to each clients, then wait for all VS send to one client, wait for it, move to next)
	w.waitForAck(client.remoteAddr, counter)

	return true
}

func (w *webSocketSender) waitForAck(remoteAddr string, counter uint64) {
	for {
		acknowledges := w.getAcknowledges()
		acksForAddress, ok := acknowledges[remoteAddr]
		if !ok {
			w.log.Warn("waiting acknowledge for an address that isn't present anymore in clients map", "remote addr", remoteAddr)
			return
		}

		_, ok = acksForAddress[counter]
		if ok {
			w.removeAcknowledge(remoteAddr, counter)
			return
		}

		time.Sleep(time.Millisecond)
	}
}

func (w *webSocketSender) getAcknowledges() map[string]map[uint64]struct{} {
	w.mutAcknowledges.RLock()
	defer w.mutAcknowledges.RUnlock()

	return w.acknowledges
}

func (w *webSocketSender) addCounterToAcknowledges(remoteAddr string, counter uint64) {
	w.mutAcknowledges.Lock()
	_, exists := w.acknowledges[remoteAddr]
	if exists {
		w.acknowledges[remoteAddr][counter] = struct{}{}
		w.mutAcknowledges.Unlock()

		return
	}

	w.log.Warn("adding counter to non-existing remote addr", "remote addr", remoteAddr, "counter", counter)
	w.acknowledges[remoteAddr] = make(map[uint64]struct{}) // should never reach here
}

func (w *webSocketSender) removeAcknowledge(remoteAddr string, counter uint64) {
	// for a better performance (avoid existing checks), this function relies on the fact that the value exists in the map
	w.mutAcknowledges.Lock()
	acks := w.acknowledges[remoteAddr]
	delete(acks, counter)
	w.mutAcknowledges.Unlock()
}

// SendOnRoute will make the request accordingly to the received arguments
func (w *webSocketSender) SendOnRoute(args data.WsSendArgs) error {
	assignedCounter := w.counter.Get()
	w.log.Info("counter", "value", assignedCounter)
	w.counter.Set(assignedCounter + 1)
	ackData := prefixWithoutAck

	if w.withAcknowledge {
		ackData = append(prefixWithAck, w.uint64ByteSliceConverter.ToByteSlice(assignedCounter)...)
	}

	newPayload := append(ackData, args.Payload...)

	w.sendWithRetrial(newPayload, ackData, assignedCounter)

	return nil
}

// IsInterfaceNil returns true if there is no value under the interface
func (w *webSocketSender) IsInterfaceNil() bool {
	return w == nil
}
