package surroundedRegions

func solve(board [][]byte) {

	for i, v := range board {
		for j, _ := range v {
			shouldFlip := isSurrounded(board, i, j)
			flip(board, i, j, shouldFlip)
		}
	}
}

func isSurrounded(board [][]byte, i, j int) bool {
	m, n := len(board), 0
	if m == 0 {
		n = 0
	} else {
		n = len(board[0])
	}

	if i < 0 || j < 0 || i >= m || j >= n {
		return false
	}

	if board[i][j] == byte('X') {
		return true
	}

	if board[i][j] == byte('O') {
		board[i][j] = byte('A')
		return isSurrounded(board, i+1, j) && isSurrounded(board, i-1, j) && isSurrounded(board, i, j+1) && isSurrounded(board, i, j-1)
	}

	return true
}

func flip(board [][]byte, i, j int, shouldFlip bool) {
	m, n := len(board), 0
	if m == 0 {
		n = 0
	} else {
		n = len(board[0])
	}

	if i >= 0 && j >= 0 && i < m && j < n && board[i][j] == byte('A') {
		if shouldFlip {
			board[i][j] = byte('X')
		} else {
			board[i][j] = byte('O')
		}
		flip(board, i+1, j, shouldFlip)
		flip(board, i-1, j, shouldFlip)
		flip(board, i, j+1, shouldFlip)
		flip(board, i, j-1, shouldFlip)
	}
}
