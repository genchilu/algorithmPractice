package searchA2DMatrixII

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchMatrix(t *testing.T) {
	testCases := []struct {
		matrix [][]int
		target int
		expect bool
	}{
		{
			[][]int{
				[]int{1, 4, 7, 11, 15},
				[]int{2, 5, 8, 12, 19},
				[]int{3, 6, 9, 16, 22},
				[]int{10, 13, 14, 17, 24},
				[]int{18, 21, 23, 26, 30},
			},
			5,
			true,
		},
		{
			[][]int{
				[]int{1, 4, 7, 11, 15},
				[]int{2, 5, 8, 12, 19},
				[]int{3, 6, 9, 16, 22},
				[]int{10, 13, 14, 17, 24},
				[]int{18, 21, 23, 26, 30},
			},
			20,
			false,
		},
		{
			[][]int{
				[]int{1, 3, 5},
			},
			2,
			false,
		},
		{
			[][]int{
				[]int{1},
				[]int{3},
				[]int{5},
			},
			2,
			false,
		},
	}

	for _, testCase := range testCases {
		actual := searchMatrix(testCase.matrix, testCase.target)
		assert.Equal(t, testCase.expect, actual)
	}
}
