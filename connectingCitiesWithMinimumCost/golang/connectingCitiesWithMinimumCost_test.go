package connectingCitiesWithMinimumCost

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//func minimumCost(n int, connections [][]int) int {
func TestMinimumCost(t *testing.T) {
	testCases := []struct {
		n           int
		connections [][]int
		expect      int
	}{
		{
			3,
			[][]int{
				{1, 2, 5},
				{1, 3, 6},
				{2, 3, 1},
			},
			6,
		},
		{
			4,
			[][]int{
				{1, 2, 3},
				{3, 4, 4},
			},
			-1,
		},
		{
			5,
			[][]int{
				{2, 1, 3267},
				{3, 2, 25910},
				{4, 1, 30518},
			},
			-1,
		},
	}

	for _, testCase := range testCases {
		actual := minimumCost(testCase.n, testCase.connections)
		assert.Equal(t, testCase.expect, actual)
	}
}
