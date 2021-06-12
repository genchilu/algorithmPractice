package integerToRoman

func intToRoman(num int) string {
	result := []byte{}

	for num >= 1000 {
		result = append(result, 'M')
		num -= 1000
	}

	if num >= 900 {
		result = append(result, []byte{'C', 'M'}...)
		num -= 900
	}

	if num >= 500 {
		result = append(result, 'D')
		num -= 500
	}

	if num >= 400 {
		result = append(result, []byte{'C', 'D'}...)
		num -= 400
	}

	for num >= 100 {
		result = append(result, 'C')
		num -= 100
	}

	if num >= 90 {
		result = append(result, []byte{'X', 'C'}...)
		num -= 90
	}

	if num >= 50 {
		result = append(result, 'L')
		num -= 50
	}

	if num >= 40 {
		result = append(result, []byte{'X', 'L'}...)
		num -= 40
	}

	for num >= 10 {
		result = append(result, 'X')
		num -= 10
	}

	if num >= 9 {
		result = append(result, []byte{'I', 'X'}...)
		num -= 9
	}

	if num >= 5 {
		result = append(result, 'V')
		num -= 5
	}

	if num >= 4 {
		result = append(result, []byte{'I', 'V'}...)
		num -= 4
	}

	for num >= 1 {
		result = append(result, 'I')
		num -= 1
	}

	return string(result)
}
