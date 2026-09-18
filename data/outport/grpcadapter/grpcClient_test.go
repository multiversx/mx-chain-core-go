package grpcadapter

import (
	"context"
	"errors"
	"testing"

	"github.com/multiversx/mx-chain-core-go/data/outport"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

type outportServiceClientStub struct {
	saveBlockCalled          func(ctx context.Context, in *outport.OutportBlock) (*outport.ResponseData, error)
	revertIndexedBlockCalled func(ctx context.Context, in *outport.BlockData) (*outport.ResponseData, error)
	saveRoundsInfoCalled     func(ctx context.Context, in *outport.RoundsInfo) (*outport.ResponseData, error)
	saveValidatorsPubKeys    func(ctx context.Context, in *outport.ValidatorsPubKeys) (*outport.ResponseData, error)
	saveValidatorsRating     func(ctx context.Context, in *outport.ValidatorsRating) (*outport.ResponseData, error)
	saveAccountsCalled       func(ctx context.Context, in *outport.Accounts) (*outport.ResponseData, error)
	finalizedBlockEvent      func(ctx context.Context, in *outport.FinalizedBlock) (*outport.ResponseData, error)
}

func (stub *outportServiceClientStub) SaveAccounts(ctx context.Context, in *outport.Accounts, _ ...grpc.CallOption) (*outport.ResponseData, error) {
	return stub.saveAccountsCalled(ctx, in)
}

func (stub *outportServiceClientStub) SaveBlock(ctx context.Context, in *outport.OutportBlock, _ ...grpc.CallOption) (*outport.ResponseData, error) {
	return stub.saveBlockCalled(ctx, in)
}

func (stub *outportServiceClientStub) RevertIndexedBlock(ctx context.Context, in *outport.BlockData, _ ...grpc.CallOption) (*outport.ResponseData, error) {
	return stub.revertIndexedBlockCalled(ctx, in)
}

func (stub *outportServiceClientStub) SaveRoundsInfo(ctx context.Context, in *outport.RoundsInfo, _ ...grpc.CallOption) (*outport.ResponseData, error) {
	return stub.saveRoundsInfoCalled(ctx, in)
}

func (stub *outportServiceClientStub) SaveValidatorsPubKeys(ctx context.Context, in *outport.ValidatorsPubKeys, _ ...grpc.CallOption) (*outport.ResponseData, error) {
	return stub.saveValidatorsPubKeys(ctx, in)
}

func (stub *outportServiceClientStub) SaveValidatorsRating(ctx context.Context, in *outport.ValidatorsRating, _ ...grpc.CallOption) (*outport.ResponseData, error) {
	return stub.saveValidatorsRating(ctx, in)
}

func (stub *outportServiceClientStub) FinalizedBlockEvent(ctx context.Context, in *outport.FinalizedBlock, _ ...grpc.CallOption) (*outport.ResponseData, error) {
	return stub.finalizedBlockEvent(ctx, in)
}

type outportHandlerStub struct {
	saveBlockCalled          func(in *outport.OutportBlock) error
	revertIndexedBlockCalled func(in *outport.BlockData) error
	saveRoundsInfoCalled     func(in *outport.RoundsInfo) error
	saveValidatorsPubKeys    func(in *outport.ValidatorsPubKeys) error
	saveValidatorsRating     func(in *outport.ValidatorsRating) error
	saveAccountsCalled       func(in *outport.Accounts) error
	finalizedBlockCalled     func(in *outport.FinalizedBlock) error
}

func (stub *outportHandlerStub) SaveBlock(in *outport.OutportBlock) error {
	return stub.saveBlockCalled(in)
}

func (stub *outportHandlerStub) RevertIndexedBlock(in *outport.BlockData) error {
	return stub.revertIndexedBlockCalled(in)
}

func (stub *outportHandlerStub) SaveRoundsInfo(in *outport.RoundsInfo) error {
	return stub.saveRoundsInfoCalled(in)
}

func (stub *outportHandlerStub) SaveValidatorsPubKeys(in *outport.ValidatorsPubKeys) error {
	return stub.saveValidatorsPubKeys(in)
}

func (stub *outportHandlerStub) SaveValidatorsRating(in *outport.ValidatorsRating) error {
	return stub.saveValidatorsRating(in)
}

func (stub *outportHandlerStub) SaveAccounts(in *outport.Accounts) error {
	return stub.saveAccountsCalled(in)
}

func (stub *outportHandlerStub) FinalizedBlock(in *outport.FinalizedBlock) error {
	return stub.finalizedBlockCalled(in)
}

func (stub *outportHandlerStub) IsInterfaceNil() bool {
	return stub == nil
}

func TestNewOutportClient(t *testing.T) {
	t.Run("nil client should error", func(t *testing.T) {
		client, err := NewOutportClient(nil)

		require.Nil(t, client)
		require.True(t, errors.Is(err, ErrNilOutportServiceClient))
	})

	t.Run("should work", func(t *testing.T) {
		client, err := NewOutportClient(&outportServiceClientStub{})

		require.NoError(t, err)
		require.NotNil(t, client)
	})

	t.Run("typed nil client should error", func(t *testing.T) {
		var clientStub *outportServiceClientStub

		client, err := NewOutportClient(clientStub)

		require.Nil(t, client)
		require.True(t, errors.Is(err, ErrNilOutportServiceClient))
	})
}

func TestOutportClientDelegation(t *testing.T) {
	expectedResponse := &outport.ResponseData{IndexingTimeInMs: 37}
	expectedErr := errors.New("expected error")
	expectedBlock := &outport.OutportBlock{ShardID: 1}
	expectedBlockData := &outport.BlockData{ShardID: 2}
	expectedRounds := &outport.RoundsInfo{}
	expectedValidatorsPubKeys := &outport.ValidatorsPubKeys{ShardID: 3}
	expectedValidatorsRating := &outport.ValidatorsRating{ShardID: 4}
	expectedAccounts := &outport.Accounts{ShardID: 5}
	expectedFinalizedBlock := &outport.FinalizedBlock{ShardID: 6}

	createClient := func(overrides *outportServiceClientStub) *outportClient {
		client, err := NewOutportClient(&outportServiceClientStub{
			saveBlockCalled: func(ctx context.Context, in *outport.OutportBlock) (*outport.ResponseData, error) {
				return nil, nil
			},
			revertIndexedBlockCalled: func(ctx context.Context, in *outport.BlockData) (*outport.ResponseData, error) {
				return nil, nil
			},
			saveRoundsInfoCalled: func(ctx context.Context, in *outport.RoundsInfo) (*outport.ResponseData, error) {
				return nil, nil
			},
			saveValidatorsPubKeys: func(ctx context.Context, in *outport.ValidatorsPubKeys) (*outport.ResponseData, error) {
				return nil, nil
			},
			saveValidatorsRating: func(ctx context.Context, in *outport.ValidatorsRating) (*outport.ResponseData, error) {
				return nil, nil
			},
			saveAccountsCalled: func(ctx context.Context, in *outport.Accounts) (*outport.ResponseData, error) {
				return nil, nil
			},
			finalizedBlockEvent: func(ctx context.Context, in *outport.FinalizedBlock) (*outport.ResponseData, error) {
				return nil, nil
			},
		})
		require.NoError(t, err)

		if overrides.saveBlockCalled != nil {
			client.client.(*outportServiceClientStub).saveBlockCalled = overrides.saveBlockCalled
		}
		if overrides.revertIndexedBlockCalled != nil {
			client.client.(*outportServiceClientStub).revertIndexedBlockCalled = overrides.revertIndexedBlockCalled
		}
		if overrides.saveRoundsInfoCalled != nil {
			client.client.(*outportServiceClientStub).saveRoundsInfoCalled = overrides.saveRoundsInfoCalled
		}
		if overrides.saveValidatorsPubKeys != nil {
			client.client.(*outportServiceClientStub).saveValidatorsPubKeys = overrides.saveValidatorsPubKeys
		}
		if overrides.saveValidatorsRating != nil {
			client.client.(*outportServiceClientStub).saveValidatorsRating = overrides.saveValidatorsRating
		}
		if overrides.saveAccountsCalled != nil {
			client.client.(*outportServiceClientStub).saveAccountsCalled = overrides.saveAccountsCalled
		}
		if overrides.finalizedBlockEvent != nil {
			client.client.(*outportServiceClientStub).finalizedBlockEvent = overrides.finalizedBlockEvent
		}

		return client
	}

	t.Run("SaveBlock", func(t *testing.T) {
		client := createClient(&outportServiceClientStub{
			saveBlockCalled: func(ctx context.Context, in *outport.OutportBlock) (*outport.ResponseData, error) {
				require.Equal(t, expectedBlock, in)
				require.NotNil(t, ctx)

				return expectedResponse, expectedErr
			},
		})

		response, returnedErr := client.SaveBlock(context.Background(), expectedBlock)

		require.Equal(t, expectedResponse, response)
		require.Equal(t, expectedErr, returnedErr)
	})

	t.Run("RevertIndexedBlock", func(t *testing.T) {
		client := createClient(&outportServiceClientStub{
			revertIndexedBlockCalled: func(ctx context.Context, in *outport.BlockData) (*outport.ResponseData, error) {
				require.Equal(t, expectedBlockData, in)
				require.NotNil(t, ctx)

				return expectedResponse, expectedErr
			},
		})

		response, returnedErr := client.RevertIndexedBlock(context.Background(), expectedBlockData)

		require.Equal(t, expectedResponse, response)
		require.Equal(t, expectedErr, returnedErr)
	})

	t.Run("SaveRoundsInfo", func(t *testing.T) {
		client := createClient(&outportServiceClientStub{
			saveRoundsInfoCalled: func(ctx context.Context, in *outport.RoundsInfo) (*outport.ResponseData, error) {
				require.Equal(t, expectedRounds, in)
				require.NotNil(t, ctx)

				return expectedResponse, expectedErr
			},
		})

		response, returnedErr := client.SaveRoundsInfo(context.Background(), expectedRounds)

		require.Equal(t, expectedResponse, response)
		require.Equal(t, expectedErr, returnedErr)
	})

	t.Run("SaveValidatorsPubKeys", func(t *testing.T) {
		client := createClient(&outportServiceClientStub{
			saveValidatorsPubKeys: func(ctx context.Context, in *outport.ValidatorsPubKeys) (*outport.ResponseData, error) {
				require.Equal(t, expectedValidatorsPubKeys, in)
				require.NotNil(t, ctx)

				return expectedResponse, expectedErr
			},
		})

		response, returnedErr := client.SaveValidatorsPubKeys(context.Background(), expectedValidatorsPubKeys)

		require.Equal(t, expectedResponse, response)
		require.Equal(t, expectedErr, returnedErr)
	})

	t.Run("SaveValidatorsRating", func(t *testing.T) {
		client := createClient(&outportServiceClientStub{
			saveValidatorsRating: func(ctx context.Context, in *outport.ValidatorsRating) (*outport.ResponseData, error) {
				require.Equal(t, expectedValidatorsRating, in)
				require.NotNil(t, ctx)

				return expectedResponse, expectedErr
			},
		})

		response, returnedErr := client.SaveValidatorsRating(context.Background(), expectedValidatorsRating)

		require.Equal(t, expectedResponse, response)
		require.Equal(t, expectedErr, returnedErr)
	})

	t.Run("SaveAccounts", func(t *testing.T) {
		client := createClient(&outportServiceClientStub{
			saveAccountsCalled: func(ctx context.Context, in *outport.Accounts) (*outport.ResponseData, error) {
				require.Equal(t, expectedAccounts, in)
				require.NotNil(t, ctx)

				return expectedResponse, expectedErr
			},
		})

		response, returnedErr := client.SaveAccounts(context.Background(), expectedAccounts)

		require.Equal(t, expectedResponse, response)
		require.Equal(t, expectedErr, returnedErr)
	})

	t.Run("FinalizedBlockEvent", func(t *testing.T) {
		client := createClient(&outportServiceClientStub{
			finalizedBlockEvent: func(ctx context.Context, in *outport.FinalizedBlock) (*outport.ResponseData, error) {
				require.Equal(t, expectedFinalizedBlock, in)
				require.NotNil(t, ctx)

				return expectedResponse, expectedErr
			},
		})

		response, returnedErr := client.FinalizedBlockEvent(context.Background(), expectedFinalizedBlock)

		require.Equal(t, expectedResponse, response)
		require.Equal(t, expectedErr, returnedErr)
	})
}
