package maximumNumberOfPointsWithCost

func maxPoints(points [][]int) int64 {

	dp := make([][]int, len(points))
	for i := range points {
		dp[i] = make([]int, len(points[0]))
		if i == 0 {
			for j := range points[0] {
				dp[i][j] = points[i][j]
			}
		}
	}

	r := 0
	for i := 1; i < len(points); i++ {
		for j := range points[i] {
			max := 0
			for k := range points[i-1] {
				v := dp[i-1][k] - abs(k-j)
				if v > max {
					max = v
				}
			}
			dp[i][j] = max + points[i][j]
			if dp[i][j] > r {
				r = dp[i][j]
			}
		}
	}

	max := 0
	for i := range dp[len(dp)-1] {
		if dp[len(points)-1][i] > max {
			max = dp[len(dp)-1][i]
		}
	}

	return int64(max)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}
