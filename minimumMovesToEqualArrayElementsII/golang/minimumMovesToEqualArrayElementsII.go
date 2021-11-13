package minimumMovesToEqualArrayElementsII

import "sort"

func minMoves2(nums []int) int {

	sort.Ints(nums)

	m := nums[len(nums)/2]

	sum := 0
	for _, v := range nums {
		sum += abs(v - m)
	}
	return sum
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}
