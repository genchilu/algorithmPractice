package letterCombinations

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	m := make(map[byte][]byte)
	m['2'] = []byte{'a', 'b', 'c'}
	m['3'] = []byte{'d', 'e', 'f'}
	m['4'] = []byte{'g', 'h', 'i'}
	m['5'] = []byte{'j', 'k', 'l'}
	m['6'] = []byte{'m', 'n', 'o'}
	m['7'] = []byte{'p', 'q', 'r', 's'}
	m['8'] = []byte{'t', 'u', 'v'}
	m['9'] = []byte{'w', 'x', 'y', 'z'}

	result := [][]byte{}
	for _, v := range m[digits[0]] {
		result = append(result, []byte{v})
	}

	for i := 1; i < len(digits); i++ {
		newResult := [][]byte{}

		for j, _ := range result {
			v := make([]byte, len(result[j]))
			copy(v, result[j])

			for _, vv := range m[digits[i]] {
				tmp := append(v, vv)
				newResult = append(newResult, tmp)
			}
		}
		result = newResult
	}

	r := []string{}
	for _, v := range result {
		r = append(r, string(v))
	}
	return r
}
