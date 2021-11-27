package threeSumClosest

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {

	sort.Ints(nums)

	min, result := math.MaxInt64, 0
	i, j, k := 0, 1, len(nums)-1
	for i < len(nums) {

		for j < k {
			sum := (nums[i] + nums[j] + nums[k])
			tmp := myabs(target - sum)
			if tmp < min {
				result = sum
				min = tmp
			}

			if sum > target {
				k--
			} else if sum < target {
				j++
			} else {
				return sum
			}
		}
		i++
		j = i + 1
		k = len(nums) - 1
	}

	return result

}

func myabs(a int) int {
	if a > 0 {
		return a
	}

	return -a
}
