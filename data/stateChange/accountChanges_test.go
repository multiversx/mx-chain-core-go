package stateChange

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddAccountChange(t *testing.T) {
	t.Parallel()

	changes := NoChange

	changes = AddAccountChange(changes, NonceChanged)
	assert.Equal(t, uint32(1), changes)
	changes = AddAccountChange(changes, BalanceChanged)
	assert.Equal(t, uint32(3), changes)
	changes = AddAccountChange(changes, CodeHashChanged)
	assert.Equal(t, uint32(7), changes)
	changes = AddAccountChange(changes, RootHashChanged)
	assert.Equal(t, uint32(15), changes)
	changes = AddAccountChange(changes, DeveloperRewardChanged)
	assert.Equal(t, uint32(31), changes)
	changes = AddAccountChange(changes, OwnerAddressChanged)
	assert.Equal(t, uint32(63), changes)
	changes = AddAccountChange(changes, UserNameChanged)
	assert.Equal(t, uint32(127), changes)
	changes = AddAccountChange(changes, CodeMetadataChanged)
	assert.Equal(t, uint32(255), changes)
}
