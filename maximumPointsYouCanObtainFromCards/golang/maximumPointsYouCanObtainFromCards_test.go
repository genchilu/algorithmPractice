package maximumPointsYouCanObtainFromCards

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//maxScore(cardPoints []int, k int) int

func TestMaxScore(t *testing.T) {
	testCases := []struct {
		cardPoints []int
		k          int
		expect     int
	}{
		{
			[]int{1, 2, 3, 4, 5, 6, 1},
			3,
			12,
		},
		{
			[]int{2, 2, 2},
			2,
			4,
		},
		{
			[]int{9, 7, 7, 9, 7, 7, 9},
			7, 55,
		},
		{
			[]int{1, 1000, 1},
			1,
			1,
		}, {
			[]int{1, 79, 80, 1, 1, 1, 200, 1},
			3,
			202,
		},
		{
			[]int{100, 40, 17, 9, 73, 75},
			3,
			248,
		},
	}

	for _, testCase := range testCases {
		actual := maxScore(testCase.cardPoints, testCase.k)
		assert.Equal(t, testCase.expect, actual)
	}
}
