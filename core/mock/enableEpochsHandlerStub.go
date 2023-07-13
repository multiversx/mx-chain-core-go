package mock

// EnableEpochsHandlerStub -
type EnableEpochsHandlerStub struct {
	GetCurrentEpochCalled                      func() uint32
	IsAutoBalanceDataTriesEnabledInEpochCalled func(epoch uint32) bool
}

// GetCurrentEpoch -
func (stub *EnableEpochsHandlerStub) GetCurrentEpoch() uint32 {
	if stub.GetCurrentEpochCalled != nil {
		return stub.GetCurrentEpochCalled()
	}
	return 0
}

// IsAutoBalanceDataTriesEnabledInEpoch -
func (stub *EnableEpochsHandlerStub) IsAutoBalanceDataTriesEnabledInEpoch(epoch uint32) bool {
	if stub.IsAutoBalanceDataTriesEnabledInEpochCalled != nil {
		return stub.IsAutoBalanceDataTriesEnabledInEpochCalled(epoch)
	}

	return false
}

// IsInterfaceNil returns true if there is no value under the interface
func (stub *EnableEpochsHandlerStub) IsInterfaceNil() bool {
	return stub == nil
}
