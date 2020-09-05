package binarySearchTreeIterator

import (
	"testing"
	"github.com/stretchr/testify/assert"
	//"fmt"
)

func TestSearchRange(t *testing.T) {

	actualResult := Constructor(prepareNode())
	
	assert.Equal(t, 3, actualResult.Next())
	assert.Equal(t, 7, actualResult.Next())
	assert.Equal(t, true, actualResult.HasNext())
	assert.Equal(t, 9, actualResult.Next())
	assert.Equal(t, true, actualResult.HasNext())
	assert.Equal(t, 15, actualResult.Next())
	assert.Equal(t, true, actualResult.HasNext())
	assert.Equal(t, 20, actualResult.Next())
	assert.Equal(t, false, actualResult.HasNext())
}

func prepareNode() *TreeNode{
	treeNode3 := TreeNode{}
	treeNode3.Val = 3

	treeNode7 := TreeNode{}
	treeNode7.Val = 7

	treeNode9 := TreeNode{}
	treeNode9.Val = 9

	treeNode15 := TreeNode{}
	treeNode15.Val = 15

	treeNode20:= TreeNode{}
	treeNode20.Val = 20

	treeNode7.Left = &treeNode3
	treeNode7.Right = &treeNode15

	treeNode15.Left = &treeNode9
	treeNode15.Right = &treeNode20

	return &treeNode7

}