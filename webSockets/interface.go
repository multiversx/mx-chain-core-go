package webSockets

import (
	"github.com/multiversx/mx-chain-core-go/data/outport"
	"github.com/multiversx/mx-chain-core-go/marshal"
	outportSenderData "github.com/multiversx/mx-chain-core-go/webSockets/data"
	"github.com/multiversx/mx-chain-core-go/webSockets/utils"
)

// Driver is an interface for saving node specific data to other storage.
// This could be an elastic search index, a MySql database or any other external services.
type Driver interface {
	SaveBlock(outportBlock *outport.OutportBlock) error
	RevertIndexedBlock(blockData *outport.BlockData) error
	SaveRoundsInfo(roundsInfos *outport.RoundsInfo) error
	SaveValidatorsPubKeys(validatorsPubKeys *outport.ValidatorsPubKeys) error
	SaveValidatorsRating(validatorsRating *outport.ValidatorsRating) error
	SaveAccounts(accounts *outport.Accounts) error
	FinalizedBlock(finalizedBlock *outport.FinalizedBlock) error
	GetMarshaller() marshal.Marshalizer
	Close() error
	IsInterfaceNil() bool
}

// WebSocketSenderHandler defines what the actions that a web socket sender should do
type WebSocketSenderHandler interface {
	Send(args outportSenderData.WsSendArgs) error
	Close() error
	IsInterfaceNil() bool
}

// HostWebSockets -
type HostWebSockets interface {
	Send(args outportSenderData.WsSendArgs) error
	RegisterPayloadHandler(handler utils.PayloadHandler)
	Listen()
	Close() error
	IsInterfaceNil() bool
}
