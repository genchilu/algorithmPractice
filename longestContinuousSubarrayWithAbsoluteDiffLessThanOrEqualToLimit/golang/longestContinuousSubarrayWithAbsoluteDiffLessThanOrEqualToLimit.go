package longestContinuousSubarrayWithAbsoluteDiffLessThanOrEqualToLimit

func longestSubarray(nums []int, limit int) int {

	maxl := 0
	for i, _ := range nums {

		if len(nums)-i > maxl {
			max, min := nums[i], nums[i]
			for j := i; j < len(nums); j++ {
				if nums[j] > max {
					max = nums[j]
				} else if nums[j] < min {
					min = nums[j]
				}
				if (max - min) > limit {
					if j-i > maxl {
						maxl = j - i
					}
					break
				}

				if j == len(nums)-1 {
					maxl = j - i + 1
				}
			}
		}
	}
	return maxl
}
