package guessNumberHigherOrLowerII

import "math"

func getMoneyAmount(n int) int {

	dp := make([][]int, n+1)
	for i, _ := range dp {
		dp[i] = make([]int, n+1)
	}

	m := math.MaxInt64

	return cacul(1, n, m, &dp)
}

func cacul(i, j, m int, dp *[][]int) int {
	if i >= j {
		return 0
	}

	if (*dp)[i][j] != 0 {
		return (*dp)[i][j]
	}

	res := m

	for k := i; k <= j; k++ {

		tmp := k + max(cacul(i, k-1, m, dp), cacul(k+1, j, m, dp))
		res = min(res, tmp)

	}

	(*dp)[i][j] = res

	//fmt.Printf("i: %d, j: %d, db: %d\n", i, j, (*dp)[i][j])
	return res
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
