package mock

// EnableEpochsHandlerStub -
type EnableEpochsHandlerStub struct {
	IsAutoBalanceDataTriesEnabledCalled func() bool
}

// IsAutoBalanceDataTriesEnabled -
func (e *EnableEpochsHandlerStub) IsAutoBalanceDataTriesEnabled() bool {
	if e.IsAutoBalanceDataTriesEnabledCalled != nil {
		return e.IsAutoBalanceDataTriesEnabledCalled()
	}

	return false
}

// IsInterfaceNil returns true if there is no value under the interface
func (e *EnableEpochsHandlerStub) IsInterfaceNil() bool {
	return e == nil
}
