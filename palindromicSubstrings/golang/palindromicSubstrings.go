package palindromicSubstrings

func countSubstrings(s string) int {
	count := 0
	for i := 0; i < len(s); i++ {
		l, r := i-1, i+1
		for r < len(s) {
			if s[r] != s[i] {
				break
			}
			r++
		}

		count += (r - l) * (r - l - 1) / 2

		i = r - 1

		for l >= 0 && r < len(s) {
			if s[l] == s[r] {
				count++
				l--
				r++
			} else {
				break
			}
		}

	}

	return count
}
