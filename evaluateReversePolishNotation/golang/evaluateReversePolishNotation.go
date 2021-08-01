package evaluateReversePolishNotation

import "strconv"

func evalRPN(tokens []string) int {

	stack := []int{}
	for _, token := range tokens {

		switch token {
		case "+":
			l := len(stack) - 1
			stack[l-1] = stack[l-1] + stack[l]
			stack = stack[0:l]
		case "-":
			l := len(stack) - 1
			stack[l-1] = stack[l-1] - stack[l]
			stack = stack[0:l]
		case "*":
			l := len(stack) - 1
			stack[l-1] = stack[l-1] * stack[l]
			stack = stack[0:l]
		case "/":
			l := len(stack) - 1
			stack[l-1] = stack[l-1] / stack[l]
			stack = stack[0:l]
		default:
			d, _ := strconv.Atoi(token)
			stack = append(stack, d)
		}
	}
	return stack[0]
}
