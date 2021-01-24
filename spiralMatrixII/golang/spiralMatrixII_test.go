package spiralMatrixII

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//func generateMatrix(n int) [][]int {
func TestGenerateMatrix(t *testing.T) {
	testCases := []struct {
		n      int
		expect [][]int
	}{
		{
			3,
			[][]int{
				[]int{1, 2, 3},
				[]int{8, 9, 4},
				[]int{7, 6, 5},
			},
		},
		{
			1,
			[][]int{
				[]int{1},
			},
		},
	}

	for _, testCase := range testCases {
		actual := generateMatrix(testCase.n)
		assert.Equal(t, testCase.expect, actual)
	}
}
