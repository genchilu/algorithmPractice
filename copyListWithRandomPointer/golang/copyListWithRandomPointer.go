package copyListWithRandomPointer

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	m := make(map[*Node]*Node)

	cur := head
	for cur != nil {
		n := Node{Val: cur.Val}
		m[cur] = &n
		cur = cur.Next
	}

	cur = head
	for cur != nil {
		cloneCur := m[cur]
		if cur.Next != nil {
			cloneCur.Next = m[cur.Next]
		}
		if cur.Random != nil {
			cloneCur.Random = m[cur.Random]
		}

		cur = cur.Next
	}

	return m[head]
}
