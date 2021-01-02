package permutations

func permute(nums []int) [][]int {
	if len(nums) <= 1 {
		return [][]int{nums}
	}

	if len(nums) == 2 {
		return [][]int{[]int{nums[0], nums[1]}, []int{nums[1], nums[0]}}
	}

	result := [][]int{}
	for i, v := range nums {
		tmp := []int{}
		for j, vv := range nums {
			if i != j {
				tmp = append([]int{vv}, tmp...)
			}
		}
		sps := permute(tmp)
		for _, sp := range sps {
			sp = append([]int{v}, sp...)
			result = append(result, sp)
		}

	}
	return result
}
