package tree

import "fmt"

type interval struct {
	low  uint64
	high uint64
}

func newInterval(low, high uint64) *interval {
	if low <= high {
		return &interval{
			low:  low,
			high: high,
		}
	}
	return &interval{
		low:  high,
		high: low,
	}
}

func (i *interval) contains(value uint64) bool {
	return i.low <= value && i.high >= value
}

func (i *interval) toString() string {
	return fmt.Sprintf("[%d,%d]", i.low, i.high)
}
