package core_test

import (
	"testing"

	"github.com/multiversx/mx-chain-core-go/core"
	"github.com/multiversx/mx-chain-core-go/core/mock"
	"github.com/stretchr/testify/require"
)

func TestCheckHandlerCompatibility(t *testing.T) {
	t.Parallel()

	err := core.CheckHandlerCompatibility(nil)
	require.Equal(t, core.ErrNilEnableEpochsHandler, err)

	allFlagsDefinedHandler := &mock.EnableEpochsHandlerStub{
		IsFlagDefinedCalled: func(flag core.EnableEpochFlag) bool {
			return true
		},
	}
	err = core.CheckHandlerCompatibility(allFlagsDefinedHandler)
	require.Nil(t, err)

	allFlagsUndefinedHandler := &mock.EnableEpochsHandlerStub{
		IsFlagDefinedCalled: func(flag core.EnableEpochFlag) bool {
			return false
		},
	}
	err = core.CheckHandlerCompatibility(allFlagsUndefinedHandler)
	require.Equal(t, core.ErrInvalidEnableEpochsHandler, err)

	oneFlagUndefinedHandler := &mock.EnableEpochsHandlerStub{
		IsFlagDefinedCalled: func(flag core.EnableEpochFlag) bool {
			return flag != core.SetGuardianFlag
		},
	}
	err = core.CheckHandlerCompatibility(oneFlagUndefinedHandler)
	require.Equal(t, core.ErrInvalidEnableEpochsHandler, err)
}
