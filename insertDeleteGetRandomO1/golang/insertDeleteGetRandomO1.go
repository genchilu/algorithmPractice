package insertDeleteGetRandomO1

import (
	"math/rand"
)

type RandomizedSet struct {
	m   map[int]int
	idx *[]int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	m := make(map[int]int)
	idx := []int{}
	return RandomizedSet{m, &idx}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.m[val]; ok {
		return false
	} else {
		this.m[val] = len(*this.idx)
		(*this.idx) = append(*this.idx, val)
		return true
	}
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	if i, ok := this.m[val]; ok {
		li := len(*this.idx) - 1
		le := (*this.idx)[li]

		this.m[le] = i
		(*this.idx)[i] = le

		delete(this.m, val)
		(*this.idx) = (*this.idx)[0:li]

		return true
	} else {
		return false
	}
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	picki := rand.Intn(len(*this.idx))

	return (*this.idx)[picki]
}
