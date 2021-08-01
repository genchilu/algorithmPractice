package findPeakElement

func findPeakElement(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	l, r := 0, len(nums)-1
	for l < r {
		//fmt.Printf("l: %d, r: %d\n", l, r)
		m := (l + r) / 2
		if (m == 0 || nums[m] > nums[m-1]) && (m == len(nums)-1 || nums[m] > nums[m+1]) {
			return m
		}

		if m == 0 || (nums[m-1] < nums[m]) {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}
