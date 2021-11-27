package threeSumClosest

func threeSumClosest(nums []int, target int) int {

	return findClosedum(nums, 0, 3, target)

}

func findClosedum(nums []int, i, l, target int) int {

	if l == 1 {
		min := myabs(target - nums[i])
		sum := nums[i]
		for j := i + 1; j < len(nums); j++ {
			tmp := myabs(target - nums[j])
			if tmp < min {
				min = tmp
				sum = nums[j]
			}
		}
		return sum

	} else if len(nums)-i == l {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
		}

		return sum
	} else {
		s1 := nums[i] + findClosedum(nums, i+1, l-1, target-nums[i])
		s2 := findClosedum(nums, i+1, l, target)
		if myabs(target-s1) < myabs(target-s2) {
			return s1
		} else {
			return s2
		}
	}
}

func myabs(a int) int {
	if a > 0 {
		return a
	}

	return -a
}
