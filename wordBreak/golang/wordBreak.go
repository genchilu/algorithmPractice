package wordBreak

func wordBreak(s string, wordDict []string) bool {

	m := make(map[string]bool)
	a := make(map[string]bool)
	for _, word := range wordDict {
		m[word] = true
		for i := 0; i < len(word); i++ {
		}
	}

	return search(s, m, a)
}

func search(s string, m map[string]bool, a map[string]bool) bool {
	if len(s) == 0 {
		return true
	}

	if _, ok := a[s]; ok {
		return a[s]
	}

	for i := len(s); i > 0; i-- {
		if _, ok := m[s[0:i]]; ok {
			if search(s[i:], m, a) {
				a[s] = true
				return true
			}
		}
	}

	a[s] = false
	return false
}
