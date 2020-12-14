package sequenceReconstruction

type Node struct {
	val  int
	next map[int]*Node
}

func sequenceReconstruction(org []int, seqs [][]int) bool {
	m := make(map[int]*Node)

	for _, v := range seqs {
		for j, _ := range v {
			if _, ok := m[v[j]]; !ok {
				vv := v[j]
				m[vv] = &Node{vv, make(map[int]*Node)}
			}
			if j > 0 {
				vv := v[j]
				par := v[j-1]
				m[par].next[vv] = m[vv]
			}
		}
	}

	if len(m) != len(org) {
		return false
	}

	n := m[org[0]]
	if n == nil {
		return false
	}

	if !isDepthNotExceed(n, 0, len(org)) {
		return false
	}
	for i, v := range org {
		if n.val == v {
			if i < len(org)-1 {
				if nn, ok := n.next[org[i+1]]; !ok {
					return false
				} else {
					n = nn
				}
			}
		} else {
			return false
		}
	}

	if len(n.next) > 0 {
		return false
	}
	return true
}

func isDepthNotExceed(n *Node, d, maxd int) bool {
	if d > maxd {
		return false
	}

	for _, v := range n.next {
		if !isDepthNotExceed(v, d+1, maxd) {
			return false
		}
	}

	return true
}
