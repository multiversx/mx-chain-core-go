package marshal

import (
	gproto "github.com/gogo/protobuf/proto"
	proto "github.com/golang/protobuf/proto" //nolint TODO:deprecated
)

// GogoProtoObj groups the necessary of a gogo protobuf marshalizeble object
type GogoProtoObj interface {
	Size() int
	MarshalToSizedBuffer(dAtA []byte) (int, error)
	gproto.Marshaler
	gproto.Unmarshaler
	proto.Message
}

// Marshalizer defines the 2 basic operations: serialize (marshal) and deserialize (unmarshal)
type Marshalizer interface {
	Marshal(obj interface{}) ([]byte, error)
	Unmarshal(obj interface{}, buff []byte) error
	IsInterfaceNil() bool
}

type MarshalizerWithExtraSize interface {
	MarshalWithExtraCapacity(obj interface{}, extraCapacity int) ([]byte, error)
	Marshal(obj interface{}) ([]byte, error)
	Unmarshal(obj interface{}, buff []byte) error
	IsInterfaceNil() bool
}
