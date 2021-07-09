package connectingCitiesWithMinimumCost

import (
	"container/heap"
)

type EdgeHeap [][]int

func (h EdgeHeap) Len() int           { return len(h) }
func (h EdgeHeap) Less(i, j int) bool { return h[i][2] < h[j][2] }
func (h EdgeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *EdgeHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *EdgeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func minimumCost(n int, connections [][]int) int {
	n2g := make(map[int]int)
	g2n := make(map[int][]int)

	h := &EdgeHeap{}
	heap.Init(h)

	for i := 1; i <= n; i++ {
		n2g[i] = i
		g2n[i] = []int{i}
	}

	for _, conn := range connections {
		heap.Push(h, conn)
	}

	newGroup := n + 1
	cost := 0
	for h.Len() > 0 {
		e := heap.Pop(h).([]int)

		g1 := n2g[e[0]]
		g2 := n2g[e[1]]

		if g1 != g2 {
			cost += e[2]
			g2n[newGroup] = []int{}

			for _, v := range g2n[g1] {
				n2g[v] = newGroup
				g2n[newGroup] = append(g2n[newGroup], v)
			}
			for _, v := range g2n[g2] {
				n2g[v] = newGroup
				g2n[newGroup] = append(g2n[newGroup], v)
			}

			newGroup++
			delete(g2n, g1)
			delete(g2n, g2)

			if len(g2n) == 1 {
				return cost
			}
		}
	}
	return -1
}
