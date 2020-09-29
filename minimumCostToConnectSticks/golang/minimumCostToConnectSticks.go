package minimumCostToConnectSticks

import (
	"container/heap"
	"sort"
)

type Interface interface {
	sort.Interface
	Push(x interface{})
	Pop() interface{}
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[0 : len(*h)-1]
	return x
}

func connectSticks(sticks []int) int {
	if len(sticks) < 2 {
		return 0
	}

	h := &IntHeap{}
	heap.Init(h)
	for _, v := range sticks {
		heap.Push(h, v)
	}

	c := 0
	for h.Len() > 1 {
		sum := heap.Pop(h).(int) + heap.Pop(h).(int)
		heap.Push(h, sum)
		c += sum
	}
	return c
}
