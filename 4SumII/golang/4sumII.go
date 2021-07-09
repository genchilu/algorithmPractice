package fourSumCount

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {

	m := make(map[int]int)

	for _, i := range nums1 {
		for _, j := range nums2 {
			s := i + j
			if _, ok := m[s]; !ok {
				m[s] = 0
			}
			m[s] = m[s] + 1
		}
	}

	c := 0
	for _, i := range nums3 {
		for _, j := range nums4 {
			s := 0 - (i + j)
			if _, ok := m[s]; ok {
				c += m[s]
			}
		}
	}
	return c
}
