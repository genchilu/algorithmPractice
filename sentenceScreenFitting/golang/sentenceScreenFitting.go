package sentenceScreenFitting

func wordsTyping(sentence []string, rows int, cols int) int {
	idx, count, cur_col := 0, 0, cols

	for rows > 0 {
		if idx == 0 {
			count++
		}

		if cur_col >= len(sentence[idx]) {
			cur_col -= (len(sentence[idx]) + 1)
		} else {
			rows--
			if cols < len(sentence[idx]) {
				count = 1
				break
			}
			cur_col = cols - (len(sentence[idx]) + 1)

		}

		idx++
		if idx == len(sentence) {
			idx = 0
		}
		//fmt.Printf("rows: %d, cur_col: %d, idx:%d, count:%d\n", rows, cur_col, idx, count)
	}

	count -= 1
	return count
}
