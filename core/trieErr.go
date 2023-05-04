package core

import (
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
)

// GetNodeFromDBErrWithKey defines a custom error for trie get node
type GetNodeFromDBErrWithKey struct {
	getErr       error
	key          []byte
	dbIdentifier string
}

// NewGetNodeFromDBErrWithKey will create a new instance of GetNodeFromDBErrWithKey
func NewGetNodeFromDBErrWithKey(key []byte, err error, id string) *GetNodeFromDBErrWithKey {
	return &GetNodeFromDBErrWithKey{
		getErr:       err,
		key:          key,
		dbIdentifier: id,
	}
}

// Error returns the error as string
func (e *GetNodeFromDBErrWithKey) Error() string {
	return fmt.Sprintf(
		"%s: %s for key %v",
		GetNodeFromDBErrorString,
		e.getErr.Error(),
		hex.EncodeToString(e.key),
	)
}

// GetKey will return the key that generated the error
func (e *GetNodeFromDBErrWithKey) GetKey() []byte {
	return e.key
}

// GetIdentifier will return the db identifier corresponding to the db
func (e *GetNodeFromDBErrWithKey) GetIdentifier() string {
	return e.dbIdentifier
}

// IsInterfaceNil returns true if there is no value under the interface
func (e *GetNodeFromDBErrWithKey) IsInterfaceNil() bool {
	return e == nil
}

// GetNodeFromDBErrorString represents the string which is returned when a getting node from DB returns an error
const GetNodeFromDBErrorString = "getNodeFromDB error"

// IsGetNodeFromDBError returns true if the provided error is of type getNodeFromDB
func IsGetNodeFromDBError(err error) bool {
	if err == nil {
		return false
	}

	if IsClosingError(err) {
		return false
	}

	if strings.Contains(err.Error(), GetNodeFromDBErrorString) {
		return true
	}

	return false
}

// IsClosingError returns true if the provided error is used whenever the node is in the closing process
func IsClosingError(err error) bool {
	if err == nil {
		return false
	}

	return strings.Contains(err.Error(), ErrDBIsClosed.Error()) ||
		strings.Contains(err.Error(), ErrContextClosing.Error())
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
