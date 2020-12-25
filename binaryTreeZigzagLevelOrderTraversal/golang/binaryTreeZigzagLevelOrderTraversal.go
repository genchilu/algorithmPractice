package binaryTreeZigzagLevelOrderTraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	l2rq := []*TreeNode{}
	r2lq := []*TreeNode{}
	l2r := true

	if root != nil {
		l2rq = append(l2rq, root)
	}

	result := [][]int{}
	for len(l2rq) > 0 || len(r2lq) > 0 {
		tmp := []int{}
		if l2r {
			for _, v := range l2rq {
				tmp = append(tmp, v.Val)
				if v.Left != nil {
					r2lq = append([]*TreeNode{v.Left}, r2lq...)
				}
				if v.Right != nil {
					r2lq = append([]*TreeNode{v.Right}, r2lq...)
				}
			}
			l2rq = []*TreeNode{}
			l2r = !l2r
		} else {
			for _, v := range r2lq {
				tmp = append(tmp, v.Val)
				if v.Right != nil {
					l2rq = append([]*TreeNode{v.Right}, l2rq...)
				}
				if v.Left != nil {
					l2rq = append([]*TreeNode{v.Left}, l2rq...)
				}
			}
			r2lq = []*TreeNode{}
			l2r = !l2r
		}
		result = append(result, tmp)
	}

	return result
}
