package RLEIterator

type RLEIterator struct {
	A []int
}

func Constructor(A []int) RLEIterator {
	return RLEIterator{A}
}

func (this *RLEIterator) Next(n int) int {
	if len(this.A) == 0 {
		return -1
	}

	if this.A[0] >= n {
		this.A[0] -= n
		return this.A[1]
	} else {
		c := this.A[0]
		this.A = this.A[2:]
		return this.Next(n - c)
	}
}
