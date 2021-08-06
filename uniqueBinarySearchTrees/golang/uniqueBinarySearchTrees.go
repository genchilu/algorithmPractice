package uniqueBinarySearchTrees

func numTrees(n int) int {
	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	dp := make([]int, n+1)
	dp[0], dp[1], dp[2] = 1, 1, 2

	r := num(n, &dp)
	return r
}

func num(n int, dp *[]int) int {
	if n > 2 && (*dp)[n] == 0 {
		for i := 0; i < n; i++ {
			(*dp)[n] += num(i, dp) * num(n-i-1, dp)
		}
	}

	return (*dp)[n]
}
