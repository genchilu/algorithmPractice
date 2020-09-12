package binaryTreeVerticalOrderTraversal

import (
	"sort"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

type Msg struct {
	treeNode *TreeNode
	horizontalIdx int
}

func verticalOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	m := Msg{root, 0}
	queue := []*Msg{&m}
	horizontalMap := make(map[int][]int)
	bfs(&queue, horizontalMap)

	keys := make([]int, 0, len(horizontalMap))
	for k,_:=range horizontalMap {
		keys = append(keys, k)
	}

	sort.Ints(keys)
	result := make([][]int, 0, len(keys))
	for _,v:=range keys {
		result = append(result, horizontalMap[v])
	}

    return result
}

func bfs(queue *[]*Msg, horizontalMap map[int][]int) {
	m := (*queue)[0]
	(*queue) = (*queue)[1:]
	treeNode := m.treeNode
	horizontalIdx := m.horizontalIdx

	if _,ok:=horizontalMap[horizontalIdx];!ok{
		horizontalMap[horizontalIdx] = []int{};
	}
	horizontalSlice := horizontalMap[horizontalIdx]
	horizontalSlice = append(horizontalSlice, treeNode.Val)
	horizontalMap[horizontalIdx] = horizontalSlice


	if(treeNode.Left!=nil) {
		m := Msg{treeNode.Left, horizontalIdx-1}
		*queue = append(*queue, &m)
	}

	if(treeNode.Right!=nil) {
		m := Msg{treeNode.Right, horizontalIdx+1}
		*queue = append(*queue, &m)
	}

	if (len(*queue) > 0) {
		bfs(queue, horizontalMap)
	}
}

