package leastNumbeOfUniqueIntegersAfterKRemovals

import (
	"sort"
)

func findLeastNumOfUniqueInts(arr []int, k int) int {
	m := make(map[int]int)
	for _, v := range arr {
		m[v]++
	}

	t := []int{}
	for _, v := range m {

		t = append(t, v)
	}

	sort.Ints(t)

	for _, v := range t {
		if k-v < 0 {
			break
		} else {
			k -= v
			t = t[1:]
		}
	}

	return len(t)
}
