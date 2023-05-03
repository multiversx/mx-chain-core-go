package mock

import "fmt"

// LoggerStub -
type LoggerStub struct {
	WarnCalled func(message string, args ...interface{})
	InfoCalled func(message string, args ...interface{})
}

// Trace will print a trace log
func (c LoggerStub) Trace(message string, args ...interface{}) {
	fmt.Printf("[TRACE] %s %v\n", message, getPrintableArgs(args))
}

// Debug will print a debug log
func (c LoggerStub) Debug(message string, args ...interface{}) {
	fmt.Printf("[DEBUG] %s %v\n", message, getPrintableArgs(args))
}

// Info will print an info log
func (c LoggerStub) Info(message string, args ...interface{}) {
	if c.InfoCalled != nil {
		c.InfoCalled(message, args)
	}

	fmt.Printf("[INFO] %s %v\n", message, getPrintableArgs(args))
}

// Warn will print a warn log
func (c LoggerStub) Warn(message string, args ...interface{}) {
	if c.WarnCalled != nil {
		c.WarnCalled(message, args)
		return
	}

	fmt.Printf("[WARN] %s %v\n", message, getPrintableArgs(args))
}

// Error will print an error log
func (c LoggerStub) Error(message string, args ...interface{}) {
	fmt.Printf("[ERROR] %s %v\n", message, getPrintableArgs(args))
}

// LogIfError will print an error if it is not nil
func (c LoggerStub) LogIfError(err error, args ...interface{}) {
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
func (c LoggerStub) IsInterfaceNil() bool {
	return false
}
