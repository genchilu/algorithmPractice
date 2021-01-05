package searchMatrix

import "sort"

func searchMatrix(matrix [][]int, target int) bool {
	tmp := []int{}
	for _, v := range matrix {
		tmp = append(tmp, v...)
	}

	idx := sort.SearchInts(tmp, target)

	if idx < len(tmp) && tmp[idx] == target {
		return true
	}
	return false
}
