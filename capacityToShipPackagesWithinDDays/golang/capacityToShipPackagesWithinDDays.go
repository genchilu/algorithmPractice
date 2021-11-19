package capacityToShipPackagesWithinDDays

func shipWithinDays(weights []int, days int) int {

	sum := 0
	max := 0
	for _, v := range weights {
		sum += v
		if v > max {
			max = v
		}
	}

	for max < sum {
		split := 1
		mid := (max + sum) / 2
		cur := 0

		for _, w := range weights {
			if (cur + w) > mid {
				cur = w
				split++
			} else {
				cur += w
			}
		}

		if split > days {
			max = mid + 1
		} else {
			sum = mid
		}
	}

	return max
}
