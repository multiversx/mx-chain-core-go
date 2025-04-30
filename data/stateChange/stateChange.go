//go:generate protoc -I=. -I=$GOPATH/src -I=$GOPATH/src/github.com/multiversx/protobuf/protobuf  --gogoslick_out=$GOPATH/src stateChange.proto

package stateChange

import "bytes"

const (
	// NotSet is the default value for state access operation
	NotSet = uint32(0)
	// GetCode is the location of the bit that represents the GetCode operation
	GetCode = uint32(1)
	// SaveAccount is the location of the bit that represents the SaveAccount operation
	SaveAccount = uint32(2)
	// GetAccount is the location of the bit that represents the GetAccount operation
	GetAccount = uint32(4)
	// WriteCode is the location of the bit that represents the WriteCode operation
	WriteCode = uint32(8)
	// RemoveDataTrie is the location of the bit that represents the RemoveDataTrie operation
	RemoveDataTrie = uint32(16)
	// GetDataTrieValue is the location of the bit that represents the GetDataTrieValue operation
	GetDataTrieValue = uint32(32)
)

var operations = []uint32{
	NotSet,
	GetCode,
	SaveAccount,
	GetAccount,
	WriteCode,
	RemoveDataTrie,
	GetDataTrieValue,
}

const (
	// NotSetString is the string representation of the NotSet operation
	NotSetString = "NotSet"
	// GetCodeString is the string representation of the GetCode operation
	GetCodeString = "GetCode"
	// SaveAccountString is the string representation of the SaveAccount operation
	SaveAccountString = "SaveAccount"
	// GetAccountString is the string representation of the GetAccount operation
	GetAccountString = "GetAccount"
	// WriteCodeString is the string representation of the WriteCode operation
	WriteCodeString = "WriteCode"
	// RemoveDataTrieString is the string representation of the RemoveDataTrie operation
	RemoveDataTrieString = "RemoveDataTrie"
	// GetDataTrieValueString is the string representation of the GetDataTrieValue operation
	GetDataTrieValueString = "GetDataTrieValue"
)

// GetOperationFromString converts a string representation of an operation to its corresponding uint32 value.
func GetOperationFromString(operation string) uint32 {
	switch operation {
	case NotSetString:
		return NotSet
	case GetCodeString:
		return GetCode
	case SaveAccountString:
		return SaveAccount
	case GetAccountString:
		return GetAccount
	case WriteCodeString:
		return WriteCode
	case RemoveDataTrieString:
		return RemoveDataTrie
	case GetDataTrieValueString:
		return GetDataTrieValue
	default:
		return NotSet
	}
}

func getOperationStringFromUint32(operation uint32) string {
	switch operation {
	case NotSet:
		return NotSetString
	case GetCode:
		return GetCodeString
	case SaveAccount:
		return SaveAccountString
	case GetAccount:
		return GetAccountString
	case WriteCode:
		return WriteCodeString
	case RemoveDataTrie:
		return RemoveDataTrieString
	case GetDataTrieValue:
		return GetDataTrieValueString
	default:
		return NotSetString
	}
}

// GetOperationString converts a uint32 operation to its string representation.
func GetOperationString(operation uint32) string {
	operationString := ""
	if operation == 0 {
		return NotSetString
	}

	for _, op := range operations {
		if operation&op != 0 {
			operationString += getOperationStringFromUint32(op) + ", "
		}
	}

	return operationString
}

// MergeOperations combines two uint32 operations using a bitwise OR and returns the resulting value.
func MergeOperations(operation1, operation2 uint32) uint32 {
	return operation1 | operation2
}

// MergeDataTrieChanges combines two lists of DataTrieChange items into a single list, merging changes with overlapping keys.
func MergeDataTrieChanges(oldDataTrieChanges, newDataTrieChanges []*DataTrieChange) []*DataTrieChange {
	if oldDataTrieChanges == nil && newDataTrieChanges == nil {
		return nil
	}
	mergedDataTrieChanges := make([]*DataTrieChange, 0, len(oldDataTrieChanges)+len(newDataTrieChanges))
	for i := range oldDataTrieChanges {
		mergedDataTrieChanges = mergeChangeInExistingChanges(mergedDataTrieChanges, oldDataTrieChanges[i])
	}
	for i := range newDataTrieChanges {
		mergedDataTrieChanges = mergeChangeInExistingChanges(mergedDataTrieChanges, newDataTrieChanges[i])
	}

	return mergedDataTrieChanges
}

func mergeChangeInExistingChanges(
	collectedDataTrieChanges []*DataTrieChange,
	newChange *DataTrieChange,
) []*DataTrieChange {
	duplicatedEntry := false
	for i := range collectedDataTrieChanges {
		if !bytes.Equal(collectedDataTrieChanges[i].Key, newChange.Key) {
			continue
		}
		if collectedDataTrieChanges[i].Type != newChange.Type {
			continue
		}
		duplicatedEntry = true
		if newChange.Type == Read {
			continue
		}
		collectedDataTrieChanges[i].Val = newChange.Val
		collectedDataTrieChanges[i].Version = newChange.Version
	}

	if !duplicatedEntry {
		collectedDataTrieChanges = append(collectedDataTrieChanges, newChange)
	}

	return collectedDataTrieChanges
}
