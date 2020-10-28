package longestPalindromicSubstring

func longestPalindrome(s string) string {
	result := ""

	for i := 0; i < len(s); i++ {
		l := i - 1
		r := i + 1
		for r < len(s) {
			if s[r] != s[i] {
				break
			}
			r++
		}
		i = r - 1

		for l >= 0 && r < len(s) {
			if s[l] != s[r] {
				break
			}
			l--
			r++
		}

		t := s[l+1 : r]
		if len(t) > len(result) {
			result = t
		}
	}

	return result
}
