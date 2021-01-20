package boundaryOfBinaryTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func boundaryOfBinaryTree(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}
	if root.Left != nil || root.Right != nil {
		result = append(result, root.Val)
	}

	left := []int{}
	leaf := []int{}
	right := []int{}
	if root.Left != nil {
		findLeftBoundrydfs(root.Left, &left)
	}
	findLeafBoundrydfs(root, &leaf)
	if root.Right != nil {
		findRightBoundrydfs(root.Right, &right)
	}

	result = append(result, left...)
	result = append(result, leaf...)
	result = append(result, right...)

	return result
}

func findLeftBoundrydfs(node *TreeNode, boundries *[]int) {
	if node.Left != nil || node.Right != nil {
		*boundries = append(*boundries, node.Val)
		if node.Left != nil {
			findLeftBoundrydfs(node.Left, boundries)
		} else {
			findLeftBoundrydfs(node.Right, boundries)
		}
	}

}

func findLeafBoundrydfs(node *TreeNode, boundries *[]int) {
	if node.Left == nil && node.Right == nil {
		*boundries = append(*boundries, node.Val)
	}

	if node.Left != nil {
		findLeafBoundrydfs(node.Left, boundries)
	}
	if node.Right != nil {
		findLeafBoundrydfs(node.Right, boundries)
	}
}

func findRightBoundrydfs(node *TreeNode, boundries *[]int) {
	if node.Left != nil || node.Right != nil {
		*boundries = append([]int{node.Val}, *boundries...)

		if node.Right != nil {
			findRightBoundrydfs(node.Right, boundries)
		} else {
			findRightBoundrydfs(node.Left, boundries)
		}
	}
}
