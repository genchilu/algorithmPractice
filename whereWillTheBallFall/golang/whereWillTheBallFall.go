package whereWillTheBallFall

func findBall(grid [][]int) []int {
	m, n := len(grid), len(grid[0])
	ans := make([]int, n)

	for i := 0; i < n; i++ {
		p := i
		for j := 0; j < m; j++ {
			if grid[j][p] == 1 {
				if p == n-1 || grid[j][p+1] == -1 {
					p = -1
					break
				} else {
					p++
				}
			} else {
				if p == 0 || grid[j][p-1] == 1 {
					p = -1
					break
				} else {
					p--
				}
			}
		}
		ans[i] = p
	}

	return ans
}
