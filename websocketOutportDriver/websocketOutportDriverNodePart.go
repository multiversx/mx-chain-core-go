package websocketOutportDriver

import (
	"fmt"
	"time"

	"github.com/ElrondNetwork/elrond-go-core/core"
	"github.com/ElrondNetwork/elrond-go-core/core/check"
	"github.com/ElrondNetwork/elrond-go-core/data"
	"github.com/ElrondNetwork/elrond-go-core/data/indexer"
	"github.com/ElrondNetwork/elrond-go-core/marshal"
	outportSenderData "github.com/ElrondNetwork/elrond-go-core/websocketOutportDriver/data"
)

const sleepBetweenRetries = 200 * time.Millisecond

// WebsocketOutportDriverNodePartArgs holds the arguments needed for creating a new websocketOutportDriverNodePart
type WebsocketOutportDriverNodePartArgs struct {
	Enabled                  bool
	Marshalizer              marshal.Marshalizer
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
}

// NewWebsocketOutportDriverNodePart will create a new instance of websocketOutportDriverNodePart
func NewWebsocketOutportDriverNodePart(args WebsocketOutportDriverNodePartArgs) (*websocketOutportDriverNodePart, error) {
	if check.IfNil(args.Marshalizer) {
		return nil, ErrNilMarshalizer
	}
	if check.IfNil(args.WebsocketSender) {
		return nil, ErrNilWebSocketSender
	}
	if check.IfNil(args.Uint64ByteSliceConverter) {
		return nil, ErrNilUint64ByteSliceConverter
	}
	if check.IfNil(args.Log) {
		return nil, ErrNilLogger
	}

	return &websocketOutportDriverNodePart{
		marshalizer:              args.Marshalizer,
		webSocketSender:          args.WebsocketSender,
		uint64ByteSliceConverter: args.Uint64ByteSliceConverter,
		log:                      args.Log,
	}, nil
}

// SaveBlock will send the provided block saving arguments within the websocket
func (o *websocketOutportDriverNodePart) SaveBlock(args *indexer.ArgsSaveBlockData) {
	err := o.handleAction(args, outportSenderData.OperationSaveBlock)
	if err != nil {
		panic(fmt.Sprintf("SaveBlock: %s", err.Error()))
	}
}

// RevertIndexedBlock will handle the action of reverting the indexed block
func (o *websocketOutportDriverNodePart) RevertIndexedBlock(header data.HeaderHandler, body data.BodyHandler) {
	args := outportSenderData.ArgsRevertIndexedBlock{
		Header: header,
		Body:   body,
	}
	err := o.handleAction(args, outportSenderData.OperationRevertIndexedBlock)
	if err != nil {
		panic(fmt.Sprintf("RevertIndexedBlock: %s", err.Error()))
	}
}

// SaveRoundsInfo will handle the saving of rounds
func (o *websocketOutportDriverNodePart) SaveRoundsInfo(roundsInfos []*indexer.RoundInfo) {
	args := outportSenderData.ArgsSaveRoundsInfo{
		RoundsInfos: roundsInfos,
	}
	err := o.handleAction(args, outportSenderData.OperationSaveRoundsInfo)
	if err != nil {
		panic(fmt.Sprintf("SaveRoundsInfo: %s", err.Error()))
	}
}

// SaveValidatorsPubKeys will handle the saving of the validators' public keys
func (o *websocketOutportDriverNodePart) SaveValidatorsPubKeys(validatorsPubKeys map[uint32][][]byte, epoch uint32) {
	args := outportSenderData.ArgsSaveValidatorsPubKeys{
		ValidatorsPubKeys: validatorsPubKeys,
		Epoch:             epoch,
	}
	err := o.handleAction(args, outportSenderData.OperationSaveValidatorsPubKeys)
	if err != nil {
		panic(fmt.Sprintf("SaveValidatorsPubKeys: %s", err.Error()))
	}
}

// SaveValidatorRating will handle the saving of the validators' rating
func (o *websocketOutportDriverNodePart) SaveValidatorsRating(indexID string, infoRating []*indexer.ValidatorRatingInfo) {
	args := outportSenderData.ArgsSaveValidatorsRating{
		IndexID:    indexID,
		InfoRating: infoRating,
	}
	err := o.handleAction(args, outportSenderData.OperationSaveValidatorsRating)
	if err != nil {
		panic(fmt.Sprintf("SaveValidatorsRating: %s", err.Error()))
	}
}

// SaveAccounts will handle the accounts' saving
func (o *websocketOutportDriverNodePart) SaveAccounts(blockTimestamp uint64, acc []data.UserAccountHandler) {
	args := outportSenderData.ArgsSaveAccounts{
		BlockTimestamp: blockTimestamp,
		Acc:            acc,
	}
	err := o.handleAction(args, outportSenderData.OperationSaveAccounts)
	if err != nil {
		panic(fmt.Sprintf("SaveAccounts: %s", err.Error()))
	}
}

// FinalizedBlock will handle the finalized block
func (o *websocketOutportDriverNodePart) FinalizedBlock(headerHash []byte) {
	args := outportSenderData.ArgsFinalizedBlock{
		HeaderHash: headerHash,
	}
	err := o.handleAction(args, outportSenderData.OperationFinalizedBlock)
	if err != nil {
		panic(fmt.Sprintf("FinalizedBlock: %s", err.Error()))
	}
}

// Close will handle the closing of the outport driver web socket sender
func (o *websocketOutportDriverNodePart) Close() error {
	// TODO: close the web socket here
	return nil
}

// IsInterfaceNil returns true if there is no value under the interface
func (o *websocketOutportDriverNodePart) IsInterfaceNil() bool {
	return o == nil
}

func (o *websocketOutportDriverNodePart) handleAction(args interface{}, operation outportSenderData.OperationType) error {
	marshaledBlock, err := o.marshalizer.Marshal(args)
	if err != nil {
		o.log.Error("cannot marshal block", "operation", operation.String(), "error", err)
		return err
	}

	payload := o.preparePayload(operation, marshaledBlock)

	for {
		err = o.webSocketSender.SendOnRoute(outportSenderData.WsSendArgs{
			Payload:   payload,
			Operation: operation,
		})
		if err != nil {
			o.log.Error("cannot send on route. Retrying", "operation", operation.String(), "error", err)
			time.Sleep(sleepBetweenRetries)
			continue
		}

		return nil
	}
}

func (o *websocketOutportDriverNodePart) preparePayload(operation outportSenderData.OperationType, data []byte) []byte {
	opBytes := o.uint32ToBytes(operation.Uint32())
	messageLength := uint32(len(data))
	messageLengthBytes := o.uint32ToBytes(messageLength)

	payload := append(opBytes, messageLengthBytes...)
	payload = append(payload, data...)

	return payload
}

func (o *websocketOutportDriverNodePart) uint32ToBytes(value uint32) []byte {
	result := o.uint64ByteSliceConverter.ToByteSlice(uint64(value))
	if len(result) != 8 {
		return make([]byte, 4)
	}

	return result[4:]
}
