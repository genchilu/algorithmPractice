package theKthFactorOfN

import "math"

func kthFactor(n int, k int) int {

	sqrt := int(math.Sqrt(float64(n)))

	d := []int{}
	for i := 1; i <= sqrt; i++ {
		if n%i == 0 {
			d = append(d, i)
			k--

			if k == 0 {
				return i
			}
		}
	}

	if sqrt*sqrt == n {
		k++
	}

	if k > len(d) {
		return -1
	}

	idx := len(d) - k

	return n / d[idx]
}
