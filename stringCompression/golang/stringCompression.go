package stringCompression

import (
	"strconv"
)

func compress(chars []byte) int {
	cur := chars[0]
	c := 1
	s := ""
	for i := 1; i < len(chars); i++ {
		if chars[i] == cur {
			c++
		} else {
			s += string(cur)
			if c > 1 {
				s += strconv.Itoa(c)
			}

			cur = chars[i]
			c = 1
		}
	}

	s += string(cur)
	if c > 1 {
		s += strconv.Itoa(c)
	}

	for i := 0; i < len(s); i++ {
		chars[i] = s[i]
	}
	return len(s)
}
