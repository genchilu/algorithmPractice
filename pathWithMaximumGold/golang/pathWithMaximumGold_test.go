package pathWithMaximumGold

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//func getMaximumGold(grid [][]int) int {

func TestGetMaximumGold(t *testing.T) {
	testCases := []struct {
		grid   [][]int
		expect int
	}{
		{
			[][]int{
				{0, 6, 0},
				{5, 8, 7},
				{0, 9, 0},
			},
			24,
		},
		{
			[][]int{
				{1, 0, 7},
				{2, 0, 6},
				{3, 4, 5},
				{0, 3, 0},
				{9, 0, 20},
			},
			28,
		},
	}

	for _, testCase := range testCases {
		actual := getMaximumGold(testCase.grid)
		assert.Equal(t, testCase.expect, actual)
	}
}
