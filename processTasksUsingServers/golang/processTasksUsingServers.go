package processTasksUsingServers

import (
	"container/heap"
)

type Server struct {
	idx    int
	weight int
}

type ServerHeap []Server

func (h ServerHeap) Len() int { return len(h) }
func (h ServerHeap) Less(i, j int) bool {
	if h[i].weight < h[j].weight {
		return true
	} else if h[i].weight > h[j].weight {
		return false
	} else {
		return h[i].idx < h[j].idx
	}
}
func (h ServerHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *ServerHeap) Push(x interface{}) {
	*h = append(*h, x.(Server))
}

func (h *ServerHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type IntHeap []int

func (h IntHeap) Len() int { return len(h) }
func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h IntHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

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

func assignTasks(servers []int, tasks []int) []int {

	h := &ServerHeap{}
	heap.Init(h)

	th := &IntHeap{}
	heap.Init(th)

	m := make(map[int][]Server)
	m[0] = []Server{}

	for i, v := range servers {
		s := Server{i, v}
		//heap.Push(h, s)
		m[0] = append(m[0], s)
	}

	i := 0

	r := make([]int, len(tasks))

	heap.Push(th, 0)

	t := 0
	for i < len(tasks) {

		if h.Len() == 0 {
			t = heap.Pop(th).(int)
		} else {
			t++
		}

		for _, myserver := range m[t] {
			heap.Push(h, myserver)
		}
		delete(m, t)

		for h.Len() > 0 && i <= t && i < len(tasks) {
			finishtime := t + tasks[i]

			server := heap.Pop(h).(Server)

			r[i] = server.idx

			if _, ok := m[finishtime]; !ok {
				m[finishtime] = []Server{}
			}
			m[finishtime] = append(m[finishtime], server)

			i++
			heap.Push(th, finishtime)
		}
	}
	return r
}
