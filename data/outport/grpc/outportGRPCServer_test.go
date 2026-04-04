package grpc

import (
	"context"
	"errors"
	"net"
	"strings"
	"testing"
	"time"

	"github.com/multiversx/mx-chain-core-go/data/outport"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

type outportServiceServerStub struct {
	outport.UnimplementedOutportServiceServer
}

func TestNewOutportGRPCServer(t *testing.T) {
	t.Run("empty address should error", func(t *testing.T) {
		server, err := NewOutportGRPCServer("", &outportHandlerStub{})

		require.Nil(t, server)
		require.True(t, errors.Is(err, ErrEmptyOutportGRPCAddress))
	})

	t.Run("empty address with adapter should error", func(t *testing.T) {
		server, err := NewOutportGRPCServerWithAdapter("", &outportServiceServerStub{})

		require.Nil(t, server)
		require.True(t, errors.Is(err, ErrEmptyOutportGRPCAddress))
	})

	t.Run("nil listener should error", func(t *testing.T) {
		server, err := NewOutportGRPCServerOnListener(nil, &outportHandlerStub{})

		require.Nil(t, server)
		require.True(t, errors.Is(err, ErrNilOutportGRPCListener))
	})

	t.Run("nil adapter should error", func(t *testing.T) {
		server, err := NewOutportGRPCServerWithAdapter("127.0.0.1:0", nil)

		require.Nil(t, server)
		require.True(t, errors.Is(err, ErrNilOutportServiceServer))
	})

	t.Run("typed nil adapter should error", func(t *testing.T) {
		var adapter *outportServiceServerStub

		server, err := NewOutportGRPCServerWithAdapter("127.0.0.1:0", adapter)

		require.Nil(t, server)
		require.True(t, errors.Is(err, ErrNilOutportServiceServer))
	})
}

func TestOutportGRPCRoundTrip(t *testing.T) {
	expectedBlock := &outport.OutportBlock{ShardID: 7}
	expectedBlockData := &outport.BlockData{ShardID: 8}
	expectedRounds := &outport.RoundsInfo{}
	expectedValidatorsPubKeys := &outport.ValidatorsPubKeys{ShardID: 9}
	expectedValidatorsRating := &outport.ValidatorsRating{ShardID: 10}
	expectedAccounts := &outport.Accounts{ShardID: 11}
	expectedFinalizedBlock := &outport.FinalizedBlock{ShardID: 12}
	listener := bufconn.Listen(1024 * 1024)

	server, err := NewOutportGRPCServerOnListener(listener, &outportHandlerStub{
		saveBlockCalled: func(in *outport.OutportBlock) error {
			require.Equal(t, expectedBlock, in)
			return nil
		},
		revertIndexedBlockCalled: func(in *outport.BlockData) error {
			require.Equal(t, expectedBlockData, in)
			return nil
		},
		saveRoundsInfoCalled: func(in *outport.RoundsInfo) error {
			require.Equal(t, expectedRounds, in)
			return nil
		},
		saveValidatorsPubKeys: func(in *outport.ValidatorsPubKeys) error {
			require.Equal(t, expectedValidatorsPubKeys, in)
			return nil
		},
		saveValidatorsRating: func(in *outport.ValidatorsRating) error {
			require.Equal(t, expectedValidatorsRating, in)
			return nil
		},
		saveAccountsCalled: func(in *outport.Accounts) error {
			require.Equal(t, expectedAccounts, in)
			return nil
		},
		finalizedBlockCalled: func(in *outport.FinalizedBlock) error {
			require.Equal(t, expectedFinalizedBlock, in)
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
		grpc.WithContextDialer(func(ctx context.Context, s string) (net.Conn, error) {
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

	_, err = client.SaveBlock(ctx, expectedBlock)
	require.NoError(t, err)

	_, err = client.RevertIndexedBlock(ctx, expectedBlockData)
	require.NoError(t, err)

	_, err = client.SaveRoundsInfo(ctx, expectedRounds)
	require.NoError(t, err)

	_, err = client.SaveValidatorsPubKeys(ctx, expectedValidatorsPubKeys)
	require.NoError(t, err)

	_, err = client.SaveValidatorsRating(ctx, expectedValidatorsRating)
	require.NoError(t, err)

	_, err = client.SaveAccounts(ctx, expectedAccounts)
	require.NoError(t, err)

	_, err = client.FinalizedBlockEvent(ctx, expectedFinalizedBlock)
	require.NoError(t, err)

	_ = server.Close()
	select {
	case serveErr := <-serveErrChan:
		require.True(t, serveErr == nil || errors.Is(serveErr, grpc.ErrServerStopped))
	case <-time.After(time.Second):
		t.Fatal("timed out waiting for grpc server to stop")
	}
}

func TestOutportGRPCRoundTripRealPort9876(t *testing.T) {
	expectedResponse := &outport.ResponseData{}
	expectedBlock := &outport.OutportBlock{ShardID: 9}
	expectedBlockData := &outport.BlockData{ShardID: 10}
	expectedRounds := &outport.RoundsInfo{}
	expectedValidatorsPubKeys := &outport.ValidatorsPubKeys{ShardID: 11}
	expectedValidatorsRating := &outport.ValidatorsRating{ShardID: 12}
	expectedAccounts := &outport.Accounts{ShardID: 13}
	expectedFinalizedBlock := &outport.FinalizedBlock{ShardID: 14}

	listener, err := net.Listen("tcp", "127.0.0.1:9876")
	if err != nil {
		errMsg := err.Error()
		if strings.Contains(errMsg, "operation not permitted") || strings.Contains(errMsg, "address already in use") {
			t.Skipf("skipping real-port grpc test: %v", err)
		}

		require.NoError(t, err)
	}
	t.Cleanup(func() {
		_ = listener.Close()
	})

	server, err := NewOutportGRPCServerOnListener(listener, &outportHandlerStub{
		saveBlockCalled: func(in *outport.OutportBlock) error {
			require.Equal(t, expectedBlock, in)
			return nil
		},
		revertIndexedBlockCalled: func(in *outport.BlockData) error {
			require.Equal(t, expectedBlockData, in)
			return nil
		},
		saveRoundsInfoCalled: func(in *outport.RoundsInfo) error {
			require.Equal(t, expectedRounds, in)
			return nil
		},
		saveValidatorsPubKeys: func(in *outport.ValidatorsPubKeys) error {
			require.Equal(t, expectedValidatorsPubKeys, in)
			return nil
		},
		saveValidatorsRating: func(in *outport.ValidatorsRating) error {
			require.Equal(t, expectedValidatorsRating, in)
			return nil
		},
		saveAccountsCalled: func(in *outport.Accounts) error {
			require.Equal(t, expectedAccounts, in)
			return nil
		},
		finalizedBlockCalled: func(in *outport.FinalizedBlock) error {
			require.Equal(t, expectedFinalizedBlock, in)
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
		"127.0.0.1:9876",
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

	_ = server.Close()
	select {
	case serveErr := <-serveErrChan:
		require.True(t, serveErr == nil || errors.Is(serveErr, grpc.ErrServerStopped))
	case <-time.After(time.Second):
		t.Fatal("timed out waiting for grpc server to stop")
	}
}
