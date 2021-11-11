package ticker

import "time"

// Reset -
func (ct *cleanTicker) Reset(duration time.Duration) {
	ct.ticker.Reset(duration)
}
