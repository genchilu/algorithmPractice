package maximumNumberOfPointsWithCost

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDev(t *testing.T) {
	testCases := []struct {
		points [][]int
		expect int64
	}{
		{
			[][]int{
				{1, 2, 3},
				{1, 5, 1},
				{3, 1, 1},
			},
			9,
		},
		{
			[][]int{
				{1, 5},
				{2, 3},
				{4, 2},
			},
			11,
		},
	}

	for _, testCase := range testCases {
		actual := maxPoints(testCase.points)
		assert.Equal(t, testCase.expect, actual)
	}
}
