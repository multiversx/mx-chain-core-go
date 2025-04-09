//go:generate protoc -I=. -I=$GOPATH/src -I=$GOPATH/src/github.com/multiversx/protobuf/protobuf  --gogoslick_out=$GOPATH/src stateChange.proto

package stateChange

// SetTxHash will set the tx hash with a provided value
func (sa *StateAccess) SetTxHash(txHash []byte) {
	sa.TxHash = txHash
}

// SetIndex will set the index with a provided value
func (sa *StateAccess) SetIndex(index int32) {
	sa.Index = index
}
