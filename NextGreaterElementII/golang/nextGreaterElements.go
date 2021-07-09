package nextGreaterElements

func nextGreaterElements(nums []int) []int {

	result := make([]int, len(nums))
	for i := range result {
		result[i] = -1
	}
	stack := []int{}
	for i := 0; i < 2; i++ {
		for j := len(nums) - 1; j >= 0; j-- {

			for len(stack) > 0 {
				l := len(stack) - 1
				if nums[stack[l]] > nums[j] {
					if result[j] == -1 {
						result[j] = nums[stack[l]]
					}
					break
				} else {
					stack = stack[0:l]
				}
			}

			stack = append(stack, j)
		}
	}
	return result
}
