package flattenBinaryTreeToLinkedList

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	flat(root, nil)

}
func flat(cur, next *TreeNode) {
	if cur == nil {
		return
	}

	//fmt.Printf("000 %d\n", cur.Val)
	if cur.Left == nil && cur.Right == nil {
		cur.Right = next
		return
	}

	if cur.Right == nil {
		//fmt.Printf("111 %d\n", cur.Val)
		cur.Right = cur.Left
		cur.Left = nil
	}

	flat(cur.Left, cur.Right)
	flat(cur.Right, next)

	if cur.Left != nil {
		cur.Right = cur.Left
		cur.Left = nil
	}
}
