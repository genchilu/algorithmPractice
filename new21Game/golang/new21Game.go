package new21Game

func new21Game(n int, k int, maxPts int) float64 {
	dp := make([]float64, k+maxPts)
	suffixSum := make([]float64, k+maxPts+1)
	m := float64(maxPts)

	for i := len(dp) - 1; i >= 0; i-- {
		if i >= k {
			if i > n {
				dp[i] = 0
			} else {
				dp[i] = 1
			}
		} else {
			dp[i] = (suffixSum[i+1] - suffixSum[i+maxPts+1]) / m
		}
		suffixSum[i] = dp[i] + suffixSum[i+1]
	}

	return dp[0]
}
