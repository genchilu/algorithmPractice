package numberOfGoodWaysToSplitAString

func numSplits(s string) int {
	m1 := make(map[byte]int)
	m2 := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		m1[s[i]]++
	}

	count := 0
	for i := 0; i < len(s); i++ {
		m2[s[i]]++
		m1[s[i]]--

		if m1[s[i]] == 0 {
			delete(m1, s[i])
		}

		if len(m1) == len(m2) {
			count++
		}
	}

	return count
}
