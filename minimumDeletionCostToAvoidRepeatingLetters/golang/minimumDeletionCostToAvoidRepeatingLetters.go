package minimumDeletionCostToAvoidRepeatingLetters

func minCost(s string, cost []int) int {

	totalCost := 0
	var latest byte
	max := 0
	costsum := 0
	c := 0
	for i := 0; i < len(s); i++ {
		if s[i] != latest {
			if c > 1 {
				totalCost += costsum - max
			}

			max = cost[i]
			costsum = cost[i]
			latest = s[i]
			c = 1
		} else {
			c++
			if cost[i] > max {
				max = cost[i]
			}
			costsum += cost[i]
		}

	}

	if c > 1 {
		totalCost += costsum - max
	}
	return totalCost
}
