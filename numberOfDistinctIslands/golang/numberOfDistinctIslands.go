package numberOfDistinctIslands

import (
	"sort"
)

func numDistinctIslands(grid [][]int) int {

	islands := [][][]int{}

	for i, _ := range grid {
		for j, _ := range grid[i] {
			if grid[i][j] == 1 {
				island := [][]int{}
				caculateIsland(&grid, &island, i, j, 0)
				found := false
				for _, v := range islands {
					if compareIsland(island, v) {
						found = true
						break
					}
				}

				if !found {
					islands = append(islands, island)
				}
			}
		}
	}

	return len(islands)
}

func caculateIsland(grid *[][]int, island *[][]int, i, j, d int) {
	if len(*island) == d {
		(*island) = append(*island, []int{})
	}

	(*island)[d] = append((*island)[d], j)
	(*grid)[i][j] = 0

	if j < len((*grid)[i])-1 && (*grid)[i][j+1] == 1 {
		caculateIsland(grid, island, i, j+1, d)
	}
	if j > 0 && (*grid)[i][j-1] == 1 {
		caculateIsland(grid, island, i, j-1, d)
	}

	if i < len(*grid)-1 && (*grid)[i+1][j] == 1 {
		caculateIsland(grid, island, i+1, j, d+1)
	}
	if i > 0 && (*grid)[i-1][j] == 1 {
		caculateIsland(grid, island, i-1, j, d-1)
	}
}

func compareIsland(i1, i2 [][]int) bool {
	if len(i1) != len(i2) {
		return false
	}

	for i, _ := range i1 {
		if len(i1[i]) != len(i2[i]) {
			return false
		}

		sort.Ints(i1[i])
		sort.Ints(i2[i])

		if i1[i][0]-i1[0][0] != i2[i][0]-i2[0][0] {
			return false
		}
		for j, _ := range i1[i] {
			if i1[i][j]-i1[i][0] != i2[i][j]-i2[i][0] {
				return false
			}
		}
	}

	return true
}
