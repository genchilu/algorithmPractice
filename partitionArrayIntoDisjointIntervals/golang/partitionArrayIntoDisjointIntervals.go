package partitionArrayIntoDisjointIntervals

func partitionDisjoint(nums []int) int {
	value, max, idx, spidx, reset := nums[0], nums[0], 1, 0, true

	for ; idx < len(nums); idx++ {
		if nums[idx] > max {
			max = nums[idx]
		}

		if nums[idx] < value {
			reset = true
			value = max
		} else if nums[idx] >= value {

			if reset {
				spidx = idx
				reset = false
			}
		}
	}

	return spidx
}
