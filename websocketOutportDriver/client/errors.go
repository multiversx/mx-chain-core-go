package client

import "errors"

var errNilPayloadProcessor = errors.New("nil payload processor provided")

var errNilPayloadParser = errors.New("nil payload parser provided")

var errNilUint64ByteSliceConverter = errors.New("nil uint64 byte slice converter provided")

var errNilWsConnReceiver = errors.New("nil ws connection receiver provided")

var errEmptyUrlProvided = errors.New("empty ws url provided")

var errNilSafeCloser = errors.New("nil safe closer provided")
