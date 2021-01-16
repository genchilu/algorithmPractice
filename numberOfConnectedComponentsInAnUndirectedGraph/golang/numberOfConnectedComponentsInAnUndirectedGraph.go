package numberOfConnectedComponentsInAnUndirectedGraph

func countComponents(n int, edges [][]int) int {
	isVistites := make(map[int]bool)

	m := make(map[int][]int)
	for _, edge := range edges {
		v0 := edge[0]
		v1 := edge[1]

		if _, ok := m[v0]; !ok {
			m[v0] = []int{}
		}
		if _, ok := m[v1]; !ok {
			m[v1] = []int{}
		}

		m[v0] = append(m[v0], v1)
		m[v1] = append(m[v1], v0)
	}

	c := 0

	for i := 0; i < n; i++ {
		component := []int{}
		dfs(i, m, isVistites, &component)
		if len(component) > 0 {
			c++
		}
	}

	return c
}

func dfs(cur int, m map[int][]int, isVistites map[int]bool, component *[]int) {
	if _, ok := isVistites[cur]; !ok {
		isVistites[cur] = true

		for _, next := range m[cur] {
			dfs(next, m, isVistites, component)
		}
		*component = append(*component, cur)
	}
}
