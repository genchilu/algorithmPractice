package numMatchingSubseq

import (
	"sort"
)

func numMatchingSubseq(S string, words []string) int {
	c := 0
	m := make(map[byte][]int)
	for i := 0; i < len(S); i++ {
		if _, ok := m[S[i]]; !ok {
			m[S[i]] = []int{}
		}
		m[S[i]] = append(m[S[i]], i)
	}
	for _, s := range words {
		if len(s) < len(S) && isSub(s, m) {
			c++
		}
	}
	return c
}

func isSub(sub string, m map[byte][]int) bool {

	p := -1
	for i := 0; i < len(sub); i++ {
		if ps, ok := m[sub[i]]; ok {
			j := sort.SearchInts(ps, p)
			if j == len(ps) {
				return false
			}
			newp := ps[j]

			if newp == p {
				if j < len(ps)-1 {
					newp = ps[j+1]
				} else {
					return false
				}
			}

			p = newp
		} else {
			return false
		}
	}

	return true
}
