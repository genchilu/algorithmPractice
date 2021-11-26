package zigzagConversion

func convert(s string, numRows int) string {

	if numRows == 1 {
		return s
	}
	result := make([][]byte, numRows)

	i, j := 0, 0
	back := false

	for idx := 0; idx < len(s); idx++ {
		if j == len(result[0]) {
			appendLeft(result)
		}

		result[i][j] = s[idx]

		if back {
			i -= 1
			j += 1

			if i == 0 {
				i = 0
				back = false
			} else if i < 0 {
				i = 1
				j--
				back = true
			}
		} else {
			i++
			if i == numRows {
				i -= 2
				j += 1
				back = true
			}
		}

	}

	bs := []byte{}

	for _, r := range result {
		for _, c := range r {
			if c != '-' {
				bs = append(bs, c)
			}
		}
	}

	return string(bs)
}

func appendLeft(result [][]byte) [][]byte {
	for i := range result {
		result[i] = append(result[i], '-')
	}

	return result
}
