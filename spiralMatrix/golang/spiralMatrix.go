package spiralMatrix

func spiralOrder(matrix [][]int) []int {
	move := [][]int{[]int{0, 1}, []int{1, 0}, []int{0, -1}, []int{-1, 0}}
	vistited := make([][]bool, len(matrix))
	for i, _ := range vistited {
		vistited[i] = make([]bool, len(matrix[i]))
	}

	curMove := 0
	i, j := 0, 0
	result := []int{}
	for true {
		result = append(result, matrix[i][j])
		vistited[i][j] = true

		nexti, nextj := i+move[curMove][0], j+move[curMove][1]
		if isInValidStep(nexti, nextj, len(matrix), len(matrix[0])) || vistited[nexti][nextj] {
			curMove = (curMove + 1) % 4
			nexti, nextj = i+move[curMove][0], j+move[curMove][1]
			if isInValidStep(nexti, nextj, len(matrix), len(matrix[0])) || vistited[nexti][nextj] {
				break
			}
		}

		i, j = nexti, nextj

	}

	return result
}

func isInValidStep(i, j, m, n int) bool {
	return i < 0 || j < 0 || i >= m || j >= n
}
