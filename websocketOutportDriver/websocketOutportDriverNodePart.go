package websocketOutportDriver

import (
	"fmt"

	"github.com/multiversx/mx-chain-core-go/core"
	"github.com/multiversx/mx-chain-core-go/core/atomic"
	"github.com/multiversx/mx-chain-core-go/core/check"
	"github.com/multiversx/mx-chain-core-go/data/outport"
	"github.com/multiversx/mx-chain-core-go/marshal"
	"github.com/multiversx/mx-chain-core-go/websocketOutportDriver/common"
	outportSenderData "github.com/multiversx/mx-chain-core-go/websocketOutportDriver/data"
)

// WebsocketOutportDriverNodePartArgs holds the arguments needed for creating a new websocketOutportDriverNodePart
type WebsocketOutportDriverNodePartArgs struct {
	Enabled                  bool
	Marshaller               marshal.Marshalizer
	WebsocketSender          WebSocketSenderHandler
	WebSocketConfig          outportSenderData.WebSocketConfig
	Uint64ByteSliceConverter common.Uint64ByteSliceConverter
	Log                      core.Logger
}

type websocketOutportDriverNodePart struct {
	marshalizer              marshal.Marshalizer
	log                      core.Logger
	uint64ByteSliceConverter common.Uint64ByteSliceConverter
	webSocketSender          WebSocketSenderHandler
	isClosed                 atomic.Flag
}

// NewWebsocketOutportDriverNodePart will create a new instance of websocketOutportDriverNodePart
func NewWebsocketOutportDriverNodePart(args WebsocketOutportDriverNodePartArgs) (*websocketOutportDriverNodePart, error) {
	if check.IfNil(args.Marshaller) {
		return nil, outportSenderData.ErrNilMarshaller
	}
	if check.IfNil(args.WebsocketSender) {
		return nil, outportSenderData.ErrNilWebSocketSender
	}
	if check.IfNil(args.Uint64ByteSliceConverter) {
		return nil, outportSenderData.ErrNilUint64ByteSliceConverter
	}
	if check.IfNil(args.Log) {
		return nil, outportSenderData.ErrNilLogger
	}

	isClosedFlag := atomic.Flag{}
	isClosedFlag.SetValue(false)

	return &websocketOutportDriverNodePart{
		marshalizer:              args.Marshaller,
		webSocketSender:          args.WebsocketSender,
		uint64ByteSliceConverter: args.Uint64ByteSliceConverter,
		log:                      args.Log,
		isClosed:                 isClosedFlag,
	}, nil
}

// SaveBlock will send the provided block saving arguments within the websocket
func (o *websocketOutportDriverNodePart) SaveBlock(outportBlock *outport.OutportBlock) error {
	return o.handleAction(outportBlock, outportSenderData.OperationSaveBlock)
}

// RevertIndexedBlock will handle the action of reverting the indexed block
func (o *websocketOutportDriverNodePart) RevertIndexedBlock(blockData *outport.BlockData) error {
	return o.handleAction(blockData, outportSenderData.OperationRevertIndexedBlock)
}

// SaveRoundsInfo will handle the saving of rounds
func (o *websocketOutportDriverNodePart) SaveRoundsInfo(roundsInfos *outport.RoundsInfo) error {
	return o.handleAction(roundsInfos, outportSenderData.OperationSaveRoundsInfo)
}

// SaveValidatorsPubKeys will handle the saving of the validators' public keys
func (o *websocketOutportDriverNodePart) SaveValidatorsPubKeys(validatorsPubKeys *outport.ValidatorsPubKeys) error {
	return o.handleAction(validatorsPubKeys, outportSenderData.OperationSaveValidatorsPubKeys)
}

// SaveValidatorsRating will handle the saving of the validators' rating
func (o *websocketOutportDriverNodePart) SaveValidatorsRating(validatorsRating *outport.ValidatorsRating) error {
	return o.handleAction(validatorsRating, outportSenderData.OperationSaveValidatorsRating)
}

// SaveAccounts will handle the accounts' saving
func (o *websocketOutportDriverNodePart) SaveAccounts(accounts *outport.Accounts) error {
	return o.handleAction(accounts, outportSenderData.OperationSaveAccounts)
}

// FinalizedBlock will handle the finalized block
func (o *websocketOutportDriverNodePart) FinalizedBlock(finalizedBlock *outport.FinalizedBlock) error {
	return o.handleAction(finalizedBlock, outportSenderData.OperationFinalizedBlock)
}

// GetMarshaller returns the internal marshaller
func (o *websocketOutportDriverNodePart) GetMarshaller() marshal.Marshalizer {
	return o.marshalizer
}

// Close will handle the closing of the outport driver web socket sender
func (o *websocketOutportDriverNodePart) Close() error {
	o.isClosed.SetValue(true)
	return o.webSocketSender.Close()
}

// IsInterfaceNil returns true if there is no value under the interface
func (o *websocketOutportDriverNodePart) IsInterfaceNil() bool {
	return o == nil
}

func (o *websocketOutportDriverNodePart) handleAction(args interface{}, operation outportSenderData.OperationType) error {
	if o.isClosed.IsSet() {
		return outportSenderData.ErrWebSocketServerIsClosed
	}

	marshaledBlock, err := o.marshalizer.Marshal(args)
	if err != nil {
		o.log.Error("cannot marshal block", "operation", operation.String(), "error", err)
		return fmt.Errorf("%w while marshaling block for operation %s", err, operation.String())
	}

	payload := o.preparePayload(operation, marshaledBlock)

	err = o.webSocketSender.Send(outportSenderData.WsSendArgs{
		Payload: payload,
	})
	if err != nil {
		o.log.Error("cannot send on route", "operation", operation.String(), "error", err)
		return fmt.Errorf("%w while sending data on route for operation %s", err, operation.String())
	}

	return nil
}

func (o *websocketOutportDriverNodePart) preparePayload(operation outportSenderData.OperationType, data []byte) []byte {
	opBytes := o.uint64ByteSliceConverter.ToByteSlice(uint64(operation.Uint32()))
	opBytes = opBytes[uint32NumBytes:]

	messageLength := uint64(len(data))
	messageLengthBytes := o.uint64ByteSliceConverter.ToByteSlice(messageLength)
	messageLengthBytes = messageLengthBytes[uint32NumBytes:]

	payload := append(opBytes, messageLengthBytes...)
	payload = append(payload, data...)

	return payload
}
