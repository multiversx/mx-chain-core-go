//go:generate protoc -I=. -I=$GOPATH/src/github.com/ElrondNetwork/elrond-go-core/data/block -I=$GOPATH/src -I=$GOPATH/src/github.com/ElrondNetwork/protobuf/protobuf --gogoslick_out=. slash.proto
package slash

import (
	"github.com/ElrondNetwork/elrond-go-core/data"
	"github.com/ElrondNetwork/elrond-go-core/data/block"
)

func (m *Headers) SetHeaders(headers []data.HeaderHandler) error {
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
