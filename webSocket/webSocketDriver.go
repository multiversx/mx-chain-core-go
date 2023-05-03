package webSocket

import (
	"fmt"

	"github.com/multiversx/mx-chain-core-go/core"
	"github.com/multiversx/mx-chain-core-go/core/atomic"
	"github.com/multiversx/mx-chain-core-go/core/check"
	"github.com/multiversx/mx-chain-core-go/data/outport"
	"github.com/multiversx/mx-chain-core-go/marshal"
	outportSenderData "github.com/multiversx/mx-chain-core-go/webSocket/data"
)

// ArgsWebSocketDriver holds the arguments needed for creating a new webSocketDriver
type ArgsWebSocketDriver struct {
	Marshaller      marshal.Marshalizer
	WebsocketSender WebSocketSenderHandler
	Log             core.Logger
}

type webSocketDriver struct {
	marshalizer      marshal.Marshalizer
	log              core.Logger
	payloadConverter PayloadConverter
	webSocketSender  WebSocketSenderHandler
	isClosed         atomic.Flag
}

// NewWebsocketDriver will create a new instance of webSocketDriver
func NewWebsocketDriver(args ArgsWebSocketDriver) (*webSocketDriver, error) {
	if check.IfNil(args.Marshaller) {
		return nil, outportSenderData.ErrNilMarshaller
	}
	if check.IfNil(args.WebsocketSender) {
		return nil, outportSenderData.ErrNilWebSocketSender
	}
	if check.IfNil(args.Log) {
		return nil, outportSenderData.ErrNilLogger
	}

	isClosedFlag := atomic.Flag{}
	isClosedFlag.SetValue(false)

	return &webSocketDriver{
		marshalizer:     args.Marshaller,
		webSocketSender: args.WebsocketSender,
		log:             args.Log,
		isClosed:        isClosedFlag,
	}, nil
}

// SaveBlock will send the provided block saving arguments within the websocket
func (o *webSocketDriver) SaveBlock(outportBlock *outport.OutportBlock) error {
	return o.handleAction(outportBlock, outportSenderData.OperationSaveBlock)
}

// RevertIndexedBlock will handle the action of reverting the indexed block
func (o *webSocketDriver) RevertIndexedBlock(blockData *outport.BlockData) error {
	return o.handleAction(blockData, outportSenderData.OperationRevertIndexedBlock)
}

// SaveRoundsInfo will handle the saving of rounds
func (o *webSocketDriver) SaveRoundsInfo(roundsInfos *outport.RoundsInfo) error {
	return o.handleAction(roundsInfos, outportSenderData.OperationSaveRoundsInfo)
}

// SaveValidatorsPubKeys will handle the saving of the validators' public keys
func (o *webSocketDriver) SaveValidatorsPubKeys(validatorsPubKeys *outport.ValidatorsPubKeys) error {
	return o.handleAction(validatorsPubKeys, outportSenderData.OperationSaveValidatorsPubKeys)
}

// SaveValidatorsRating will handle the saving of the validators' rating
func (o *webSocketDriver) SaveValidatorsRating(validatorsRating *outport.ValidatorsRating) error {
	return o.handleAction(validatorsRating, outportSenderData.OperationSaveValidatorsRating)
}

// SaveAccounts will handle the accounts' saving
func (o *webSocketDriver) SaveAccounts(accounts *outport.Accounts) error {
	return o.handleAction(accounts, outportSenderData.OperationSaveAccounts)
}

// FinalizedBlock will handle the finalized block
func (o *webSocketDriver) FinalizedBlock(finalizedBlock *outport.FinalizedBlock) error {
	return o.handleAction(finalizedBlock, outportSenderData.OperationFinalizedBlock)
}

// GetMarshaller returns the internal marshaller
func (o *webSocketDriver) GetMarshaller() marshal.Marshalizer {
	return o.marshalizer
}

func (o *webSocketDriver) handleAction(args interface{}, operation outportSenderData.OperationType) error {
	if o.isClosed.IsSet() {
		return outportSenderData.ErrWebSocketServerIsClosed
	}

	marshalledPayload, err := o.marshalizer.Marshal(args)
	if err != nil {
		o.log.Error("cannot marshal block", "operation", operation.String(), "error", err)
		return fmt.Errorf("%w while marshaling block for operation %s", err, operation.String())
	}

	err = o.webSocketSender.Send(outportSenderData.WsSendArgs{
		Payload: marshalledPayload,
		OpType:  operation,
	})
	if err != nil {
		o.log.Error("cannot send on route", "operation", operation.String(), "error", err)
		return fmt.Errorf("%w while sending data on route for operation %s", err, operation.String())
	}

	return nil
}

// Close will handle the closing of the outport driver web socket sender
func (o *webSocketDriver) Close() error {
	o.isClosed.SetValue(true)
	return o.webSocketSender.Close()
}

// IsInterfaceNil returns true if there is no value under the interface
func (o *webSocketDriver) IsInterfaceNil() bool {
	return o == nil
}