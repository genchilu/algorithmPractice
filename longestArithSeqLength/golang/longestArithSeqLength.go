package longestArithSeqLength

import (
	"sort"
)

func longestArithSeqLength(A []int) int {

	if len(A) < 2 {
		return len(A)
	}

	max := 0
	m := make(map[int][]int)

	for i, v := range A {
		if _, ok := m[v]; !ok {
			m[v] = []int{}
		}
		m[v] = append(m[v], i)
	}

	m2 := make(map[int]map[int]bool)
	for i, _ := range A {
		for j := i + 1; j < len(A); j++ {
			d := A[j] - A[i]
			if _, ok := m2[d]; !ok {
				m2[d] = make(map[int]bool)
			}
			if _, ok := m2[d][j]; !ok {
				p := j
				v := A[j]
				// fmt.Printf("[%d, %d]", i, A[i])
				// fmt.Printf("[%d, %d]", j, A[j])
				c := 2
				for m[v+d] != nil {
					ii := sort.SearchInts(m[v+d], p)
					if ii == len(m[v+d]) {
						break
					}
					newp := m[v+d][ii]
					if newp == p {
						if ii == len(m[v+d])-1 {
							break
						} else {
							ii++
							newp = m[v+d][ii]
						}
					}
					v += d
					p = newp
					//fmt.Printf("[%d, %d]", p, v)
					c++
				}
				//fmt.Printf("\n")
				if c > max {
					max = c
				}
			}

			m2[d][j] = true
		}
	}
	return max
}
