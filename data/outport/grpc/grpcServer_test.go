package grpc

import (
	"context"
	"errors"
	"testing"

	"github.com/multiversx/mx-chain-core-go/data/outport"
	"github.com/stretchr/testify/require"
)

func TestNewOutportServer(t *testing.T) {
	t.Run("nil handler should error", func(t *testing.T) {
		server, err := NewOutportServer(nil)

		require.Nil(t, server)
		require.True(t, errors.Is(err, ErrNilOutportServiceHandler))
	})

	t.Run("typed nil handler should error", func(t *testing.T) {
		var handler *outportHandlerStub

		server, err := NewOutportServer(handler)

		require.Nil(t, server)
		require.True(t, errors.Is(err, ErrNilOutportServiceHandler))
	})

	t.Run("should work", func(t *testing.T) {
		server, err := NewOutportServer(&outportHandlerStub{})

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
	expectedOutportConfig := &outport.OutportConfig{ShardID: 8, IsInImportDBMode: true}

	createServer := func(overrides *outportHandlerStub) *outportServer {
		server, err := NewOutportServer(&outportHandlerStub{
			saveBlockCalled: func(in *outport.OutportBlock) error {
				return nil
			},
			revertIndexedBlockCalled: func(in *outport.BlockData) error {
				return nil
			},
			saveRoundsInfoCalled: func(in *outport.RoundsInfo) error {
				return nil
			},
			saveValidatorsPubKeys: func(in *outport.ValidatorsPubKeys) error {
				return nil
			},
			saveValidatorsRating: func(in *outport.ValidatorsRating) error {
				return nil
			},
			saveAccountsCalled: func(in *outport.Accounts) error {
				return nil
			},
			finalizedBlockCalled: func(in *outport.FinalizedBlock) error {
				return nil
			},
			setOutportConfigCalled: func(in *outport.OutportConfig) error {
				return nil
			},
		})
		require.NoError(t, err)

		handler := server.handler.(*outportHandlerStub)
		if overrides.saveBlockCalled != nil {
			handler.saveBlockCalled = overrides.saveBlockCalled
		}
		if overrides.revertIndexedBlockCalled != nil {
			handler.revertIndexedBlockCalled = overrides.revertIndexedBlockCalled
		}
		if overrides.saveRoundsInfoCalled != nil {
			handler.saveRoundsInfoCalled = overrides.saveRoundsInfoCalled
		}
		if overrides.saveValidatorsPubKeys != nil {
			handler.saveValidatorsPubKeys = overrides.saveValidatorsPubKeys
		}
		if overrides.saveValidatorsRating != nil {
			handler.saveValidatorsRating = overrides.saveValidatorsRating
		}
		if overrides.saveAccountsCalled != nil {
			handler.saveAccountsCalled = overrides.saveAccountsCalled
		}
		if overrides.finalizedBlockCalled != nil {
			handler.finalizedBlockCalled = overrides.finalizedBlockCalled
		}
		if overrides.setOutportConfigCalled != nil {
			handler.setOutportConfigCalled = overrides.setOutportConfigCalled
		}

		return server
	}

	t.Run("SaveBlock", func(t *testing.T) {
		server := createServer(&outportHandlerStub{
			saveBlockCalled: func(in *outport.OutportBlock) error {
				require.Equal(t, expectedBlock, in)
				return expectedErr
			},
		})

		_, returnedErr := server.SaveBlock(context.Background(), expectedBlock)

		require.Equal(t, expectedErr, returnedErr)
	})

	t.Run("RevertIndexedBlock", func(t *testing.T) {
		server := createServer(&outportHandlerStub{
			revertIndexedBlockCalled: func(in *outport.BlockData) error {
				require.Equal(t, expectedBlockData, in)
				return expectedErr
			},
		})

		_, returnedErr := server.RevertIndexedBlock(context.Background(), expectedBlockData)

		require.Equal(t, expectedErr, returnedErr)
	})

	t.Run("SaveRoundsInfo", func(t *testing.T) {
		server := createServer(&outportHandlerStub{
			saveRoundsInfoCalled: func(in *outport.RoundsInfo) error {
				require.Equal(t, expectedRounds, in)
				return expectedErr
			},
		})

		_, returnedErr := server.SaveRoundsInfo(context.Background(), expectedRounds)

		require.Equal(t, expectedErr, returnedErr)
	})

	t.Run("SaveValidatorsPubKeys", func(t *testing.T) {
		server := createServer(&outportHandlerStub{
			saveValidatorsPubKeys: func(in *outport.ValidatorsPubKeys) error {
				require.Equal(t, expectedValidatorsPubKeys, in)
				return expectedErr
			},
		})

		_, returnedErr := server.SaveValidatorsPubKeys(context.Background(), expectedValidatorsPubKeys)

		require.Equal(t, expectedErr, returnedErr)
	})

	t.Run("SaveValidatorsRating", func(t *testing.T) {
		server := createServer(&outportHandlerStub{
			saveValidatorsRating: func(in *outport.ValidatorsRating) error {
				require.Equal(t, expectedValidatorsRating, in)
				return expectedErr
			},
		})

		_, returnedErr := server.SaveValidatorsRating(context.Background(), expectedValidatorsRating)

		require.Equal(t, expectedErr, returnedErr)
	})

	t.Run("SaveAccounts", func(t *testing.T) {
		server := createServer(&outportHandlerStub{
			saveAccountsCalled: func(in *outport.Accounts) error {
				require.Equal(t, expectedAccounts, in)
				return expectedErr
			},
		})

		_, returnedErr := server.SaveAccounts(context.Background(), expectedAccounts)

		require.Equal(t, expectedErr, returnedErr)
	})

	t.Run("FinalizedBlockEvent", func(t *testing.T) {
		server := createServer(&outportHandlerStub{
			finalizedBlockCalled: func(in *outport.FinalizedBlock) error {
				require.Equal(t, expectedFinalizedBlock, in)
				return expectedErr
			},
		})

		_, returnedErr := server.FinalizedBlockEvent(context.Background(), expectedFinalizedBlock)

		require.Equal(t, expectedErr, returnedErr)
	})

	t.Run("SetOutportConfig", func(t *testing.T) {
		server := createServer(&outportHandlerStub{
			setOutportConfigCalled: func(in *outport.OutportConfig) error {
				require.Equal(t, expectedOutportConfig, in)
				return expectedErr
			},
		})

		_, returnedErr := server.SetOutportConfig(context.Background(), expectedOutportConfig)

		require.Equal(t, expectedErr, returnedErr)
	})
}
