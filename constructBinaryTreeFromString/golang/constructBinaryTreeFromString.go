package constructBinaryTreeFromString

import (
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func str2tree(s string) *TreeNode {
	if len(s) == 0 {
		return nil
	}

	i := 0
	stack := []*TreeNode{}
	for i < len(s) {

		if s[i] == ')' {
			stack = stack[:len(stack)-1]
		} else if s[i] != '(' {
			tmp := i
			for i < len(s)-1 {
				if s[i+1] == '(' || s[i+1] == ')' {
					break
				}
				i++
			}
			val, _ := strconv.Atoi(s[tmp : i+1])

			node := &TreeNode{Val: val}

			if len(stack) > 0 {
				r := stack[len(stack)-1]
				if r.Left == nil {
					r.Left = node
				} else {
					r.Right = node
				}

			}

			stack = append(stack, node)
		}
		i++
	}

	return stack[0]
}
