package lowestCommonAncestorOfABinaryTreeIII

type Node struct {
	Val    int
	Left   *Node
	Right  *Node
	Parent *Node
}

func lowestCommonAncestor(p *Node, q *Node) *Node {
	mp := make(map[int]bool)
	mq := make(map[int]bool)

	for {

		mp[p.Val] = true
		mq[q.Val] = true

		if p.Parent != nil {
			p = p.Parent
		}

		if q.Parent != nil {
			q = q.Parent
		}

		if p.Val == q.Val {
			return p
		} else if _, ok := mq[p.Val]; ok {
			return p
		} else if _, ok := mp[q.Val]; ok {
			return q
		}
	}

}
