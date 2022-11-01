package data

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

// ErrServerIsClosed represents the error thrown by the server's ListenAndServe() function when the server is closed
var ErrServerIsClosed = errors.New("http: Server closed")

// ErrNilMarshaller signals that a nil marshaller has been provided
var ErrNilMarshaller = errors.New("nil marshaller")

// ErrNilWebSocketSender signals that a nil web socket sender has been provided
var ErrNilWebSocketSender = errors.New("nil sender sender")

// ErrWebSocketServerIsClosed signals that the web socket server was closed while trying to perform actions
var ErrWebSocketServerIsClosed = errors.New("server is closed")

// ErrWebSocketClientNotFound signals that the provided websocket client was not found
var ErrWebSocketClientNotFound = errors.New("websocket client not found")

// ErrNilWebSocketClient signals that a nil websocket client has been provided
var ErrNilWebSocketClient = errors.New("nil websocket client")
