package validSquare

func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	m := make(map[int]int)

	m[cacu(p1, p2)]++
	m[cacu(p1, p3)]++
	m[cacu(p1, p4)]++
	m[cacu(p2, p3)]++
	m[cacu(p2, p4)]++
	m[cacu(p3, p4)]++

	if len(m) != 2 {
		return false
	}

	ks := []int{}
	for k, _ := range m {
		ks = append(ks, k)
	}
	if ks[0] > ks[1] {
		return m[ks[0]] == 2 && m[ks[1]] == 4
	} else {
		return m[ks[1]] == 2 && m[ks[0]] == 4
	}
}

func cacu(p1, p2 []int) int {
	d1 := p1[0] - p2[0]
	d2 := p1[1] - p2[1]

	return (d1 * d1) + (d2 * d2)
}
