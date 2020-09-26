package MaximumDifferenceBetweenNodeAndAncestor

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxAncestorDiff(root *TreeNode) int {
	max := 0
	ancestors := []int{}

	return dfs(root, &ancestors, max)
}

func dfs(node *TreeNode, ancestors *[]int, max int) int {
	for _, v := range *ancestors {
		tmp := node.Val - v
		if tmp < 0 {
			tmp = -tmp
		}

		if tmp > max {
			max = tmp
		}
	}

	*ancestors = append(*ancestors, node.Val)
	if node.Left != nil {
		max = dfs(node.Left, ancestors, max)
	}

	if node.Right != nil {
		max = dfs(node.Right, ancestors, max)
	}

	*ancestors = (*ancestors)[0 : len(*ancestors)-1]

	return max
}
