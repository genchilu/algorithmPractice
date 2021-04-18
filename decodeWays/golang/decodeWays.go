package decodeWays

import (
	"strconv"
)

func numDecodings(s string) int {
	m := make(map[int]int)
	return dive(s, 0, m)
}

func dive(s string, idx int, m map[int]int) int {

	if v, ok := m[idx]; ok {
		return v
	}

	if idx > len(s)-1 {
		return 1
	}

	if s[idx] == '0' {
		return 0
	}

	if idx == len(s)-1 {
		return 1
	}

	c := dive(s, idx+1, m)

	tmp, _ := strconv.Atoi(s[idx : idx+2])
	if tmp <= 26 {
		c += dive(s, idx+2, m)
	}

	m[idx] = c

	return c
}
