package grpc

import (
	"net"
	"strings"

	"github.com/multiversx/mx-chain-core-go/core/check"
	"github.com/multiversx/mx-chain-core-go/data/outport"
	"google.golang.org/grpc"
)

// OutportGRPCServer is a bootstrapped gRPC server for the outport service.
type OutportGRPCServer struct {
	listener net.Listener
	server   *grpc.Server
}

// NewOutportGRPCServer creates a listening gRPC server and registers the outport service.
func NewOutportGRPCServer(address string, handler OutportHandler, opts ...grpc.ServerOption) (*OutportGRPCServer, error) {
	adapter, err := NewOutportServer(handler)
	if err != nil {
		return nil, err
	}

	return NewOutportGRPCServerWithAdapter(address, adapter, opts...)
}

func NewOutportGRPCServerWithAdapter(address string, adapter outport.OutportServiceServer, opts ...grpc.ServerOption) (*OutportGRPCServer, error) {
	if strings.TrimSpace(address) == "" {
		return nil, ErrEmptyOutportGRPCAddress
	}

	listener, err := net.Listen("tcp", address)
	if err != nil {
		return nil, err
	}

	server, err := newOutportGRPCServerOnListenerWithAdapter(listener, adapter, opts...)
	if err != nil {
		_ = listener.Close()
		return nil, err
	}

	return server, nil
}

// NewOutportGRPCServerOnListener creates a gRPC server on top of an existing listener and registers the outport service.
func NewOutportGRPCServerOnListener(listener net.Listener, handler OutportHandler, opts ...grpc.ServerOption) (*OutportGRPCServer, error) {
	adapter, err := NewOutportServer(handler)
	if err != nil {
		return nil, err
	}

	return newOutportGRPCServerOnListenerWithAdapter(listener, adapter, opts...)
}

func newOutportGRPCServerOnListenerWithAdapter(listener net.Listener, adapter outport.OutportServiceServer, opts ...grpc.ServerOption) (*OutportGRPCServer, error) {
	if listener == nil {
		return nil, ErrNilOutportGRPCListener
	}

	if adapter == nil || check.IfNilReflect(adapter) {
		return nil, ErrNilOutportServiceServer
	}

	server := grpc.NewServer(opts...)
	outport.RegisterOutportServiceServer(server, adapter)

	return &OutportGRPCServer{
		listener: listener,
		server:   server,
	}, nil
}

// Start starts serving requests on the configured listener.
func (ogs *OutportGRPCServer) Start() error {
	return ogs.server.Serve(ogs.listener)
}

// Close stops the underlying gRPC server and closes active listeners.
func (ogs *OutportGRPCServer) Close() error {
	if ogs == nil || ogs.server == nil {
		return nil
	}

	ogs.server.Stop()
	return nil
}

// Address returns the resolved listening address.
func (ogs *OutportGRPCServer) Address() string {
	if ogs == nil || ogs.listener == nil {
		return ""
	}

	return ogs.listener.Addr().String()
}

// IsInterfaceNil returns true if there is no value under the interface.
func (ogs *OutportGRPCServer) IsInterfaceNil() bool {
	return ogs == nil
}
