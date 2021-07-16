package shortestPathToGetFood

func getFood(grid [][]byte) int {
	vistited := make([][]int, len(grid))
	for i := range vistited {
		vistited[i] = make([]int, len(grid[i]))
		for j := range vistited[i] {
			vistited[i][j] = -1
		}
	}

	q := [][]int{}
	q = append(q, findLocaltion(grid))
	step := 0
	for len(q) > 0 {
		newq := [][]int{}
		for _, g := range q {
			x, y := g[0], g[1]
			if vistited[x][y] == -1 {
				if grid[x][y] == '#' {
					return step
				}
				vistited[x][y] = step

				if isValidPos(grid, x+1, y) {
					newq = append(newq, []int{x + 1, y})
				}
				if isValidPos(grid, x-1, y) {
					newq = append(newq, []int{x - 1, y})
				}
				if isValidPos(grid, x, y+1) {
					newq = append(newq, []int{x, y + 1})
				}
				if isValidPos(grid, x, y-1) {
					newq = append(newq, []int{x, y - 1})
				}
			}
		}
		step++
		q = newq
	}

	return -1
}

func isValidPos(grid [][]byte, x, y int) bool {
	if x >= len(grid) || x < 0 {
		return false
	}

	if y >= len(grid[0]) || y < 0 {
		return false
	}

	if grid[x][y] == 'X' {
		return false
	}

	return true

}

func findLocaltion(grid [][]byte) []int {
	for i, v := range grid {
		for j := range v {
			if grid[i][j] == '*' {
				return []int{i, j}
			}
		}
	}

	return []int{-1, -1}
}
