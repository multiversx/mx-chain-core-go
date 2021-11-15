package mock

import "github.com/ElrondNetwork/elrond-go-core/data"

type HeaderInfoStub struct {
	Header data.HeaderHandler
	Hash   []byte
}

func (his *HeaderInfoStub) GetHeaderHandler() data.HeaderHandler {
	return his.Header
}

func (his *HeaderInfoStub) GetHash() []byte {
	return his.Hash
}
