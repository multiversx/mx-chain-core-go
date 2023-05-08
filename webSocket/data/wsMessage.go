//go:generate protoc -I=. -I=$GOPATH/src -I=$GOPATH/src/github.com/multiversx/protobuf/protobuf  --gogoslick_out=$GOPATH/src wsMessage.proto

package data

const (
	// AckMessage holds the identifier for an ack messages
	AckMessage = 1
	// PayloadMessage holds the identifier for a payload message
	PayloadMessage = 2
)
