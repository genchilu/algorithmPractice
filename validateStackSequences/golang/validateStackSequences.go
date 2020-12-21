package validateStackSequences

func validateStackSequences(pushed []int, popped []int) bool {
	stack := []int{}

	c := 0
	i := 0
	for i < len(pushed) {
		stack = append(stack, pushed[i])
		for len(stack) != 0 && c < len(popped) {
			if stack[len(stack)-1] == popped[c] {
				stack = stack[0 : len(stack)-1]
				c++
			} else {
				break
			}
		}
		i++
	}
	return c == len(popped)
}
