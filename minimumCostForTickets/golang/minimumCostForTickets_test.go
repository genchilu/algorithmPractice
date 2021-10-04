package minimumCostForTickets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//func mincostTickets(days []int, costs []int) int {
func TestDev(t *testing.T) {
	testCases := []struct {
		days   []int
		costs  []int
		expect int
	}{
		{
			[]int{1, 4, 6, 7, 8, 20},
			[]int{2, 7, 15},
			11,
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 30, 31},
			[]int{2, 7, 15},
			17,
		},
	}

	for _, testCase := range testCases {
		actual := mincostTickets(testCase.days, testCase.costs)
		assert.Equal(t, testCase.expect, actual)
	}
}
