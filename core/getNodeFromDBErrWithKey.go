package core

import (
	"encoding/hex"
	"fmt"
)

// GetNodeFromDBErrorString represents the string which is returned when a getting node from DB returns an error
const GetNodeFromDBErrorString = "getNodeFromDB error"

// getNodeFromDBErrWithKey defines a custom error for trie get node
type getNodeFromDBErrWithKey struct {
	getErr       error
	key          []byte
	dbIdentifier string
}

// NewGetNodeFromDBErrWithKey will create a new instance of GetNodeFromDBErrWithKey
func NewGetNodeFromDBErrWithKey(key []byte, err error, id string) *getNodeFromDBErrWithKey {
	return &getNodeFromDBErrWithKey{
		getErr:       err,
		key:          key,
		dbIdentifier: id,
	}
}

// Error returns the error as string
func (e *getNodeFromDBErrWithKey) Error() string {
	return fmt.Sprintf(
		"%s: %s for key %v",
		GetNodeFromDBErrorString,
		e.getErr.Error(),
		hex.EncodeToString(e.key),
	)
}

// GetKey will return the key that generated the error
func (e *getNodeFromDBErrWithKey) GetKey() []byte {
	return e.key
}

// GetIdentifier will return the db identifier corresponding to the db
func (e *getNodeFromDBErrWithKey) GetIdentifier() string {
	return e.dbIdentifier
}

// IsInterfaceNil returns true if there is no value under the interface
func (e *getNodeFromDBErrWithKey) IsInterfaceNil() bool {
	return e == nil
}
