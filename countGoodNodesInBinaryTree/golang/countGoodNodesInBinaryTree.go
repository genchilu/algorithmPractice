package countGoodNodesInBinaryTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func goodNodes(root *TreeNode) int {

	c := 1
	c += dfs(root.Left, root.Val)
	c += dfs(root.Right, root.Val)
	return c
}

func dfs(n *TreeNode, max int) int {
	if n == nil {
		return 0
	}

	c := 0
	if n.Val >= max {
		c++
		max = n.Val
	}

	c += dfs(n.Left, max)
	c += dfs(n.Right, max)

	return c
}
