package distributeCoinsInBinaryTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func distributeCoins(root *TreeNode) int {

	_, ans := dfs(root, 0)
	return ans
}

func dfs(n *TreeNode, a int) (val, ans int) {
	if n == nil {
		return 0, 0
	}

	lv, ans := dfs(n.Left, ans)
	rv, ans := dfs(n.Right, ans)

	ans += ans
	v := lv + rv + n.Val - 1

	return v, ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}
