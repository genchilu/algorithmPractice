package maxChunksToSorted

import (
	"sort"
)

func maxChunksToSorted(arr []int) int {
	cur := arr[0]
	maxs := []int{}
	for i := 1; i < len(arr); i++ {
		if arr[i] > cur {

			if len(maxs) == 0 || cur != maxs[len(maxs)-1] {
				maxs = append(maxs, cur)
			}
			cur = arr[i]
		} else if arr[i] < cur {
			idx := sort.SearchInts(maxs, arr[i])
			if idx != len(maxs) {
				maxs[idx] = cur
				maxs = maxs[0 : idx+1]
			}
		}
	}

	idx := sort.SearchInts(maxs, arr[len(arr)-1])
	if idx != len(maxs) {
		maxs[idx] = cur
		maxs = maxs[0 : idx+1]
	} else {
		maxs = append(maxs, cur)
	}

	//fmt.Printf("%v\n", maxs)
	return len(maxs)
}
