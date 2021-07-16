package core

import "fmt"

type consoleLogger struct {
}

// NewConsoleLogger returns a new instance of consoleLogger
func NewConsoleLogger() *consoleLogger {
	return &consoleLogger{}
}

// Trace will print a trace log
func (c consoleLogger) Trace(message string, args ...interface{}) {
	fmt.Printf("[TRACE] %s %v", message, getPrintableArgs(args))
}

// Debug will print a debug log
func (c consoleLogger) Debug(message string, args ...interface{}) {
	fmt.Printf("[DEBUG] %s %v", message, getPrintableArgs(args))
}

// Info will print an info log
func (c consoleLogger) Info(message string, args ...interface{}) {
	fmt.Printf("[INFO] %s %v", message, getPrintableArgs(args))
}

// Warn will print a warn log
func (c consoleLogger) Warn(message string, args ...interface{}) {
	fmt.Printf("[WARN] %s %v", message, getPrintableArgs(args))
}

// Error will print an error log
func (c consoleLogger) Error(message string, args ...interface{}) {
	fmt.Printf("[ERROR] %s %v", message, getPrintableArgs(args))
}

// LogIfError will print an error if it is not nil
func (c consoleLogger) LogIfError(err error, args ...interface{}) {
	if err != nil {
		fmt.Printf("[ERROR] %s %v", err.Error(), getPrintableArgs(args))
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
func (c consoleLogger) IsInterfaceNil() bool {
	return false
}
