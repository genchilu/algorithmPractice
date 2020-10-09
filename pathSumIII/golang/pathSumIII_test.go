package pathSumIII

import (
	"math"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPathSum(t *testing.T) {
	testCases := []struct {
		root   *TreeNode
		sum    int
		expect int
	}{
		{
			prepareTreeNode([]string{"10", "5", "-3", "3", "2", "null", "11", "3", "-2", "null", "1"}, 0, 0),
			8,
			3,
		},
		{
			prepareTreeNode([]string{"1"}, 0, 0),
			0,
			0,
		},
	}

	for _, testCase := range testCases {
		actual := pathSum(testCase.root, testCase.sum)
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
