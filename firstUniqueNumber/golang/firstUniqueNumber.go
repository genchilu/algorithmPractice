package firstUniqueNumber

type FirstUnique struct {
	m map[int]int
	q *[]int
}

func Constructor(nums []int) FirstUnique {
	m := make(map[int]int)
	q := []int{}

	for _, v := range nums {
		if _, ok := m[v]; !ok {
			q = append(q, v)
			m[v] = 0
		}
		m[v]++
	}

	return FirstUnique{m, &q}
}

func (this *FirstUnique) ShowFirstUnique() int {
	for len(*this.q) > 0 {
		h := (*this.q)[0]
		if this.m[h] == 1 {
			return h
		} else {
			(*this.q) = (*this.q)[1:]
		}
	}
	return -1
}

func (this *FirstUnique) Add(value int) {
	if _, ok := this.m[value]; !ok {
		(*this.q) = append((*this.q), value)
		this.m[value] = 0
	}
	this.m[value]++
}
