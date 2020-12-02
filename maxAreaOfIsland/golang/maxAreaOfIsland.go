package maxAreaOfIsland

func maxAreaOfIsland(grid [][]int) int {

	h := len(grid)
	w := 0
	if h > 0 {
		w = len(grid[0])
	}

	max := 0
	for i, _ := range grid {
		for j, _ := range grid[i] {
			c := dfs(i, j, h, w, grid)
			if c > max {
				max = c
			}
		}
	}
	return max
}

func dfs(i, j, h, w int, grid [][]int) int {
	if i < 0 || i >= h || j < 0 || j >= w || grid[i][j] == 0 {
		return 0
	} else {
		grid[i][j] = 0
		c := 1
		c += dfs(i+1, j, h, w, grid)
		c += dfs(i-1, j, h, w, grid)
		c += dfs(i, j+1, h, w, grid)
		c += dfs(i, j-1, h, w, grid)
		return c
	}
}
