package optimalAccountBalancing

import "math"

func minTransfers(transactions [][]int) int {

	debts := make([]int, 21)

	for _, t := range transactions {
		debts[t[0]] -= t[2]
		debts[t[1]] += t[2]
	}

	return dfs(0, &debts)
}

func dfs(n int, debts *[]int) int {
	for n < len(*debts) && (*debts)[n] == 0 {
		n++
	}

	mint := math.MaxInt64
	pre := 0
	for i := n + 1; i < len(*debts); i++ {
		if (*debts)[i] != pre && (*debts)[i]*(*debts)[n] < 0 {

			(*debts)[i] += (*debts)[n]
			mint = min(mint, 1+dfs(n+1, debts))
			(*debts)[i] -= (*debts)[n]
			pre = (*debts)[i]
		}
	}

	if mint == math.MaxInt64 {
		return 0
	}
	return mint

}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
