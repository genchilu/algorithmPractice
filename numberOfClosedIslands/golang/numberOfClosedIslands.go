package numberOfClosedIslands

func closedIsland(grid [][]int) int {
	c := 0
	for i, v := range grid {
		for j, _ := range v {
			if grid[i][j] == 0 && isClose(i, j, &grid) {
				// fmt.Printf("%d, %d\n", i, j)
				// for _, vv := range grid {
				// 	fmt.Printf("%v\n", vv)
				// }
				c++
			}
		}
	}
	return c
}

func isClose(i, j int, grid *[][]int) bool {
	if i >= len(*grid) || i < 0 || j >= len((*grid)[0]) || j < 0 {
		return false
	}

	if (*grid)[i][j] == 0 {
		(*grid)[i][j] = 1
		u := isClose(i+1, j, grid)
		d := isClose(i-1, j, grid)
		r := isClose(i, j+1, grid)
		l := isClose(i, j-1, grid)
		return u && d && l && r
	} else {
		return true
	}
}
