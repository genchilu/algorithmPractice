package addTwoNumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	v, c := add(l1.Val, l2.Val)
	l1 = l1.Next
	l2 = l2.Next

	n := &ListNode{v, nil}
	cur := n
	for l1 != nil && l2 != nil {
		v, c = add(l1.Val+c, l2.Val)
		nn := &ListNode{v, nil}
		cur.Next = nn
		cur = nn
		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil {
		v, c = add(l1.Val, c)
		nn := &ListNode{v, nil}
		cur.Next = nn
		cur = nn
		l1 = l1.Next
	}

	for l2 != nil {
		v, c = add(l2.Val, c)
		nn := &ListNode{v, nil}
		cur.Next = nn
		cur = nn
		l2 = l2.Next
	}

	if c != 0 {
		nn := &ListNode{c, nil}
		cur.Next = nn
		cur = nn
	}

	return n
}

func add(i, j int) (v, c int) {
	v = i + j
	c = 0
	if v > 9 {
		v = v - 10
		c = 1
	}

	return v, c
}

// func showList(l *ListNode) {
// 	for l != nil {
// 		fmt.Printf("%d ", l.Val)
// 		l = l.Next
// 	}
// 	fmt.Printf("\n")
// }
