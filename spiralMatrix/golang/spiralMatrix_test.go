package spiralMatrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//func spiralOrder(matrix [][]int) []int {
func TestSpiralOrder(t *testing.T) {
	testCases := []struct {
		matrix [][]int
		expect []int
	}{
		{
			[][]int{
				[]int{1, 2, 3},
				[]int{4, 5, 6},
				[]int{7, 8, 9},
			},
			[]int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			[][]int{[]int{1}},
			[]int{1},
		},
		{
			[][]int{
				[]int{1, 2, 3, 4},
				[]int{5, 6, 7, 8},
				[]int{9, 10, 11, 12},
			},
			[]int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
	}

	for _, testCase := range testCases {
		actual := spiralOrder(testCase.matrix)
		assert.Equal(t, testCase.expect, actual)
	}
}
