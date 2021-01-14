package nextPermutation

func nextPermutation(nums []int) {
	if len(nums) > 1 {
		firstDecreasing := len(nums) - 1
		for i := len(nums) - 1; i > 0; i-- {
			if nums[i] > nums[i-1] {
				firstDecreasing = i - 1
				break
			}
		}

		if firstDecreasing == (len(nums) - 1) {
			reverse(nums)
		} else {
			for i := len(nums) - 1; i > 0; i-- {
				if nums[i] > nums[firstDecreasing] {
					nums[i], nums[firstDecreasing] = nums[firstDecreasing], nums[i]
					reverse(nums[firstDecreasing+1:])
					break
				}
			}
		}
	}
}

func reverse(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
