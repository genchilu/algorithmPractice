package maximumPointsYouCanObtainFromCards

func maxScore(cardPoints []int, k int) int {
	s := 0
	l := len(cardPoints)
	for i := 0; i < k; i++ {
		s += cardPoints[i]
	}

	cs := s
	for i := 0; i < k; i++ {
		cs -= cardPoints[k-i-1]
		cs += cardPoints[l-i-1]
		if cs > s {
			s = cs
		}
	}
	return s
}
