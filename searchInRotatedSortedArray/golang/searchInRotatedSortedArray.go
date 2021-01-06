package searchInRotatedSortedArray

import (
	"sort"
)

func search(nums []int, target int) int {
	pivot := findpivot(nums, 0, len(nums))
	left := nums[0:pivot]
	right := nums[pivot:]
	// fmt.Printf("%d\n", pivot)
	if len(left) > 0 && target >= left[0] {
		// fmt.Printf("11111\n")
		idx := sort.SearchInts(left, target)
		if idx < pivot && left[idx] == target {
			return idx
		}
	} else {
		// fmt.Printf("22222\n")
		idx := sort.SearchInts(right, target)
		if idx < len(right) && right[idx] == target {
			return idx + pivot
		}
	}
	return -1
}

func findpivot(nums []int, begin, end int) int {
	if end-begin == 1 {
		return begin
	}

	t := (begin + end) / 2
	if t < len(nums)-1 && nums[t+1] < nums[t] {
		return t + 1
	}

	if t > 0 && nums[t-1] > nums[t] {
		return t
	}

	if nums[t] < nums[begin] {
		//find left
		return findpivot(nums, begin, t)
	} else {
		if nums[begin] < nums[end-1] {
			return findpivot(nums, begin, t)
		} else {
			//find right
			return findpivot(nums, t, end)
		}
	}
}
