package uniquePaths

func uniquePaths(m int, n int) int {
	total := (m - 1) + (n - 1)
	result := total

	min := m
	if n < m {
		min = n
	}

	if min-1 == 0 {
		return 1
	}

	for i := 1; i < min-1; i++ {
		result *= (total - i)
		result /= (i + 1)

	}

	return result
}
