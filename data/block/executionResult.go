package block

import "github.com/multiversx/mx-chain-core-go/data"

// GetHeaderHash returns the header hash
func (eer *ExecutionResult) GetHeaderHash() []byte {
	if eer == nil {
		return nil
	}

	return eer.BaseExecutionResult.GetHeaderHash()
}

// GetHeaderNonce returns the header nonce
func (eer *ExecutionResult) GetHeaderNonce() uint64 {
	if eer == nil {
		return 0
	}

	return eer.BaseExecutionResult.GetHeaderNonce()
}

// GetHeaderRound returns the header round
func (eer *ExecutionResult) GetHeaderRound() uint64 {
	if eer == nil {
		return 0
	}

	return eer.BaseExecutionResult.GetHeaderRound()
}

// GetRootHash returns the root hash
func (eer *ExecutionResult) GetRootHash() []byte {
	if eer == nil {
		return nil
	}

	return eer.BaseExecutionResult.GetRootHash()
}

// GetMiniBlockHeadersHandlers returns the miniblock headers handlers
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
