package setMatrixZeroes

func setZeroes(matrix [][]int) {

	m := len(matrix)
	n := len(matrix[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {

				for k := 0; k < m; k++ {
					if matrix[k][j] != 0 {
						matrix[k][j] = 9223372036854775807
					}
				}

				for k := 0; k < n; k++ {
					if matrix[i][k] != 0 {
						matrix[i][k] = 9223372036854775807
					}
				}
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 9223372036854775807 {
				matrix[i][j] = 0
			}
		}
	}
}
