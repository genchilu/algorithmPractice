package minDeletions

import (
	"sort"
)

func minDeletions(s string) int {
	ss := make([]int, 26)
	for i := 0; i < len(s); i++ {
		idx := s[i] - 97

		ss[idx]++
	}

	if len(ss) <= 1 {
		return 0
	}

	c := 0
	sort.Ints(ss)

	for i := len(ss) - 2; i >= 0; i-- {
		if ss[i] != 0 && ss[i] >= ss[i+1] {
			n := ss[i+1] - 1
			if ss[i+1] == 0 {
				n = 0
			}
			c += ss[i] - n
			ss[i] = n
		}
	}
	return c
}
