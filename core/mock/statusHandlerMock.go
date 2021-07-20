package mock

// StatusHandlerMock -
type StatusHandlerMock struct {
}

// Increment -
func (s *StatusHandlerMock) Increment(_ string) {
}

// AddUint64 -
func (s *StatusHandlerMock) AddUint64(_ string, _ uint64) {
}

// Decrement -
func (s *StatusHandlerMock) Decrement(_ string) {
}

// SetInt64Value -
func (s *StatusHandlerMock) SetInt64Value(_ string, _ int64) {
}

// SetUInt64Value -
func (s *StatusHandlerMock) SetUInt64Value(_ string, _ uint64) {
}

// SetStringValue -
func (s *StatusHandlerMock) SetStringValue(_ string, _ string) {
}

// Close -
func (s *StatusHandlerMock) Close() {
}

// IsInterfaceNil -
func (s *StatusHandlerMock) IsInterfaceNil() bool {
	return false
}
