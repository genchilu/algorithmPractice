package addTwoNumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	stack1, stack2 := []*ListNode{}, []*ListNode{}

	cur := l1
	for cur != nil {
		stack1 = append(stack1, cur)
		cur = cur.Next
	}

	cur = l2
	for cur != nil {
		stack2 = append(stack2, cur)
		cur = cur.Next
	}

	c, v := 0, 0
	var result *ListNode
	result = nil
	for len(stack1) > 0 && len(stack2) > 0 {
		n1 := stack1[len(stack1)-1]
		stack1 = stack1[0 : len(stack1)-1]
		n2 := stack2[len(stack2)-1]
		stack2 = stack2[0 : len(stack2)-1]
		c, v = add(n1.Val+c, n2.Val)
		n := &ListNode{v, result}
		result = n
	}

	for len(stack1) != 0 {
		n1 := stack1[len(stack1)-1]
		stack1 = stack1[0 : len(stack1)-1]

		c, v = add(n1.Val, c)
		n := &ListNode{v, result}
		result = n
	}

	for len(stack2) != 0 {
		n2 := stack2[len(stack2)-1]
		stack2 = stack2[0 : len(stack2)-1]
		c, v = add(n2.Val, c)
		n := &ListNode{v, result}
		result = n
	}

	if c > 0 {
		n := &ListNode{c, result}
		result = n
	}

	return result
}

func add(i, j int) (c, v int) {
	v = i + j
	if v >= 10 {
		c = 1
		v = v - 10
	}

	return c, v
}
