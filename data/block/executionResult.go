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
		mbs = append(mbs, &mb)
	}

	return mbs
}

// GetExecutionResultHandler returns the execution result handler
func (eri *ExecutionResultInfo) GetExecutionResultHandler() data.BaseExecutionResultHandler {
	if eri == nil {
		return nil
	}

	return eri.ExecutionResult
}
