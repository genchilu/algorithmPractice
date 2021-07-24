package kokoEatingBananas

import (
	"math"
	"sort"
)

func minEatingSpeed(piles []int, h int) int {
	sum := 0
	for _, v := range piles {
		sum += v
	}

	sort.Ints(piles)

	min := int(math.Ceil(float64(sum) / float64(h)))

	for {
		tmp := h
		for i := len(piles) - 1; i >= 0; i-- {
			c := int(math.Ceil(float64(piles[i]) / float64(min)))
			tmp -= c
			if tmp < 0 {
				min++
				break
			}
		}

		if tmp >= 0 {
			break
		}
	}
	return min
}
