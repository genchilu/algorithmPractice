package spiralMatrixII

func generateMatrix(n int) [][]int {
	result := make([][]int, n)
	for i, _ := range result {
		result[i] = make([]int, n)
	}
	moves := [][]int{
		[]int{0, 1},
		[]int{1, 0},
		[]int{0, -1},
		[]int{-1, 0},
	}

	c := 1
	i := 0
	j := 0
	moveidx := 0
	for c <= n*n {
		result[i][j] = c
		c++
		i, j, moveidx = nextMove(moves, result, i, j, n, moveidx)
	}

	return result
}

func nextMove(moves, result [][]int, i, j, n, moveidx int) (ni, nj, idx int) {
	ni = i + moves[moveidx][0]
	nj = j + moves[moveidx][1]
	if ni < 0 || ni >= n || nj < 0 || nj >= n || result[ni][nj] != 0 {
		moveidx += 1
		if moveidx == 4 {
			moveidx = 0
		}

		ni = i + moves[moveidx][0]
		nj = j + moves[moveidx][1]
	}

	return ni, nj, moveidx
}
