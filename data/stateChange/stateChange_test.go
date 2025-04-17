package stateChange

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetOperationFromString(t *testing.T) {
	t.Parallel()

	assert.Equal(t, NotSet, GetOperationFromString(NotSetString))
	assert.Equal(t, GetCode, GetOperationFromString(GetCodeString))
	assert.Equal(t, SaveAccount, GetOperationFromString(SaveAccountString))
	assert.Equal(t, GetAccount, GetOperationFromString(GetAccountString))
	assert.Equal(t, WriteCode, GetOperationFromString(WriteCodeString))
	assert.Equal(t, RemoveDataTrie, GetOperationFromString(RemoveDataTrieString))
	assert.Equal(t, GetDataTrieValue, GetOperationFromString(GetDataTrieValueString))
	assert.Equal(t, NotSet, GetOperationFromString("UnknownOperation"))
}

func TestGetOperationStringFromUint32(t *testing.T) {
	t.Parallel()

	assert.Equal(t, NotSetString, getOperationStringFromUint32(NotSet))
	assert.Equal(t, GetCodeString, getOperationStringFromUint32(GetCode))
	assert.Equal(t, SaveAccountString, getOperationStringFromUint32(SaveAccount))
	assert.Equal(t, GetAccountString, getOperationStringFromUint32(GetAccount))
	assert.Equal(t, WriteCodeString, getOperationStringFromUint32(WriteCode))
	assert.Equal(t, RemoveDataTrieString, getOperationStringFromUint32(RemoveDataTrie))
	assert.Equal(t, GetDataTrieValueString, getOperationStringFromUint32(GetDataTrieValue))
	assert.Equal(t, NotSetString, getOperationStringFromUint32(999))
}

func TestGetOperationString(t *testing.T) {
	t.Parallel()

	operation := SaveAccount + GetAccount + GetDataTrieValue
	expectedString := SaveAccountString + ", " + GetAccountString + ", " + GetDataTrieValueString + ", "

	operationString := GetOperationString(operation)
	assert.Equal(t, expectedString, operationString)
}

func TestMergeOperations(t *testing.T) {
	t.Parallel()

	operation1 := SaveAccount + GetAccount
	operation2 := GetDataTrieValue + WriteCode
	mergedOperation := MergeOperations(operation1, operation2)
	expectedMergedOperation := SaveAccount + GetAccount + GetDataTrieValue + WriteCode
	assert.Equal(t, expectedMergedOperation, mergedOperation)
}
