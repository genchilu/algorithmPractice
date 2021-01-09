package validateBinarySearchTree

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	min := math.MinInt64
	max := math.MaxInt64

	return dfs(root, min, max)
}

func dfs(node *TreeNode, min, max int) bool {
	if node == nil {
		return true
	}

	if node.Left != nil && (node.Left.Val >= node.Val || node.Left.Val <= min) {
		return false
	}

	if node.Right != nil && (node.Right.Val <= node.Val || node.Right.Val >= max) {
		return false
	}

	return dfs(node.Left, min, node.Val) && dfs(node.Right, node.Val, max)
}
