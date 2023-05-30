package core

import (
	"encoding/hex"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testErr = errors.New("test error")
var testKey = []byte("test key")
var testId = "test id"

func TestNewStopWatch(t *testing.T) {
	t.Parallel()

	err := NewGetNodeFromDBErrWithKey(testKey, testErr, testId)
	assert.Equal(t, testErr, err.getErr)
	assert.Equal(t, testKey, err.key)
	assert.Equal(t, testId, err.dbIdentifier)
}

func TestGetNodeFromDBErrWithKey_Error(t *testing.T) {
	t.Parallel()

	expectedResult := fmt.Sprintf(
		"%s: %s for key %v",
		GetNodeFromDBErrorString,
		testErr.Error(),
		hex.EncodeToString(testKey),
	)
	err := NewGetNodeFromDBErrWithKey(testKey, testErr, testId)
	assert.Equal(t, expectedResult, err.Error())
}

func TestGetNodeFromDBErrWithKey_GetKey(t *testing.T) {
	t.Parallel()

	err := NewGetNodeFromDBErrWithKey(testKey, testErr, testId)
	assert.Equal(t, testKey, err.GetKey())
}

func TestGetNodeFromDBErrWithKey_GetIdentifier(t *testing.T) {
	t.Parallel()

	err := NewGetNodeFromDBErrWithKey(testKey, testErr, testId)
	assert.Equal(t, testId, err.GetIdentifier())
}

func TestGetNodeFromDBErrWithKey_IsInterfaceNil(t *testing.T) {
	t.Parallel()

	var err *getNodeFromDBErrWithKey
	assert.True(t, err.IsInterfaceNil())

	err = NewGetNodeFromDBErrWithKey(testKey, testErr, testId)
	assert.False(t, err.IsInterfaceNil())
}
