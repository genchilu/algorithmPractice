package maximumAverageSubtree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maximumAverageSubtree(root *TreeNode) float64 {
	_, _, max := dfs(root)

	return max
}

func dfs(n *TreeNode) (sum, count int, max float64) {
	sum = 0
	count = 0
	max = 0
	if n != nil {
		ls, lc, lmax := dfs(n.Left)
		rs, rc, rmax := dfs(n.Right)

		sum = ls + rs + n.Val
		count = lc + rc + 1

		avg := float64(sum) / float64(count)

		if max < lmax {
			max = lmax
		}

		if max < rmax {
			max = rmax
		}

		if max < avg {
			max = avg
		}
	}

	return sum, count, max
}
