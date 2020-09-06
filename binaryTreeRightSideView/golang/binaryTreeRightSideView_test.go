package binaryTreeRightSideView

import (
	"math"
	"testing"
	"github.com/stretchr/testify/assert"
	"strconv"
	//"fmt"
)

func TestRightSideView(t *testing.T) {
	testCases := []struct{
		treeNode *TreeNode
		expect []int
	} {
		{
			prepareTreeNode([]string{"1","2","3","null","5", "null","4"}, 0 ,0),
			[]int{1,3,4},
		},
		{
			prepareTreeNode([]string{}, 0 ,0),
			[]int{},
		},
		{
			nil,
			[]int{},
		},
	}

	for _, testCase := range testCases {

		actualResult := rightSideView(testCase.treeNode)
		assert.Equal(t, testCase.expect, actualResult)
		
		// assert.Equal(t, len(testCase.expectResult), len(actualResult))
		// for idx := range(testCase.expectResult) {
		// 	assert.Equal(t, testCase.expectResult[idx], actualResult[idx])
		// }
	}
}


func prepareTreeNode(input []string, idx int, h int) *TreeNode {

	
	if(idx < len(input)) {
		if(input[idx] == "null") {
			return nil
		}

		node := &TreeNode{}
		node.Val, _ = strconv.Atoi(input[idx])

		leftIdx := (int)(math.Pow(2, float64(h+1))) - 1 + 2 * ((idx+1) % ((int)(math.Pow(2,float64(h)))))
		node.Left = prepareTreeNode(input, leftIdx, h+1)

		rightIdx := leftIdx + 1
		node.Right = prepareTreeNode(input, rightIdx, h+1)

		return node
	} else {
		return nil
	}
		
}