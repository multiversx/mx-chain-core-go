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
	if bm == nil || bm.BaseExecutionResult == nil {
		return nil
	}

	return bm.BaseExecutionResult.GetHeaderHash()
}

// GetHeaderNonce returns the header nonce
func (bm *BaseMetaExecutionResult) GetHeaderNonce() uint64 {
	if bm == nil || bm.BaseExecutionResult == nil {
		return 0
	}

	return bm.BaseExecutionResult.GetHeaderNonce()
}

// GetHeaderRound returns the header round
func (bm *BaseMetaExecutionResult) GetHeaderRound() uint64 {
	if bm == nil || bm.BaseExecutionResult == nil {
		return 0
	}

	return bm.BaseExecutionResult.GetHeaderRound()
}

// GetHeaderEpoch return the header epoch
func (bm *BaseMetaExecutionResult) GetHeaderEpoch() uint32 {
	if bm == nil || bm.BaseExecutionResult == nil {
		return 0
	}

	return bm.BaseExecutionResult.HeaderEpoch
}

// GetRootHash returns the header root hash
func (bm *BaseMetaExecutionResult) GetRootHash() []byte {
	if bm == nil || bm.BaseExecutionResult == nil {
		return nil
	}

	return bm.BaseExecutionResult.GetRootHash()
}

// GetGasUsed returns the gas used
func (bm *BaseMetaExecutionResult) GetGasUsed() uint64 {
	if bm == nil || bm.BaseExecutionResult == nil {
		return 0
	}

	return bm.BaseExecutionResult.GetGasUsed()
}

// IsInterfaceNil returns true if there is no value under the interface
func (bm *BaseMetaExecutionResult) IsInterfaceNil() bool {
	return bm == nil
}

// GetHeaderHash returns the header hash
func (mes *MetaExecutionResult) GetHeaderHash() []byte {
	if mes == nil || mes.ExecutionResult == nil {
		return nil
	}

	return mes.ExecutionResult.GetHeaderHash()
}

// GetHeaderNonce returns the header nonce
func (mes *MetaExecutionResult) GetHeaderNonce() uint64 {
	if mes == nil || mes.ExecutionResult == nil {
		return 0
	}

	return mes.ExecutionResult.GetHeaderNonce()
}

// GetMiniBlockHeadersHandlers returns the mini block headers handlers
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

// SetMiniBlockHeadersHandlers sets the mini block headers handlers
func (mes *MetaExecutionResult) SetMiniBlockHeadersHandlers(mbs []data.MiniBlockHeaderHandler) error {
	if mes == nil {
		return data.ErrNilPointerReceiver
	}
	if len(mbs) == 0 {
		mes.MiniBlockHeaders = nil
		return nil
	}

	miniBlockHeaders := make([]MiniBlockHeader, len(mbs))
	for i, mb := range mbs {
		mbHeader, ok := mb.(*MiniBlockHeader)
		if !ok {
			return data.ErrInvalidTypeAssertion
		}
		if mbHeader == nil {
			return data.ErrNilPointerDereference
		}
		miniBlockHeaders[i] = *mbHeader
	}

	mes.MiniBlockHeaders = miniBlockHeaders
	return nil
}

// GetHeaderRound returns the header round
func (mes *MetaExecutionResult) GetHeaderRound() uint64 {
	if mes == nil || mes.ExecutionResult == nil {
		return 0
	}

	return mes.ExecutionResult.GetHeaderRound()
}

// GetRootHash returns the header root hash
func (mes *MetaExecutionResult) GetRootHash() []byte {
	if mes == nil || mes.ExecutionResult == nil {
		return nil
	}

	return mes.ExecutionResult.GetRootHash()
}

// GetValidatorStatsRootHash returns the validators statistics root hash
func (mes *MetaExecutionResult) GetValidatorStatsRootHash() []byte {
	if mes == nil || mes.ExecutionResult == nil {
		return nil
	}

	return mes.ExecutionResult.GetValidatorStatsRootHash()
}

// GetAccumulatedFeesInEpoch returns the accumulated fees in epoch
func (mes *MetaExecutionResult) GetAccumulatedFeesInEpoch() *big.Int {
	if mes == nil || mes.ExecutionResult == nil {
		return nil
	}

	return mes.ExecutionResult.GetAccumulatedFeesInEpoch()
}

// GetDevFeesInEpoch returns the developer fees in epoch
func (mes *MetaExecutionResult) GetDevFeesInEpoch() *big.Int {
	if mes == nil || mes.ExecutionResult == nil {
		return nil
	}

	return mes.ExecutionResult.GetDevFeesInEpoch()
}

// GetGasUsed returns the gas used
func (mes *MetaExecutionResult) GetGasUsed() uint64 {
	if mes == nil || mes.ExecutionResult == nil {
		return 0
	}

	return mes.ExecutionResult.GetGasUsed()
}

// GetHeaderEpoch return the header epoch
func (mes *MetaExecutionResult) GetHeaderEpoch() uint32 {
	if mes == nil || mes.ExecutionResult == nil {
		return 0
	}

	return mes.ExecutionResult.GetHeaderEpoch()
}

// IsInterfaceNil returns true if there is no value under the interface
func (mes *MetaExecutionResult) IsInterfaceNil() bool {
	return mes == nil
}
