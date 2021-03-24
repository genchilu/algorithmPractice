package pathWithMaximumGold

func getMaximumGold(grid [][]int) int {

	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	max := 0
	for i, _ := range grid {
		for j, _ := range grid[i] {
			tmp := findMax(grid, i, j)
			if tmp > max {
				max = tmp
			}
		}
	}
	return max
}

func findMax(grid [][]int, i, j int) int {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) || grid[i][j] == 0 {
		return 0
	}

	value := grid[i][j]
	grid[i][j] = 0
	max := 0

	tmp := findMax(grid, i-1, j)
	if tmp > max {
		max = tmp
	}

	tmp = findMax(grid, i+1, j)
	if tmp > max {
		max = tmp
	}

	tmp = findMax(grid, i, j-1)
	if tmp > max {
		max = tmp
	}

	tmp = findMax(grid, i, j+1)
	if tmp > max {
		max = tmp
	}

	grid[i][j] = value
	return max + value

}
