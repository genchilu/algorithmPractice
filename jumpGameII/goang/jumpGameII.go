package jumpGameII

func jump(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}

	dp := make([]int, len(nums))

	r := back(&nums, &dp, len(nums)-1)

	return r
}

func back(nums, dp *[]int, pos int) int {

	if pos > 0 && (*dp)[pos] == 0 {
		for i := 1; (pos - i) >= 0; i++ {
			if (*nums)[pos-i] >= i {
				(*dp)[pos] = 1 + back(nums, dp, pos-i)
			}
		}
	}

	//fmt.Printf("pos: %d,  %v\n", pos, dp)

	return (*dp)[pos]
}
