package clientServerSender

import (
	"sync/atomic"
	"time"

	"github.com/multiversx/mx-chain-core-go/websocketOutportDriver/clientServerSender/client"
	"github.com/multiversx/mx-chain-core-go/websocketOutportDriver/clientServerSender/server"
	logger "github.com/multiversx/mx-chain-logger-go"
)

var (
	prefixWithoutAck = []byte{0}
	log              = logger.GetOrCreate("websocketOutportDriver/clientServerSender")
)

type ArgsWSClientServerSender struct {
	IsServer                 bool
	Url                      string
	Uint64ByteSliceConverter server.Uint64ByteSliceConverter
	RetryDuration            time.Duration
}

type sender struct {
	messageSender            MessageSender
	uint64ByteSliceConverter server.Uint64ByteSliceConverter
	counter                  uint64
}

func NewClientServerSender(args ArgsWSClientServerSender) (*sender, error) {
	messageSender, err := createMessageSender(args)
	if err != nil {
		return nil, err
	}

	wsSender := &sender{
		counter:                  0,
		uint64ByteSliceConverter: args.Uint64ByteSliceConverter,
		messageSender:            messageSender,
	}

	return wsSender, nil
}

func (s *sender) SendMessage(message []byte) error {
	assignedCounter := atomic.AddUint64(&s.counter, 1)

	ackData := prefixWithoutAck

	newPayload := append(ackData, s.uint64ByteSliceConverter.ToByteSlice(assignedCounter)...)
	newPayload = append(newPayload, message...)

	return s.messageSender.Send(assignedCounter, message)
}

func (s *sender) Close() error {
	return s.messageSender.Close()
}

func createMessageSender(args ArgsWSClientServerSender) (MessageSender, error) {
	if args.IsServer {
		return server.NewWebSocketSender(server.WebSocketSenderArgs{
			Uint64ByteSliceConverter: args.Uint64ByteSliceConverter,
			Log:                      log,
			URL:                      args.Url,
			WithAcknowledge:          false,
		})
	}

	return client.NewClient(client.WebSocketClientSenderArgs{
		Uint64ByteSliceConverter: args.Uint64ByteSliceConverter,
		RetryDuration:            args.RetryDuration,
		URL:                      args.Url,
	})
}
