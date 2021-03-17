package sellDiminishingValuedColoredBalls

import (
	"container/heap"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[0 : len(*h)-1]
	return x
}

func maxProfit(inventory []int, orders int) int {
	h := &IntHeap{}
	heap.Init(h)
	heap.Push(h, 0)
	for _, v := range inventory {
		heap.Push(h, v)
	}

	count := 1
	target := heap.Pop(h).(int)
	sum := 0
	for orders > 0 {
		p := heap.Pop(h).(int)
		d := target - p
		r := d * count
		if orders < r {
			d = orders / count
			sum += count * ((target + target - d + 1) * d / 2)

			target -= d
			sum += (orders % count) * target

			break
		}

		orders -= r

		sum += (d * (target + p + 1) / 2) * count

		target = p
		count++
	}

	return sum % 1000000007
}
