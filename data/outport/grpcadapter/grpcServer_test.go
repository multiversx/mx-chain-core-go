package grpcadapter

import (
	"context"
	"errors"
	"testing"

	"github.com/multiversx/mx-chain-core-go/data/mock"
	"github.com/multiversx/mx-chain-core-go/data/outport"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestNewOutportServer(t *testing.T) {
	t.Run("nil handler should error", func(t *testing.T) {
		server, err := NewOutportServer(nil)

		require.Nil(t, server)
		require.True(t, errors.Is(err, ErrNilOutportServiceHandler))
	})

	t.Run("typed nil handler should error", func(t *testing.T) {
		var handler *mock.OutportHandlerStub

		server, err := NewOutportServer(handler)

		require.Nil(t, server)
		require.True(t, errors.Is(err, ErrNilOutportServiceHandler))
	})

	t.Run("should work", func(t *testing.T) {
		server, err := NewOutportServer(&mock.OutportHandlerStub{})

		require.NoError(t, err)
		require.NotNil(t, server)
	})
}

func TestOutportServerDelegation(t *testing.T) {
	expectedErr := errors.New("expected error")
	expectedBlock := &outport.OutportBlock{ShardID: 2}
	expectedBlockData := &outport.BlockData{ShardID: 3}
	expectedRounds := &outport.RoundsInfo{}
	expectedValidatorsPubKeys := &outport.ValidatorsPubKeys{ShardID: 4}
	expectedValidatorsRating := &outport.ValidatorsRating{ShardID: 5}
	expectedAccounts := &outport.Accounts{ShardID: 6}
	expectedFinalizedBlock := &outport.FinalizedBlock{ShardID: 7}

	createServer := func(overrides *mock.OutportHandlerStub) *outportServer {
		server, err := NewOutportServer(&mock.OutportHandlerStub{
			SaveBlockCalled: func(in *outport.OutportBlock) error {
				return nil
			},
			RevertIndexedBlockCalled: func(in *outport.BlockData) error {
				return nil
			},
			SaveRoundsInfoCalled: func(in *outport.RoundsInfo) error {
				return nil
			},
			SaveValidatorsPubKeysCalled: func(in *outport.ValidatorsPubKeys) error {
				return nil
			},
			SaveValidatorsRatingCalled: func(in *outport.ValidatorsRating) error {
				return nil
			},
			SaveAccountsCalled: func(in *outport.Accounts) error {
				return nil
			},
			FinalizedBlockCalled: func(in *outport.FinalizedBlock) error {
				return nil
			},
		})
		require.NoError(t, err)

		handler := server.handler.(*mock.OutportHandlerStub)
		if overrides.SaveBlockCalled != nil {
			handler.SaveBlockCalled = overrides.SaveBlockCalled
		}
		if overrides.RevertIndexedBlockCalled != nil {
			handler.RevertIndexedBlockCalled = overrides.RevertIndexedBlockCalled
		}
		if overrides.SaveRoundsInfoCalled != nil {
			handler.SaveRoundsInfoCalled = overrides.SaveRoundsInfoCalled
		}
		if overrides.SaveValidatorsPubKeysCalled != nil {
			handler.SaveValidatorsPubKeysCalled = overrides.SaveValidatorsPubKeysCalled
		}
		if overrides.SaveValidatorsRatingCalled != nil {
			handler.SaveValidatorsRatingCalled = overrides.SaveValidatorsRatingCalled
		}
		if overrides.SaveAccountsCalled != nil {
			handler.SaveAccountsCalled = overrides.SaveAccountsCalled
		}
		if overrides.FinalizedBlockCalled != nil {
			handler.FinalizedBlockCalled = overrides.FinalizedBlockCalled
		}

		return server
	}

	t.Run("SaveBlock", func(t *testing.T) {
		server := createServer(&mock.OutportHandlerStub{
			SaveBlockCalled: func(in *outport.OutportBlock) error {
				require.Equal(t, expectedBlock, in)
				return expectedErr
			},
		})

		_, returnedErr := server.SaveBlock(context.Background(), expectedBlock)

		require.Equal(t, codes.Internal, status.Code(returnedErr))
		require.Equal(t, "outport handler failed", status.Convert(returnedErr).Message())
	})

	t.Run("RevertIndexedBlock", func(t *testing.T) {
		server := createServer(&mock.OutportHandlerStub{
			RevertIndexedBlockCalled: func(in *outport.BlockData) error {
				require.Equal(t, expectedBlockData, in)
				return expectedErr
			},
		})

		_, returnedErr := server.RevertIndexedBlock(context.Background(), expectedBlockData)

		require.Equal(t, codes.Internal, status.Code(returnedErr))
		require.Equal(t, "outport handler failed", status.Convert(returnedErr).Message())
	})

	t.Run("SaveRoundsInfo", func(t *testing.T) {
		server := createServer(&mock.OutportHandlerStub{
			SaveRoundsInfoCalled: func(in *outport.RoundsInfo) error {
				require.Equal(t, expectedRounds, in)
				return expectedErr
			},
		})

		_, returnedErr := server.SaveRoundsInfo(context.Background(), expectedRounds)

		require.Equal(t, codes.Internal, status.Code(returnedErr))
		require.Equal(t, "outport handler failed", status.Convert(returnedErr).Message())
	})

	t.Run("SaveValidatorsPubKeys", func(t *testing.T) {
		server := createServer(&mock.OutportHandlerStub{
			SaveValidatorsPubKeysCalled: func(in *outport.ValidatorsPubKeys) error {
				require.Equal(t, expectedValidatorsPubKeys, in)
				return expectedErr
			},
		})

		_, returnedErr := server.SaveValidatorsPubKeys(context.Background(), expectedValidatorsPubKeys)

		require.Equal(t, codes.Internal, status.Code(returnedErr))
		require.Equal(t, "outport handler failed", status.Convert(returnedErr).Message())
	})

	t.Run("SaveValidatorsRating", func(t *testing.T) {
		server := createServer(&mock.OutportHandlerStub{
			SaveValidatorsRatingCalled: func(in *outport.ValidatorsRating) error {
				require.Equal(t, expectedValidatorsRating, in)
				return expectedErr
			},
		})

		_, returnedErr := server.SaveValidatorsRating(context.Background(), expectedValidatorsRating)

		require.Equal(t, codes.Internal, status.Code(returnedErr))
		require.Equal(t, "outport handler failed", status.Convert(returnedErr).Message())
	})

	t.Run("SaveAccounts", func(t *testing.T) {
		server := createServer(&mock.OutportHandlerStub{
			SaveAccountsCalled: func(in *outport.Accounts) error {
				require.Equal(t, expectedAccounts, in)
				return expectedErr
			},
		})

		_, returnedErr := server.SaveAccounts(context.Background(), expectedAccounts)

		require.Equal(t, codes.Internal, status.Code(returnedErr))
		require.Equal(t, "outport handler failed", status.Convert(returnedErr).Message())
	})

	t.Run("FinalizedBlockEvent", func(t *testing.T) {
		server := createServer(&mock.OutportHandlerStub{
			FinalizedBlockCalled: func(in *outport.FinalizedBlock) error {
				require.Equal(t, expectedFinalizedBlock, in)
				return expectedErr
			},
		})

		_, returnedErr := server.FinalizedBlockEvent(context.Background(), expectedFinalizedBlock)

		require.Equal(t, codes.Internal, status.Code(returnedErr))
		require.Equal(t, "outport handler failed", status.Convert(returnedErr).Message())
	})
}
