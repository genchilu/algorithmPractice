package constructBinaryTreeFromPreorderAndInorderTraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := preorder[0]
	count := 0
	for count < len(inorder) {
		if inorder[count] == root {
			break
		}
		count++
	}

	linorder := inorder[0:count]
	rinorder := inorder[count+1:]

	lpreorder := preorder[1 : count+1]
	rpreorder := preorder[count+1:]

	n := &TreeNode{root, buildTree(lpreorder, linorder), buildTree(rpreorder, rinorder)}

	return n
}
