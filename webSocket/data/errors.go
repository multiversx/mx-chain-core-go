package data

import "errors"

// ErrNilLogger signals that a nil instance of logger has been provided
var ErrNilLogger = errors.New("nil logger")

// ErrServerIsClosed represents the error thrown by the server's ListenAndServe() function when the server is closed
var ErrServerIsClosed = errors.New("http: Server closed")

// ErrNilMarshaller signals that a nil marshaller has been provided
var ErrNilMarshaller = errors.New("nil marshaller")

// ErrNilWebSocketSender signals that a nil web socket sender has been provided
var ErrNilWebSocketSender = errors.New("nil sender sender")

// ErrWebSocketServerIsClosed signals that the web socket server was closed while trying to perform actions
var ErrWebSocketServerIsClosed = errors.New("server is closed")

// ErrNilPayloadProcessor signals that a nil payload processor has been provided
var ErrNilPayloadProcessor = errors.New("nil payload processor provided")

// ErrNilPayloadConverter signals that a nil payload converter has been provided
var ErrNilPayloadConverter = errors.New("nil payload converter provided")

// ErrEmptyUrl signals that an empty websocket url has been provided
var ErrEmptyUrl = errors.New("empty websocket url provided")

// ErrZeroValueRetryDuration signals that a zero value for retry duration has been provided
var ErrZeroValueRetryDuration = errors.New("zero value provided for retry duration")

// ErrConnectionAlreadyOpen signal that the WebSocket connection was already open
var ErrConnectionAlreadyOpen = errors.New("connection already open")

// ErrConnectionNotOpen signal that the WebSocket connection is not open
var ErrConnectionNotOpen = errors.New("connection not open")

// ErrInvalidPayloadForAckMessage signal that an invalid payload for ack message has been provided
var ErrInvalidPayloadForAckMessage = errors.New("invalid payload for ack message")

// ErrExpectedAckWasNotReceivedOnClose signals that the acknowledgment message was not received at close
var ErrExpectedAckWasNotReceivedOnClose = errors.New("expected ack message was not received on close")
