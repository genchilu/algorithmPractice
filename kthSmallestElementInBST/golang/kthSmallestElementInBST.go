package kthSmallestElementInBST

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {

	count := 0
	return find(root, k, &count)
}

func find(node *TreeNode, k int, c *int) int {
	if node.Left != nil {
		r := find(node.Left, k, c)
		if r != -1 {
			return r
		}
	}

	*c++
	if *c == k {
		return node.Val
	}

	if node.Right != nil {
		r := find(node.Right, k, c)
		if r != -1 {
			return r
		}
	}

	return -1
}
