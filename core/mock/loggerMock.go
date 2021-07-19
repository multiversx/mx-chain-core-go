package mock

import "fmt"

// LoggerMock -
type LoggerMock struct {
}

// Trace will print a trace log
func (c LoggerMock) Trace(message string, args ...interface{}) {
	fmt.Printf("[TRACE] %s %v\n", message, getPrintableArgs(args))
}

// Debug will print a debug log
func (c LoggerMock) Debug(message string, args ...interface{}) {
	fmt.Printf("[DEBUG] %s %v\n", message, getPrintableArgs(args))
}

// Info will print an info log
func (c LoggerMock) Info(message string, args ...interface{}) {
	fmt.Printf("[INFO] %s %v\n", message, getPrintableArgs(args))
}

// Warn will print a warn log
func (c LoggerMock) Warn(message string, args ...interface{}) {
	fmt.Printf("[WARN] %s %v\n", message, getPrintableArgs(args))
}

// Error will print an error log
func (c LoggerMock) Error(message string, args ...interface{}) {
	fmt.Printf("[ERROR] %s %v\n", message, getPrintableArgs(args))
}

// LogIfError will print an error if it is not nil
func (c LoggerMock) LogIfError(err error, args ...interface{}) {
	if err != nil {
		fmt.Printf("[ERROR] %s %v\n", err.Error(), getPrintableArgs(args))
	}
}

func getPrintableArgs(args ...interface{}) string {
	if len(args)%2 != 0 {
		return fmt.Sprintf("%v", args)
	}

	printableArgs := ""

	for idx, arg := range args {
		if idx%2 == 0 {
			printableArgs += fmt.Sprintf("%s = ", arg)
			continue
		}
		printableArgs += fmt.Sprintf("%v ", arg)
	}

	return printableArgs
}

// IsInterfaceNil returns false as the struct doesn't use pointer receivers
func (c LoggerMock) IsInterfaceNil() bool {
	return false
}
