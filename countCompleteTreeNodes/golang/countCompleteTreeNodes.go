package countCompleteTreeNodes

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	d := calculateDepth(root)
	if d == 0 {
		return 1
	}

	tmp := int(math.Pow(float64(2), float64(d)))
	left := 1
	right := tmp

	for left <= right {
		pivot := left + (right-left)/2
		if exist(root, d, pivot) {
			left = pivot + 1
		} else {
			right = pivot - 1
		}
	}

	if left > tmp {
		left = tmp
	}
	return tmp - 1 + left
}

func exist(n *TreeNode, d, idx int) bool {
	tmp := n
	left := 0
	right := int(math.Pow(float64(2), float64(d))) - 1

	for i := 0; i < d; i++ {
		pivot := left + (right-left)/2
		if idx <= pivot {
			right = pivot
			tmp = tmp.Left
		} else {
			left = pivot + 1
			tmp = tmp.Right
		}
	}

	return tmp != nil
}

func calculateDepth(n *TreeNode) int {
	d := 0

	tmp := n
	for tmp.Left != nil {
		d++
		tmp = tmp.Left
	}

	return d
}
