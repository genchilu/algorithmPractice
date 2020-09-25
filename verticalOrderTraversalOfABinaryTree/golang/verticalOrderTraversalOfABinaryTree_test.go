package verticalOrderTraversalOfABinaryTree

import (
	"math"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	// "math"
	//"fmt"
)

func TestVerticalTraversal(t *testing.T) {

	testCases := []struct {
		treeNode *TreeNode
		expect   [][]int
	}{
		{
			prepareTreeNode([]string{"3", "9", "20", "null", "null", "15", "7"}, 0, 0),
			[][]int{
				[]int{9},
				[]int{3, 15},
				[]int{20},
				[]int{7},
			},
		},
		{
			prepareTreeNode([]string{"1", "2", "3", "4", "5", "6", "7"}, 0, 0),
			[][]int{
				[]int{4},
				[]int{2},
				[]int{1, 5, 6},
				[]int{3},
				[]int{7},
			},
		},
	}

	for _, testCase := range testCases {
		actual := verticalTraversal(testCase.treeNode)
		assert.Equal(t, testCase.expect, actual)
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
