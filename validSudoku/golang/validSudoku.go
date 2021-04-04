package validSudoku

func isValidSudoku(board [][]byte) bool {
	rm := make([]map[byte]bool, 9)
	cm := make([]map[byte]bool, 9)
	sm := make([]map[byte]bool, 9)

	for i := 0; i < 9; i++ {
		rm[i] = make(map[byte]bool)
		cm[i] = make(map[byte]bool)
		sm[i] = make(map[byte]bool)
	}

	for i, _ := range board {
		for j, _ := range board[i] {
			if board[i][j] != '.' {
				if _, ok := rm[i][board[i][j]]; ok {
					return false
				} else {
					rm[i][board[i][j]] = true
				}

				if _, ok := cm[j][board[i][j]]; ok {
					return false
				} else {
					cm[j][board[i][j]] = true
				}

				subboxIdx := (i/3)*3 + j/3
				if _, ok := sm[subboxIdx][board[i][j]]; ok {
					return false
				} else {
					sm[subboxIdx][board[i][j]] = true
				}
			}
		}
	}
	return true
}
