package maxAreaOfIsland

func maxAreaOfIsland(grid [][]int) int {
	m := make(map[int]map[int]bool)

	h := len(grid)
	w := 0
	if h > 0 {
		w = len(grid[0])
	}

	max := 0
	for i, _ := range grid {
		for j, _ := range grid[i] {
			c := dfs(i, j, h, w, grid, m)
			if c > max {
				max = c
			}
		}
	}
	return max
}

func dfs(i, j, h, w int, grid [][]int, m map[int]map[int]bool) int {
	if i < 0 || i >= h || j < 0 || j >= w || grid[i][j] == 0 {
		return 0
	}

	if _, ok := m[i]; !ok {
		m[i] = make(map[int]bool)
	}
	if _, ok := m[i][j]; !ok {
		m[i][j] = true
		c := 1
		c += dfs(i+1, j, h, w, grid, m)
		c += dfs(i-1, j, h, w, grid, m)
		c += dfs(i, j+1, h, w, grid, m)
		c += dfs(i, j-1, h, w, grid, m)
		return c
	} else {
		return 0
	}
}
