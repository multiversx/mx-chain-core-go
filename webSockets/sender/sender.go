package sender

import (
	"fmt"
	"sync/atomic"
	"time"

	"github.com/gorilla/websocket"
	"github.com/multiversx/mx-chain-core-go/core"
	"github.com/multiversx/mx-chain-core-go/core/check"
	"github.com/multiversx/mx-chain-core-go/core/closing"
	"github.com/multiversx/mx-chain-core-go/webSockets"
	"github.com/multiversx/mx-chain-core-go/webSockets/connection"
	outportData "github.com/multiversx/mx-chain-core-go/webSockets/data"
)

// ArgsSender holds the arguments needed for creating a web-sockets sender
type ArgsSender struct {
	WithAcknowledge          bool
	RetryDurationInSeconds   int
	Uint64ByteSliceConverter webSockets.Uint64ByteSliceConverter
	Log                      core.Logger
}

type sender struct {
	counter         uint64
	withAcknowledge bool
	connections     ConnectionsHandler
	payloadParser   webSockets.PayloadParser
	log             core.Logger
	safeCloser      core.SafeCloser
	retryDuration   time.Duration
}

func NewSender(args ArgsSender) (*sender, error) {
	err := checkArgs(args)
	if err != nil {
		return nil, err
	}
	payloadConverter, err := webSockets.NewWebSocketPayloadParser(args.Uint64ByteSliceConverter)
	if err != nil {
		return nil, err
	}

	return &sender{
		counter:         0,
		safeCloser:      closing.NewSafeChanCloser(),
		connections:     connection.NewWebsocketClientsHolder(),
		log:             args.Log,
		payloadParser:   payloadConverter,
		withAcknowledge: args.WithAcknowledge,
		retryDuration:   time.Duration(args.RetryDurationInSeconds) * time.Second,
	}, nil
}

func (s *sender) AddConnection(client webSockets.WSConClient) error {
	return s.connections.AddClient(client)
}

func (s *sender) Send(payload []byte) error {
	assignedCounter := atomic.AddUint64(&s.counter, 1)
	newPayload := s.payloadParser.ExtendPayloadWithCounter(payload, assignedCounter, s.withAcknowledge)

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
	var lastError error

	for _, client := range clients {
		err := s.sendPayload(payload, assignedCounter, client)
		if err != nil {
			s.log.Debug("sender.send(): couldn't send data to client", "id", client.GetID(), "error", err)
			lastError = err
			continue
		}

		numSent++
	}

	if numSent == 0 {
		return fmt.Errorf("data wasn't sent to any client. last known error: %w", lastError)
	}

	return nil

}

func (s *sender) sendPayload(payload []byte, assignedCounter uint64, connection webSockets.WSConClient) error {
	errSend := connection.WriteMessage(websocket.BinaryMessage, payload)
	if errSend != nil {
		return errSend
	}

	if !s.withAcknowledge {
		return nil
	}

	return s.waitForAck(connection, assignedCounter)
}

func (s *sender) waitForAck(connection webSockets.WSConClient, assignedCounter uint64) error {
	timer := time.NewTimer(s.retryDuration)
	defer timer.Stop()

	for {
		mType, message, err := connection.ReadMessage()
		if err != nil {
			s.log.Debug("s.waitForAck(): cannot read message", "id", connection.GetID(), "error", err)

			err = s.connections.CloseAndRemove(connection.GetID())
			s.log.LogIfError(err)
			break
		}

		if mType != websocket.BinaryMessage {
			s.log.Debug("received message is not binary message", "id", connection.GetID(), "message type", mType)
			continue
		}

		s.log.Trace("received ack", "remote addr", connection.GetID(), "message", message)

		receivedCounter, err := s.payloadParser.DecodeCounter(message)
		if err != nil {
			s.log.Warn("cannot decode counter: bytes to uint64",
				"id", connection.GetID(),
				"counter bytes", message,
				"error", err,
			)
		}

		if receivedCounter == assignedCounter {
			return nil
		}

		timer.Reset(s.retryDuration)
		s.log.Debug("s.waitForAck invalid counter", "expected", assignedCounter, "received", receivedCounter, "id", connection.GetID())

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

func checkArgs(args ArgsSender) error {
	if check.IfNil(args.Log) {
		return core.ErrNilLogger
	}
	if check.IfNil(args.Uint64ByteSliceConverter) {
		return outportData.ErrNilUint64ByteSliceConverter
	}
	if args.RetryDurationInSeconds == 0 {
		return outportData.ErrZeroValueRetryDuration
	}
	return nil
}
