package stateChange

const (
	// NoChange is the default value for account changes
	NoChange = uint32(0)
	// NonceChanged is the location of the bit that represents the NonceChanged change
	NonceChanged = uint32(1)
	// BalanceChanged is the location of the bit that represents the BalanceChanged change
	BalanceChanged = uint32(2)
	// CodeHashChanged is the location of the bit that represents the CodeHashChanged change
	CodeHashChanged = uint32(4)
	// RootHashChanged is the location of the bit that represents the RootHashChanged change
	RootHashChanged = uint32(8)
	// DeveloperRewardChanged is the location of the bit that represents the DeveloperRewardChanged change
	DeveloperRewardChanged = uint32(16)
	// OwnerAddressChanged is the location of the bit that represents the OwnerAddressChanged change
	OwnerAddressChanged = uint32(32)
	// UserNameChanged is the location of the bit that represents the UserNameChanged change
	UserNameChanged = uint32(64)
	// CodeMetadataChanged is the location of the bit that represents the CodeMetadataChanged change
	CodeMetadataChanged = uint32(128)
)

// AddAccountChange adds a change to the current changes and returns the updated changes
func AddAccountChange(changes uint32, change uint32) uint32 {
	return changes | change
}
