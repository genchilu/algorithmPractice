package nextGreaterElements

func nextGreaterElements(nums []int) []int {
	result := make([]int, len(nums))
	for i, _ := range result {
		result[i] = -1
	}

	for i, v := range nums {
		if i > 0 && nums[i-1] == v {
			result[i] = result[i-1]
		} else {
			for j := 1; j <= len(nums); j++ {
				idx := j + i
				if idx >= len(nums) {
					idx -= len(nums)
				}
				if nums[idx] > v {
					result[i] = nums[idx]
					break
				}
			}
		}
	}
	return result
}
