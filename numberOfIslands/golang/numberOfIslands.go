package numberOfIslands

func numIslands(grid [][]byte) int {

	m := make(map[int]map[int]bool)

	for i, _ := range grid {
		for j, _ := range grid[i] {
			if grid[i][j] == byte('1') {
				if _, ok := m[i]; !ok {
					m[i] = make(map[int]bool)
				}
				m[i][j] = true
			}
		}
	}

	c := 0
	for len(m) > 0 {
		c++
		for i, v := range m {
			for j, _ := range v {
				scan(i, j, m)
				break
			}
			break
		}
	}

	return c
}

func scan(i, j int, m map[int]map[int]bool) {
	if _, ok := m[i][j]; ok {
		delete(m[i], j)
		if len(m[i]) == 0 {
			delete(m, i)
		}
		scan(i-1, j, m)
		scan(i+1, j, m)
		scan(i, j-1, m)
		scan(i, j+1, m)
	}
}
