package searchMatrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//func searchMatrix(matrix [][]int, target int) bool {
func TestSearchMatrix(t *testing.T) {
	testCases := []struct {
		matrix [][]int
		target int
		expect bool
	}{
		{
			[][]int{
				[]int{1, 3, 5, 7},
				[]int{10, 11, 16, 20},
				[]int{23, 30, 34, 60},
			},
			3,
			true,
		},
		{
			[][]int{
				[]int{1, 3, 5, 7},
				[]int{10, 11, 16, 20},
				[]int{23, 30, 34, 60},
			},
			13,
			false,
		},
	}

	for _, testCase := range testCases {
		actual := searchMatrix(testCase.matrix, testCase.target)
		assert.Equal(t, testCase.expect, actual)
	}
}
