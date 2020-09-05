package binarySearchTreeIterator

// import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	q []*TreeNode
	latestSmallest int
}

func Constructor(root *TreeNode) BSTIterator {
	bstIterator := BSTIterator{[]*TreeNode{}, 0}
	bstIterator.travelToLeftestNode(root)
    return bstIterator
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	if (len(this.q) == 0) {
		return this.latestSmallest
	}

	treeNode := this.q[len(this.q)-1]
	this.q = this.q[0:(len(this.q)-1)]
	if(treeNode.Right != nil) {
		this.travelToLeftestNode(treeNode.Right)
	}

	this.latestSmallest = treeNode.Val
	
    return treeNode.Val
}


/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
    return len(this.q) != 0
}

func (this *BSTIterator) travelToLeftestNode(treeNode *TreeNode) {
	if(treeNode!=(*TreeNode)(nil)) {
		this.q = append(this.q, treeNode)
		this.travelToLeftestNode(treeNode.Left)
	}
}