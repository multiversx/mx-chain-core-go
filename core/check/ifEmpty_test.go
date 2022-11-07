package check_test

import (
	"testing"

	"github.com/ElrondNetwork/elrond-go-core/core/check"
	"github.com/stretchr/testify/assert"
)

func TestCheckIfEmpty_NoSpaceShouldRetTrue(t *testing.T) {
	t.Parallel()

	assert.True(t, check.IfEmpty(""))
}

func TestCheckIfEmpty_OneSpaceShouldRetTrue(t *testing.T) {
	t.Parallel()

	assert.True(t, check.IfEmpty(" "))
}

func TestCheckIfEmpty_TwoSpacesShouldRetTrue(t *testing.T) {
	t.Parallel()

	assert.True(t, check.IfEmpty("  "))
}
