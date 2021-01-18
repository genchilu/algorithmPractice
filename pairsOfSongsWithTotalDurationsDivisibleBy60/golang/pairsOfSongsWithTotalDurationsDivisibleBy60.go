package pairsOfSongsWithTotalDurationsDivisibleBy60

func numPairsDivisibleBy60(time []int) int {

	m := make([]int, 60)

	c := 0
	for _, v := range time {

		r := v % 60
		if r == 0 {
			c += m[0]
		} else {
			c += m[60-r]
		}
		m[r]++
	}
	return c

}
