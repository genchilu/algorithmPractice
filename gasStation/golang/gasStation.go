package gasStation

func canCompleteCircuit(gas []int, cost []int) int {

	total, cur, s := 0, 0, 0

	for i := range gas {
		total += (gas[i] - cost[i])
		cur += (gas[i] - cost[i])
		if cur < 0 {
			s = i + 1
			cur = 0
		}
	}

	if total >= 0 {
		return s
	}

	return -1
}
