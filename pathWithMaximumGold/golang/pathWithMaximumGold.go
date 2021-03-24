package pathWithMaximumGold

func getMaximumGold(grid [][]int) int {

	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	max := 0
	for i, _ := range grid {
		for j, _ := range grid[i] {
			clone := getClone(grid)
			tmp := findMax(clone, i, j)
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
	max := 0
	cloneUp := getClone(grid)
	cloneUp[i][j] = 0

	cloneDown := getClone(cloneUp)

	cloneLeft := getClone(cloneUp)

	cloneRight := getClone(cloneUp)

	tmp := findMax(cloneUp, i-1, j)
	if tmp > max {
		max = tmp
		grid = cloneUp
	}

	tmp = findMax(cloneDown, i+1, j)
	if tmp > max {
		max = tmp
		grid = cloneDown
	}

	tmp = findMax(cloneLeft, i, j-1)
	if tmp > max {
		max = tmp
		grid = cloneLeft
	}

	tmp = findMax(cloneRight, i, j+1)
	if tmp > max {
		max = tmp
		grid = cloneRight
	}

	return max + value

}

func getClone(grid [][]int) [][]int {
	clone := make([][]int, len(grid))

	for i, _ := range grid {
		tmp := make([]int, len(grid[i]))
		copy(tmp, grid[i])
		clone[i] = tmp
	}

	return clone
}
