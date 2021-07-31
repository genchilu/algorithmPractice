package sortList

import "fmt"

func sortList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	size := 1
	c := count(head)

	for size <= c {
		var l, r, cur, newh, tail *ListNode
		newh, tail = nil, nil
		cur = head

		//fmt.Printf("0000 cur: %s, size: %d\n", nodt2Str(cur), size)
		for {
			l, r, cur = splic2(cur, size)
			//fmt.Printf("1111 l: %s, r: %s, cur: %s\n", nodt2Str(l), nodt2Str(r), nodt2Str(cur))
			if l == nil {
				break
			}
			if r == nil {
				if tail != nil {
					tail.Next = l
				}
				break
			}

			p, t := merge2(l, r)

			if newh == nil {
				newh = p
			}

			if tail != nil {
				tail.Next = p
			}

			tail = t

			//fmt.Printf("2222 p: %s, newh: %s, tail: %s\n", nodt2Str(p), nodt2Str(newh), nodt2Str(tail))
		}
		if newh != nil {
			head = newh
		}
		//fmt.Printf("3333 head: %s\n", nodt2Str(head))
		size *= 2
	}

	return head
}

func count(head *ListNode) int {
	count := 1
	for head.Next != nil {
		head = head.Next
		count++
	}

	return count
}

func splic2(heap *ListNode, size int) (*ListNode, *ListNode, *ListNode) {
	if heap == nil {
		return nil, nil, nil
	}
	var l, r, pre *ListNode
	l = heap
	r = heap
	for i := 0; i < size; i++ {
		pre = r
		if r.Next != nil {
			r = r.Next
		} else {
			r = nil
			break
		}
	}

	if r == nil {
		return l, nil, nil
	}

	pre.Next = nil
	heap = r
	for i := 0; i < size; i++ {
		pre = heap
		if heap.Next != nil {
			heap = heap.Next
		} else {
			heap = nil
			break
		}
	}

	pre.Next = nil

	return l, r, heap
}

func merge2(l, r *ListNode) (*ListNode, *ListNode) {

	result := &ListNode{}
	cur := result

	for l != nil || r != nil {
		if l == nil {
			cur.Next = r
			r = nil
		} else if r == nil {
			cur.Next = l
			l = nil
		} else if l.Val < r.Val {
			cur.Next = l
			l = l.Next
		} else {
			cur.Next = r
			r = r.Next
		}
		cur = cur.Next
	}

	for cur.Next != nil {
		cur = cur.Next
	}

	return result.Next, cur
}

func nodt2Str(n *ListNode) string {
	if n == nil {
		return ""
	}
	tmp := ""
	for n != nil {
		tmp += fmt.Sprintf("%d ", n.Val)
		n = n.Next
	}

	return tmp
}
