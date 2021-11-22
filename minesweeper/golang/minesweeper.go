package minesweeper

func updateBoard(board [][]byte, click []int) [][]byte {

	q := [][]int{click}

	for len(q) > 0 {
		c := q[0]
		q = q[1:]
		i, j := c[0], c[1]

		switch board[i][j] {
		case 'M':
			board[i][j] = 'X'
		case 'E':
			count := countAdjMine(board, i, j)
			if count > 0 {
				board[i][j] = byte(count + 48)
			} else {
				board[i][j] = 'B'
				if i > 0 && j > 0 && board[i-1][j-1] == 'E' {
					q = append(q, []int{i - 1, j - 1})
				}

				if i > 0 && board[i-1][j] == 'E' {
					q = append(q, []int{i - 1, j})
				}

				if i > 0 && j < len(board[0])-1 && board[i-1][j+1] == 'E' {
					q = append(q, []int{i - 1, j + 1})
				}

				if j > 0 && board[i][j-1] == 'E' {
					q = append(q, []int{i, j - 1})
				}

				if j < len(board[0])-1 && board[i][j+1] == 'E' {
					q = append(q, []int{i, j + 1})
				}

				if i < len(board)-1 && j > 0 && board[i+1][j-1] == 'E' {
					q = append(q, []int{i + 1, j - 1})
				}

				if i < len(board)-1 && board[i+1][j] == 'E' {
					q = append(q, []int{i + 1, j})
				}

				if i < len(board)-1 && j < len(board[0])-1 && board[i+1][j+1] == 'E' {
					q = append(q, []int{i + 1, j + 1})
				}
			}
		}
	}

	return board
}

func countAdjMine(board [][]byte, i, j int) int {
	count := 0
	if i > 0 && j > 0 && board[i-1][j-1] == 'M' {
		count++
	}

	if i > 0 && board[i-1][j] == 'M' {
		count++
	}

	if i > 0 && j < len(board[0])-1 && board[i-1][j+1] == 'M' {
		count++
	}

	if j > 0 && board[i][j-1] == 'M' {
		count++
	}

	if j < len(board[0])-1 && board[i][j+1] == 'M' {
		count++
	}

	if i < len(board)-1 && j > 0 && board[i+1][j-1] == 'M' {
		count++
	}

	if i < len(board)-1 && board[i+1][j] == 'M' {
		count++
	}

	if i < len(board)-1 && j < len(board[0])-1 && board[i+1][j+1] == 'M' {
		count++
	}

	return count
}
