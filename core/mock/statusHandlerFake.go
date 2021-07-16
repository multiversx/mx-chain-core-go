package mock

// StatusHandlerFake -
type StatusHandlerFake struct {
}

// Increment -
func (s *StatusHandlerFake) Increment(_ string) {
}

// AddUint64 -
func (s *StatusHandlerFake) AddUint64(_ string, _ uint64) {
}

// Decrement -
func (s *StatusHandlerFake) Decrement(_ string) {
}

// SetInt64Value -
func (s *StatusHandlerFake) SetInt64Value(_ string, _ int64) {
}

// SetUInt64Value -
func (s *StatusHandlerFake) SetUInt64Value(_ string, _ uint64) {
}

// SetStringValue -
func (s *StatusHandlerFake) SetStringValue(_ string, _ string) {
}

// Close -
func (s *StatusHandlerFake) Close() {
}

// IsInterfaceNil -
func (s *StatusHandlerFake) IsInterfaceNil() bool {
	return false
}
