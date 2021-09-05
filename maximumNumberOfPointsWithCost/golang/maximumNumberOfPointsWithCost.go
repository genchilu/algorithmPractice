package maximumNumberOfPointsWithCost

func maxPoints(points [][]int) int64 {

	dp := make([]int, len(points[0]))
	for i := range points[0] {
		dp[i] = points[0][i]
	}

	//fmt.Printf("111 %v\n", dp)

	for i := 1; i < len(points); i++ {
		tmp := make([]int, len(dp))
		copy(tmp, dp)
		for j := range points[i] {
			max := 0
			//fmt.Printf("222 %v\n", tmp)
			for k := range tmp {
				v := tmp[k] - abs(k-j)
				if v > max {
					max = v
				}
			}
			dp[j] = max + points[i][j]
		}
		//fmt.Printf("333 %v\n", dp)
	}

	max := 0
	for i := range dp {
		if dp[i] > max {
			max = dp[i]
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
