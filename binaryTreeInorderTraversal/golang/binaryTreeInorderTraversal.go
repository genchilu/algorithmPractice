package binaryTreeInorderTraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	result := []int{}
	stack := []*TreeNode{}

	cur := root

	for cur != nil {
		if cur.Left != nil {
			stack = append([]*TreeNode{cur}, stack...)
			cur = cur.Left
			stack[0].Left = nil
		} else {
			result = append(result, cur.Val)

			if cur.Right != nil {
				cur = cur.Right
			} else {
				if len(stack) > 0 {
					cur = stack[0]
					stack = stack[1:]
				} else {
					cur = nil
				}
			}
		}

	}

	return result
}
