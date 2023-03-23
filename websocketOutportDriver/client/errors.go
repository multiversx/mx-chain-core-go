package client

import "errors"

var errNilOperationHandler = errors.New("nil operation handler provided")

var errEmptyUrlProvided = errors.New("empty ws url provided")
