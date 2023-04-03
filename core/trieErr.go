package core

import "strings"

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
