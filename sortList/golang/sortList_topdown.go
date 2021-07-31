package sortList

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	l, r := split(head)

	sortl := sortList(l)
	sortr := sortList(r)

	return merge(sortl, sortr)
}

func merge(l, r *ListNode) *ListNode {

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

	return result.Next
}

func split(head *ListNode) (l, r *ListNode) {
	mid, tail := head, head
	for tail.Next != nil && tail.Next.Next != nil {
		mid = mid.Next
		tail = tail.Next.Next
	}

	preMid := mid
	mid = mid.Next
	preMid.Next = nil
	return head, mid
}
