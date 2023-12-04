//go:generate protoc -I=. --go_out=$GOPATH/src --go-grpc_out=$GOPATH/src outGoingBridgeData.proto
package sovereign
