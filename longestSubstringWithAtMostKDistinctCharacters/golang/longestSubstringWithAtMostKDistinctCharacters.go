package longestSubstringWithAtMostKDistinctCharacters

type Node struct {
	val  byte
	idx  int
	next *Node
	pre  *Node
}

func lengthOfLongestSubstringKDistinct(s string, k int) int {
	if k == 0 || len(s) == 0 {
		return 0
	}

	m := make(map[byte]*Node)

	h := &Node{s[0], 0, nil, nil}
	m[s[0]] = h
	t := h
	lidx := 0
	max := 0

	for i := 1; i < len(s); i++ {
		if n, ok := m[s[i]]; !ok {
			n = &Node{s[i], i, nil, nil}
			t.next = n
			n.pre = t
			t = t.next
			m[s[i]] = n

			if len(m) > k {
				d := i - lidx
				if d > max {
					max = d
				}

				lidx = m[h.val].idx + 1
				delete(m, h.val)
				h = h.next
				h.pre = nil
			}
		} else {
			n.idx = i
			if t != n {
				if h == n {
					h = n.next
					h.pre = nil
				} else {
					n.next.pre = n.pre
					n.pre.next = n.next
				}

				t.next = n
				n.pre = t
				n.next = nil
				t = n
			}
		}
	}

	if len(m) <= k {
		d := len(s) - lidx
		if d > max {
			max = d
		}
	}
	return max
}
