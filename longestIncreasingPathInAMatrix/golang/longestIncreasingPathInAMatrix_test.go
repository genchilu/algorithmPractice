package longestIncreasingPathInAMatrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//func longestIncreasingPath(matrix [][]int) int {
func TestDev(t *testing.T) {
	testCases := []struct {
		matrix [][]int
		expect int
	}{
		{
			[][]int{
				{9, 9, 4},
				{6, 6, 8},
				{2, 2, 1},
			},
			4,
		},
		{
			[][]int{
				{3, 4, 5},
				{3, 2, 6},
				{2, 2, 1},
			},
			4,
		},
		{
			[][]int{
				{1},
			},
			1,
		},
	}

	for _, testCase := range testCases {
		actual := longestIncreasingPath(testCase.matrix)
		assert.Equal(t, testCase.expect, actual)
	}
}
