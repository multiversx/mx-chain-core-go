package server

import (
	"time"

	"github.com/gorilla/websocket"
	"github.com/multiversx/mx-chain-core-go/websocketOutportDriver/common"
)

func (w *serverSender) handleReceiveAck(client common.WSClient) {
	for {
		mType, message, err := client.ReadMessage()
		if err != nil {
			w.log.Error("cannot read message", "id", client.GetID(), "error", err)

			err = w.clientsHolder.CloseAndRemove(client.GetID())
			w.log.LogIfError(err)

			w.acknowledges.RemoveEntryForAddress(client.GetID())

			break
		}

		if mType != websocket.BinaryMessage {
			w.log.Warn("received message is not binary message", "id", client.GetID(), "message type", mType)
			continue
		}

		w.log.Trace("received ack", "remote addr", client.GetID(), "message", message)
		counter, err := w.uint64ByteSliceConverter.ToUint64(message)
		if err != nil {
			w.log.Warn("cannot decode counter: bytes to uint64",
				"id", client.GetID(),
				"counter bytes", message,
				"error", err,
			)
			continue
		}

		w.acknowledges.AddReceivedAcknowledge(client.GetID(), counter)
	}
}

func (w *serverSender) waitForAck(remoteAddr string, counter uint64) {
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
