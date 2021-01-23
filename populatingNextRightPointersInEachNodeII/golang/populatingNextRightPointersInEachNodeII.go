package populatingNextRightPointersInEachNodeII

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	q := []*Node{}

	if root != nil {
		q = append(q, root)
	}

	for len(q) > 0 {
		newq := []*Node{}
		for i, n := range q {
			if i+1 < len(q) {
				n.Next = q[i+1]
			}

			if n.Left != nil {
				newq = append(newq, n.Left)
			}

			if n.Right != nil {
				newq = append(newq, n.Right)
			}
		}

		q = newq
	}

	return root
}
