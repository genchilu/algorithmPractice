package longestSubstringWithoutRepeatingCharacters

func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)
	b := 0

	max := 0
	for i := 0; i < len(s); i++ {
		if _, ok := m[s[i]]; ok {
			if m[s[i]] >= b {
				l := i - b
				if l > max {
					//fmt.Printf("111  %d, %d, %d, %d\n", max, l, i, b)
					max = l
				}
				b = m[s[i]] + 1
			}
			//fmt.Printf("333  %d, %d\n", i, b)
		}
		m[s[i]] = i
	}

	l := len(s) - b
	if l > max {
		//fmt.Printf("222 %d, %d, %d, %d\n", max, l, len(s), b)
		max = l
	}
	return max
}
