package grpcadapter

import (
	"context"
	"strings"

	"github.com/multiversx/mx-chain-core-go/data/outport"
	"google.golang.org/grpc"
)

// OutportGRPCClient is a networked outport client that owns its gRPC connection.
type OutportGRPCClient struct {
	conn   *grpc.ClientConn
	client outport.OutportServiceClient
}

// NewOutportGRPCClient creates a gRPC client connected to the given target.
// Transport credentials must be provided in opts.
func NewOutportGRPCClient(target string, opts ...grpc.DialOption) (*OutportGRPCClient, error) {
	if strings.TrimSpace(target) == "" {
		return nil, ErrEmptyOutportGRPCAddress
	}

	conn, err := grpc.NewClient(target, opts...)
	if err != nil {
		return nil, err
	}

	return &OutportGRPCClient{
		conn:   conn,
		client: outport.NewOutportServiceClient(conn),
	}, nil
}

// SaveAccounts forwards the request to the remote outport service.
func (ogc *OutportGRPCClient) SaveAccounts(ctx context.Context, in *outport.Accounts) (*outport.ResponseData, error) {
	return ogc.client.SaveAccounts(ctx, in)
}

// SaveBlock forwards the request to the remote outport service.
func (ogc *OutportGRPCClient) SaveBlock(ctx context.Context, in *outport.OutportBlock) (*outport.ResponseData, error) {
	return ogc.client.SaveBlock(ctx, in)
}

// RevertIndexedBlock forwards the request to the remote outport service.
func (ogc *OutportGRPCClient) RevertIndexedBlock(ctx context.Context, in *outport.BlockData) (*outport.ResponseData, error) {
	return ogc.client.RevertIndexedBlock(ctx, in)
}

// SaveRoundsInfo forwards the request to the remote outport service.
func (ogc *OutportGRPCClient) SaveRoundsInfo(ctx context.Context, in *outport.RoundsInfo) (*outport.ResponseData, error) {
	return ogc.client.SaveRoundsInfo(ctx, in)
}

// SaveValidatorsPubKeys forwards the request to the remote outport service.
func (ogc *OutportGRPCClient) SaveValidatorsPubKeys(ctx context.Context, in *outport.ValidatorsPubKeys) (*outport.ResponseData, error) {
	return ogc.client.SaveValidatorsPubKeys(ctx, in)
}

// SaveValidatorsRating forwards the request to the remote outport service.
func (ogc *OutportGRPCClient) SaveValidatorsRating(ctx context.Context, in *outport.ValidatorsRating) (*outport.ResponseData, error) {
	return ogc.client.SaveValidatorsRating(ctx, in)
}

// FinalizedBlockEvent forwards the request to the remote outport service.
func (ogc *OutportGRPCClient) FinalizedBlockEvent(ctx context.Context, in *outport.FinalizedBlock) (*outport.ResponseData, error) {
	return ogc.client.FinalizedBlockEvent(ctx, in)
}

// Close closes the underlying gRPC connection.
func (ogc *OutportGRPCClient) Close() error {
	if ogc == nil || ogc.conn == nil {
		return nil
	}

	return ogc.conn.Close()
}

// IsInterfaceNil returns true if there is no value under the interface.
func (ogc *OutportGRPCClient) IsInterfaceNil() bool {
	return ogc == nil
}
