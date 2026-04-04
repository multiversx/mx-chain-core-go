package grpc

import "errors"

// ErrNilOutportServiceClient signals that a nil outport service client has been provided
var ErrNilOutportServiceClient = errors.New("nil outport service client")

// ErrNilOutportServiceHandler signals that a nil outport service handler has been provided
var ErrNilOutportServiceHandler = errors.New("nil outport service handler")

// ErrNilOutportServiceServer signals that a nil outport service server has been provided
var ErrNilOutportServiceServer = errors.New("nil outport service server")

// ErrEmptyOutportGRPCAddress signals that an empty gRPC address has been provided
var ErrEmptyOutportGRPCAddress = errors.New("empty outport grpc address")

// ErrNilOutportGRPCListener signals that a nil gRPC listener has been provided
var ErrNilOutportGRPCListener = errors.New("nil outport grpc listener")
