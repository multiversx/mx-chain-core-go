package mock

// LoggerFake -
type LoggerFake struct {
}

// Trace -
func (c LoggerFake) Trace(_ string, _ ...interface{}) {
}

// Debug -
func (c LoggerFake) Debug(_ string, _ ...interface{}) {
}

// Info -
func (c LoggerFake) Info(_ string, _ ...interface{}) {
}

// Warn -
func (c LoggerFake) Warn(_ string, _ ...interface{}) {
}

// Error -
func (c LoggerFake) Error(_ string, _ ...interface{}) {
}

// LogIfError -
func (c LoggerFake) LogIfError(_ error, _ ...interface{}) {
}

// IsInterfaceNil -
func (c LoggerFake) IsInterfaceNil() bool {
	return false
}
