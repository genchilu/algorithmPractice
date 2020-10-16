package shortestPathInBinaryMatrix

func shortestPathBinaryMatrix(grid [][]int) int {
	n := len(grid)

	if grid[0][0] == 0 {
		q := [][]int{[]int{0, 0, 1}}

		for len(q) > 0 {
			i := q[0][0]
			j := q[0][1]
			cs := q[0][2]

			if i == n-1 && j == n-1 {
				return cs
			}

			q = q[1:]

			cs++
			addToQueue(i+1, j, n, cs, grid, &q)
			addToQueue(i-1, j, n, cs, grid, &q)
			addToQueue(i, j+1, n, cs, grid, &q)
			addToQueue(i, j-1, n, cs, grid, &q)
			addToQueue(i+1, j+1, n, cs, grid, &q)
			addToQueue(i+1, j-1, n, cs, grid, &q)
			addToQueue(i-1, j+1, n, cs, grid, &q)
			addToQueue(i-1, j-1, n, cs, grid, &q)
		}

	}
	return -1
}

func addToQueue(i, j, n, cs int, grid [][]int, q *[][]int) {
	if i < n && j < n && i >= 0 && j >= 0 && grid[i][j] == 0 {
		*q = append(*q, []int{i, j, cs})
		grid[i][j] = 1
	}
}
