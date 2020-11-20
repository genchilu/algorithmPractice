package countSquareSubmatricesWithAllOnes

func countSquares(matrix [][]int) int {
	h := len(matrix)
	w := 0
	if h > 0 {
		w = len(matrix[0])
	}

	for i := 0; i < h; i++ {
		c := 0
		for j := w - 1; j >= 0; j-- {
			if matrix[i][j] == 1 {
				c++
				matrix[i][j] = c
			} else {
				c = 0
			}
		}
	}

	sum := 0
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if matrix[i][j] > 0 {
				min := matrix[i][j]
				for k := 0; k < matrix[i][j] && (k+i) < h; k++ {
					if matrix[k+i][j] < min {
						min = matrix[k+i][j]
					}

					if min > k {
						sum++
					}
				}
			}
		}
	}
	return sum
}
