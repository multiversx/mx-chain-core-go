package core_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/multiversx/mx-chain-core-go/core"
	"github.com/multiversx/mx-chain-core-go/core/mock"
	"github.com/stretchr/testify/require"
)

func TestCheckHandlerCompatibility(t *testing.T) {
	t.Parallel()

	err := core.CheckHandlerCompatibility(nil, []core.EnableEpochFlag{})
	require.Equal(t, core.ErrNilEnableEpochsHandler, err)

	testFlags := []core.EnableEpochFlag{"f0", "f1", "f2"}
	allFlagsDefinedHandler := &mock.EnableEpochsHandlerStub{
		IsFlagDefinedCalled: func(flag core.EnableEpochFlag) bool {
			return true
		},
	}
	err = core.CheckHandlerCompatibility(allFlagsDefinedHandler, testFlags)
	require.Nil(t, err)

	allFlagsUndefinedHandler := &mock.EnableEpochsHandlerStub{
		IsFlagDefinedCalled: func(flag core.EnableEpochFlag) bool {
			return false
		},
	}
	err = core.CheckHandlerCompatibility(allFlagsUndefinedHandler, testFlags)
	require.True(t, errors.Is(err, core.ErrInvalidEnableEpochsHandler))

	missingFlag := testFlags[1]
	oneFlagUndefinedHandler := &mock.EnableEpochsHandlerStub{
		IsFlagDefinedCalled: func(flag core.EnableEpochFlag) bool {
			return flag != missingFlag
		},
	}
	err = core.CheckHandlerCompatibility(oneFlagUndefinedHandler, testFlags)
	require.True(t, errors.Is(err, core.ErrInvalidEnableEpochsHandler))
	require.True(t, strings.Contains(err.Error(), string(missingFlag)))
}
