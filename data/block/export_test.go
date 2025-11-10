package block

import "github.com/multiversx/mx-chain-core-go/data"

func (m *MiniBlockHeader) SetMiniBlockHeaderReserved(mbhr *MiniBlockHeaderReserved) error {
	return m.setMiniBlockHeaderReserved(mbhr)
}

func (m *MiniBlockHeader) GetMiniBlockHeaderReserved() (*MiniBlockHeaderReserved, error) {
	return m.getMiniBlockHeaderReserved()
}

// CheckBaseExecutionResultIntegrity -
func (hv3 *HeaderV3) CheckBaseExecutionResultIntegrity(ownBaseExecutionResult data.BaseExecutionResultHandler) error {
	return hv3.checkBaseExecutionResultIntegrity(ownBaseExecutionResult)
}

// CheckExecutionResultsIntegrity -
func (hv3 *HeaderV3) CheckExecutionResultsIntegrity() error {
	return hv3.checkExecutionResultsIntegrity()
}

// CheckLastExecutionResultIntegrity -
func (hv3 *HeaderV3) CheckLastExecutionResultIntegrity() error {
	return hv3.checkLastExecutionResultIntegrity()
}

// CheckBaseExecutionResultIntegrity -
func (m *MetaBlockV3) CheckBaseExecutionResultIntegrity(ownBaseExecutionResult data.BaseExecutionResultHandler) error {
	return m.checkBaseExecutionResultIntegrity(ownBaseExecutionResult)
}

func (m *MetaBlockV3) CheckBaseMetaExecutionResultIntegrity(ownBaseMetaExecutionResult *BaseMetaExecutionResult) error {
	return m.checkBaseMetaExecutionResultIntegrity(ownBaseMetaExecutionResult)
}

// CheckExecutionResultsIntegrity -
func (m *MetaBlockV3) CheckExecutionResultsIntegrity() error {
	return m.checkExecutionResultsIntegrity()
}

// CheckLastExecutionResultIntegrity -
func (m *MetaBlockV3) CheckLastExecutionResultIntegrity() error {
	return m.checkLastExecutionResultIntegrity()
}
