package core

import "testing"

func TestConsoleLogger(t *testing.T) {
	cl := consoleLogger{}

	cl.Info("test message", "key0", "value0")
}
