//go:generate sh -c "protoc -I=. -I=$(go env GOPATH)/src/github.com/multiversx/mx-chain-core-go/data/block -I=$(go env GOPATH)/src -I=$(go env GOPATH)/src/github.com/multiversx/protobuf/protobuf --gogoslick_out=plugins=grpc:$(go env GOPATH)/src outportService.proto"
package outport
