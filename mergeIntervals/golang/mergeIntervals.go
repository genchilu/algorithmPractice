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

	first := intervals[0][0]
	last :=intervals[0][1]
	ans := [][]int{}
	for idx, _ := range intervals {
		if (idx==len(intervals)-1) {
			ans = append(ans, []int{first, last})
		} else if (last >= intervals[idx+1][0]) {
			if(last < intervals[idx+1][1]) {
				last = intervals[idx+1][1]
			} 
		} else {
			ans = append(ans, []int{first, last})
			first = intervals[idx+1][0]
			last = intervals[idx+1][1]
		}
	}
	return ans
}