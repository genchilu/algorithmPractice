package minimumSwapsToGroupAll1Together

func minSwaps(data []int) int {
	c := 0
	for _, v := range data {
		if v == 1 {
			c++
		}
	}

	sc := 0
	for i := 0; i < c; i++ {
		if data[i] == 1 {
			sc++
		}
	}

	max := sc
	for i := c; i < len(data); i++ {
		if data[i] == 1 {
			sc++
		}

		if data[i-c] == 1 {
			sc--
		}

		if sc > max {
			max = sc
		}
	}

	return (c - max)
}
