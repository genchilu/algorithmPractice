package snakesAndLadders

func snakesAndLadders(board [][]int) int {
	n := len(board)
	maxIdx := n * n
	c := 0
	q := []int{1}

	vistited := make([]bool, maxIdx+1)
	for i := 1; i <= maxIdx; i++ {
		vistited[i] = false
	}

	for len(q) > 0 {
		newq := []int{}
		for _, v := range q {
			if vistited[v] {
				continue
			}

			vistited[v] = true
			if v == maxIdx {
				return c
			}
			for i := 1; i <= 6; i++ {

				step := v + i

				if step > maxIdx {
					break
				}

				ii, jj := caculPostion(step, n)
				//fmt.Printf("ii %d, jj:%d, step: %d, n: %d\n", ii, jj, step, n)
				if board[ii][jj] != -1 {
					newq = append(newq, board[ii][jj])
				} else {
					newq = append(newq, step)
				}
			}
		}

		c++
		q = newq
	}

	return -1
}

func caculPostion(idx, n int) (i, j int) {
	reverIidx := (idx - 1) / n
	i = n - reverIidx - 1
	if reverIidx&1 == 1 {
		j = n - (idx-1)%n - 1
	} else {
		j = (idx - 1) % n
	}

	return
}
