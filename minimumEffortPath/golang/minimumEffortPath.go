package minimumEffortPath

import "container/heap"

type Node struct {
	effort int
	i      int
	j      int
}

type NodeHeap []*Node

func (h NodeHeap) Len() int           { return len(h) }
func (h NodeHeap) Less(i, j int) bool { return h[i].effort < h[j].effort }
func (h NodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *NodeHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(*Node))
}

func (h *NodeHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[0 : len(*h)-1]
	return x
}

func minimumEffortPath(heights [][]int) int {
	h := len(heights)
	w := len(heights[0])
	isVistited := make([][]bool, h)
	for i, _ := range heights {
		isVistited[i] = make([]bool, w)
		for j, _ := range heights[i] {
			isVistited[i][j] = false
		}
	}

	n := Node{0, 0, 0}
	nh := &NodeHeap{}
	heap.Init(nh)
	heap.Push(nh, &n)
	for nh.Len() > 0 {
		curNode := heap.Pop(nh).(*Node)
		if curNode.i == h-1 && curNode.j == w-1 {
			return curNode.effort
		}
		isVistited[curNode.i][curNode.j] = true
		vistited(curNode.i-1, curNode.j, h, w, curNode, heights, nh, isVistited)
		vistited(curNode.i+1, curNode.j, h, w, curNode, heights, nh, isVistited)
		vistited(curNode.i, curNode.j-1, h, w, curNode, heights, nh, isVistited)
		vistited(curNode.i, curNode.j+1, h, w, curNode, heights, nh, isVistited)
	}

	return -1
}

func vistited(i, j, h, w int, n *Node, heights [][]int, nh *NodeHeap, isVistited [][]bool) {
	if isValidCell(i, j, h, w) && !isVistited[i][j] {
		abs := heights[n.i][n.j] - heights[i][j]
		if abs < 0 {
			abs = -abs
		}
		if n.effort > abs {
			abs = n.effort
		}
		newn := &Node{abs, i, j}
		heap.Push(nh, newn)
	}
}
func isValidCell(i, j, h, w int) bool {
	if i >= 0 && i < h && j >= 0 && j < w {
		return true
	}

	return false
}
