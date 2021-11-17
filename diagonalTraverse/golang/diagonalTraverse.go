package diagonalTraverse

func findDiagonalOrder(mat [][]int) []int {

	ru := true

	m := len(mat)
	n := len(mat[0])

	result := []int{}
	i, j := 0, 0
	for {
		//fmt.Printf("%d, %d\n", i, j)
		result = append(result, mat[i][j])
		if ru {
			i -= 1
			j += 1
			if j >= n {
				i += 2
				j -= 1
				ru = false
			} else if i < 0 {
				i += 1
				ru = false
			}

			if i < 0 || i >= m || j < 0 || j >= n {
				break
			}
		} else {
			j -= 1
			i += 1

			if i >= m {
				j += 2
				i -= 1
				ru = true
			} else if j < 0 {
				j += 1
				ru = true
			}

			if i < 0 || i >= m || j < 0 || j >= n {
				break
			}
		}
	}

	return result
}
