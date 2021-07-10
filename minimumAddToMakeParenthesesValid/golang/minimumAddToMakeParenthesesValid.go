package minimumAddToMakeParenthesesValid

func minAddToMakeValid(s string) int {

	c := 0
	b := 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(':
			b++
		case ')':
			if b == 0 {
				c++
			} else {
				b--
			}
		}
	}

	if b < 0 {
		b = -b
	}
	return b + c
}
