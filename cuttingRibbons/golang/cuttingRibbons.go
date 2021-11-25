package cuttingRibbons

func maxLength(ribbons []int, k int) int {

	h := 0
	for _, v := range ribbons {
		if v > h {
			h = v
		}
	}

	m := 1

	result := 0
	for m <= h {
		t := (m + h) / 2
		count := 0
		for _, v := range ribbons {
			count += v / t
		}

		if count >= k {
			result = t
			m = t + 1
		} else {
			h = t - 1
		}
	}

	return result
}
