package clientServerSender

import (
	"sync/atomic"

	"github.com/multiversx/mx-chain-core-go/core"
	"github.com/multiversx/mx-chain-core-go/websocketOutportDriver/clientServerSender/client"
	"github.com/multiversx/mx-chain-core-go/websocketOutportDriver/clientServerSender/server"
	"github.com/multiversx/mx-chain-core-go/websocketOutportDriver/common"
	outportSenderData "github.com/multiversx/mx-chain-core-go/websocketOutportDriver/data"
)

var (
	prefixWithoutAck = []byte{0}
	prefixWithAck    = []byte{1}
)

type ArgsWSClientServerSender struct {
	IsServer                 bool
	WithAcknowledge          bool
	Url                      string
	RetryDurationInSec       int
	Uint64ByteSliceConverter common.Uint64ByteSliceConverter
	Log                      core.Logger
}

type sender struct {
	messageSender            MessageSender
	uint64ByteSliceConverter common.Uint64ByteSliceConverter
	counter                  uint64
	withAcknowledge          bool
}

// NewClientServerSender will create a new instance of *sender
func NewClientServerSender(args ArgsWSClientServerSender) (*sender, error) {
	messageSender, err := createMessageSender(args)
	if err != nil {
		return nil, err
	}

	wsSender := &sender{
		counter:                  0,
		uint64ByteSliceConverter: args.Uint64ByteSliceConverter,
		messageSender:            messageSender,
		withAcknowledge:          args.WithAcknowledge,
	}

	return wsSender, nil
}

func createMessageSender(args ArgsWSClientServerSender) (MessageSender, error) {
	if args.IsServer {
		return server.NewServerSender(server.ArgsServerSender{
			Uint64ByteSliceConverter: args.Uint64ByteSliceConverter,
			Log:                      args.Log,
			URL:                      args.Url,
			WithAcknowledge:          args.WithAcknowledge,
		})
	}

	return client.NewClientSender(client.ArgsWsClientSender{
		Uint64ByteSliceConverter: args.Uint64ByteSliceConverter,
		RetryDurationInSec:       args.RetryDurationInSec,
		URL:                      args.Url,
		WithAcknowledge:          args.WithAcknowledge,
		Log:                      args.Log,
	})
}

// Send will send the provided payload from the args
func (s *sender) Send(args outportSenderData.WsSendArgs) error {
	assignedCounter := atomic.AddUint64(&s.counter, 1)

	ackData := prefixWithoutAck
	if s.withAcknowledge {
		ackData = prefixWithAck
	}

	newPayload := append(ackData, s.uint64ByteSliceConverter.ToByteSlice(assignedCounter)...)
	newPayload = append(newPayload, args.Payload...)

	return s.messageSender.Send(assignedCounter, newPayload)
}

func (s *sender) Close() error {
	return s.messageSender.Close()
}

// IsInterfaceNil returns true if there is no value under the interface
func (s *sender) IsInterfaceNil() bool {
	return s == nil
}
