package jumpGameII

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//func jump(nums []int) int {
func TestDev(t *testing.T) {
	testCases := []struct {
		nums   []int
		expect int
	}{
		{
			[]int{2, 3, 1, 1, 4},
			2,
		},
		{
			[]int{2, 3, 0, 1, 4},
			2,
		},
		{
			[]int{1, 2},
			1,
		},
		{
			[]int{1, 2, 3},
			2,
		},
		{
			[]int{3, 2, 1},
			1,
		},
		{
			[]int{2, 3, 1},
			1,
		},
		{
			[]int{2, 3, 1, 1, 4},
			2,
		},
	}

	for _, testCase := range testCases {
		actual := jump(testCase.nums)
		assert.Equal(t, testCase.expect, actual)
	}
}
