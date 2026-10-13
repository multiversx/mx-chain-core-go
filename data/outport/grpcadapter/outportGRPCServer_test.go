package grpcadapter

import (
	"context"
	"errors"
	"net"
	"testing"
	"time"

	"github.com/multiversx/mx-chain-core-go/data/mock"
	"github.com/multiversx/mx-chain-core-go/data/outport"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"
)

type outportServiceServerStub struct {
	outport.UnimplementedOutportServiceServer
}

func TestNewOutportGRPCServer(t *testing.T) {
	t.Run("empty address should error", func(t *testing.T) {
		server, err := NewOutportGRPCServer("", &mock.OutportHandlerStub{})

		require.Nil(t, server)
		require.ErrorIs(t, err, ErrEmptyOutportGRPCAddress)
	})

	t.Run("empty address with adapter should error", func(t *testing.T) {
		server, err := NewOutportGRPCServerWithAdapter("", &outportServiceServerStub{})

		require.Nil(t, server)
		require.ErrorIs(t, err, ErrEmptyOutportGRPCAddress)
	})

	t.Run("nil listener should error", func(t *testing.T) {
		server, err := NewOutportGRPCServerOnListener(nil, &mock.OutportHandlerStub{})

		require.Nil(t, server)
		require.ErrorIs(t, err, ErrNilOutportGRPCListener)
	})

	t.Run("nil adapter should error", func(t *testing.T) {
		server, err := NewOutportGRPCServerWithAdapter("127.0.0.1:0", nil)

		require.Nil(t, server)
		require.ErrorIs(t, err, ErrNilOutportServiceServer)
	})

	t.Run("typed nil adapter should error", func(t *testing.T) {
		var adapter *outportServiceServerStub

		server, err := NewOutportGRPCServerWithAdapter("127.0.0.1:0", adapter)

		require.Nil(t, server)
		require.ErrorIs(t, err, ErrNilOutportServiceServer)
	})

	t.Run("listens on an ephemeral address", func(t *testing.T) {
		server, err := NewOutportGRPCServer("127.0.0.1:0", &mock.OutportHandlerStub{})

		require.NoError(t, err)
		require.NotEmpty(t, server.Address())
		require.NoError(t, server.Close())
	})
}

func TestOutportGRPCRoundTrip(t *testing.T) {
	expectedResponse := &outport.ResponseData{}
	expectedBlock := &outport.OutportBlock{ShardID: 7, HighestFinalBlockHash: make([]byte, 5*1024*1024)}
	expectedBlockData := &outport.BlockData{ShardID: 8}
	expectedRounds := &outport.RoundsInfo{}
	expectedValidatorsPubKeys := &outport.ValidatorsPubKeys{ShardID: 9}
	expectedValidatorsRating := &outport.ValidatorsRating{ShardID: 10}
	expectedAccounts := &outport.Accounts{ShardID: 11}
	expectedFinalizedBlock := &outport.FinalizedBlock{ShardID: 12}
	listener := bufconn.Listen(16 * 1024 * 1024)

	server, err := NewOutportGRPCServerOnListener(listener, &mock.OutportHandlerStub{
		SaveBlockCalled: func(in *outport.OutportBlock) error {
			if in.ShardID == 99 {
				return errors.New("sensitive handler detail")
			}
			assert.Equal(t, expectedBlock, in)
			return nil
		},
		RevertIndexedBlockCalled: func(in *outport.BlockData) error {
			assert.Equal(t, expectedBlockData, in)
			return nil
		},
		SaveRoundsInfoCalled: func(in *outport.RoundsInfo) error {
			assert.Equal(t, expectedRounds, in)
			return nil
		},
		SaveValidatorsPubKeysCalled: func(in *outport.ValidatorsPubKeys) error {
			assert.Equal(t, expectedValidatorsPubKeys, in)
			return nil
		},
		SaveValidatorsRatingCalled: func(in *outport.ValidatorsRating) error {
			assert.Equal(t, expectedValidatorsRating, in)
			return nil
		},
		SaveAccountsCalled: func(in *outport.Accounts) error {
			assert.Equal(t, expectedAccounts, in)
			return nil
		},
		FinalizedBlockCalled: func(in *outport.FinalizedBlock) error {
			assert.Equal(t, expectedFinalizedBlock, in)
			return nil
		},
	})
	require.NoError(t, err)

	serveErrChan := make(chan error, 1)
	go func() {
		serveErrChan <- server.Start()
	}()
	t.Cleanup(func() {
		_ = server.Close()
	})

	client, err := NewOutportGRPCClient(
		"passthrough:///bufnet",
		grpc.WithContextDialer(func(ctx context.Context, _ string) (net.Conn, error) {
			return listener.DialContext(ctx)
		}),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	require.NoError(t, err)
	t.Cleanup(func() {
		_ = client.Close()
	})

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	response, err := client.SaveBlock(ctx, expectedBlock)
	require.NoError(t, err)
	require.Equal(t, expectedResponse, response)

	_, err = client.SaveBlock(ctx, &outport.OutportBlock{ShardID: 99})
	require.Equal(t, codes.Internal, status.Code(err))
	require.Equal(t, "outport handler failed", status.Convert(err).Message())

	response, err = client.RevertIndexedBlock(ctx, expectedBlockData)
	require.NoError(t, err)
	require.Equal(t, expectedResponse, response)

	response, err = client.SaveRoundsInfo(ctx, expectedRounds)
	require.NoError(t, err)
	require.Equal(t, expectedResponse, response)

	response, err = client.SaveValidatorsPubKeys(ctx, expectedValidatorsPubKeys)
	require.NoError(t, err)
	require.Equal(t, expectedResponse, response)

	response, err = client.SaveValidatorsRating(ctx, expectedValidatorsRating)
	require.NoError(t, err)
	require.Equal(t, expectedResponse, response)

	response, err = client.SaveAccounts(ctx, expectedAccounts)
	require.NoError(t, err)
	require.Equal(t, expectedResponse, response)

	response, err = client.FinalizedBlockEvent(ctx, expectedFinalizedBlock)
	require.NoError(t, err)
	require.Equal(t, expectedResponse, response)

	require.NoError(t, server.Close())
	select {
	case serveErr := <-serveErrChan:
		require.True(t, serveErr == nil || errors.Is(serveErr, grpc.ErrServerStopped))
	case <-time.After(time.Second):
		t.Fatal("timed out waiting for grpc server to stop")
	}
}
