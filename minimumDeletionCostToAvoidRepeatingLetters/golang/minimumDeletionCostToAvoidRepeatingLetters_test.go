package minimumDeletionCostToAvoidRepeatingLetters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//func minCost(s string, cost []int) int {
func TestMinCost(t *testing.T) {
	testCases := []struct {
		s      string
		cost   []int
		expect int
	}{
		{
			"abaac",
			[]int{1, 2, 3, 4, 5},
			3,
		},
		{
			"abc",
			[]int{1, 2, 3},
			0,
		},
		{
			"aabaa",
			[]int{1, 2, 3, 4, 1},
			2,
		},
		{
			"bbbaaa",
			[]int{4, 9, 3, 8, 8, 9},
			23,
		},
	}

	for _, testCase := range testCases {
		actual := minCost(testCase.s, testCase.cost)
		assert.Equal(t, testCase.expect, actual)
	}
}
