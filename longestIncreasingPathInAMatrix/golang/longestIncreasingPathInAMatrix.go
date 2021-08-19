package longestIncreasingPathInAMatrix

func longestIncreasingPath(matrix [][]int) int {
	m := len(matrix)
	n := len(matrix[0])

	l := make([][]int, m)
	for i := range l {
		l[i] = make([]int, n)
	}

	max := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			tmp := dfs(i, j, m, n, &l, matrix)
			if tmp > max {
				max = tmp
			}
		}
	}
	//fmt.Printf("%v\n", l)
	return max
}

func dfs(i, j, m, n int, l *[][]int, matrix [][]int) int {
	if (*l)[i][j] != 0 {
		return (*l)[i][j]
	}

	max := 1
	if i != 0 && matrix[i][j] < matrix[i-1][j] {
		up := dfs(i-1, j, m, n, l, matrix) + 1
		if up > max {
			max = up
		}
	}

	if i != m-1 && matrix[i][j] < matrix[i+1][j] {
		down := dfs(i+1, j, m, n, l, matrix) + 1
		if down > max {
			max = down
		}
	}

	if j != 0 && matrix[i][j] < matrix[i][j-1] {
		left := dfs(i, j-1, m, n, l, matrix) + 1
		if left > max {
			max = left
		}
	}

	if j != n-1 && matrix[i][j] < matrix[i][j+1] {
		right := dfs(i, j+1, m, n, l, matrix) + 1
		if right > max {
			max = right
		}
	}

	(*l)[i][j] = max
	return (*l)[i][j]

}
