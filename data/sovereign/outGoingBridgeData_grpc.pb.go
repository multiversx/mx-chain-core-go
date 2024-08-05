// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: outGoingBridgeData.proto

package sovereign

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// BridgeTxSenderClient is the client API for BridgeTxSender service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BridgeTxSenderClient interface {
	Send(ctx context.Context, in *BridgeOperations, opts ...grpc.CallOption) (*BridgeOperationsResponse, error)
}

type bridgeTxSenderClient struct {
	cc grpc.ClientConnInterface
}

func NewBridgeTxSenderClient(cc grpc.ClientConnInterface) BridgeTxSenderClient {
	return &bridgeTxSenderClient{cc}
}

func (c *bridgeTxSenderClient) Send(ctx context.Context, in *BridgeOperations, opts ...grpc.CallOption) (*BridgeOperationsResponse, error) {
	out := new(BridgeOperationsResponse)
	err := c.cc.Invoke(ctx, "/sovereign.BridgeTxSender/Send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BridgeTxSenderServer is the server API for BridgeTxSender service.
// All implementations must embed UnimplementedBridgeTxSenderServer
// for forward compatibility
type BridgeTxSenderServer interface {
	Send(context.Context, *BridgeOperations) (*BridgeOperationsResponse, error)
	mustEmbedUnimplementedBridgeTxSenderServer()
}

// UnimplementedBridgeTxSenderServer must be embedded to have forward compatible implementations.
type UnimplementedBridgeTxSenderServer struct {
}

func (UnimplementedBridgeTxSenderServer) Send(context.Context, *BridgeOperations) (*BridgeOperationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (UnimplementedBridgeTxSenderServer) mustEmbedUnimplementedBridgeTxSenderServer() {}

// UnsafeBridgeTxSenderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BridgeTxSenderServer will
// result in compilation errors.
type UnsafeBridgeTxSenderServer interface {
	mustEmbedUnimplementedBridgeTxSenderServer()
}

func RegisterBridgeTxSenderServer(s grpc.ServiceRegistrar, srv BridgeTxSenderServer) {
	s.RegisterService(&BridgeTxSender_ServiceDesc, srv)
}

func _BridgeTxSender_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BridgeOperations)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BridgeTxSenderServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sovereign.BridgeTxSender/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BridgeTxSenderServer).Send(ctx, req.(*BridgeOperations))
	}
	return interceptor(ctx, in, info, handler)
}

// BridgeTxSender_ServiceDesc is the grpc.ServiceDesc for BridgeTxSender service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BridgeTxSender_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sovereign.BridgeTxSender",
	HandlerType: (*BridgeTxSenderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _BridgeTxSender_Send_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "outGoingBridgeData.proto",
}