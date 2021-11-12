//go:generate protoc -I=. -I=$GOPATH/src/github.com/ElrondNetwork/elrond-go-core/data/block -I=$GOPATH/src -I=$GOPATH/src/github.com/ElrondNetwork/protobuf/protobuf --gogoslick_out=. slash.proto
package slash

import (
	"github.com/ElrondNetwork/elrond-go-core/data"
	"github.com/ElrondNetwork/elrond-go-core/data/block"
)

// GetHeaderHandlers returns a slice with all data.HeaderHandler.
// Used so we can work with interfaces instead of defined types
func (m *HeadersV2) GetHeaderHandlers() []data.HeaderHandler {
	if m == nil {
		return nil
	}

	ret := make([]data.HeaderHandler, 0, len(m.Headers))

	for _, header := range m.Headers {
		ret = append(ret, header)
	}

	return ret
}

// SetHeaders sets internal header structs to a given slice of data.HeaderHandler.
// Used so we can work with interfaces instead of defined types
func (m *HeadersV2) SetHeaders(headers []data.HeaderHandler) error {
	if m == nil {
		return data.ErrNilPointerReceiver
	}

	for _, header := range headers {
		hdr, castOk := header.(*block.HeaderV2)
		if !castOk {
			return data.ErrInvalidTypeAssertion
		}

		m.Headers = append(m.Headers, hdr)
	}

	return nil
}
