package russianDollEnvelopes

import (
	"sort"
)

func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(x, y int) bool {
		if envelopes[x][0] < envelopes[y][0] {
			return true
		} else if envelopes[x][0] == envelopes[y][0] {
			return envelopes[x][1] > envelopes[y][1]
		} else {
			return false
		}
	})

	result := []int{}
	for _, e := range envelopes {
		idx := sort.SearchInts(result, e[1])
		if idx == len(result) {
			result = append(result, e[1])
		} else {
			result[idx] = e[1]
		}
	}

	return len(result)
}
