package mock

import "github.com/multiversx/mx-chain-core-go/core"

// EnableEpochsHandlerStub -
type EnableEpochsHandlerStub struct {
	IsFlagDefinedCalled               func(flag core.EnableEpochFlag) bool
	IsFlagEnabledInCurrentEpochCalled func(flag core.EnableEpochFlag) bool
}

// IsFlagDefined -
func (stub *EnableEpochsHandlerStub) IsFlagDefined(flag core.EnableEpochFlag) bool {
	if stub.IsFlagDefinedCalled != nil {
		return stub.IsFlagDefinedCalled(flag)
	}
	return false
}

// IsFlagEnabledInCurrentEpoch -
func (stub *EnableEpochsHandlerStub) IsFlagEnabledInCurrentEpoch(flag core.EnableEpochFlag) bool {
	if stub.IsFlagEnabledInCurrentEpochCalled != nil {
		return stub.IsFlagEnabledInCurrentEpochCalled(flag)
	}
	return false
}

// IsInterfaceNil -
func (stub *EnableEpochsHandlerStub) IsInterfaceNil() bool {
	return stub == nil
}
