package arrayOfDoubledPairs

import "sort"

func canReorderDoubled(arr []int) bool {
	m := make(map[int]int)

	sort.Ints(arr)
	for _, v := range arr {
		m[v] = m[v] + 1
	}

	for _, v := range arr {
		if m[v] > 0 && m[2*v] > 0 {
			m[v] = m[v] - 1
			m[2*v] = m[2*v] - 1
		}
	}

	for _, v := range m {
		if v > 0 {
			return false
		}
	}
	return true
}
