package partitionLabels

func partitionLabels(S string) []int {
	if len(S) == 0 {
		return []int{}
	}

	m := make(map[byte]int)
	for i := 0; i < len(S); i++ {
		m[S[i]] = i
	}

	l := 0
	r := m[S[l]]

	result := []int{}
	for r < len(S) {
		for i := l; i <= r; i++ {
			if m[S[i]] > r {
				r = m[S[i]]
			}
		}
		result = append(result, r-l+1)

		l = r + 1
		if l == len(S) {
			break
		} else {
			r = m[S[l]]
		}
	}

	return result
}
