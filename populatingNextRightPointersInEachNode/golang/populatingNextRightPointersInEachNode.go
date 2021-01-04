package populatingNextRightPointersInEachNode

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	q := []*Node{}
	q = append(q, root)
	for len(q) > 0 {
		tmp := []*Node{}
		pre := q[0]
		if pre.Left != nil {
			tmp = append(tmp, pre.Left)
			tmp = append(tmp, pre.Right)
		}
		for i := 1; i < len(q); i++ {
			cur := q[i]
			if cur.Left != nil {
				tmp = append(tmp, cur.Left)
				tmp = append(tmp, cur.Right)
			}
			pre.Next = cur
			pre = cur
		}
		pre.Next = nil

		q = tmp
	}
	return root
}
