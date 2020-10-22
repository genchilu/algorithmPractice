package lruCache

type LRUCache struct {
	m  map[int]int
	l  *[]int
	lm map[int]int
	c  int
}

func Constructor(capacity int) LRUCache {
	m := make(map[int]int)
	c := capacity
	l := []int{}
	lm := make(map[int]int)

	return LRUCache{m, &l, lm, c}
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.m[key]; ok {
		this.updateLRU(key)
		return v
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.m[key]; !ok {
		if len(this.m) == this.c {
			ll := (*this.l)[len(*this.l)-1]
			(*this.l) = (*this.l)[:len(*this.l)-1]
			delete(this.m, ll)
			delete(this.lm, ll)
		}
	}
	this.m[key] = value
	this.updateLRU(key)
}

func (this *LRUCache) updateLRU(key int) {
	if v, ok := this.lm[key]; ok {
		for i := 0; i < v; i++ {
			this.lm[(*this.l)[i]]++
		}
		(*this.l) = append((*this.l)[0:v], (*this.l)[v+1:]...)
	} else {
		for k, _ := range this.lm {
			this.lm[k]++
		}
	}
	(*this.l) = append([]int{key}, (*this.l)...)
	this.lm[key] = 0
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
