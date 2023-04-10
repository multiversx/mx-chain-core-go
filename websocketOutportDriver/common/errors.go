package common

import "errors"

var errConnectionAlreadyOpened = errors.New("connection already opened")

var errConnectionNotOpened = errors.New("connection not opened")
