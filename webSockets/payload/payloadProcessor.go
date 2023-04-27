package payload

import (
	"github.com/multiversx/mx-chain-core-go/core"
	"github.com/multiversx/mx-chain-core-go/webSockets/connection"
	"github.com/multiversx/mx-chain-core-go/webSockets/data"
)

// ArgsPayloadProcessor holds the components needed in order to create a payloadProcessor
type ArgsPayloadProcessor struct {
	Log           core.Logger
	PayloadParser connection.PayloadParser
	PayloadProc   connection.PayloadProcessor
}

type payloadProcessor struct {
	log           core.Logger
	payloadParser connection.PayloadParser
	payloadProc   connection.PayloadProcessor
}

// NewPayloadProcessor will create a new instance of payloadProcessor
func NewPayloadProcessor(args ArgsPayloadProcessor) (*payloadProcessor, error) {
	return &payloadProcessor{
		log:           args.Log,
		payloadParser: args.PayloadParser,
		payloadProc:   args.PayloadProc,
	}, nil
}

// HandlePayload will handler the provided payload
func (pp *payloadProcessor) HandlePayload(payload []byte) (*data.PayloadData, error) {
	payloadData, err := pp.payloadParser.ExtractPayloadData(payload)
	if err != nil {
		pp.log.Error("error while extracting payload data: ", "error", err)
		return nil, err
	}

	pp.log.Info("processing payload",
		"counter", payloadData.Counter,
		"operation type", payloadData.OperationType,
		"message length", len(payloadData.Payload),
	)

	pp.log.Trace("processing payload data", "payload", payloadData.Payload)

	err = pp.payloadProc.ProcessPayload(payloadData)
	if err != nil {
		return nil, err
	}

	return payloadData, nil
}

// Close will close the payload processor
func (pp *payloadProcessor) Close() error {
	return pp.payloadProc.Close()
}
