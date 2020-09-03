package mergeIntervals

import (
    //"fmt"
    "sort"
)

type byFirstEle [][]int

func (s byFirstEle) Len() int {
    return len(s)
}
func (s byFirstEle) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}
func (s byFirstEle) Less(i, j int) bool {
    return s[i][0] < s[j][0]
}

func merge(intervals [][]int) [][]int {
	
	if(intervals == nil || len(intervals) <= 1) {
		return intervals
	}

	sort.Sort(byFirstEle(intervals))

	idx := 0
	for idx<(len(intervals)-1) {
		if (intervals[idx][1] >= intervals[idx+1][0]) {
			if(intervals[idx][1] < intervals[idx+1][1]) {
				intervals[idx][1] = intervals[idx+1][1]
			}
			intervals = append(intervals[0:idx+1], intervals[idx+2:]...)
		} else {
			idx++
		}
	}
	return intervals
}