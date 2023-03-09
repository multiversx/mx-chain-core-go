//go:generate protoc -I=. -I=$GOPATH/src/github.com/multiversx/mx-chain-core-go/data/block -I=$GOPATH/src -I=$GOPATH/src/github.com/multiversx/protobuf/protobuf --gogoslick_out=$GOPATH/src outportBlock.proto

package outport

import "github.com/multiversx/mx-chain-core-go/data"

// OutportBlockWithHeader will extend the OutportBlock structure
type OutportBlockWithHeader struct {
	OutportBlock
	Header data.HeaderHandler
}
