package core

import (
	"bytes"
	"runtime"
	"runtime/pprof"
)

// DumpGoRoutinesToLog will print the currently running go routines in the log
func DumpGoRoutinesToLog(goRoutinesNumberStart int, log Logger) {
	if log == nil {
		return
	}

	buf := new(bytes.Buffer)
	err := pprof.Lookup("goroutine").WriteTo(buf, 2)
	if err != nil {
		log.Error("could not dump goroutines", "error", err)
	}
	log.Debug("go routines number",
		"start", goRoutinesNumberStart,
		"end", runtime.NumGoroutine())

	log.Debug(buf.String())
}

// GetRunningGoRoutines gets the currently running go routines stack trace as a bytes.Buffer
func GetRunningGoRoutines(log Logger) *bytes.Buffer {
	buffer := new(bytes.Buffer)
	err := pprof.Lookup("goroutine").WriteTo(buffer, 2)
	if err != nil {
		log.Error("could not dump goroutines", "error", err)
	}
	return buffer
}
