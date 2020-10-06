package allNodesDistanceKInBinaryTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func distanceK(root *TreeNode, target *TreeNode, K int) []int {
	p := make(map[int]*TreeNode)
	dfs1(root, p)

	result := []int{}
	dfs2(target, nil, 0, K, &result, p)

	return result
}

func dfs1(n *TreeNode, p map[int]*TreeNode) {

	if n.Left != nil {
		p[n.Left.Val] = n
		dfs1(n.Left, p)
	}

	if n.Right != nil {
		p[n.Right.Val] = n
		dfs1(n.Right, p)
	}

}

func dfs2(n, s *TreeNode, d, k int, result *[]int, p map[int]*TreeNode) {
	if d == k {
		(*result) = append(*result, n.Val)
	}

	if n.Left != nil && n.Left != s && d+1 <= k {
		dfs2(n.Left, n, d+1, k, result, p)
	}

	if n.Right != nil && n.Right != s && d+1 <= k {
		dfs2(n.Right, n, d+1, k, result, p)
	}

	if pp, ok := p[n.Val]; ok {
		if pp != s && d+1 <= k {
			dfs2(pp, n, d+1, k, result, p)
		}
	}
}
