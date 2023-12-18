package core_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/multiversx/mx-chain-core-go/core"
	"github.com/multiversx/mx-chain-core-go/core/check"
	"github.com/multiversx/mx-chain-core-go/core/mock"
	"github.com/stretchr/testify/assert"
)

func TestNewTrieNodeVersionVerifier(t *testing.T) {
	t.Parallel()

	t.Run("nil enableEpochsHandler", func(t *testing.T) {
		t.Parallel()

		vv, err := core.NewTrieNodeVersionVerifier(nil)
		assert.Nil(t, vv)
		assert.Equal(t, core.ErrNilEnableEpochsHandler, err)
	})
	t.Run("incompatible enableEpochsHandler", func(t *testing.T) {
		t.Parallel()

		vv, err := core.NewTrieNodeVersionVerifier(&mock.EnableEpochsHandlerStub{
			IsFlagDefinedCalled: func(flag core.EnableEpochFlag) bool {
				assert.Equal(t, core.TestAutoBalanceDataTriesFlag, flag)
				return false
			},
		})
		assert.Nil(t, vv)
		assert.True(t, errors.Is(err, core.ErrInvalidEnableEpochsHandler))
		assert.True(t, strings.Contains(err.Error(), string(core.TestAutoBalanceDataTriesFlag)))
	})
	t.Run("new trieNodeVersionVerifier", func(t *testing.T) {
		t.Parallel()

		vv, err := core.NewTrieNodeVersionVerifier(&mock.EnableEpochsHandlerStub{
			IsFlagDefinedCalled: func(flag core.EnableEpochFlag) bool {
				return flag == core.TestAutoBalanceDataTriesFlag
			},
		})
		assert.Nil(t, err)
		assert.False(t, check.IfNil(vv))
	})
}

func TestTrieNodeVersionVerifier_IsValidVersion(t *testing.T) {
	t.Parallel()

	t.Run("auto balance enabled", func(t *testing.T) {
		t.Parallel()

		vv, _ := core.NewTrieNodeVersionVerifier(
			&mock.EnableEpochsHandlerStub{
				IsFlagDefinedCalled: func(flag core.EnableEpochFlag) bool {
					return flag == core.TestAutoBalanceDataTriesFlag
				},
				IsFlagEnabledCalled: func(flag core.EnableEpochFlag) bool {
					return flag == core.TestAutoBalanceDataTriesFlag
				},
			},
		)
		assert.True(t, vv.IsValidVersion(core.NotSpecified))
		assert.True(t, vv.IsValidVersion(core.AutoBalanceEnabled))
		assert.False(t, vv.IsValidVersion(core.AutoBalanceEnabled+1))
	})

	t.Run("auto balance disabled", func(t *testing.T) {
		t.Parallel()

		vv, _ := core.NewTrieNodeVersionVerifier(
			&mock.EnableEpochsHandlerStub{
				IsFlagDefinedCalled: func(flag core.EnableEpochFlag) bool {
					return flag == core.TestAutoBalanceDataTriesFlag
				},
				IsFlagEnabledCalled: func(flag core.EnableEpochFlag) bool {
					return false
				},
			},
		)
		assert.True(t, vv.IsValidVersion(core.NotSpecified))
		assert.False(t, vv.IsValidVersion(core.AutoBalanceEnabled))
		assert.False(t, vv.IsValidVersion(core.AutoBalanceEnabled+1))
	})
}

func TestTrieNodeVersion_String(t *testing.T) {
	t.Parallel()

	assert.Equal(t, core.NotSpecifiedString, core.NotSpecified.String())
	assert.Equal(t, core.AutoBalanceEnabledString, core.AutoBalanceEnabled.String())
	assert.Equal(t, "unknown: 100", core.TrieNodeVersion(100).String())
}

func TestGetVersionForNewData(t *testing.T) {
	t.Parallel()

	t.Run("auto balance enabled", func(t *testing.T) {
		t.Parallel()

		getVersionForNewData := core.GetVersionForNewData(
			&mock.EnableEpochsHandlerStub{
				IsFlagDefinedCalled: func(flag core.EnableEpochFlag) bool {
					return flag == core.TestAutoBalanceDataTriesFlag
				},
				IsFlagEnabledCalled: func(flag core.EnableEpochFlag) bool {
					return flag == core.TestAutoBalanceDataTriesFlag
				},
			},
		)
		assert.Equal(t, core.AutoBalanceEnabled, getVersionForNewData)
	})

	t.Run("auto balance disabled", func(t *testing.T) {
		t.Parallel()

		getVersionForNewData := core.GetVersionForNewData(
			&mock.EnableEpochsHandlerStub{
				IsFlagDefinedCalled: func(flag core.EnableEpochFlag) bool {
					return flag == core.TestAutoBalanceDataTriesFlag
				},
				IsFlagEnabledCalled: func(flag core.EnableEpochFlag) bool {
					return false
				},
			},
		)
		assert.Equal(t, core.NotSpecified, getVersionForNewData)
	})
}
