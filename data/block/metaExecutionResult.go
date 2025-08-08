package block

import (
	"math/big"

	"github.com/multiversx/mx-chain-core-go/data"
)

var _ = data.MetaExecutionResultInfoHandler(&MetaExecutionResultInfo{})
var _ = data.BaseMetaExecutionResultHandler(&BaseMetaExecutionResult{})
var _ = data.MetaExecutionResultHandler(&MetaExecutionResult{})

// GetNotarizedOnHeaderHash returns the hash of the header at the moment the execution result was notarized.
func (mm *MetaExecutionResultInfo) GetNotarizedOnHeaderHash() []byte {
	return mm.NotarizedAtHeaderHash
}

// GetExecutionResultHandler return the execution result handler
func (mm *MetaExecutionResultInfo) GetExecutionResultHandler() data.BaseMetaExecutionResultHandler {
	return mm.ExecutionResult
}

// IsInterfaceNil returns true if there is no value under the interface
func (mm *MetaExecutionResultInfo) IsInterfaceNil() bool {
	return mm == nil
}

// GetHeaderHash will return the header hash
func (bm *BaseMetaExecutionResult) GetHeaderHash() []byte {
	return bm.BaseExecutionResult.GetHeaderHash()
}

// GetHeaderNonce will return the header nonce
func (bm *BaseMetaExecutionResult) GetHeaderNonce() uint64 {
	return bm.BaseExecutionResult.GetHeaderNonce()
}

// GetHeaderRound will return the header round
func (bm *BaseMetaExecutionResult) GetHeaderRound() uint64 {
	return bm.BaseExecutionResult.GetHeaderRound()
}

// GetRootHash will return the header root hash
func (bm *BaseMetaExecutionResult) GetRootHash() []byte {
	return bm.BaseExecutionResult.GetRootHash()
}

// GetHeaderHash returns the header hash
func (mes *MetaExecutionResult) GetHeaderHash() []byte {
	return mes.ExecutionResult.GetHeaderHash()
}

// GetHeaderNonce will return the header nonce
func (mes *MetaExecutionResult) GetHeaderNonce() uint64 {
	return mes.ExecutionResult.GetHeaderNonce()
}

// GetHeaderRound will return the header round
func (mes *MetaExecutionResult) GetHeaderRound() uint64 {
	return mes.ExecutionResult.GetHeaderRound()
}

// GetRootHash will return the header root hash
func (mes *MetaExecutionResult) GetRootHash() []byte {
	return mes.ExecutionResult.GetRootHash()
}

// GetValidatorStatsRootHash will return the validatos statistics root hash
func (mes *MetaExecutionResult) GetValidatorStatsRootHash() []byte {
	return mes.ExecutionResult.GetValidatorStatsRootHash()
}

// GetAccumulatedFeesInEpoch will return the accumulated fees in epoch
func (mes *MetaExecutionResult) GetAccumulatedFeesInEpoch() *big.Int {
	return mes.ExecutionResult.GetAccumulatedFeesInEpoch()
}

func (mes *MetaExecutionResult) GetDevFeesInEpoch() *big.Int {
	return mes.ExecutionResult.GetDevFeesInEpoch()
}

// IsInterfaceNil returns true if there is no value under the interface
func (mes *MetaExecutionResult) IsInterfaceNil() bool {
	return mes == nil
}

// IsInterfaceNil returns true if there is no value under the interface
func (bme *BaseMetaExecutionResult) IsInterfaceNil() bool {
	return bme == nil
}
