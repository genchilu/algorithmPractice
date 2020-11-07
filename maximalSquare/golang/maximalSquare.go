package maximalSquare

func maximalSquare(matrix [][]byte) int {
	var h, w int
	h = len(matrix)
	cloneMatrix := make([][]int, h)
	w = 0
	if h > 0 {
		w = len(matrix[0])
	}

	for i := 0; i < h; i++ {
		c := 0
		cloneMatrix[i] = make([]int, w)
		for j := w - 1; j >= 0; j-- {
			if matrix[i][j] == '1' {
				c++
			} else {
				c = 0
			}
			cloneMatrix[i][j] = c
		}
	}

	max := 0
	for i, r := range cloneMatrix {
		for j, c := range r {
			if c > max && h-i > max {
				min := c
				if h-i < min {
					min = h - i
				}

				for k := 0; k < min; k++ {
					if k <= min {
						if cloneMatrix[k+i][j] < min {
							min = cloneMatrix[k+i][j]
						}
					}
				}

				if min > max {
					max = min
				}
			}
		}
	}
	return max * max
}
