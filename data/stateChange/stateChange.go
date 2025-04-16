//go:generate protoc -I=. -I=$GOPATH/src -I=$GOPATH/src/github.com/multiversx/protobuf/protobuf  --gogoslick_out=$GOPATH/src stateChange.proto

package stateChange

// GetAccessTypeBasedOnOperation returns the access type based on the operation.
func GetAccessTypeBasedOnOperation(operation Operation) ActionType {
	if operation == GetCode || operation == GetAccount {
		return Read
	}

	return Write
}
