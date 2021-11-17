package websocketOutportDriver

import "errors"

// ErrNilMarshalizer signals that a nil marshalizer has been provided
var ErrNilMarshalizer = errors.New("nil marshalizer")

// ErrNilWebSocketSender signals that a nil web socket sender has been provided
var ErrNilWebSocketSender = errors.New("nil sender sender")

// ErrNilUint64ByteSliceConverter signals that a nil uint64 byte slice converter has been provided
var ErrNilUint64ByteSliceConverter = errors.New("nil uint64 byte slice converter")

// ErrNilLogger signals that a nil logger instance has been provided
var ErrNilLogger = errors.New("nil logger")
