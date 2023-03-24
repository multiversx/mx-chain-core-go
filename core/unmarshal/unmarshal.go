package unmarshal

import (
	"errors"

	"github.com/multiversx/mx-chain-core-go/core"
	"github.com/multiversx/mx-chain-core-go/data"
	"github.com/multiversx/mx-chain-core-go/data/block"
	"github.com/multiversx/mx-chain-core-go/marshal"
)

var errInvalidHeaderType = errors.New("invalid header type")

// GetHeaderFromBytes will unmarshal the header bytes based on the header type
func GetHeaderFromBytes(marshaller marshal.Marshalizer, headerType core.HeaderType, headerBytes []byte) (data.HeaderHandler, error) {
	var header data.HeaderHandler

	switch headerType {
	case core.MetaHeader:
		header = &block.MetaBlock{}
	case core.ShardHeaderV1:
		header = &block.Header{}
	case core.ShardHeaderV2:
		header = &block.HeaderV2{}
	default:
		return nil, errInvalidHeaderType
	}

	err := marshaller.Unmarshal(header, headerBytes)
	return header, err
}
