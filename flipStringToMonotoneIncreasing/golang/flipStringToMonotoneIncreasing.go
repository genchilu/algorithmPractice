package flipStringToMonotoneIncreasing

func minFlipsMonoIncr(s string) int {

	zc1 := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			zc1++
		}
	}

	min := findmin(zc1, len(s)-zc1)
	zc2 := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			zc1--
			zc2++
		}
		f := (i + 1) - zc2 + zc1
		if f < min {
			min = f
		}
	}

	return min
}

func findmin(a, b int) int {
	if a < b {
		return a
	}

	return b
}
