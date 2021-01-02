package permutations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//func permute(nums []int) [][]int {
func TestPermute(t *testing.T) {
	testCases := []struct {
		nums   []int
		expect [][]int
	}{
		{
			[]int{0},
			[][]int{
				[]int{0},
			},
		},
		{
			[]int{0, 1},
			[][]int{
				[]int{0, 1},
				[]int{1, 0},
			},
		},
		{
			[]int{0, 1, 2},
			[][]int{
				[]int{0, 1, 2},
				[]int{0, 2, 1},
				[]int{1, 0, 2},
				[]int{1, 2, 0},
				[]int{2, 0, 1},
				[]int{2, 1, 0},
			},
		},
		{
			[]int{0, 1, 2, 3},
			[][]int{
				[]int{0, 1, 2, 3},
				[]int{0, 1, 3, 2},
				[]int{0, 2, 1, 3},
				[]int{0, 2, 3, 1},
				[]int{0, 3, 1, 2},
				[]int{0, 3, 2, 1},

				[]int{1, 0, 2, 3},
				[]int{1, 0, 3, 2},
				[]int{1, 2, 0, 3},
				[]int{1, 2, 3, 0},
				[]int{1, 3, 0, 2},
				[]int{1, 3, 2, 0},

				[]int{2, 0, 1, 3},
				[]int{2, 0, 3, 1},
				[]int{2, 1, 0, 3},
				[]int{2, 1, 3, 0},
				[]int{2, 3, 0, 1},
				[]int{2, 3, 1, 0},

				[]int{3, 0, 1, 2},
				[]int{3, 0, 2, 1},
				[]int{3, 1, 0, 2},
				[]int{3, 1, 2, 0},
				[]int{3, 2, 0, 1},
				[]int{3, 2, 1, 0},
			},
		},
	}

	for _, testCase := range testCases {
		actual := permute(testCase.nums)
		for _, v := range actual {
			assert.Contains(t, testCase.expect, v)
		}
	}
}
