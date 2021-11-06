package singleNumberIII

func singleNumber(nums []int) []int {

	bitmask := 0
	for _, v := range nums {
		bitmask ^= v
	}

	diff := -bitmask & bitmask

	a := 0
	for _, v := range nums {
		if v&diff != 0 {
			a ^= v
		}
	}
	return []int{a, a ^ bitmask}
}
