package carFleet

import (
	"sort"
)

type Fleet struct {
	position    int
	arrive_time float32
}

func carFleet(target int, position []int, speed []int) int {

	if len(position) == 0 {
		return 0
	}
	fleets := make([]Fleet, len(position))
	for i, _ := range position {
		fleet := Fleet{position[i], float32(target-position[i]) / float32(speed[i])}
		fleets[i] = fleet
	}
	sort.SliceStable(fleets, func(i, j int) bool { return fleets[i].position > fleets[j].position })

	count := len(fleets)
	cur_arrive_time := fleets[0].arrive_time

	for i := 1; i < len(fleets); i++ {
		if fleets[i].arrive_time <= cur_arrive_time {
			count--
		} else {
			cur_arrive_time = fleets[i].arrive_time
		}
	}

	return count
}
