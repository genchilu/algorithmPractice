package shortestBridge

func shortestBridge(grid [][]int) int {

	m := len(grid)
	n := len(grid[0])

	no := 2

	coastals := make([][][]int, 4)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				coastals[no] = findIsland(i, j, m, n, no, &grid)
				no++
			}
		}
	}

	// fmt.Printf("%v\n", coastals)
	// fmt.Printf("%v\n", grid)
	min := m + n
	for _, c1 := range coastals[2] {
		for _, c2 := range coastals[3] {
			dis := caculDist(c1, c2)

			if dis < min {
				min = dis
			}
		}
	}

	return min
}

func findIsland(i, j, m, n, no int, grid *[][]int) [][]int {
	coastal := [][]int{}

	(*grid)[i][j] = no
	isCoastal := false

	if isValid(i+1, j, m, n) {
		if (*grid)[i+1][j] == 0 {
			isCoastal = true
		} else if (*grid)[i+1][j] == 1 {
			coastal = append(coastal, findIsland(i+1, j, m, n, no, grid)...)
		}
	}

	if isValid(i-1, j, m, n) {
		if (*grid)[i-1][j] == 0 {
			isCoastal = true
		} else if (*grid)[i-1][j] == 1 {
			coastal = append(coastal, findIsland(i-1, j, m, n, no, grid)...)
		}
	}

	if isValid(i, j+1, m, n) {
		if (*grid)[i][j+1] == 0 {
			isCoastal = true
		} else if (*grid)[i][j+1] == 1 {
			coastal = append(coastal, findIsland(i, j+1, m, n, no, grid)...)
		}
	}

	if isValid(i, j-1, m, n) {
		if (*grid)[i][j-1] == 0 {
			isCoastal = true
		} else if (*grid)[i][j-1] == 1 {
			coastal = append(coastal, findIsland(i, j-1, m, n, no, grid)...)
		}
	}

	// fmt.Printf("i: %d, j: %d, g: %d, %v\n", i, j, (*grid)[i][j], isCoastal)
	// fmt.Printf("before: %v\n", coastal)
	if isCoastal {
		coastal = append(coastal, []int{i, j})
	}
	// fmt.Printf("after: %v\n", coastal)
	return coastal
}

func isValid(i, j, m, n int) bool {
	return (i >= 0 && i < m && j >= 0 && j < n)
}

func caculDist(g1, g2 []int) int {
	x := g1[0] - g2[0]
	if x < 0 {
		x = -x
	}

	y := g1[1] - g2[1]
	if y < 0 {
		y = -y
	}

	return x + y - 1
}
