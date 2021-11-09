package factory

import "errors"

// ErrNilMarshaller signals a nil marshaller has been provided
var ErrNilMarshaller = errors.New("nil marshaller")

// ErrNilUint64ByteSliceConverter signals that a nil uint64 byte slice converter has been provided
var ErrNilUint64ByteSliceConverter = errors.New("nil uint64 byte slice converter")

// ErrNilLogger signals that a nil logger instance has been provided
var ErrNilLogger = errors.New("nil logger")
