package kClosestPointsToOrigin

import (
	"container/heap"
	//"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[0:len(*h)-1]
	return x
}

func KClosest(points [][]int, K int) [][]int {
	if len(points) == 0 {
		return [][]int{}
	}

	h := &IntHeap{}
	heap.Init(h)

	dmap := make(map[int][][]int)
	for _, p:=range points {
		e := p[0] * p[0] + p[1] * p[1]
		if _,ok:=dmap[e];ok{
			dmap[e] = append(dmap[e], p)
		} else {
			dmap[e] = [][]int{p}
		}
		heap.Push(h,e)
	}


	result := make([][]int, 0, K)

	for h.Len() >0 && len(result) < K{
		v := heap.Pop(h).(int)
		result = append(result, dmap[v]...)
	}

    return result
}