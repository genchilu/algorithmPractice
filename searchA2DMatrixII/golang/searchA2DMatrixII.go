package searchA2DMatrixII

import (
	"sort"
)

func searchMatrix(matrix [][]int, target int) bool {

	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	u, d, l, r := 0, len(matrix), 0, len(matrix[0])

	for l < r {
		// fmt.Printf("======\n")
		// fmt.Printf("ooooo: r %d, l %d, d %d, u %d\n", r, l, d, u)
		// for i := u; i < d; i++ {
		// 	fmt.Printf("%v\n", matrix[i][l:r])
		// }
		rr := sort.SearchInts(matrix[u][l:r], target)
		//fmt.Printf("rrr %d\n", rr)
		if rr != len(matrix[u][l:r]) && matrix[u][l:r][rr] == target {
			return true
		}

		ll := sort.SearchInts(matrix[d-1][l:r], target)
		//fmt.Printf("lll %d\n", ll)
		if ll != len(matrix[d-1][l:r]) && matrix[d-1][l:r][ll] == target {
			return true
		}

		tmp := []int{}
		for i := u; i < d; i++ {
			tmp = append(tmp, matrix[i][l])
		}
		dd := sort.SearchInts(tmp, target)
		//fmt.Printf("ddd %d\n", dd)
		if dd != len(tmp) && tmp[dd] == target {
			return true
		}

		tmp = []int{}
		for i := u; i < d; i++ {
			tmp = append(tmp, matrix[i][r-1])
		}
		uu := sort.SearchInts(tmp, target)
		//fmt.Printf("uuu %d\n", uu)
		if uu != len(tmp) && tmp[uu] == target {
			return true
		}

		r = rr + l
		l = ll + l
		d = dd + u
		u = uu + u
	}

	return false
}
