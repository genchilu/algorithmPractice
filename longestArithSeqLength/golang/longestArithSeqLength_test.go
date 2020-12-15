package longestArithSeqLength

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestArithSeqLength(t *testing.T) {
	teseCases := []struct {
		A      []int
		expect int
	}{
		{
			[]int{0},
			1,
		},
		{
			[]int{0, 0},
			2,
		},
		{
			[]int{3, 6, 9, 12},
			4,
		},
		{
			[]int{9, 4, 7, 2, 10},
			3,
		},
		{
			[]int{20, 1, 15, 3, 10, 5, 8},
			4,
		},
		{
			[]int{12, 28, 13, 6, 34, 36, 53, 24, 29, 2, 23, 0, 22, 25, 53, 34, 23, 50, 35, 43, 53, 11, 48, 56, 44, 53, 31, 6, 31, 57, 46, 6, 17, 42, 48, 28, 5, 24, 0, 14, 43, 12, 21, 6, 30, 37, 16, 56, 19, 45, 51, 10, 22, 38, 39, 23, 8, 29, 60, 18},
			4,
		},
		{
			[]int{44, 46, 22, 68, 45, 66, 43, 9, 37, 30, 50, 67, 32, 47, 44, 11, 15, 4, 11, 6, 20, 64, 54, 54, 61, 63, 23, 43, 3, 12, 51, 61, 16, 57, 14, 12, 55, 17, 18, 25, 19, 28, 45, 56, 29, 39, 52, 8, 1, 21, 17, 21, 23, 70, 51, 61, 21, 52, 25, 28},
			6,
		},
	}

	for _, testCase := range teseCases {
		actual := longestArithSeqLength(testCase.A)
		assert.Equal(t, testCase.expect, actual)
	}

}
