//go:generate protoc -I=. -I=$GOPATH/src -I=$GOPATH/src/github.com/ElrondNetwork/protobuf/protobuf  --gogoslick_out=. headerInfo.proto
package block

import "github.com/ElrondNetwork/elrond-go-core/data"

func (m *HeaderInfo) GetHeaderHandler() data.HeaderHandler {
	if m == nil {
		return nil
	}

	return m.Header
}

func (m *HeaderInfoList) SetHeadersInfo(headers []data.HeaderInfoHandler) error {
	if m == nil {
		return data.ErrNilPointerReceiver
	}

	m.Headers = nil
	for _, header := range headers {
		headerHandler := header.GetHeaderHandler()
		hdr, castOk := headerHandler.(*HeaderV2)
		if !castOk {
			return data.ErrInvalidTypeAssertion
		}

		m.Headers = append(m.Headers, &HeaderInfo{Header: hdr, Hash: header.GetHash()})
	}

	return nil
}
