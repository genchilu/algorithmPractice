package longestIncreasingSubsequence

import "sort"

func lengthOfLIS(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	sub := []int{}

	for i := range nums {
		idx := sort.SearchInts(sub, nums[i])
		if idx == len(sub) {
			sub = append(sub, nums[i])
		} else {
			sub[idx] = nums[i]
		}
	}

	return len(sub)
}
