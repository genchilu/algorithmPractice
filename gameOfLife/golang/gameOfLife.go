package gameOfLife

/*
0 dead
1 live
2 dead to live
3 live to dead
*/
func gameOfLife(board [][]int) {

	for i, _ := range board {
		for j, _ := range board[i] {
			c := countLiveNeightbor(i, j, board)
			if board[i][j] == 1 {
				if c <= 1 {
					board[i][j] = 3
				} else if c > 3 {
					board[i][j] = 3
				}
			} else if board[i][j] == 0 {
				if c == 3 {
					board[i][j] = 2
				}
			}
		}
	}

	for i, _ := range board {
		for j, _ := range board[i] {
			if board[i][j] == 2 {
				board[i][j] = 1
			} else if board[i][j] == 3 {
				board[i][j] = 0
			}
		}
	}

}

func countLiveNeightbor(i, j int, board [][]int) int {
	c := 0

	if i > 0 && j > 0 {
		if board[i-1][j-1] == 1 || board[i-1][j-1] == 3 {
			c++
		}
	}

	if i > 0 {
		if board[i-1][j] == 1 || board[i-1][j] == 3 {
			c++
		}
	}

	if i > 0 && j < len(board[i-1])-1 {
		if board[i-1][j+1] == 1 || board[i-1][j+1] == 3 {
			c++
		}
	}

	if j > 0 {
		if board[i][j-1] == 1 || board[i][j-1] == 3 {
			c++
		}
	}

	if j < len(board[i])-1 {
		if board[i][j+1] == 1 || board[i][j+1] == 3 {
			c++
		}
	}

	if i < len(board)-1 && j > 0 {
		if board[i+1][j-1] == 1 || board[i+1][j-1] == 3 {
			c++
		}
	}

	if i < len(board)-1 {
		if board[i+1][j] == 1 || board[i+1][j] == 3 {
			c++
		}
	}

	if i < len(board)-1 && j < len(board[i+1])-1 {
		if board[i+1][j+1] == 1 || board[i+1][j+1] == 3 {
			c++
		}
	}

	return c
}
