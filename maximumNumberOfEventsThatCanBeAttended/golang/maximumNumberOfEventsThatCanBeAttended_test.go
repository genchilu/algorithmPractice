package maximumNumberOfEventsThatCanBeAttended

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//func maxEvents(events [][]int) int {
func TestDev(t *testing.T) {
	testCases := []struct {
		events [][]int
		expect int
	}{
		{
			[][]int{
				{1, 2}, {2, 3}, {3, 4},
			},
			3,
		},
		{
			[][]int{
				{1, 2}, {2, 3}, {3, 4}, {1, 2},
			},
			4,
		},
		{
			[][]int{
				{1, 4}, {4, 4}, {2, 2}, {3, 4}, {1, 1},
			},
			4,
		},
		{
			[][]int{
				{1, 100000},
			},
			1,
		},
		{
			[][]int{
				{1, 1}, {1, 2}, {1, 3}, {1, 4}, {1, 5}, {1, 6}, {1, 7},
			},
			7,
		},
		{
			[][]int{
				{1, 2}, {1, 2}, {3, 3}, {1, 5}, {1, 5},
			},
			5,
		},
	}

	for _, testCase := range testCases {
		actual := maxEvents(testCase.events)
		assert.Equal(t, testCase.expect, actual)
	}
}
