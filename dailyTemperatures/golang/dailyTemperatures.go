package dailyTemperatures

func dailyTemperatures(temperatures []int) []int {
	result := make([]int, len(temperatures))
	stack := [][]int{}
	for i := 0; i < len(temperatures)-1; i++ {
		if temperatures[i] < temperatures[i+1] {
			result[i] = 1

			for len(stack) > 0 && stack[len(stack)-1][0] < temperatures[i+1] {
				idx := stack[len(stack)-1][1]
				result[idx] = i + 1 - idx
				stack = stack[0 : len(stack)-1]
			}
		} else {
			stack = append(stack, []int{temperatures[i], i})
		}
	}

	temperatures[len(temperatures)-1] = 0
	for i := range stack {
		idx := stack[len(stack)-1-i][1]
		result[idx] = 0
	}

	return result

}
