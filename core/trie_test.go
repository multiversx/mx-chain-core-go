package core

import (
	"testing"

	"github.com/multiversx/mx-chain-core-go/core/check"
	"github.com/multiversx/mx-chain-core-go/core/mock"
	"github.com/stretchr/testify/assert"
)

func TestNewTrieNodeVersionVerifier(t *testing.T) {
	t.Parallel()

	t.Run("nil enableEpochsHandler", func(t *testing.T) {
		t.Parallel()

		vv, err := NewTrieNodeVersionVerifier(nil)
		assert.Nil(t, vv)
		assert.Equal(t, ErrNilEnableEpochsHandler, err)
	})
	t.Run("new trieNodeVersionVerifier", func(t *testing.T) {
		t.Parallel()

		vv, err := NewTrieNodeVersionVerifier(&mock.EnableEpochsHandlerStub{})
		assert.Nil(t, err)
		assert.False(t, check.IfNil(vv))
	})
}

func TestTrieNodeVersionVerifier_IsValidVersion(t *testing.T) {
	t.Parallel()

	t.Run("auto balance enabled", func(t *testing.T) {
		t.Parallel()

		vv, _ := NewTrieNodeVersionVerifier(
			&mock.EnableEpochsHandlerStub{
				IsAutoBalanceDataTriesEnabledCalled: func() bool {
					return true
				},
			},
		)
		assert.True(t, vv.IsValidVersion(NotSpecified))
		assert.True(t, vv.IsValidVersion(AutoBalanceEnabled))
		assert.False(t, vv.IsValidVersion(AutoBalanceEnabled+1))
	})

	t.Run("auto balance disabled", func(t *testing.T) {
		t.Parallel()

		vv, _ := NewTrieNodeVersionVerifier(
			&mock.EnableEpochsHandlerStub{
				IsAutoBalanceDataTriesEnabledCalled: func() bool {
					return false
				},
			},
		)
		assert.True(t, vv.IsValidVersion(NotSpecified))
		assert.False(t, vv.IsValidVersion(AutoBalanceEnabled))
		assert.False(t, vv.IsValidVersion(AutoBalanceEnabled+1))
	})
}

func TestTrieNodeVersion_String(t *testing.T) {
	t.Parallel()

	assert.Equal(t, NotSpecifiedString, NotSpecified.String())
	assert.Equal(t, AutoBalanceEnabledString, AutoBalanceEnabled.String())
	assert.Equal(t, "unknown: 100", TrieNodeVersion(100).String())
}

func TestGetVersionForNewData(t *testing.T) {
	t.Parallel()

	t.Run("auto balance enabled", func(t *testing.T) {
		t.Parallel()

		getVersionForNewData := GetVersionForNewData(
			&mock.EnableEpochsHandlerStub{
				IsAutoBalanceDataTriesEnabledCalled: func() bool {
					return true
				},
			},
		)
		assert.Equal(t, AutoBalanceEnabled, getVersionForNewData)
	})

	t.Run("auto balance disabled", func(t *testing.T) {
		t.Parallel()

		getVersionForNewData := GetVersionForNewData(
			&mock.EnableEpochsHandlerStub{
				IsAutoBalanceDataTriesEnabledCalled: func() bool {
					return false
				},
			},
		)
		assert.Equal(t, NotSpecified, getVersionForNewData)
	})
}
