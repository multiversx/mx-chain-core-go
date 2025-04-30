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

	t.Run("should work", func(t *testing.T) {
		t.Parallel()

		operation := SaveAccount + GetAccount + GetDataTrieValue
		expectedString := SaveAccountString + ", " + GetAccountString + ", " + GetDataTrieValueString + ", "

		operationString := GetOperationString(operation)
		assert.Equal(t, expectedString, operationString)
	})

	t.Run("should return with early exit", func(t *testing.T) {
		t.Parallel()

		operationString := GetOperationString(0)
		expectedOperationString := NotSetString
		assert.Equal(t, expectedOperationString, operationString)
	})
}

func TestMergeOperations(t *testing.T) {
	t.Parallel()

	operation1 := SaveAccount + GetAccount
	operation2 := GetDataTrieValue + WriteCode
	mergedOperation := MergeOperations(operation1, operation2)
	expectedMergedOperation := SaveAccount + GetAccount + GetDataTrieValue + WriteCode
	assert.Equal(t, expectedMergedOperation, mergedOperation)
}

func TestMergeDataTrieChanges(t *testing.T) {
	t.Parallel()
	t.Run("should return nil for nil slices", func(t *testing.T) {
		t.Parallel()

		mergedDataTrieChanges := MergeDataTrieChanges(nil, nil)
		assert.Nil(t, mergedDataTrieChanges)
	})

	t.Run("should work for not nil slices", func(t *testing.T) {
		t.Parallel()

		firstDataTrieChanges := []*DataTrieChange{
			{
				Type: Read,
				Key:  []byte("key1"),
			},
			{
				Type: Write,
				Key:  []byte("key2"),
				Val:  []byte("value1"),
			},
			{
				Type: Write,
				Key:  []byte("key1"),
				Val:  []byte("value1"),
			},
			{
				Type: Write,
				Key:  []byte("key1"),
				Val:  []byte("value2"),
			},
		}
		secondDataTrieChanges := []*DataTrieChange{
			{
				Type: Write,
				Key:  []byte("key1"),
				Val:  []byte("value2"),
			},
			{
				Type: Read,
				Key:  []byte("key2"),
			},
			{
				Type:    Write,
				Key:     []byte("key3"),
				Val:     []byte("value3"),
				Version: 1,
			},
			{
				Type: Read,
				Key:  []byte("key1"),
			},
			{
				Type:    Write,
				Key:     []byte("key2"),
				Val:     []byte("value1"),
				Version: 1,
			},
		}

		mergedDataTrieChanges := MergeDataTrieChanges(firstDataTrieChanges, secondDataTrieChanges)
		assert.Len(t, mergedDataTrieChanges, 5)
		assert.Equal(t, firstDataTrieChanges[0], mergedDataTrieChanges[0])
		assert.Equal(t, secondDataTrieChanges[4], mergedDataTrieChanges[1])
		assert.Equal(t, secondDataTrieChanges[0], mergedDataTrieChanges[2])
		assert.Equal(t, secondDataTrieChanges[1], mergedDataTrieChanges[3])
		assert.Equal(t, secondDataTrieChanges[2], mergedDataTrieChanges[4])
	})

}
