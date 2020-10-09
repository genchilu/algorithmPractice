package pathSumIII

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) int {
	m := make(map[int]int)
	return dfs(root, sum, 0, m)
}

func dfs(node *TreeNode, t, s int, m map[int]int) int {
	c := 0
	if node == nil {
		return c
	}
	s += node.Val
	if _, ok := m[s]; !ok {
		m[s] = 0
	}

	if s == t {
		c++
	}

	if _, ok := m[s-t]; ok {
		c += m[s-t]
	}

	m[s]++
	c += dfs(node.Left, t, s, m) + dfs(node.Right, t, s, m)
	m[s]--

	return c
}
