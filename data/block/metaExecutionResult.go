package block

import (
	"math/big"

	"github.com/multiversx/mx-chain-core-go/data"
)

var _ = data.LastMetaExecutionResultHandler(&MetaExecutionResultInfo{})
var _ = data.BaseMetaExecutionResultHandler(&BaseMetaExecutionResult{})
var _ = data.MetaExecutionResultHandler(&MetaExecutionResult{})

// GetExecutionResultHandler return the execution result handler
func (mm *MetaExecutionResultInfo) GetExecutionResultHandler() data.BaseMetaExecutionResultHandler {
	if mm == nil {
		return nil
	}

	return mm.ExecutionResult
}

// IsInterfaceNil returns true if there is no value under the interface
func (mm *MetaExecutionResultInfo) IsInterfaceNil() bool {
	return mm == nil
}

// GetHeaderHash returns the header hash
func (bm *BaseMetaExecutionResult) GetHeaderHash() []byte {
	if bm == nil {
		return nil
	}

	return bm.BaseExecutionResult.GetHeaderHash()
}

// GetHeaderNonce returns the header nonce
func (bm *BaseMetaExecutionResult) GetHeaderNonce() uint64 {
	if bm == nil {
		return 0
	}

	return bm.BaseExecutionResult.GetHeaderNonce()
}

// GetHeaderRound returns the header round
func (bm *BaseMetaExecutionResult) GetHeaderRound() uint64 {
	if bm == nil {
		return 0
	}

	return bm.BaseExecutionResult.GetHeaderRound()
}

// GetRootHash returns the header root hash
func (bm *BaseMetaExecutionResult) GetRootHash() []byte {
	if bm == nil {
		return nil
	}

	return bm.BaseExecutionResult.GetRootHash()
}

// GetHeaderEpoch return the header epoch
func (bm *BaseMetaExecutionResult) GetHeaderEpoch() uint32 {
	if bm == nil {
		return 0
	}

	return bm.BaseExecutionResult.HeaderEpoch
}

// IsInterfaceNil returns true if there is no value under the interface
func (bm *BaseMetaExecutionResult) IsInterfaceNil() bool {
	return bm == nil
}

// GetHeaderHash returns the header hash
func (mes *MetaExecutionResult) GetHeaderHash() []byte {
	if mes == nil {
		return nil
	}

	return mes.ExecutionResult.GetHeaderHash()
}

// GetMiniBlockHeadersHandlers returns the miniblock headers handlers
func (mes *MetaExecutionResult) GetMiniBlockHeadersHandlers() []data.MiniBlockHeaderHandler {
	if mes == nil {
		return nil
	}

	mbs := make([]data.MiniBlockHeaderHandler, 0, len(mes.GetMiniBlockHeaders()))
	for _, mb := range mes.GetMiniBlockHeaders() {
		mbCopy := mb
		mbs = append(mbs, &mbCopy)
	}

	return mbs
}

// GetHeaderNonce returns the header nonce
func (mes *MetaExecutionResult) GetHeaderNonce() uint64 {
	if mes == nil {
		return 0
	}

	return mes.ExecutionResult.GetHeaderNonce()
}

// GetHeaderRound returns the header round
func (mes *MetaExecutionResult) GetHeaderRound() uint64 {
	if mes == nil {
		return 0
	}

	return mes.ExecutionResult.GetHeaderRound()
}

// GetRootHash returns the header root hash
func (mes *MetaExecutionResult) GetRootHash() []byte {
	if mes == nil {
		return nil
	}

	return mes.ExecutionResult.GetRootHash()
}

// GetHeaderEpoch return the header epoch
func (mes *MetaExecutionResult) GetHeaderEpoch() uint32 {
	if mes == nil {
		return 0
	}

	return mes.ExecutionResult.GetHeaderEpoch()
}

// GetValidatorStatsRootHash returns the validators statistics root hash
func (mes *MetaExecutionResult) GetValidatorStatsRootHash() []byte {
	if mes == nil {
		return nil
	}

	return mes.ExecutionResult.GetValidatorStatsRootHash()
}

// GetAccumulatedFeesInEpoch returns the accumulated fees in epoch
func (mes *MetaExecutionResult) GetAccumulatedFeesInEpoch() *big.Int {
	if mes == nil {
		return nil
	}

	return mes.ExecutionResult.GetAccumulatedFeesInEpoch()
}

// GetDevFeesInEpoch returns the developer fees in epoch
func (mes *MetaExecutionResult) GetDevFeesInEpoch() *big.Int {
	if mes == nil {
		return nil
	}

	return mes.ExecutionResult.GetDevFeesInEpoch()
}

// IsInterfaceNil returns true if there is no value under the interface
func (mes *MetaExecutionResult) IsInterfaceNil() bool {
	return mes == nil
}
