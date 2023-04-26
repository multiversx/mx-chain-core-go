package sender

import (
	"fmt"
	"sync/atomic"
	"time"

	"github.com/gorilla/websocket"
	"github.com/multiversx/mx-chain-core-go/core"
	"github.com/multiversx/mx-chain-core-go/core/closing"
	"github.com/multiversx/mx-chain-core-go/webSockets/connection"
	outportData "github.com/multiversx/mx-chain-core-go/webSockets/data"
)

var (
	prefixWithoutAck = []byte{0}
	prefixWithAck    = []byte{1}
)

type ArgsSender struct {
	WithAcknowledge          bool
	RetryDurationInSeconds   int
	Uint64ByteSliceConverter connection.Uint64ByteSliceConverter
	Log                      core.Logger
}

type sender struct {
	counter                  uint64
	withAcknowledge          bool
	connections              ConnectionsHandler
	uint64ByteSliceConverter connection.Uint64ByteSliceConverter
	log                      core.Logger
	safeCloser               core.SafeCloser
	retryDuration            time.Duration
}

func NewSender(args ArgsSender) (*sender, error) {
	return &sender{
		counter:                  0,
		safeCloser:               closing.NewSafeChanCloser(),
		connections:              connection.NewWebsocketClientsHolder(),
		log:                      args.Log,
		uint64ByteSliceConverter: args.Uint64ByteSliceConverter,
		withAcknowledge:          args.WithAcknowledge,
		retryDuration:            time.Duration(args.RetryDurationInSeconds) * time.Second,
	}, nil
}

func (s *sender) AddConnection(client connection.WSConClient) error {
	return s.connections.AddClient(client)
}

func (s *sender) Send(payload []byte) error {
	assignedCounter := atomic.AddUint64(&s.counter, 1)

	ackData := prefixWithoutAck
	if s.withAcknowledge {
		ackData = prefixWithAck
	}

	newPayload := append(ackData, s.uint64ByteSliceConverter.ToByteSlice(assignedCounter)...)
	newPayload = append(newPayload, payload...)

	return s.send(newPayload, assignedCounter)
}

func (s *sender) send(payload []byte, assignedCounter uint64) error {
	clients := s.connections.GetAll()
	if len(clients) == 0 {
		return outportData.ErrNoClientToSendTo
	}
	if len(payload) == 0 {
		return outportData.ErrEmptyDataToSend
	}

	numSent := 0
	var err error

	for _, client := range clients {
		err = s.sendPayload(payload, assignedCounter, client)
		if err != nil {
			s.log.Error("couldn't send data to client", "error", err)
			continue
		}

		numSent++
	}

	if numSent == 0 {
		return fmt.Errorf("data wasn't sent to any client. last known error: %w", err)
	}

	return nil

}

func (s *sender) sendPayload(payload []byte, assignedCounter uint64, connection connection.WSConClient) error {
	errSend := connection.WriteMessage(websocket.BinaryMessage, payload)
	if errSend != nil {
		s.log.Warn("could not send data to client", "remote addr", connection.GetID(), "error", errSend)
		return fmt.Errorf("%w while writing message to client %s", errSend, connection.GetID())
	}

	if !s.withAcknowledge {
		return nil
	}

	return s.waitForAck(connection, assignedCounter)
}

func (s *sender) waitForAck(connection connection.WSConClient, assignedCounter uint64) error {
	timer := time.NewTimer(s.retryDuration)
	defer timer.Stop()

	for {
		mType, message, err := connection.ReadMessage()
		if err != nil {
			s.log.Error("cannot read message", "id", connection.GetID(), "error", err)

			err = s.connections.CloseAndRemove(connection.GetID())
			s.log.LogIfError(err)
			break
		}

		if mType != websocket.BinaryMessage {
			s.log.Warn("received message is not binary message", "id", connection.GetID(), "message type", mType)
			continue
		}

		s.log.Trace("received ack", "remote addr", connection.GetID(), "message", message)

		receivedCounter, err := s.uint64ByteSliceConverter.ToUint64(message)
		if err != nil {
			s.log.Warn("cannot decode counter: bytes to uint64",
				"id", connection.GetID(),
				"counter bytes", message,
				"error", err,
			)
			continue
		}

		if receivedCounter == assignedCounter {
			return nil
		}

		timer.Reset(s.retryDuration)
		s.log.Warn("s.waitForAck invalid counter", "expected", assignedCounter, "received", receivedCounter, "id", connection.GetID())

		select {
		case <-timer.C:
		case <-s.safeCloser.ChanClose():
			return nil
		}
	}

	return nil
}

func (s *sender) Close() error {
	defer s.safeCloser.Close()

	connections := s.connections.GetAll()
	for _, conn := range connections {
		err := conn.Close()
		if err != nil {
			s.log.Warn("sender.Close() cannot close connection", "id", conn.GetID(), "error", err)
		}
	}

	return nil
}
