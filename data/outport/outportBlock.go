//go:generate protoc -I=. -I=$GOPATH/src/github.com/multiversx/mx-chain-core-go/data/block -I=$GOPATH/src -I=$GOPATH/src/github.com/multiversx/protobuf/protobuf --gogoslick_out=$GOPATH/src outportBlock.proto

package outport

import (
	"github.com/multiversx/mx-chain-core-go/data"
	"github.com/multiversx/mx-chain-core-go/data/block"
)

// OutportBlockWithHeader will extend the OutportBlock structure
type OutportBlockWithHeader struct {
	*OutportBlock
	Header data.HeaderHandler
}

// HeaderDataWithBody holds header and body data
type HeaderDataWithBody struct {
	Body                 data.BodyHandler
	Header               data.HeaderHandler
	IntraShardMiniBlocks []*block.MiniBlock
	HeaderHash           []byte
}

// OutportBlockWithHeaderAndBody is a wrapper for OutportBlock used for outport handler
type OutportBlockWithHeaderAndBody struct {
	*OutportBlock
	HeaderDataWithBody *HeaderDataWithBody
}

// SetExecutionOrder sets execution order
func (t *TxInfo) SetExecutionOrder(order uint32) {
	t.ExecutionOrder = order
}

// GetTxHandler returns tx handler
func (t *TxInfo) GetTxHandler() data.TransactionHandler {
	return t.Transaction
}

// SetExecutionOrder sets execution order
func (s *SCRInfo) SetExecutionOrder(order uint32) {
	s.ExecutionOrder = order
}

// GetTxHandler returns tx handler
func (s *SCRInfo) GetTxHandler() data.TransactionHandler {
	return s.SmartContractResult
}

// SetExecutionOrder sets execution order
func (r *RewardInfo) SetExecutionOrder(order uint32) {
	r.ExecutionOrder = order
}

// GetTxHandler returns tx handler
func (r *RewardInfo) GetTxHandler() data.TransactionHandler {
	return r.Reward
}
