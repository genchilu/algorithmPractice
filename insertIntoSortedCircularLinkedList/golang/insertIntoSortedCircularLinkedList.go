package insertIntoSortedCircularLinkedList

type Node struct {
	Val  int
	Next *Node
}

func insert(aNode *Node, x int) *Node {
	if aNode == nil {
		n := &Node{x, nil}
		n.Next = n
		return n
	}

	m := make(map[*Node]bool)
	cur := aNode
	max := aNode

	insert := false
	_, ok := m[cur]
	for !ok {
		m[cur] = true
		if cur.Val > cur.Next.Val {
			max = cur
		}

		if cur.Val <= x && x <= cur.Next.Val {
			insert = true
			nn := &Node{x, nil}
			nn.Next = cur.Next
			cur.Next = nn

			break
		}

		cur = cur.Next
		_, ok = m[cur]
	}

	if !insert {
		nn := &Node{x, nil}
		nn.Next = max.Next
		max.Next = nn
	}

	return aNode
}
