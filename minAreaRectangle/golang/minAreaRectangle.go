package minAreaRectangle

func minAreaRect(points [][]int) int {
	vm := make(map[int]map[int]bool)
	hm := make(map[int]map[int]bool)

	for _, p := range points {
		if _, ok := hm[p[0]]; !ok {
			hm[p[0]] = make(map[int]bool)
		}
		hm[p[0]][p[1]] = true

		if _, ok := vm[p[1]]; !ok {
			vm[p[1]] = make(map[int]bool)
		}
		vm[p[1]][p[0]] = true
	}

	min := 0
	for h1, vs := range hm {
		if len(vs) > 1 {
			for v1 := range vs {
				for v2 := range vs {
					if v1 != v2 {
						for h2 := range vm[v1] {
							if h2 != h1 && vm[v2][h2] {
								a := abs((h1 - h2) * (v1 - v2))
								if min == 0 {
									min = a
								} else if a < min {
									min = a
									//fmt.Printf("a: %d, h1: %d, h2: %d, v1: %d, v2: %d\n", a, h1, h2, v1, v2)
								}
							}
						}
					}
				}
			}
		}
	}
	return min
}

func abs(n int) int {
	if n > 0 {
		return n
	}

	return -n
}
