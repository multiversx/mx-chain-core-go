package mock

import "github.com/ElrondNetwork/elrond-go-core/data"

// HeaderInfoStub -
type HeaderInfoStub struct {
	Header data.HeaderHandler
	Hash   []byte
}

// GetHeaderHandler -
func (his *HeaderInfoStub) GetHeaderHandler() data.HeaderHandler {
	return his.Header
}

// GetHash -
func (his *HeaderInfoStub) GetHash() []byte {
	return his.Hash
}
