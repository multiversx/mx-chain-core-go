package mock

// AppStatusHandlerStub is a stub implementation of AppStatusHandler
type AppStatusHandlerStub struct {
	IncrementHandler      func(key string)
	DecrementHandler      func(key string)
	SetUInt64ValueHandler func(key string, value uint64)
	SetInt64ValueHandler  func(key string, value int64)
	GetValueHandler       func(key string) float64
	CloseHandler          func(key string)
}

// Increment will call the handler of the stub for incrementing
func (ashs *AppStatusHandlerStub) Increment(key string) {
	ashs.IncrementHandler(key)
}

// Decrement will call the handler of the stub for decrementing
func (ashs *AppStatusHandlerStub) Decrement(key string) {
	ashs.DecrementHandler(key)
}

// SetInt64Value will call the handler of the stub for setting an int64 value
func (ashs *AppStatusHandlerStub) SetInt64Value(key string, value int64) {
	ashs.SetInt64ValueHandler(key, value)
}

// SetUInt64Value will call the handler of the stub for setting an uint64 value
func (ashs *AppStatusHandlerStub) SetUInt64Value(key string, value uint64) {
	ashs.SetUInt64ValueHandler(key, value)
}

// GetValue will call the handler of the stub for getting a value
func (ashs *AppStatusHandlerStub) GetValue(key string) float64 {
	return ashs.GetValueHandler(key)
}

// Close will call the handler of the stub for closing
func (psh *AppStatusHandlerStub) Close() {
}
