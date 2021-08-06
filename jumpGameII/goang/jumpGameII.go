package jumpGameII

func jump(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}

	cur, next, step := 0, 0, 0
	for cur < len(nums)-1 {
		step++

		if next >= len(nums)-1 {
			break
		}
		idx, max := 0, 0

		for pad := 1; pad <= nums[cur] && pad+cur < len(nums); pad++ {
			d := nums[pad+cur] + cur + pad
			// fmt.Printf("111 idx: %d, d: %d, max: %d\n", pad+cur, d, max)
			if d >= max || d >= len(nums)-1 {
				max = d
				idx = cur + pad
			}
		}

		cur = idx
		next = max
		//fmt.Printf("cur: %d, next: %d\n", cur, next)

	}

	return step
}
