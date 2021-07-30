package gasStation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//func canCompleteCircuit(gas []int, cost []int) int {
func TestCanCompleteCircuit(t *testing.T) {
	testCases := []struct {
		gas    []int
		cost   []int
		expect int
	}{
		{
			[]int{1, 2, 3, 4, 5},
			[]int{3, 4, 5, 1, 2},
			3,
		},
		{
			[]int{2, 3, 4},
			[]int{3, 4, 3},
			-1,
		},
		{
			[]int{5, 1, 2, 3, 4},
			[]int{5, 4, 1, 5, 1},
			-1,
		},
		{
			[]int{1, 2, 3, 4, 5},
			[]int{3, 4, 5, 1, 2},
			3,
		},
		{
			[]int{7, 1, 0, 11, 4},
			[]int{5, 9, 1, 2, 5},
			3,
		},
	}

	for _, testCase := range testCases {
		actual := canCompleteCircuit(testCase.gas, testCase.cost)
		assert.Equal(t, testCase.expect, actual)
	}
}
