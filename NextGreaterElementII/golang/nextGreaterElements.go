package nextGreaterElements

type INum struct {
	idx int
	val int
}

func nextGreaterElements(nums []int) []int {

	result := make([]int, len(nums))
	for i := range result {
		result[i] = -1
	}
	stack := []INum{}
	for i := 0; i < 2; i++ {
		for j := len(nums) - 1; j >= 0; j-- {

			for len(stack) > 0 {
				l := len(stack) - 1
				if stack[l].val > nums[j] {
					if result[j] == -1 {
						result[j] = stack[l].val
					}
					break
				} else {
					stack = stack[0:l]
				}
			}

			stack = append(stack, INum{idx: j, val: nums[j]})
		}
	}
	return result
}
