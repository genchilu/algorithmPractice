package FlipEquivalentBinaryTrees

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {

	return bfs(root1, root2)
}

func bfs(node1 *TreeNode, node2 *TreeNode) bool {
	if node1 == nil && node2 == nil {
		return true
	}

	if node1 == nil {
		return false
	}

	if node2 == nil {
		return false
	}

	if node1.Val != node2.Val {
		return false
	}

	return (bfs(node1.Left, node2.Right) || bfs(node1.Left, node2.Left)) &&
		(bfs(node1.Right, node2.Right) || bfs(node1.Right, node2.Left))
}
