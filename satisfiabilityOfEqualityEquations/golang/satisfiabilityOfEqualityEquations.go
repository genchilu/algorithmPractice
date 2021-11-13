package satisfiabilityOfEqualityEquations

func equationsPossible(equations []string) bool {

	g := make([]int, 26)
	m := make(map[int][]int)
	for i := 0; i < 26; i++ {
		g[i] = i
		m[i] = []int{}
		m[i] = append(m[i], i)
	}

	count := 26
	for _, equation := range equations {
		a := int(equation[0] - 'a')
		b := int(equation[3] - 'a')
		s := equation[1]

		if s == '=' {
			ga := g[a]
			gb := g[b]
			if ga != gb {
				m[count] = []int{}

				m[count] = append(m[count], m[ga]...)
				for _, n := range m[ga] {
					g[n] = count
				}

				m[count] = append(m[count], m[gb]...)
				for _, n := range m[gb] {
					g[n] = count
				}

				count++
			}

		}
	}

	for _, equation := range equations {
		a := int(equation[0] - 'a')
		b := int(equation[3] - 'a')
		s := equation[1]

		if s == '!' {
			ga := g[a]
			gb := g[b]
			if ga == gb {
				return false
			}
		}
	}
	return true
}
