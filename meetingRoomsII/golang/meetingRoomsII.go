package meetingRoomsII

import (
	"sort"
)

func minMeetingRooms(intervals [][]int) int {

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	rooms := []int{}

	for _, interval := range intervals {
		found := false

		for i, _ := range rooms {
			if rooms[i] <= interval[0] {
				found = true
				rooms[i] = interval[1]
				break
			}
		}

		if !found {
			rooms = append(rooms, interval[1])
		}
	}

	return len(rooms)
}
