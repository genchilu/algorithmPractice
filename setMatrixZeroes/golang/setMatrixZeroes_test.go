package setMatrixZeroes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//func setZeroes(matrix [][]int) {
func TestSetZeroes(t *testing.T) {
	testCases := []struct {
		matrix [][]int
		expect [][]int
	}{
		{
			[][]int{
				{1, 1, 1},
				{1, 0, 1},
				{1, 1, 1},
			},
			[][]int{
				{1, 0, 1},
				{0, 0, 0},
				{1, 0, 1},
			},
		},
		{
			[][]int{
				{0, 1, 2, 0},
				{3, 4, 5, 2},
				{1, 3, 1, 5},
			},
			[][]int{
				{0, 0, 0, 0},
				{0, 4, 5, 0},
				{0, 3, 1, 0},
			},
		},
	}

	for _, testCase := range testCases {
		setZeroes(testCase.matrix)
		assert.Equal(t, testCase.expect, testCase.matrix)
	}
}
