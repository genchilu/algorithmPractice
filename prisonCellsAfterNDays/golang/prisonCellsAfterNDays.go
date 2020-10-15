package prisonCellsAfterNDays

func prisonAfterNDays(cells []int, N int) []int {

	m := make(map[int]int)
	f := false

	for N > 0 {

		if !f {
			b := tobits(cells)
			if v, ok := m[b]; ok {
				d := v - N
				N %= d
				f = true
			} else {
				m[b] = N
			}
		}
		if N > 0 {
			N--
			cells = next(cells)
		}
	}

	return cells
}

func tobits(in []int) int {
	result := 0
	for _, v := range in {
		result = result << 1
		result = result | v
	}

	return result
}

func next(in []int) []int {
	res := make([]int, 8)
	res[0] = 0
	res[7] = 0
	for i := 1; i < 7; i++ {
		res[i] = 0
		if in[i-1] == in[i+1] {
			res[i] = 1
		}
	}

	return res
}

func toslice(in int) []int {
	res := []int{}
	for in > 0 {
		res = append([]int{in & 1}, res...)
		in = in >> 1
	}

	return res
}
