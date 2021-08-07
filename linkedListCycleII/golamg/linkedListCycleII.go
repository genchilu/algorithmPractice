package linkedListCycleII

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	f, s := head, head
	for {
		if f.Next == nil || f.Next.Next == nil {
			return nil
		}
		f = f.Next.Next
		s = s.Next

		if s == f {
			break
		}
	}

	if s == head {
		return s
	}

	s = head

	for {
		f = f.Next
		s = s.Next
		if s == f {
			break
		}
	}

	return s
}
