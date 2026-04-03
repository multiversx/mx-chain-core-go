package grpc

import (
	"context"

	"github.com/multiversx/mx-chain-core-go/data/outport"
)

// OutportClient defines the client-side operations for the outport service.
type OutportClient interface {
	SaveBlock(ctx context.Context, in *outport.OutportBlock) (*outport.ResponseData, error)
	RevertIndexedBlock(ctx context.Context, in *outport.BlockData) (*outport.ResponseData, error)
	SaveRoundsInfo(ctx context.Context, in *outport.RoundsInfo) (*outport.ResponseData, error)
	SaveValidatorsPubKeys(ctx context.Context, in *outport.ValidatorsPubKeys) (*outport.ResponseData, error)
	SaveValidatorsRating(ctx context.Context, in *outport.ValidatorsRating) (*outport.ResponseData, error)
	SaveAccounts(ctx context.Context, in *outport.Accounts) (*outport.ResponseData, error)
	FinalizedBlockEvent(ctx context.Context, in *outport.FinalizedBlock) (*outport.ResponseData, error)
	IsInterfaceNil() bool
}

// OutportHandler defines the server-side operations for the outport service.
type OutportHandler interface {
	SaveBlock(outportBlock *outport.OutportBlock) error
	RevertIndexedBlock(blockData *outport.BlockData) error
	SaveRoundsInfo(roundsInfos *outport.RoundsInfo) error
	SaveValidatorsPubKeys(validatorsPubKeys *outport.ValidatorsPubKeys) error
	SaveValidatorsRating(ratingData *outport.ValidatorsRating) error
	SaveAccounts(accountsData *outport.Accounts) error
	FinalizedBlock(finalizedBlock *outport.FinalizedBlock) error
	IsInterfaceNil() bool
}
