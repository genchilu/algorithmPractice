package fourSumCount

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {

	m := make(map[int]int, len(nums1)*len(nums2))

	for _, i := range nums1 {
		for _, j := range nums2 {
			m[i+j]++
		}
	}

	c := 0
	for _, i := range nums3 {
		for _, j := range nums4 {
			c += m[0-i-j]
		}
	}
	return c
}
