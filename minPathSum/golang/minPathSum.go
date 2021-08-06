package minPathSum

import (
	"math"
)

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	back(&grid, &dp, m-1, n-1)
	//fmt.Printf("%v\n", dp)
	return dp[m-1][n-1]
}

func back(grid, dp *[][]int, i, j int) int {
	//fmt.Printf("%d, %d\n", i, j)
	if i < 0 || j < 0 {
		return math.MaxInt64
	}

	if i == 0 && j == 0 {
		(*dp)[i][j] = (*grid)[i][j]
	} else if (*dp)[i][j] == 0 {
		(*dp)[i][j] = (*grid)[i][j] + min(back(grid, dp, i-1, j), back(grid, dp, i, j-1))
	}

	return (*dp)[i][j]

}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
