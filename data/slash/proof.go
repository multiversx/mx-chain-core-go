//go:generate protoc -I=. -I=$GOPATH/src/github.com/ElrondNetwork/elrond-go-core/data/block -I=$GOPATH/src -I=$GOPATH/src/github.com/ElrondNetwork/protobuf/protobuf --gogoslick_out=. proof.proto
package slash
