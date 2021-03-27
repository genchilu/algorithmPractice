package findClosestElements

import "sort"

func findClosestElements(arr []int, k int, x int) []int {

	if k == len(arr) {
		return arr
	}

	idx := sort.SearchInts(arr, x)
	l, r := 0, 0
	if idx == 0 {
		l = idx
		r = idx + 1
	} else {
		l = idx - 1
		r = idx
	}

	result := []int{}
	for k > 0 {
		if l < 0 {
			result = append(result, arr[r])
			r++
		} else if r >= len(arr) {
			result = append([]int{arr[l]}, result...)
			l--
		} else if x-arr[l] <= arr[r]-x {
			result = append([]int{arr[l]}, result...)
			l--
		} else {
			result = append(result, arr[r])
			r++
		}
		k--
	}
	return result
}
