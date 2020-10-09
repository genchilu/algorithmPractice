package pathSumIII

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	return dfs(root, sum, []int{})
}

func dfs(node *TreeNode, t int, ps []int) int {
	cps := make([]int, len(ps))
	copy(cps, ps)

	c := 0
	for i, _ := range cps {
		cps[i] += node.Val
		if cps[i] == t {
			c++
		}
	}

	cps = append(cps, node.Val)
	if node.Val == t {
		c++
	}

	if node.Left != nil {
		c += dfs(node.Left, t, cps)
	}

	if node.Right != nil {
		c += dfs(node.Right, t, cps)
	}

	return c
}
