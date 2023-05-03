package mock

// LoggerStub -
type LoggerStub struct {
	TraceCalled      func(message string, args ...interface{})
	DebugCalled      func(message string, args ...interface{})
	InfoCalled       func(message string, args ...interface{})
	WarnCalled       func(message string, args ...interface{})
	ErrorCalled      func(message string, args ...interface{})
	LogIfErrorCalled func(err error, args ...interface{})
}

// Trace -
func (l LoggerStub) Trace(message string, args ...interface{}) {
	if l.TraceCalled != nil {
		l.TraceCalled(message, args)
	}
}

// Debug -
func (l LoggerStub) Debug(message string, args ...interface{}) {
	if l.DebugCalled != nil {
		l.DebugCalled(message, args)
	}
}

// Info -
func (l LoggerStub) Info(message string, args ...interface{}) {
	if l.InfoCalled != nil {
		l.InfoCalled(message, args)
	}
}

// Warn -
func (l LoggerStub) Warn(message string, args ...interface{}) {
	if l.WarnCalled != nil {
		l.WarnCalled(message, args)
	}
}

// Error -
func (l LoggerStub) Error(message string, args ...interface{}) {
	if l.ErrorCalled != nil {
		l.ErrorCalled(message, args)
	}
}

// LogIfError -
func (l LoggerStub) LogIfError(err error, args ...interface{}) {
	if l.LogIfErrorCalled != nil {
		l.LogIfErrorCalled(err, args)
	}
}

// IsInterfaceNil -
func (l LoggerStub) IsInterfaceNil() bool {
	return false
}
