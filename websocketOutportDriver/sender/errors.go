package sender

import "errors"

// ErrNilHttpServer signals that a nil http server has been provided
var ErrNilHttpServer = errors.New("nil http server")

// ErrNilUint64ByteSliceConverter signals that a nil uint64 byte slice converter has been provided
var ErrNilUint64ByteSliceConverter = errors.New("nil uint64 byte slice converter")

// ErrNilLogger signals that a nil instance of logger has been provided
var ErrNilLogger = errors.New("nil logger")

// ErrEmptyDataToSend signals that the data that should be sent via websocket is empty
var ErrEmptyDataToSend = errors.New("empty data to send")

// ErrNoClientToSendTo signals that the list of clients listening to messages is empty
var ErrNoClientToSendTo = errors.New("no client to send to")
