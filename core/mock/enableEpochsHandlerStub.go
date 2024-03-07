package mock

import "github.com/multiversx/mx-chain-core-go/core"

// EnableEpochsHandlerStub -
type EnableEpochsHandlerStub struct {
	IsFlagDefinedCalled        func(flag core.EnableEpochFlag) bool
	IsFlagEnabledCalled        func(flag core.EnableEpochFlag) bool
	IsFlagEnabledInEpochCalled func(flag core.EnableEpochFlag, epoch uint32) bool
	GetActivationEpochCalled   func(flag core.EnableEpochFlag) uint32
}

// IsFlagDefined -
func (stub *EnableEpochsHandlerStub) IsFlagDefined(flag core.EnableEpochFlag) bool {
	if stub.IsFlagDefinedCalled != nil {
		return stub.IsFlagDefinedCalled(flag)
	}
	return false
}

// IsFlagEnabled -
func (stub *EnableEpochsHandlerStub) IsFlagEnabled(flag core.EnableEpochFlag) bool {
	if stub.IsFlagEnabledCalled != nil {
		return stub.IsFlagEnabledCalled(flag)
	}
	return false
}

// IsFlagEnabledInEpoch -
func (stub *EnableEpochsHandlerStub) IsFlagEnabledInEpoch(flag core.EnableEpochFlag, epoch uint32) bool {
	if stub.IsFlagEnabledInEpochCalled != nil {
		return stub.IsFlagEnabledInEpochCalled(flag, epoch)
	}
	return false
}

// GetActivationEpoch -
func (stub *EnableEpochsHandlerStub) GetActivationEpoch(flag core.EnableEpochFlag) uint32 {
	if stub.GetActivationEpochCalled != nil {
		return stub.GetActivationEpochCalled(flag)
	}
	return 0
}

// IsInterfaceNil -
func (stub *EnableEpochsHandlerStub) IsInterfaceNil() bool {
	return stub == nil
}
