package sender

import (
	"fmt"
	"sync/atomic"
	"time"

	"github.com/gorilla/websocket"
	"github.com/multiversx/mx-chain-core-go/core"
	"github.com/multiversx/mx-chain-core-go/core/check"
	"github.com/multiversx/mx-chain-core-go/core/closing"
	"github.com/multiversx/mx-chain-core-go/webSocket"
	"github.com/multiversx/mx-chain-core-go/webSocket/connection"
	outportData "github.com/multiversx/mx-chain-core-go/webSocket/data"
)

// ArgsSender holds the arguments needed for creating a WebSocket sender
type ArgsSender struct {
	WithAcknowledge        bool
	RetryDurationInSeconds int
	PayloadConverter       webSocket.PayloadConverter
	Log                    core.Logger
}

type sender struct {
	counter         uint64
	withAcknowledge bool
	connections     ConnectionsHandler
	payloadParser   webSocket.PayloadConverter
	log             core.Logger
	safeCloser      core.SafeCloser
	retryDuration   time.Duration
}

// NewSender will create a new instance of WebSocket sender
func NewSender(args ArgsSender) (*sender, error) {
	err := checkArgs(args)
	if err != nil {
		return nil, err
	}

	return &sender{
		counter:         0,
		safeCloser:      closing.NewSafeChanCloser(),
		connections:     connection.NewWebsocketClientsHolder(),
		log:             args.Log,
		payloadParser:   args.PayloadConverter,
		withAcknowledge: args.WithAcknowledge,
		retryDuration:   time.Duration(args.RetryDurationInSeconds) * time.Second,
	}, nil
}

// AddConnection will add in the connections map the provided connection
func (s *sender) AddConnection(client webSocket.WSConClient) error {
	return s.connections.AddClient(client)
}

// Send will prepare and send the provided WsSendArgs
func (s *sender) Send(args outportData.WsSendArgs) error {
	assignedCounter := atomic.AddUint64(&s.counter, 1)
	newPayload := s.payloadParser.ConstructPayloadData(args, assignedCounter, s.withAcknowledge)

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

func (s *sender) sendPayload(payload []byte, assignedCounter uint64, connection webSocket.WSConClient) error {
	errSend := connection.WriteMessage(websocket.BinaryMessage, payload)
	if errSend != nil {
		return errSend
	}

	if !s.withAcknowledge {
		return nil
	}

	return s.waitForAck(connection, assignedCounter)
}

func (s *sender) waitForAck(connection webSocket.WSConClient, assignedCounter uint64) error {
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

		receivedCounter, err := s.payloadParser.DecodeUint64(message)
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
	if check.IfNil(args.PayloadConverter) {
		return outportData.ErrNilPayloadConverter
	}
	if args.RetryDurationInSeconds == 0 {
		return outportData.ErrZeroValueRetryDuration
	}
	return nil
}