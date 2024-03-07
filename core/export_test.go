package core

import "time"

// GetContainingDuration -
func (sw *StopWatch) GetContainingDuration() (map[string]time.Duration, []string) {
	return sw.getContainingDuration()
}

// GetIdentifiers -
func (sw *StopWatch) GetIdentifiers() []string {
	return sw.identifiers
}

// SetIdentifiers -
func (sw *StopWatch) SetIdentifiers(identifiers []string) {
	sw.identifiers = identifiers
}

// GetStarted -
func (sw *StopWatch) GetStarted(identifier string) (time.Time, bool) {
	s, has := sw.started[identifier]
	return s, has
}

// GetElapsed -
func (sw *StopWatch) GetElapsed(identifier string) (time.Duration, bool) {
	e, has := sw.elapsed[identifier]
	return e, has
}

// SetElapsed -
func (sw *StopWatch) SetElapsed(identifier string, duration time.Duration) {
	sw.elapsed[identifier] = duration
}

// SplitExponentFraction -
func SplitExponentFraction(val string) (string, string) {
	return splitExponentFraction(val)
}

// TestAutoBalanceDataTriesFlag -
const TestAutoBalanceDataTriesFlag = autoBalanceDataTriesFlag
