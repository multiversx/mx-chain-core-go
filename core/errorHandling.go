package core

import (
	"errors"
	"strings"
)

// IsGetNodeFromDBError returns true if the provided error is of type getNodeFromDB
func IsGetNodeFromDBError(err error) bool {
	if err == nil {
		return false
	}

	if IsClosingError(err) {
		return false
	}

	return strings.Contains(err.Error(), GetNodeFromDBErrorString)
}

// IsClosingError returns true if the provided error is used whenever the node is in the closing process
func IsClosingError(err error) bool {
	if err == nil {
		return false
	}

	errString := err.Error()
	return strings.Contains(errString, ErrDBIsClosed.Error()) ||
		strings.Contains(errString, ErrContextClosing.Error())
}

// UnwrapGetNodeFromDBErr unwraps the provided error until it finds a GetNodeFromDbErrHandler
func UnwrapGetNodeFromDBErr(wrappedErr error) GetNodeFromDbErrHandler {
	errWithKeyHandler, ok := wrappedErr.(GetNodeFromDbErrHandler)
	for !ok {
		if wrappedErr == nil {
			return nil
		}

		err := errors.Unwrap(wrappedErr)
		errWithKeyHandler, ok = err.(GetNodeFromDbErrHandler)
		wrappedErr = err
	}

	return errWithKeyHandler
}
