//go:generate protoc -I=. -I=$GOPATH/src -I=$GOPATH/src/github.com/multiversx/protobuf/protobuf  --gogoslick_out=. executionResult.proto

package block

import (
	"github.com/multiversx/mx-chain-core-go/data"
)

// GetHeaderHash returns the header hash
func (eer *ExecutionResult) GetHeaderHash() []byte {
	if eer == nil || eer.BaseExecutionResult == nil {
		return nil
	}

	return eer.BaseExecutionResult.GetHeaderHash()
}

// GetHeaderNonce returns the header nonce
func (eer *ExecutionResult) GetHeaderNonce() uint64 {
	if eer == nil || eer.BaseExecutionResult == nil {
		return 0
	}

	return eer.BaseExecutionResult.GetHeaderNonce()
}

// GetHeaderRound returns the header round
func (eer *ExecutionResult) GetHeaderRound() uint64 {
	if eer == nil || eer.BaseExecutionResult == nil {
		return 0
	}

	return eer.BaseExecutionResult.GetHeaderRound()
}

// GetRootHash returns the root hash
func (eer *ExecutionResult) GetRootHash() []byte {
	if eer == nil || eer.BaseExecutionResult == nil {
		return nil
	}

	return eer.BaseExecutionResult.GetRootHash()
}

// GetGasUsed returns the gas used
func (eer *ExecutionResult) GetGasUsed() uint64 {
	if eer == nil || eer.BaseExecutionResult == nil {
		return 0
	}

	return eer.BaseExecutionResult.GasUsed
}

// GetHeaderEpoch returns the header epoch
func (eer *ExecutionResult) GetHeaderEpoch() uint32 {
	if eer == nil || eer.BaseExecutionResult == nil {
		return 0
	}

	return eer.BaseExecutionResult.HeaderEpoch
}

// GetMiniBlockHeadersHandlers returns the mini block headers handlers
func (eer *ExecutionResult) GetMiniBlockHeadersHandlers() []data.MiniBlockHeaderHandler {
	if eer == nil {
		return nil
	}

	mbs := make([]data.MiniBlockHeaderHandler, 0, len(eer.GetMiniBlockHeaders()))
	for _, mb := range eer.GetMiniBlockHeaders() {
		mbCopy := mb
		mbs = append(mbs, &mbCopy)
	}

	return mbs
}

// SetMiniBlockHeadersHandlers sets the mini block headers handlers
func (eer *ExecutionResult) SetMiniBlockHeadersHandlers(mbs []data.MiniBlockHeaderHandler) error {
	if eer == nil {
		return data.ErrNilPointerReceiver
	}
	if len(mbs) == 0 {
		eer.MiniBlockHeaders = nil
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

	eer.MiniBlockHeaders = miniBlockHeaders
	return nil
}

// IsInterfaceNil returns true if there is no value under the interface
func (eer *ExecutionResult) IsInterfaceNil() bool {
	return eer == nil
}

// GetExecutionResultHandler returns the execution result handler
func (eri *ExecutionResultInfo) GetExecutionResultHandler() data.BaseExecutionResultHandler {
	if eri == nil {
		return nil
	}

	return eri.ExecutionResult
}

// IsInterfaceNil returns true if there is no value under the interface
func (eri *ExecutionResultInfo) IsInterfaceNil() bool {
	return eri == nil
}

// IsInterfaceNil returns true if there is no value under the interface
func (ber *BaseExecutionResult) IsInterfaceNil() bool {
	return ber == nil
}
