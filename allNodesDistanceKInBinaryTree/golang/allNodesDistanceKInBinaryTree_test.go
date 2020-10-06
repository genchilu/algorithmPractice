package allNodesDistanceKInBinaryTree

import (
	"math"
	"sort"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	//"fmt"
)

func TestDistanceK(t *testing.T) {
	testCases := []struct {
		tree   *TreeNode
		target int
		k      int
		expect []int
	}{
		{
			prepareTreeNode([]string{"3", "5", "1", "6", "2", "0", "8", "null", "null", "7", "4"}, 0, 0),
			5,
			2,
			[]int{7, 4, 1},
		},
	}

	for _, testCase := range testCases {

		tt := findTarget(testCase.tree, testCase.target)
		actualResult := distanceK(testCase.tree, tt, testCase.k)

		sort.Ints(actualResult)
		sort.Ints(testCase.expect)
		assert.Equal(t, testCase.expect, actualResult)
	}
}

func prepareTreeNode(input []string, idx int, h int) *TreeNode {

	if idx < len(input) {
		if input[idx] == "null" {
			return nil
		}

		node := &TreeNode{}
		node.Val, _ = strconv.Atoi(input[idx])

		leftIdx := (int)(math.Pow(2, float64(h+1))) - 1 + 2*((idx+1)%((int)(math.Pow(2, float64(h)))))
		node.Left = prepareTreeNode(input, leftIdx, h+1)

		rightIdx := leftIdx + 1
		node.Right = prepareTreeNode(input, rightIdx, h+1)

		return node
	} else {
		return nil
	}
}

func findTarget(n *TreeNode, target int) *TreeNode {
	if n.Val == target {
		return n
	} else {
		if n.Left != nil {
			t := findTarget(n.Left, target)
			if t != nil {
				return t
			}
		}
		if n.Right != nil {
			t := findTarget(n.Right, target)
			if t != nil {
				return t
			}
		}

		return nil
	}
}
