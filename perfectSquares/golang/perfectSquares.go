package perfectSquares

import "math"

func numSquares(n int) int {
	max_idx := int(math.Sqrt(float64(n))) + 1
	sq_nums := make([]int, max_idx)
	for i := range sq_nums {
		sq_nums[i] = i * i
	}

	dp := make([]int, n+1)

	dp[0] = 0
	for i := 1; i <= n; i++ {
		for j := 1; j < max_idx; j++ {
			if i >= sq_nums[j] {
				if dp[i] == 0 {
					dp[i] = dp[i-sq_nums[j]] + 1
				} else {
					dp[i] = min(dp[i], dp[i-sq_nums[j]]+1)
				}
			}
		}
	}

	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
