package numMatchingSubseq

func numMatchingSubseq(S string, words []string) int {
	m := make(map[byte][]string)

	for _, w := range words {
		if _, ok := m[w[0]]; !ok {
			m[w[0]] = []string{}
		}
		m[w[0]] = append(m[w[0]], w)
	}

	for i := 0; i < len(S); i++ {
		char := S[i]
		if ws, ok := m[char]; ok {
			delete(m, char)
			for _, w := range ws {
				//fmt.Printf("%s\n", w)
				if len(w) > 1 {
					//fmt.Printf("%s put into m\n", w)
					w = w[1:]
					if _, ok2 := m[w[0]]; !ok2 {
						m[w[0]] = []string{}
					}
					m[w[0]] = append(m[w[0]], w)
				}
			}
		}
	}
	c := 0
	for _, v := range m {
		c += len(v)
	}
	return len(words) - c
}
