package minimumDifferenceBetweenLargestAndSmallestValueInThreeMoves

import (
	"sort"
)

func minDifference(nums []int) int {

	if len(nums) <= 4 {
		return 0
	}

	sort.Ints(nums)

	l := len(nums) - 1
	min := nums[l] - nums[0]

	for i := 0; i < 4; i++ {
		diff := nums[l-3+i] - nums[i]
		if diff < min {
			min = diff
		}
	}

	return min
}
