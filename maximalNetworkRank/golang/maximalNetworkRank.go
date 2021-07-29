package maximalNetworkRank

func maximalNetworkRank(n int, roads [][]int) int {

	m := make(map[int][]int)
	for _, road := range roads {
		if _, ok := m[road[0]]; !ok {
			m[road[0]] = []int{}
		}
		m[road[0]] = append(m[road[0]], road[1])

		if _, ok := m[road[1]]; !ok {
			m[road[1]] = []int{}
		}
		m[road[1]] = append(m[road[1]], road[0])
	}

	city1 := -1
	max1 := 0
	for k, v := range m {
		if len(v) > max1 {
			max1 = len(v)
			city1 = k
		}
	}

	for _, c := range m[city1] {
		m[c] = m[c][1:]
	}
	delete(m, city1)

	max2 := 0
	for _, v := range m {
		if len(v) > max2 {
			max2 = len(v)
		}
	}

	return max1 + max2
}
