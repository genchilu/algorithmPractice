package minimumRemoveToMakeValidParentheses

func minRemoveToMakeValid(s string) string {
	if len(s) < 1 {
		return ""
	}

	l := []int{}
	stack := []byte{}
	c := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 40 {
			c++
			stack = append(stack, s[i])
			l = append(l, len(stack)-1)
		} else if s[i] == 41 {
			if c != 0 {
				c--
				stack = append(stack, s[i])
				l = l[0 : len(l)-1]
			}
		} else {
			stack = append(stack, s[i])
		}
	}

	for i := len(l) - 1; i >= 0; i-- {
		if l[i] == len(stack)-1 {
			stack = stack[0:l[i]]
		} else {
			stack = append(stack[0:l[i]], stack[l[i]+1:]...)
		}
	}
	return string(stack)
}
