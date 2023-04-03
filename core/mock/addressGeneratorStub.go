package mock

// AddressGeneratorStub is a mock implementation of AddressGenerator interface
type AddressGeneratorStub struct {
	NewAddressCalled func(address []byte, nonce uint64, vmType []byte) ([]byte, error)
}

// NewAddress is a mock implementation of NewAddress method
func (ags *AddressGeneratorStub) NewAddress(address []byte, nonce uint64, vmType []byte) ([]byte, error) {
	if ags.NewAddressCalled != nil {
		return ags.NewAddressCalled(address, nonce, vmType)
	}
	return nil, nil
}

// IsInterfaceNil returns true if there is no value under the interface
func (ags *AddressGeneratorStub) IsInterfaceNil() bool {
	return ags == nil
}
