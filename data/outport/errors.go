package outport

import "errors"

var errInvalidHeaderType = errors.New("received invalid/unknown header type")

var errNilBodyHandler = errors.New("nil body handler")

var errCannotCastBlockBody = errors.New("cannot cast block body")

var errNilHeaderProof = errors.New("nil header proof")

var errCannotCastHeaderProof = errors.New("cannot cast header proof")
