package removeComments

func removeComments(source []string) []string {
	commentState := false
	result := []string{}
	t := ""
	for _, s := range source {
		b := 0
		for i := 0; i < len(s); i++ {
			if commentState {
				if s[i] == '*' {
					j := i + 1
					if j < len(s) {
						if s[j] == '/' {
							commentState = false
							b = j + 1
							i = j
						}
					}
				}
			} else {
				if s[i] == '/' {
					j := i + 1
					if j < len(s) {
						if s[j] == '*' {
							commentState = true
							t += s[b:i]

							i = j
						} else if s[j] == '/' {
							t += s[b:i]
							if len(t) > 0 {
								result = append(result, t)
								t = ""
							}
							b = len(s)
							break
						}
					}
				}
			}
		}

		if !commentState {
			if b < len(s) {
				t += s[b:]
			}
			if len(t) > 0 {
				result = append(result, t)
				t = ""
			}
		}
	}
	return result
}
