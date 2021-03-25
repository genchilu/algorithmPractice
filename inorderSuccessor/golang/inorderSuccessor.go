package inorderSuccessor

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	if p == nil {
		return nil
	}

	if root == nil {
		return nil
	}

	_, n := inorder(root, p, -1000000)
	return n

}

func inorder(t, p *TreeNode, pre int) (cur_pre int, successor *TreeNode) {
	if t == nil {
		return pre, nil
	}
	fmt.Printf("before left pre: %d\n", pre)
	pre, successor = inorder(t.Left, p, pre)
	if successor != nil {
		return pre, successor
	}
	fmt.Printf("after left pre: %d\n", pre)

	fmt.Printf("in mid!%d, pre: %d\n", t.Val, pre)
	if pre == t.Val {
		return pre, t
	}
	pre = t.Val
	fmt.Printf("after mid!%d, pre: %d\n", t.Val, pre)

	fmt.Printf("before Right pre: %d\n", pre)
	pre, successor = inorder(t.Right, p, pre)
	if successor != nil {
		return pre, successor
	}
	fmt.Printf("after Right pre: %d\n", pre)

	return pre, nil
}
