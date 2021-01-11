package convertBinarySearchTreeToSortedDoublyLinkedList

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func treeToDoublyList(root *Node) *Node {
	h, _ := convert(root)
	return h
}

func convert(node *Node) (h, t *Node) {
	if node == nil {
		return node, node
	}

	if node.Left == nil && node.Right == nil {
		node.Left, node.Right = node, node
		return node, node
	}

	lh, lt := convert(node.Left)
	rh, rt := convert(node.Right)

	if lh != nil {
		h = lh
		lt.Right = node
		node.Left = lt
	} else {
		h = node
	}

	if rh != nil {
		t = rt
		rh.Left = node
		node.Right = rh
	} else {
		t = node
	}

	h.Left = t
	t.Right = h

	return h, t
}
