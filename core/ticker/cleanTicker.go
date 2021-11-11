package ticker

import "time"

// cleanTicker wraps around the time.Ticker but with a slight modification: instead of using Reset there is
// a function called ResetAndClean that will reset the ticker and immediately will try to empty the ticker's channel
// Otherwise, it can't be used as a straightforward replacement to the time.After because a new tick might already be available
// on the channel. This might raise issues whenever using the ticker in conjunction with a context instance.
type cleanTicker struct {
	ticker *time.Ticker
}

// NewCleanTicker creates a new instance of a clean ticker
func NewCleanTicker(duration time.Duration) *cleanTicker {
	return &cleanTicker{
		ticker: time.NewTicker(duration),
	}
}

// Chan will return the channel on which periodic ticks will be output
func (ct *cleanTicker) Chan() <-chan time.Time {
	return ct.ticker.C
}

// ResetAndClean resets the ticker and tries to empty any existing values in the channel
func (ct *cleanTicker) ResetAndClean(duration time.Duration) {
	ct.ticker.Reset(duration)

	select {
	case <-ct.ticker.C:
	default:
	}
}

// Stop will stop the ticker
func (ct *cleanTicker) Stop() {
	ct.ticker.Stop()
}

// IsInterfaceNil returns true if there is no value under the interface
func (ct *cleanTicker) IsInterfaceNil() bool {
	return ct == nil
}
