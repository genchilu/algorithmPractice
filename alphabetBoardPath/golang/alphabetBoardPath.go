package alphabetBoardPath

func alphabetBoardPath(target string) string {
	if len(target) == 0 {
		return ""
	}

	m := make(map[byte][]int)
	c := 0
	s := []int{0, 0}

	result := ""
	for c < len(target) {
		e := getPosition(m, target[c])

		if e[0] < s[0] {
			for i := 0; i < s[0]-e[0]; i++ {
				result += "U"
			}
			result += moveHorizon(s[1], e[1])
			result += "!"
		} else if s[0] < e[0] {
			result += moveHorizon(s[1], e[1])
			for i := 0; i < e[0]-s[0]; i++ {
				result += "D"
			}
			result += "!"
		} else {
			result += moveHorizon(s[1], e[1])
			result += "!"
		}
		s = e
		c++
	}
	return result
}

func cacluatePosition(b byte) []int {
	idx := (int)(b - byte('a'))
	r := idx / 5
	c := idx % 5

	return []int{r, c}
}

func getPosition(m map[byte][]int, b byte) []int {
	if v, ok := m[b]; ok {
		return v
	} else {
		p := cacluatePosition(b)
		m[b] = p
		return p
	}
}

func moveHorizon(s, e int) string {
	result := ""
	if e < s {
		for i := 0; i < s-e; i++ {
			result += "L"
		}
	} else if e > s {
		for i := 0; i < e-s; i++ {
			result += "R"
		}
	}

	return result
}
