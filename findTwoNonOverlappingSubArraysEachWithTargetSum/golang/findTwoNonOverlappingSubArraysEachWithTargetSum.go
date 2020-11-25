package findTwoNonOverlappingSubArraysEachWithTargetSum

import "math"

func minSumOfLengths(arr []int, target int) int {

	prefix1 := make([]int, len(arr))
	suffix1 := make([]int, len(arr))
	prefix2 := make([]int, len(arr))
	suffix2 := make([]int, len(arr))

	for i, _ := range arr {
		prefix1[i] = math.MaxInt32
		prefix2[i] = math.MaxInt32
		suffix1[i] = math.MaxInt32
		suffix2[i] = math.MaxInt32
	}

	for i := 0; i < len(arr); i++ {
		sum := 0
		for j := i; j < len(arr); j++ {
			sum += arr[j]
			if sum > target {
				break
			} else if sum == target {
				l := (j - i) + 1
				prefix1[i] = l
				suffix1[j] = l
			}
		}
	}

	mins := math.MaxInt32
	for i := 0; i < len(suffix1); i++ {
		if suffix1[i] < mins {
			mins = suffix1[i]
		}
		suffix2[i] = mins
	}

	minp := math.MaxInt32
	for i := len(prefix1) - 2; i >= 0; i-- {
		if prefix1[i+1] < minp {
			minp = prefix1[i+1]
		}
		prefix2[i] = minp
	}

	// fmt.Printf("%v\n", suffix1)
	// fmt.Printf("%v\n", prefix1)
	// fmt.Printf("%v\n", suffix2)
	// fmt.Printf("%v\n", prefix2)
	min := math.MaxInt32
	for i, _ := range suffix2 {
		if suffix2[i]+prefix2[i] < min {
			min = suffix2[i] + prefix2[i]
		}
	}

	if min >= math.MaxInt32 {
		return -1
	}
	return min
}
