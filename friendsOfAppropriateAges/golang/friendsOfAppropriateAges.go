package friendsOfAppropriateAges

func numFriendRequests(ages []int) int {
	aggregates := make([]int, 121)
	for _, age := range ages {
		aggregates[age]++
	}

	c := 0
	for ageA, countA := range aggregates {
		for ageB, countB := range aggregates {
			if ageB <= ageA/2+7 {
				continue
			} else if ageB > ageA {
				continue
			} else if ageB > 100 && ageA < 100 {
				continue
			}
			c += countA * countB
			if ageA == ageB {
				c -= countA
			}
		}
	}
	return c
}
