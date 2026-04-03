package grpc

import (
	"context"

	"github.com/multiversx/mx-chain-core-go/core/check"
	"github.com/multiversx/mx-chain-core-go/data/outport"
)

type outportClient struct {
	client outport.OutportServiceClient
}

// NewOutportClient creates a new client wrapper over the generated gRPC client.
func NewOutportClient(client outport.OutportServiceClient) (*outportClient, error) {
	if check.IfNilReflect(client) {
		return nil, ErrNilOutportServiceClient
	}

	return &outportClient{
		client: client,
	}, nil
}

// SaveBlock forwards the call to the generated gRPC client.
func (oc *outportClient) SaveBlock(ctx context.Context, in *outport.OutportBlock) (*outport.ResponseData, error) {
	return oc.client.SaveBlock(ctx, in)
}

// RevertIndexedBlock forwards the call to the generated gRPC client.
func (oc *outportClient) RevertIndexedBlock(ctx context.Context, in *outport.BlockData) (*outport.ResponseData, error) {
	return oc.client.RevertIndexedBlock(ctx, in)
}

// SaveRoundsInfo forwards the call to the generated gRPC client.
func (oc *outportClient) SaveRoundsInfo(ctx context.Context, in *outport.RoundsInfo) (*outport.ResponseData, error) {
	return oc.client.SaveRoundsInfo(ctx, in)
}

// SaveValidatorsPubKeys forwards the call to the generated gRPC client.
func (oc *outportClient) SaveValidatorsPubKeys(ctx context.Context, in *outport.ValidatorsPubKeys) (*outport.ResponseData, error) {
	return oc.client.SaveValidatorsPubKeys(ctx, in)
}

// SaveValidatorsRating forwards the call to the generated gRPC client.
func (oc *outportClient) SaveValidatorsRating(ctx context.Context, in *outport.ValidatorsRating) (*outport.ResponseData, error) {
	return oc.client.SaveValidatorsRating(ctx, in)
}

// SaveAccounts forwards the call to the generated gRPC client.
func (oc *outportClient) SaveAccounts(ctx context.Context, in *outport.Accounts) (*outport.ResponseData, error) {
	return oc.client.SaveAccounts(ctx, in)
}

// FinalizedBlockEvent forwards the call to the generated gRPC client.
func (oc *outportClient) FinalizedBlockEvent(ctx context.Context, in *outport.FinalizedBlock) (*outport.ResponseData, error) {
	return oc.client.FinalizedBlockEvent(ctx, in)
}

// SetOutportConfig forwards the call to the generated gRPC client.
func (oc *outportClient) SetOutportConfig(ctx context.Context, in *outport.OutportConfig) (*outport.ResponseData, error) {
	return oc.client.SetOutportConfig(ctx, in)
}

// IsInterfaceNil returns true if there is no value under the interface.
func (oc *outportClient) IsInterfaceNil() bool {
	return oc == nil
}
