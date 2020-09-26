package MaximumDifferenceBetweenNodeAndAncestor

import (
	"math"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	//"fmt"
)

func TestMaxAncestorDiff(t *testing.T) {
	testCases := []struct {
		root   *TreeNode
		expect int
	}{
		{
			prepareTreeNode([]string{"8", "3", "10", "1", "6", "null", "14", "null", "null", "4", "7", "13"}, 0, 0),
			7,
		},
	}

	for _, testCase := range testCases {
		actual := maxAncestorDiff(testCase.root)
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
