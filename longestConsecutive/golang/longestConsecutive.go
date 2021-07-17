package longestConsecutive

func longestConsecutive(nums []int) int {
	m := make(map[int]bool)
	for _, v := range nums {
		m[v] = true
	}

	mc := 0
	for k, _ := range m {

		delete(m, k)
		c := 1
		f := k + 1
		for {
			if _, ok := m[f]; ok {
				delete(m, f)
				c++
				f++
			} else {
				break
			}
		}
		b := k - 1
		for {
			if _, ok := m[b]; ok {
				delete(m, b)
				c++
				b--
			} else {
				break
			}
		}

		if c > mc {
			mc = c
		}
	}

	return mc
}
