package sumOfSubarrayMinimums

import "math"

func sumSubarrayMins(arr []int) int {

	m := int(math.Pow10(9)) + 7
	stack := [][]int{}

	result, sum := 0, 0
	for _, v := range arr {
		w := 1
		for len(stack) > 0 && v < stack[0][0] {
			sum -= stack[0][0] * stack[0][1]
			w += stack[0][1]
			stack = stack[1:]
		}

		stack = append([][]int{{v, w}}, stack...)
		sum += w * v

		result = (result + sum%m) % m

	}

	return result
}
