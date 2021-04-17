package wordSearch

func exist(board [][]byte, word string) bool {

	for i, _ := range board {
		for j, _ := range board[i] {
			if dfs(board, word, i, j, 0) {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, word string, i, j, idx int) bool {
	if idx == len(word) {
		return true
	}

	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
		return false
	}

	if board[i][j] != word[idx] {
		return false
	}

	tmp := board[i][j]
	board[i][j] = '-'

	idx++

	result := dfs(board, word, i+1, j, idx) || dfs(board, word, i-1, j, idx) || dfs(board, word, i, j+1, idx) || dfs(board, word, i, j-1, idx)

	board[i][j] = tmp

	return result
}
