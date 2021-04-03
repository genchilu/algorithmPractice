package generateParentheses

func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{""}
	}

	result := []string{}
	for c := 0; c < n; c++ {
		for _, left := range generateParenthesis(c) {
			for _, right := range generateParenthesis(n - c - 1) {
				result = append(result, "("+left+")"+right)
			}

		}
	}

	return result
}
