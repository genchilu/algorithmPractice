package findDuplicateSubtrees

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	result := []*TreeNode{}
	m := make(map[int][]*TreeNode)
	dfs(root, m)
	for _, v := range m {
		for i, _ := range v {
			c := 0
			for j := i + 1; j < len(v); j++ {
				if isEqualTree(v[i], v[j]) {
					c++
					v[j] = nil
				}
			}
			if c > 0 && v[i] != nil {
				result = append(result, v[i])
			}
		}
	}

	return result
}

func dfs(n *TreeNode, m map[int][]*TreeNode) {
	if _, ok := m[n.Val]; ok {
		m[n.Val] = append(m[n.Val], n)
	} else {
		m[n.Val] = []*TreeNode{n}
	}

	if n.Left != nil {
		dfs(n.Left, m)
	}

	if n.Right != nil {
		dfs(n.Right, m)
	}
}
func isEqualTree(t1, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil {
		return false
	}

	if t1.Val != t2.Val {
		return false
	}

	return isEqualTree(t1.Left, t2.Left) && isEqualTree(t1.Right, t2.Right)
}
