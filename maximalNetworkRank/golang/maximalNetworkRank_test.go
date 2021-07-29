package maximalNetworkRank

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//maximalNetworkRank(n int, roads [][]int) int {
func TestMaximalNetworkRank(t *testing.T) {
	testCases := []struct {
		n      int
		roads  [][]int
		expect int
	}{
		{
			0,
			[][]int{
				{0, 1},
				{0, 3},
				{1, 2},
				{1, 3},
			},
			4,
		},
		{
			0,
			[][]int{
				{0, 1},
				{0, 3},
				{1, 2},
				{1, 3},
				{2, 3},
				{2, 4},
			},
			5,
		},
		{
			0,
			[][]int{
				{0, 1},
				{1, 2},
				{2, 3},
				{2, 4},
				{5, 6},
				{5, 7},
			},
			5,
		},
	}

	for _, testCase := range testCases {
		actual := maximalNetworkRank(testCase.n, testCase.roads)
		assert.Equal(t, testCase.expect, actual)
	}
}
