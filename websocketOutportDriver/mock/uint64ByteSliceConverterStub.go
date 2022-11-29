package mock

// Uint64ByteSliceConverterStub -
type Uint64ByteSliceConverterStub struct {
	ToByteSliceCalled func(u2 uint64) []byte
	ToUint64Called    func(bytes []byte) (uint64, error)
}

// ToByteSlice -
func (u *Uint64ByteSliceConverterStub) ToByteSlice(u2 uint64) []byte {
	if u.ToByteSliceCalled != nil {
		return u.ToByteSliceCalled(u2)
	}

	return nil
}

// ToUint64 -
func (u *Uint64ByteSliceConverterStub) ToUint64(bytes []byte) (uint64, error) {
	if u.ToUint64Called != nil {
		return u.ToUint64Called(bytes)
	}

	return 0, nil
}

// IsInterfaceNil -
func (u *Uint64ByteSliceConverterStub) IsInterfaceNil() bool {
	return u == nil
}
