package client

import "errors"

var errNilPayloadProcessor = errors.New("nil payload processor provided")

var errNilPayloadParser = errors.New("nil payload parser provided")

var errNilUint64ByteSliceConverter = errors.New("nil uint64 byte slice converter provided")

var errNilWsConnReceiver = errors.New("nil ws connection receiver provided")

var errEmptyUrl = errors.New("empty websocket url provided")

var errZeroValueRetryDuration = errors.New("zero value provided for retry duration")

var errNilSafeCloser = errors.New("nil safe closer provided")
