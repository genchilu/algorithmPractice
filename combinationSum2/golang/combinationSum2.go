package combinationSum2

func combinationSum2(candidates []int, target int) [][]int {
	m := make(map[int]int)
	for _, v := range candidates {
		m[v]++
	}

	c := []int{}
	for k, _ := range m {
		c = append(c, k)
	}

	return combinationSum(c, target, m)
}

func combinationSum(c []int, target int, m map[int]int) [][]int {

	result := [][]int{}

	for i, v := range c {
		t := target
		pre := []int{}

		m1 := myclone(m)

		for m1[v] > 0 && t >= v {
			pre = append(pre, v)
			if t == v {
				result = append(result, pre)
			} else if t > v {
				j := i + 1
				if j < len(c) {
					tmp := combinationSum(c[j:], t-v, m1)
					if len(tmp) > 0 {
						for i := range tmp {
							tmp[i] = append(tmp[i], pre...)
						}
					}

					result = append(result, tmp...)
				}

			}
			t -= v
			m1[v]--
		}
	}
	return result
}

func myclone(m map[int]int) map[int]int {
	m1 := make(map[int]int)

	for k, v := range m {
		m1[k] = v
	}

	return m1
}
