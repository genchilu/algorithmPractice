package DivideArrayInSetsOfKConsecutiveNumbers

import (
	"sort"
)

func isPossibleDivide(nums []int, k int) bool {
	if len(nums)%k != 0 {
		return false
	}

	m := make(map[int]int)

	for _, v := range nums {
		if _, ok := m[v]; !ok {
			m[v] = 0
		}
		m[v]++
	}

	sort.Ints(nums)

	for _, v := range nums {
		if c, ok1 := m[v]; ok1 && c > 0 {
			for vv := v + 1; vv < v+k; vv++ {
				if cc, ok2 := m[vv]; ok2 {
					if cc == 0 {
						return false
					} else {
						m[vv]--
					}
				} else {
					return false
				}
			}
			m[v]--
		}
	}

	return true
}
