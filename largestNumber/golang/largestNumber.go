package largestNumber

import (
	"sort"
	"strconv"
)

func largestNumber(nums []int) string {

	sort.Slice(nums, func(i, j int) bool {
		return compare(nums[i], nums[j])
	})

	result := ""

	for _, v := range nums {

		if len(result) != 0 || v != 0 {
			s := strconv.Itoa(v)
			result += s
		}
	}

	if len(result) == 0 {
		return "0"
	}
	return result
}

func compare(i, j int) bool {
	s1 := strconv.Itoa(i)
	s2 := strconv.Itoa(j)

	t1, _ := strconv.Atoi(s1 + s2)
	t2, _ := strconv.Atoi(s2 + s1)

	return t1 > t2
}
