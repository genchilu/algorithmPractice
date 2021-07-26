package singleThreadedCPU

import (
	"container/heap"
	"sort"
)

type Task struct {
	idx int
	t   []int
}

type TaskHeap []Task

func (h TaskHeap) Len() int { return len(h) }

func (h TaskHeap) Less(i, j int) bool {
	if h[i].t[1] < h[j].t[1] {
		return true
	} else if h[i].t[1] > h[j].t[1] {
		return false
	} else {
		return h[i].idx < h[j].idx
	}
}

func (h TaskHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *TaskHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(Task))
}

func (h *TaskHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func getOrder(tasks [][]int) []int {

	result := []int{}

	h := &TaskHeap{}
	heap.Init(h)

	q := make([]Task, len(tasks))
	for i, t := range tasks {
		task := Task{idx: i, t: t}
		q[i] = task
		//heap.Push(h, task)
	}
	sort.Slice(q, func(i, j int) bool {
		return q[i].t[0] < q[j].t[0]
	})

	idx := 0
	ct := 0
	for h.Len() > 0 || idx < len(q) {
		for idx < len(q) {
			if q[idx].t[0] <= ct {
				heap.Push(h, q[idx])
				idx++
			} else {
				break
			}
		}

		if h.Len() > 0 {
			t := heap.Pop(h).(Task)
			ct += t.t[1]
			result = append(result, t.idx)
		} else {
			ct++
		}

	}
	return result
}
