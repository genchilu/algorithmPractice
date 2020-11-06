package evaluateDivision

type Node struct {
	lab       string
	neighbors *[]*Node
	values    *[]float64
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {

	m := make(map[string]*Node)
	for i, equation := range equations {
		p := equation[0]
		c := equation[1]
		value := values[i]

		if _, ok := m[p]; !ok {
			m[p] = &Node{p, &[]*Node{}, &[]float64{}}
		}

		if _, ok := m[c]; !ok {
			m[c] = &Node{c, &[]*Node{}, &[]float64{}}
		}

		*m[p].neighbors = append(*m[p].neighbors, m[c])
		*m[p].values = append(*m[p].values, value)

		*m[c].neighbors = append(*m[c].neighbors, m[p])
		*m[c].values = append(*m[c].values, 1.0/value)
	}

	result := []float64{}
	for _, querie := range queries {
		s := querie[0]
		t := querie[1]

		vals := []float64{1.0}
		vistited := make(map[string]bool)
		if _, ok := m[s]; ok {
			result = append(result, dfs(m[s], &vals, t, vistited))
		} else {
			result = append(result, -1.0)
		}
	}
	return result
}

func dfs(n *Node, vals *[]float64, t string, vistited map[string]bool) float64 {
	vistited[n.lab] = true
	if n.lab == t {
		m := 1.0
		for _, v := range *vals {
			m = m * v
		}

		return m
	} else {
		for i, v := range *n.neighbors {

			if _, ok := vistited[v.lab]; !ok {
				(*vals) = append(*vals, (*n.values)[i])
				tmp := dfs(v, vals, t, vistited)

				if tmp != -1.0 {
					return tmp
				}

				(*vals) = (*vals)[0 : len(*vals)-1]
			}
		}
	}

	return -1.0
}
