//go:generate protoc -I=. -I=$GOPATH/src/github.com/multiversx/mx-chain-core-go/data/block -I=$GOPATH/src -I=$GOPATH/src/github.com/multiversx/protobuf/protobuf --gogoslick_out=$GOPATH/src outportBlock.proto

package outport

import (
	"github.com/multiversx/mx-chain-core-go/data"
)

// OutportBlockWithHeader will extend the OutportBlock structure
type OutportBlockWithHeader struct {
	*OutportBlock
	Header data.HeaderHandler
}

// OutportBlockWithHeaderAndBody is a wrapper for OutportBlock used for outport handler
type OutportBlockWithHeaderAndBody struct {
	*OutportBlock
	Header data.HeaderHandler
	Body   data.BodyHandler
}

type HeaderDataWithBody struct {
	Body       data.BodyHandler
	Header     data.HeaderHandler
	HeaderHash []byte
}

func (t *TxInfo) SetExecutionOrder(order uint32) {
	t.ExecutionOrder = order
}

func (t *TxInfo) GetTxHandler() data.TransactionHandler {
	return t.Transaction
}

func (s *SCRInfo) SetExecutionOrder(order uint32) {
	s.ExecutionOrder = order
}

func (s *SCRInfo) GetTxHandler() data.TransactionHandler {
	return s.SmartContractResult
}

func (r *RewardInfo) SetExecutionOrder(order uint32) {
	r.ExecutionOrder = order
}

func (r *RewardInfo) GetTxHandler() data.TransactionHandler {
	return r.Reward
}
