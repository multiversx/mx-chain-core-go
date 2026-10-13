package grpcadapter

import (
	"context"

	"github.com/multiversx/mx-chain-core-go/core/check"
	"github.com/multiversx/mx-chain-core-go/data/outport"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func responseForHandlerError(err error) (*outport.ResponseData, error) {
	if err != nil {
		return nil, status.Error(codes.Internal, "outport handler failed")
	}

	return &outport.ResponseData{}, nil
}

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
	return responseForHandlerError(os.handler.SaveBlock(in))
}

// RevertIndexedBlock forwards the call to the provided handler.
func (os *outportServer) RevertIndexedBlock(ctx context.Context, in *outport.BlockData) (*outport.ResponseData, error) {
	return responseForHandlerError(os.handler.RevertIndexedBlock(in))
}

// SaveRoundsInfo forwards the call to the provided handler.
func (os *outportServer) SaveRoundsInfo(ctx context.Context, in *outport.RoundsInfo) (*outport.ResponseData, error) {
	return responseForHandlerError(os.handler.SaveRoundsInfo(in))
}

// SaveValidatorsPubKeys forwards the call to the provided handler.
func (os *outportServer) SaveValidatorsPubKeys(ctx context.Context, in *outport.ValidatorsPubKeys) (*outport.ResponseData, error) {
	return responseForHandlerError(os.handler.SaveValidatorsPubKeys(in))
}

// SaveValidatorsRating forwards the call to the provided handler.
func (os *outportServer) SaveValidatorsRating(ctx context.Context, in *outport.ValidatorsRating) (*outport.ResponseData, error) {
	return responseForHandlerError(os.handler.SaveValidatorsRating(in))
}

// SaveAccounts forwards the call to the provided handler.
func (os *outportServer) SaveAccounts(ctx context.Context, in *outport.Accounts) (*outport.ResponseData, error) {
	return responseForHandlerError(os.handler.SaveAccounts(in))
}

// FinalizedBlockEvent forwards the call to the provided handler.
func (os *outportServer) FinalizedBlockEvent(ctx context.Context, in *outport.FinalizedBlock) (*outport.ResponseData, error) {
	return responseForHandlerError(os.handler.FinalizedBlock(in))
}

// IsInterfaceNil returns true if there is no value under the interface.
func (os *outportServer) IsInterfaceNil() bool {
	return os == nil
}
