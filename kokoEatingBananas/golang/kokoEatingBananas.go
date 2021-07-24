package kokoEatingBananas

func minEatingSpeed(piles []int, h int) int {

	l := 1
	r := 1000000000

	for l < r {
		m := (l + r) / 2

		c := 0
		for _, v := range piles {
			c += ((v - 1) / m) + 1
			if c > h {
				break
			}
		}

		if c > h {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}
