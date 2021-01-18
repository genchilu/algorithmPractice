package pairsOfSongsWithTotalDurationsDivisibleBy60

func numPairsDivisibleBy60(time []int) int {

	m := make(map[int][]int)
	for i, v := range time {

		r := v % 60
		if _, ok := m[r]; !ok {
			m[r] = []int{}
		}
		m[r] = append(m[r], i)
	}

	c := 0
	for i, v := range time {

		r := v % 60
		if r != 0 {
			r = 60 - r
		}
		if _, ok := m[r]; ok {
			for _, v := range m[r] {
				if v != i {
					c++
				}
			}
		}
	}
	return c / 2

}
