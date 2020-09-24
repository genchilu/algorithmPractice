package intervalListIntersections

func intervalIntersection(A [][]int, B [][]int) [][]int {

	if len(A) == 0 || len(B) == 0 {
		return [][]int{}
	}

	as := 0
	ae := 0
	bs := 0
	be := 0
	result := [][]int{}
	idxa := 0
	idxb := 0
	for true {
		as = A[idxa][0]
		ae = A[idxa][1]

		bs = B[idxb][0]
		be = B[idxb][1]

		if ae >= bs && as <= be {
			s := as
			if bs > as {
				s = bs
			}

			e := ae
			if be < ae {
				e = be
			}
			result = append(result, []int{s, e})
		}
		if ae < be {
			idxa++
		} else {
			idxb++
		}

		if idxa >= len(A) || idxb >= len(B) {
			break
		}
	}

	return result
}
