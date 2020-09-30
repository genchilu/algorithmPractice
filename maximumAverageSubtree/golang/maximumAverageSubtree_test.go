package maximumAverageSubtree

import (
	"math"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaximumAverageSubtree(t *testing.T) {
	testCases := []struct {
		root   *TreeNode
		expect float64
	}{
		{
			prepareTreeNode([]string{"5", "6", "1"}, 0, 0),
			6.0,
		},
	}

	for _, testCase := range testCases {
		actual := maximumAverageSubtree(testCase.root)
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
