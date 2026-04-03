package grpc

import (
	"context"

	"github.com/multiversx/mx-chain-core-go/core/check"
	"github.com/multiversx/mx-chain-core-go/data/outport"
)

type outportServer struct {
	handler OutportHandler
}

// NewOutportServer creates a new gRPC server adapter over the provided handler.
func NewOutportServer(handler OutportHandler) (*outportServer, error) {
	if check.IfNil(handler) {
		return nil, ErrNilOutportServiceHandler
	}

	return &outportServer{
		handler: handler,
	}, nil
}

// SaveBlock forwards the call to the provided handler.
func (os *outportServer) SaveBlock(ctx context.Context, in *outport.OutportBlock) (*outport.ResponseData, error) {
	err := os.handler.SaveBlock(in)

	return &outport.ResponseData{}, err
}

// RevertIndexedBlock forwards the call to the provided handler.
func (os *outportServer) RevertIndexedBlock(ctx context.Context, in *outport.BlockData) (*outport.ResponseData, error) {
	err := os.handler.RevertIndexedBlock(in)

	return &outport.ResponseData{}, err
}

// SaveRoundsInfo forwards the call to the provided handler.
func (os *outportServer) SaveRoundsInfo(ctx context.Context, in *outport.RoundsInfo) (*outport.ResponseData, error) {
	err := os.handler.SaveRoundsInfo(in)

	return &outport.ResponseData{}, err
}

// SaveValidatorsPubKeys forwards the call to the provided handler.
func (os *outportServer) SaveValidatorsPubKeys(ctx context.Context, in *outport.ValidatorsPubKeys) (*outport.ResponseData, error) {
	err := os.handler.SaveValidatorsPubKeys(in)

	return &outport.ResponseData{}, err
}

// SaveValidatorsRating forwards the call to the provided handler.
func (os *outportServer) SaveValidatorsRating(ctx context.Context, in *outport.ValidatorsRating) (*outport.ResponseData, error) {
	err := os.handler.SaveValidatorsRating(in)

	return &outport.ResponseData{}, err
}

// SaveAccounts forwards the call to the provided handler.
func (os *outportServer) SaveAccounts(ctx context.Context, in *outport.Accounts) (*outport.ResponseData, error) {
	err := os.handler.SaveAccounts(in)

	return &outport.ResponseData{}, err
}

// FinalizedBlockEvent forwards the call to the provided handler.
func (os *outportServer) FinalizedBlockEvent(ctx context.Context, in *outport.FinalizedBlock) (*outport.ResponseData, error) {
	err := os.handler.FinalizedBlock(in)

	return &outport.ResponseData{}, err
}

// IsInterfaceNil returns true if there is no value under the interface.
func (os *outportServer) IsInterfaceNil() bool {
	return os == nil
}
