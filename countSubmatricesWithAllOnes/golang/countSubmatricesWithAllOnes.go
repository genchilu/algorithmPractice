package countSubmatricesWithAllOnes

func numSubmat(mat [][]int) int {

	h := len(mat)
	if h == 0 {
		return 0
	}

	w := len(mat[0])

	for _, v := range mat {
		ts := 0
		for i := w - 1; i >= 0; i-- {
			if v[i] == 1 {
				ts += v[i]
				v[i] = ts
			} else {
				ts = 0
			}
		}
	}
	c := 0
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if mat[i][j] != 0 {
				m := mat[i][j]
				c += m
				for k := 1; (k + i) < h; k++ {
					if mat[k+i][j] == 0 {
						break
					}

					if mat[k+i][j] < m {
						m = mat[k+i][j]
					}
					c += m
				}
			}
		}
	}
	return c
}
