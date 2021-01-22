package lowestCommonAncestorOfBinaryTree

type TreeNode struct {
	Val   int
	Left  *ListNode
	Right *ListNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	queue := []*TreeNode{}
	queue = append(queue, root)

	m := make(map[int]*TreeNode)
	pd := -1
	qd := -1
	c := 0
	done := false
	for !done {
		newq := []*TreeNode{}
		for _, n := range queue {
			if n.Val == p.Val {
				pd = c
			} else if n.Val == q.Val {
				qd = c
			}

			if pd > -1 && qd > -1 {
				done = true
				break
			}

			if n.Left != nil {
				m[n.Left.Val] = n
				newq = append(newq, n.Left)
			}

			if n.Right != nil {
				m[n.Right.Val] = n
				newq = append(newq, n.Right)
			}
		}

		queue = newq
		c++
	}

	for q.Val != p.Val {
		if pd > qd {
			pd--
			p = m[p.Val]
		} else if qd > pd {
			qd--
			q = m[q.Val]
		} else {
			p = m[p.Val]
			q = m[q.Val]
		}
	}

	return p
}
