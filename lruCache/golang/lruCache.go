package lruCache

type Node struct {
	val  int
	key  int
	pre  *Node
	next *Node
}

type LRUCache struct {
	m    map[int]*Node
	c    int
	head *Node
	tail *Node
}

func Constructor(capacity int) LRUCache {
	m := make(map[int]*Node)

	return LRUCache{m, capacity, nil, nil}
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.m[key]; ok {
		this.updateLRU(v)
		return v.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if v, ok := this.m[key]; ok {
		this.updateLRU(v)
		v.val = value
	} else {
		if len(this.m) == this.c {
			l := this.head
			if this.head.next != nil {
				this.head.next.pre = nil
			}
			this.head = this.head.next

			delete(this.m, l.key)
		}
		if this.head == nil {
			node := Node{value, key, nil, nil}
			this.m[key] = &node
			this.head = &node
			this.tail = &node
		} else {
			node := Node{value, key, nil, nil}
			this.m[key] = &node
			node.pre = this.tail
			this.tail.next = &node
			this.tail = &node
		}
	}
}

func (this *LRUCache) updateLRU(v *Node) {
	if v.pre != nil && v.next != nil {
		pre := v.pre
		next := v.next

		pre.next = next
		next.pre = pre

		this.tail.next = v
		v.pre = this.tail
		v.next = nil
		this.tail = v
	} else if v.pre == nil && v.next != nil {
		next := v.next
		next.pre = nil
		this.head = next

		this.tail.next = v
		v.pre = this.tail
		v.next = nil
		this.tail = v
	}
}

// func (this *LRUCache) printList() {
// 	cur := this.head
// 	a := []int{}
// 	b := []int{}
// 	for cur != nil {
// 		a = append(a, cur.key)
// 		cur = cur.next
// 	}

// 	fmt.Printf("head to tail: %v\n", a)

// 	cur = this.tail
// 	for cur != nil {
// 		b = append(b, cur.key)
// 		cur = cur.pre
// 	}

// 	fmt.Printf("tail to head: %v\n", b)
// }

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
