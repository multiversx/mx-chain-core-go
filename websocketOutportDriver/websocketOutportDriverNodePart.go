package websocketOutportDriver

import (
	"fmt"

	"github.com/ElrondNetwork/elrond-go-core/core"
	"github.com/ElrondNetwork/elrond-go-core/core/atomic"
	"github.com/ElrondNetwork/elrond-go-core/core/check"
	"github.com/ElrondNetwork/elrond-go-core/data"
	"github.com/ElrondNetwork/elrond-go-core/data/outport"
	"github.com/ElrondNetwork/elrond-go-core/marshal"
	outportSenderData "github.com/ElrondNetwork/elrond-go-core/websocketOutportDriver/data"
)

// WebsocketOutportDriverNodePartArgs holds the arguments needed for creating a new websocketOutportDriverNodePart
type WebsocketOutportDriverNodePartArgs struct {
	Enabled                  bool
	Marshaller               marshal.Marshalizer
	WebsocketSender          WebSocketSenderHandler
	WebSocketConfig          outportSenderData.WebSocketConfig
	Uint64ByteSliceConverter Uint64ByteSliceConverter
	Log                      core.Logger
}

type websocketOutportDriverNodePart struct {
	marshalizer              marshal.Marshalizer
	log                      core.Logger
	uint64ByteSliceConverter Uint64ByteSliceConverter
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
func (o *websocketOutportDriverNodePart) SaveBlock(args *outport.ArgsSaveBlockData) error {
	return o.handleAction(args, outportSenderData.OperationSaveBlock)
}

// RevertIndexedBlock will handle the action of reverting the indexed block
func (o *websocketOutportDriverNodePart) RevertIndexedBlock(header data.HeaderHandler, body data.BodyHandler) error {
	args := outportSenderData.ArgsRevertIndexedBlock{
		Header: header,
		Body:   body,
	}

	return o.handleAction(args, outportSenderData.OperationRevertIndexedBlock)
}

// SaveRoundsInfo will handle the saving of rounds
func (o *websocketOutportDriverNodePart) SaveRoundsInfo(roundsInfos []*outport.RoundInfo) error {
	args := outportSenderData.ArgsSaveRoundsInfo{
		RoundsInfos: roundsInfos,
	}

	return o.handleAction(args, outportSenderData.OperationSaveRoundsInfo)
}

// SaveValidatorsPubKeys will handle the saving of the validators' public keys
func (o *websocketOutportDriverNodePart) SaveValidatorsPubKeys(validatorsPubKeys map[uint32][][]byte, epoch uint32) error {
	args := outportSenderData.ArgsSaveValidatorsPubKeys{
		ValidatorsPubKeys: validatorsPubKeys,
		Epoch:             epoch,
	}

	return o.handleAction(args, outportSenderData.OperationSaveValidatorsPubKeys)
}

// SaveValidatorsRating will handle the saving of the validators' rating
func (o *websocketOutportDriverNodePart) SaveValidatorsRating(indexID string, infoRating []*outport.ValidatorRatingInfo) error {
	args := outportSenderData.ArgsSaveValidatorsRating{
		IndexID:    indexID,
		InfoRating: infoRating,
	}

	return o.handleAction(args, outportSenderData.OperationSaveValidatorsRating)
}

// SaveAccounts will handle the accounts' saving
func (o *websocketOutportDriverNodePart) SaveAccounts(blockTimestamp uint64, acc map[string]*outport.AlteredAccount, shardID uint32) error {
	args := outportSenderData.ArgsSaveAccounts{
		BlockTimestamp: blockTimestamp,
		Acc:            acc,
		ShardID:        shardID,
	}

	return o.handleAction(args, outportSenderData.OperationSaveAccounts)
}

// FinalizedBlock will handle the finalized block
func (o *websocketOutportDriverNodePart) FinalizedBlock(headerHash []byte) error {
	args := outportSenderData.ArgsFinalizedBlock{
		HeaderHash: headerHash,
	}

	return o.handleAction(args, outportSenderData.OperationFinalizedBlock)
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
