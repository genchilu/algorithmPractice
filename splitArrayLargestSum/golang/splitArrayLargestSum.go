package splitArrayLargestSum

func splitArray(nums []int, m int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	dp := make([][]int, len(nums))
	for i := range dp {
		dp[i] = make([]int, m+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	return split(nums, 0, m, sum, &dp)
}

func split(nums []int, idx, m, sum int, dp *[][]int) int {
	if m == 1 {
		return sum
	}

	if (*dp)[idx][m] != -1 {
		return (*dp)[idx][m]
	}
	if m == 2 {
		lsum := nums[idx]
		rsum := sum - lsum

		min := findMax(lsum, rsum)

		for i := idx + 1; i < len(nums)-1; i++ {
			lsum += nums[i]
			rsum -= nums[i]
			tsum := findMax(lsum, rsum)
			if tsum < min {
				min = tsum
			}
		}

		return min
	}

	fsum := nums[idx]
	rmax := split(nums, idx+1, m-1, sum-fsum, dp)
	min := findMax(fsum, rmax)

	for i := idx + 1; i <= len(nums)-m; i++ {
		fsum += nums[i]
		rmax = split(nums, i+1, m-1, sum-fsum, dp)
		tsum := findMax(fsum, rmax)

		if tsum < min {
			min = tsum
		}
	}

	(*dp)[idx][m] = min

	return (*dp)[idx][m]
}

func findMax(a, b int) int {
	if a > b {
		return a
	}

	return b
}
