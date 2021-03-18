package insertInterval

import (
	"sort"
)

func insert(intervals [][]int, newInterval []int) [][]int {
	start := make([]int, len(intervals))
	end := make([]int, len(intervals))
	result := [][]int{}

	if len(intervals) == 0 {
		result = append(result, newInterval)
		return result
	}
	for i, v := range intervals {
		start[i] = v[0]
		end[i] = v[1]
	}

	idx1 := sort.SearchInts(start, newInterval[0])

	idx2 := sort.SearchInts(end, newInterval[1])

	if idx1 < len(intervals) && intervals[idx1][0] == newInterval[0] {
		result = append(result, intervals[0:idx1]...)
	} else if idx1 != 0 {
		if newInterval[0] <= intervals[idx1-1][1] {
			result = append(result, intervals[0:idx1-1]...)
			newInterval[0] = intervals[idx1-1][0]
		} else {
			result = append(result, intervals[0:idx1]...)
		}
	}

	if idx2 == len(intervals) {
		result = append(result, newInterval)
	} else if newInterval[1] < intervals[idx2][0] {
		result = append(result, newInterval)
		result = append(result, intervals[idx2:]...)
	} else {
		newInterval[1] = intervals[idx2][1]
		result = append(result, newInterval)
		result = append(result, intervals[idx2+1:]...)
	}

	return result
}
