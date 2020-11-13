package decodeString

func decodeString(s string) string {
	stack1 := []string{}
	stack2 := []int{}

	for len(s) > 0 {
		if d, ss := getDigitPrefix(s); d > 0 {
			stack2 = append(stack2, d)
			s = ss
			continue
		}

		if prefix, ss := getCharPrefix(s); len(prefix) > 0 {
			stack1 = append(stack1, prefix)
			s = ss
			continue
		}

		if s[0] == 91 {
			stack1 = append(stack1, s[0:1])
			s = s[1:]
			continue
		}

		if s[0] == 93 {
			s = s[1:]
			dd := stack2[len(stack2)-1]
			stack2 = stack2[0 : len(stack2)-1]

			ss := ""
			for stack1[len(stack1)-1] != "[" {
				ss = stack1[len(stack1)-1] + ss
				stack1 = stack1[0 : len(stack1)-1]
			}
			stack1 = stack1[0 : len(stack1)-1]

			t := ""
			for i := 0; i < dd; i++ {
				t += ss
			}

			stack1 = append(stack1, t)
			continue
		}
	}

	result := ""
	for _, v := range stack1 {
		result += v
	}

	return result
}

func getDigitPrefix(s string) (int, string) {
	d := 0
	for _, v := range s {
		if v >= 48 && v <= 57 {
			d = d*10 + int(v) - 48
			s = s[1:]
		} else {
			break
		}
	}

	return d, s
}

func getCharPrefix(s string) (string, string) {
	d := 0
	for _, v := range s {
		if (v < 48 || v > 57) && v != 91 && v != 93 {
			d++
		} else {
			break
		}
	}

	return s[0:d], s[d:]
}
