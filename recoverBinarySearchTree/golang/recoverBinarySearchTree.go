package recoverBinarySearchTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recoverTree(root *TreeNode) {

	var from, to, pre *TreeNode
	from, to, _ = inOrder(root, from, to, pre)

	swap(from, to)
	if root == to {
		root = from
	}
}

func inOrder(cur, from, to, pre *TreeNode) (*TreeNode, *TreeNode, *TreeNode) {
	if cur.Left != nil {
		from, to, pre = inOrder(cur.Left, from, to, pre)
	}

	if pre != nil {
		if pre.Val > cur.Val {
			if from == nil {
				from = pre
			}
			to = cur
		}
	}

	pre = cur

	if cur.Right != nil {
		from, to, pre = inOrder(cur.Right, from, to, pre)
	}

	return from, to, pre
}

func swap(t1, t2 *TreeNode) {
	tmp := t1.Val
	t1.Val = t2.Val
	t2.Val = tmp

}
