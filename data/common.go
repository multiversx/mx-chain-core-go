package data

import (
	"github.com/multiversx/mx-chain-core-go/core"
	"github.com/multiversx/mx-chain-core-go/core/check"
	"github.com/multiversx/mx-chain-core-go/data/block"
	"github.com/multiversx/mx-chain-core-go/marshal"
)

// TrimHeaderHandlerSlice creates a copy of the provided slice without the excess capacity
func TrimHeaderHandlerSlice(in []HeaderHandler) []HeaderHandler {
	if len(in) == 0 {
		return []HeaderHandler{}
	}
	ret := make([]HeaderHandler, len(in))
	copy(ret, in)
	return ret
}

// GetHeaderBytesAndType returns the marshalled header bytes along with header type, if known
func GetHeaderBytesAndType(marshaller marshal.Marshalizer, headerHandler HeaderHandler) ([]byte, core.HeaderType, error) {
	var err error
	var headerBytes []byte
	var headerType core.HeaderType

	switch header := headerHandler.(type) {
	case *block.MetaBlock:
		headerType = core.MetaHeader
		headerBytes, err = marshaller.Marshal(header)
	case *block.Header:
		headerType = core.ShardHeaderV1
		headerBytes, err = marshaller.Marshal(header)
	case *block.HeaderV2:
		headerType = core.ShardHeaderV2
		headerBytes, err = marshaller.Marshal(header)
	default:
		return nil, "", errInvalidHeaderType
	}

	return headerBytes, headerType, err
}

// GetBody converts the BodyHandler interface to Body struct
func GetBody(bodyHandler BodyHandler) (*block.Body, error) {
	if check.IfNil(bodyHandler) {
		return nil, errNilBodyHandler
	}

	body, castOk := bodyHandler.(*block.Body)
	if !castOk {
		return nil, errCannotCastBlockBody
	}

	return body, nil
}
