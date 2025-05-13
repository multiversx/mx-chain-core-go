//go:generate protoc -I=. -I=$GOPATH/src/github.com/multiversx/mx-chain-core-go/data/transaction -I=$GOPATH/src -I=$GOPATH/src/github.com/multiversx/protobuf/protobuf --gogoslick_out=$GOPATH/src receipt.proto

package state

// NewSerializedNodesMap will create a new instance of *SerializedNodeMap
func NewSerializedNodesMap() *SerializedNodeMap {
	return &SerializedNodeMap{
		SerializedNodes: map[string][]byte{},
	}
}
