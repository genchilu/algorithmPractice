package maximumNumberOfEventsThatCanBeAttended

import (
	"sort"
)

func maxEvents(events [][]int) int {

	sort.SliceStable(events, func(i, j int) bool {
		return events[i][0] < events[j][0]
	})

	idx := 0
	end := []int{}
	count := 0
	for i := 0; i <= 100000; i++ {
		for idx < len(events) {
			if events[idx][0] == i {
				e := events[idx][1]
				insert := sort.SearchInts(end, e)
				end = append(end, 0)
				copy(end[insert+1:], end[insert:])
				end[insert] = e
			} else if events[idx][0] > i {
				break
			}
			idx++
		}

		j := 0
		for ; j < len(end); j++ {
			if end[j] >= i {
				break
			}
		}
		end = end[j:]

		if len(end) > 0 {
			count++
			end = end[1:]
		}

		if idx >= len(events) && len(end) == 0 {
			break
		}
	}
	return count
}
