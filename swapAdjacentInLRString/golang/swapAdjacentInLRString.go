package swapAdjacentInLRString

func canTransform(start string, end string) bool {
	c := 0
	for i := 0; i < len(start); i++ {
		if start[i] == 'X' {
			c++
		}
		if end[i] == 'X' {
			c--
		}
	}
	if c != 0 {
		return false
	}

	i, j := 0, 0
	for i < len(start) || j < len(end) {

		for i < len(start) && start[i] == 'X' {
			i++
		}
		for j < len(end) && end[j] == 'X' {
			j++
		}

		if i == len(start) || j == len(end) {
			return i == j
		}

		if start[i] != end[j] {
			//fmt.Printf("111 %d, %d\n", i, j)
			return false
		}

		if start[i] == 'R' && i > j {
			//fmt.Printf("222\n")
			return false
		}

		if start[i] == 'L' && i < j {
			//fmt.Printf("333\n")
			return false
		}

		i++
		j++
	}

	return true
}
