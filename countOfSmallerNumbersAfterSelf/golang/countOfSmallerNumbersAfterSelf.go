package countOfSmallerNumbersAfterSelf

import "sort"

func countSmaller(nums []int) []int {

	sortNums := []int{}
	result := make([]int, len(nums))
	l := len(nums) - 1

	for i, _ := range nums {
		idx := sort.SearchInts(sortNums, nums[l-i])
		sortNums = insert(sortNums, idx, nums[l-i])
		result[l-i] = idx
	}
	return result
}

func insert(nums []int, i, v int) []int {
	nums = append(nums, 0)
	copy(nums[i+1:], nums[i:])
	nums[i] = v
	return nums
}
