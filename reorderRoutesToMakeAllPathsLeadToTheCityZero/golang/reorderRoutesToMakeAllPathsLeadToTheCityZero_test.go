package reorderRoutesToMakeAllPathsLeadToTheCityZero

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//func minReorder(n int, connections [][]int) int {
func TestMinReorder(t *testing.T) {
	testCases := []struct {
		n           int
		connections [][]int
		expect      int
	}{
		{
			6,
			[][]int{
				{0, 1},
				{1, 3},
				{2, 3},
				{4, 0},
				{4, 5},
			},
			3,
		},
		{
			5,
			[][]int{
				{1, 0},
				{1, 2},
				{3, 2},
				{3, 4},
			},
			2,
		},
		{
			3,
			[][]int{
				{1, 0},
				{2, 0},
			},
			0,
		},
	}

	for _, testCase := range testCases {
		actual := minReorder(testCase.n, testCase.connections)
		assert.Equal(t, testCase.expect, actual)
	}
}
