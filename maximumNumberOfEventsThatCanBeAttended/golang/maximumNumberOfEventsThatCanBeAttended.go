package maximumNumberOfEventsThatCanBeAttended

import (
	"container/heap"
	"sort"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func maxEvents(events [][]int) int {

	sort.SliceStable(events, func(i, j int) bool {
		return events[i][0] < events[j][0]
	})

	idx := 0
	h := &IntHeap{}
	heap.Init(h)
	count := 0
	for i := 0; i <= 100000; i++ {
		for idx < len(events) {
			if events[idx][0] == i {
				e := events[idx][1]
				heap.Push(h, e)
			} else if events[idx][0] > i {
				break
			}
			idx++
		}

		for h.Len() > 0 && (*h)[0] < i {
			heap.Pop(h)
		}

		if h.Len() > 0 {
			count++
			heap.Pop(h)
		}

		if idx >= len(events) && h.Len() == 0 {
			break
		}
	}
	return count
}
