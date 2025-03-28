//go:generate protoc -I=. -I=$GOPATH/src -I=$GOPATH/src/github.com/multiversx/protobuf/protobuf  --gogoslick_out=. headerProof.proto
package block

// IsInterfaceNil returns true if there is no value under the interface
func (x *HeaderProof) IsInterfaceNil() bool {
	return x == nil
}
