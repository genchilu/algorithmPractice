package sortColors

func sortColors(nums []int) {
	l, c, r := 0, 0, len(nums)-1

	for c <= r {
		if nums[c] == 2 {
			nums[c], nums[r] = nums[r], nums[c]
			r--
		} else if nums[c] == 0 {
			nums[l], nums[c] = nums[c], nums[l]
			l++
			c++
		} else {
			c++
		}
	}

}
