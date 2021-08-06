package minPathSum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//func minPathSum(grid [][]int) int {
func TestDev(t *testing.T) {
	testCases := []struct {
		grid   [][]int
		expect int
	}{
		{
			[][]int{
				{1, 3, 1},
				{1, 5, 1},
				{4, 2, 1},
			},
			7,
		},
		{
			[][]int{
				{1, 2, 3},
				{4, 5, 6},
			},
			12,
		},
		{
			[][]int{
				{0},
			},
			0,
		},
	}

	for _, testCase := range testCases {
		actual := minPathSum(testCase.grid)
		assert.Equal(t, testCase.expect, actual)
	}
}
