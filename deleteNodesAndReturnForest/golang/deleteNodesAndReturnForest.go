package deleteNodesAndReturnForest

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	pm := make(map[int]*TreeNode)
	m := make(map[int]*TreeNode)
	dfs(root, pm, m)

	fm := make(map[int]*TreeNode)
	fm[root.Val] = root

	for _, d := range to_delete {
		if node, ok := m[d]; ok {
			if parent, ok2 := pm[d]; ok2 {
				if parent.Left == node {
					parent.Left = nil
				} else {
					parent.Right = nil
				}
				delete(pm, d)
			}

			if node.Left != nil {
				fm[node.Left.Val] = node.Left
			}
			if node.Right != nil {
				fm[node.Right.Val] = node.Right
			}

			delete(m, d)
			delete(fm, d)
		}
	}

	forest := []*TreeNode{}
	for _, t := range fm {
		forest = append(forest, t)
	}

	return forest
}

func dfs(node *TreeNode, pm map[int]*TreeNode, m map[int]*TreeNode) {
	m[node.Val] = node
	if node.Left != nil {
		pm[node.Left.Val] = node
		dfs(node.Left, pm, m)
	}
	if node.Right != nil {
		pm[node.Right.Val] = node
		dfs(node.Right, pm, m)
	}
}
