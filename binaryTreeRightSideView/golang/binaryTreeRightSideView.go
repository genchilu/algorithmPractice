package binaryTreeRightSideView


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

 
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	result := []int{}
	bfs(root, 0, &result)

    return result
}

func bfs(node *TreeNode, h int, result *[]int) {
	if (node != nil) {
		if (len(*result)  == h) {
			*result = append(*result, node.Val)
		}
		bfs(node.Right, h+1, result)
		bfs(node.Left, h+1, result)
	}
}