package tree

import "sort"

func sortAndOptimizeIntervals(intervals []BlocksExceptionInterval) []BlocksExceptionInterval {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Low <= intervals[j].Low
	})

	optimizedIntervals := make([]BlocksExceptionInterval, 0, len(intervals))
	optimizeIntervalsSlice(intervals, &optimizedIntervals)

	return optimizedIntervals
}

func optimizeIntervalsSlice(intervals []BlocksExceptionInterval, finalIntervals *[]BlocksExceptionInterval) {
	midIdx := len(intervals) / 2
	*finalIntervals = append(*finalIntervals, intervals[midIdx])
	nextLeftInterval := intervals[:midIdx]
	if len(nextLeftInterval) > 0 {
		optimizeIntervalsSlice(nextLeftInterval, finalIntervals)
	}
	nextRightInterval := intervals[midIdx+1:]
	if len(nextRightInterval) > 0 {
		optimizeIntervalsSlice(nextRightInterval, finalIntervals)
	}
}
