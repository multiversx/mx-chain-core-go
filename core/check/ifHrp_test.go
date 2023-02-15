package check_test

import (
	"testing"

	"github.com/multiversx/mx-chain-core-go/core/check"
	"github.com/stretchr/testify/assert"
)

func TestCheckIfHRP_ShouldWork(t *testing.T) {
	t.Parallel()

	assert.True(t, check.IfHrp("abc"))
	assert.True(t, check.IfHrp("Abc"))
	assert.False(t, check.IfHrp("abc1"))
	assert.False(t, check.IfHrp("abc/"))
	assert.False(t, check.IfHrp("/t"))
	assert.False(t, check.IfHrp("!t"))
	assert.False(t, check.IfHrp(".t"))
	assert.False(t, check.IfHrp("`t"))
	assert.False(t, check.IfHrp(""))
	assert.False(t, check.IfHrp("123()"))
	assert.False(t, check.IfHrp(" "))
	assert.False(t, check.IfHrp("  "))
}
